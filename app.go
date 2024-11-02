package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"project-management-client/cookie"
	"project-management-client/types"
	"syscall"

	osRuntime "runtime"

	"github.com/gorilla/websocket"
	"github.com/r3labs/sse/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	httpClient   = &http.Client{}
	retries      = 3
	cookieUrl, _ = url.Parse("http://localhost:8000")
)

// App struct
type App struct {
	ctx         context.Context
	cookieStore *cookie.CookieStore
	conn        *websocket.Conn
}

// NewApp creates a new App application struct
func NewApp(cookieStore *cookie.CookieStore) *App {
	return &App{
		cookieStore: cookieStore,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	jar, _ := cookiejar.New(nil)
	httpClient = &http.Client{
		Timeout: 0,
		Jar:     jar,
	}
	a.getCookies()
}

// Greet returns a greeting for the given name
func (a *App) Login(email string, password string) int {
	var id int
	reqBody := types.LoginReq{
		Email:    email,
		Password: password,
	}
	encodeBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.Post(
		"http://localhost:8000/api/v1/login/en",
		"Content-Type: application/json",
		bytes.NewReader(encodeBody),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&id); err != nil {
		return 0
	}
	httpClient.Jar.SetCookies(cookieUrl, resp.Cookies())
	a.cookieStore.SaveCookies(resp.Cookies())
	return id
}

func (a *App) GetUser() *types.User {
	data := &types.User{}
	var res *http.Response
	if len(httpClient.Jar.Cookies(cookieUrl)) == 0 {
		return &types.User{}
	}

	//SERVERSIDE localhost SECURECOOKIES FALSE
	res = a.getMeRequest()
	defer res.Body.Close()
	if res.StatusCode == 401 {
		retries--
		a.renewCookie()
		res = a.getMeRequest()
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	go a.notify(data.ID)
	return data
}

func (a *App) Logout() {
	a.cookieStore.DeleteCookies()
	httpClient.Jar.SetCookies(cookieUrl, nil)
	RestartSelf()
}

func (a *App) Register(registerReq types.RegisterReq) {
	payload, err := json.Marshal(registerReq)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8000/api/v1/register/en", bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}

	_, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *App) CreateWebSocketRoom(roomID int) {
	encodeBody, err := json.Marshal(roomID)
	if err != nil {
		fmt.Println(err)
	}

	reader := bytes.NewReader(encodeBody)
	_, err = http.Post(
		"http://localhost:8000/api/v1/ws/create_room",
		"Content-Type: application/json",
		reader,
	)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *App) JoinWebSocketRoom(roomID int, userID int, username string) {
	msg := make(chan types.SocketMessage)
	msgData := &types.SocketMessage{}

	connectStr := fmt.Sprintf("ws://localhost:8000/api/v1/ws/join_room/%d/%d/%s", roomID, userID, username)
	c, _, err := websocket.DefaultDialer.Dial(connectStr, nil)
	if err != nil {
		fmt.Println("DAIL:", err)
	}
	a.conn = c
	defer a.conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := a.conn.ReadMessage()
			if err != nil {
				log.Println("READ:", err)
				break
			}
			if err := json.Unmarshal(message, &msgData); err != nil {
				fmt.Println(err)
			}
			runtime.EventsEmit(a.ctx, "websocket", msgData)
		}
	}()

	for {
		select {
		case <-done:
			return
		case m := <-msg:
			err := a.conn.WriteJSON(m)
			if err != nil {
				log.Println("WRITE:", err)
				return
			}
		}
	}
}

func (a *App) ReadSocketMessage() {
	_, m, err := a.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			fmt.Printf("error: %v", err)
		}
	}
	var message types.SocketMessage
	if err := json.Unmarshal(m, &message); err != nil {
		fmt.Printf("error: %v", err)
	}
}

func (a *App) DisconnectWebSocketRoom() {
	if err := a.conn.Close(); err != nil {
		fmt.Println(err)
	}
}

func (a *App) WriteSocketMessage(message types.SocketMessage) {
	msg, err := json.Marshal(message.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err := a.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		fmt.Println(err)
	}
}

func (a *App) SendNotification(id int, notfication types.Notification) {
	encodeBody, err := json.Marshal(notfication)
	if err != nil {
		fmt.Println(err)
	}

	reader := bytes.NewReader(encodeBody)
	connectStr := fmt.Sprintf("http://localhost:8000/api/v1/notifications/send/%d", id)
	_, err = http.Post(
		connectStr,
		"Content-Type: application/json",
		reader,
	)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *App) renewCookie() {
	if err := a.cookieStore.DeleteAuthCookie(); err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8000/api/v1/renew_token", nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if err := a.cookieStore.SaveCookies(res.Cookies()); err != nil {
		fmt.Println(err)
	}
	httpClient.Jar.SetCookies(cookieUrl, res.Cookies())
}

func (a *App) getMeRequest() *http.Response {
	if retries == 0 {
		return nil
	}
	a.getCookies()
	req, err := http.NewRequest("GET", "http://localhost:8000/api/v1/me", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func (a *App) getCookies() {
	cookieSlice, err := a.cookieStore.GetCookies()
	if err != nil {
		fmt.Println(err)
	}
	httpClient.Jar.SetCookies(cookieUrl, cookieSlice)
}

func (a *App) notify(id int) {
	notification := &types.Notification{}
	connectStr := fmt.Sprintf("http://localhost:8000/api/v1/notifications/stream/%d", id)

	client := sse.NewClient(connectStr)
	client.SubscribeRaw(func(msg *sse.Event) {
		if err := json.Unmarshal(msg.Data, &notification); err != nil {
			fmt.Println(err)
		}
		runtime.EventsEmit(a.ctx, "notification", notification)
	})
}

func RestartSelf() error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	args := os.Args
	env := os.Environ()
	// Windows does not support exec syscall.
	if osRuntime.GOOS == "windows" {
		cmd := exec.Command(self, args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Env = env
		err := cmd.Start()
		if err == nil {
			os.Exit(0)
		}
		return err
	}
	return syscall.Exec(self, args, env)
}
