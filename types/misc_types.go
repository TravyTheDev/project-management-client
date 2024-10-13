package types

type SearchReq struct {
	Text string `json:"text"`
}

type SocketMessage struct {
	UserID  string `json:"id"`
	Message string `json:"message"`
}
