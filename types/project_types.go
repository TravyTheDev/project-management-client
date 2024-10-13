package types

import "time"

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

type ProjectReq struct {
	ParentID    int    `json:"parentID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	AssigneeID  int    `json:"assigneeID"`
	Urgency     int    `json:"urgency"`
	Notes       string `json:"notes"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type ProjectRes struct {
	Project *Project `json:"project"`
	User    *User    `json:"user"`
}
