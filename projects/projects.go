package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ProjectsHandler struct {
	projectsStore *ProjectsStore
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserSlice struct {
	Users []User `json:"users"`
}

type SearchReq struct {
	Text string `json:"text"`
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

func NewProjectsHandler(projectsStore *ProjectsStore) *ProjectsHandler {
	return &ProjectsHandler{
		projectsStore: projectsStore,
	}
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
	layout := "2006-01-02 15:04:05"
	projectStart, err := time.Parse(layout, project.StartDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	projectEnd, err := time.Parse(layout, project.EndDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	project.StartDate = projectStart.String()
	project.EndDate = projectEnd.String()
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

func (p *ProjectsHandler) GetProjectByID(id string) *Project {
	var project *Project
	str := fmt.Sprintf("http://localhost:8000/api/v1/projects/project/%s", id)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		fmt.Println(err)
	}
	return project
}

func (p *ProjectsHandler) CreateNotes(projectID int, notes string) {
	if err := p.projectsStore.CreatePersonalNotes(projectID, notes); err != nil {
		fmt.Println(err)
	}
}

func (p *ProjectsHandler) GetNotesByProjectID(projectID int) string {
	notes, err := p.projectsStore.GetPersonalNotesByProjectID(projectID)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return notes
}

func (p *ProjectsHandler) EditPersonalNotes(projectID int, notes string) {
	if err := p.projectsStore.EditPersonalNotesByProjectID(projectID, notes); err != nil {
		fmt.Println(err)
	}
}

func (p *ProjectsHandler) SearchProjectAssignee(search SearchReq) []*User {
	users := make([]*User, 0)
	str := fmt.Sprintf("http://localhost:8000/api/v1/search_username/%s", search.Text)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Println("THERE", err)
	}
	return users
}
