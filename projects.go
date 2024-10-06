package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ProjectsHandler struct{}

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

func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{}
}

func (p *ProjectsHandler) GetAllProjects() []*Project {
	projects := []*Project{}
	resp, err := http.Get("http://localhost:8000/api/v1/projects/all_projects")
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		fmt.Println(err)
	}
	return projects
}

func (p *ProjectsHandler) CreateProject(project ProjectReq) {
	encodeBody, err := json.Marshal(project)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := http.Post(
		"http://localhost:8000/api/v1/projects/create_project",
		"Content-Type: application/json",
		bytes.NewReader(encodeBody),
	); err != nil {
		fmt.Println(err)
	}
}
