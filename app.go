package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"project-management-client/cookie"
	"time"

	"github.com/r3labs/sse/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	httpClient   = &http.Client{}
	retries      = 3
	cookieUrl, _ = url.Parse("http://localhost:8000")
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Notification struct {
	Message string `json:"message"`
}

type Project struct {
	ID          int       `json:"id"`
	ParentID    int       `json:"parentID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	AssigneeID  int       `json:"assigneeID"`
	Urgency     int       `json:"urgency"`
	Notes       string    `json:"notes"`
	StartDate   string    `json:"startDate"`
	EndDate     string    `json:"endDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// App struct
type App struct {
	ctx         context.Context
	cookieStore *cookie.CookieStore
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
func (a *App) Login(email string, password string) {
	reqBody := LoginReq{
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

	httpClient.Jar.SetCookies(cookieUrl, resp.Cookies())
	a.cookieStore.SaveCookies(resp.Cookies())

}

func (a *App) GetUser() *User {
	data := &User{}
	var res *http.Response
	if len(httpClient.Jar.Cookies(cookieUrl)) == 0 {
		return &User{}
	}

	//SERVERSIDE LOCALHOST SECURECOOKIES FALSE
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
	notification := &Notification{}
	connectStr := fmt.Sprintf("http://localhost:8000/api/v1/notifications/stream/%d", id)

	client := sse.NewClient(connectStr)
	client.SubscribeRaw(func(msg *sse.Event) {
		if err := json.Unmarshal(msg.Data, &notification); err != nil {
			fmt.Println(err)
		}
		runtime.EventsEmit(a.ctx, "notification", notification)
	})
}
