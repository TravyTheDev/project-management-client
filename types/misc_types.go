package types

type SearchReq struct {
	Text string `json:"text"`
}

type SocketMessage struct {
	Body     string `json:"body"`
	RoomID   string `json:"roomID"`
	Username string `json:"username"`
	UserID   string `json:"userID"`
}
