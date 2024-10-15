package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project-management-client/types"
)

type ProjectsHandler struct {
	projectsStore *ProjectsStore
}

func NewProjectsHandler(projectsStore *ProjectsStore) *ProjectsHandler {
	return &ProjectsHandler{
		projectsStore: projectsStore,
	}
}

func (p *ProjectsHandler) GetProjectsByStatus(status int) []*types.ProjectRes {
	projects := []*types.ProjectRes{}
	str := fmt.Sprintf("http://localhost:8000/api/v1/projects/project_status/%d", status)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return projects
}

func (p *ProjectsHandler) GetProjectsByUrgency(urgency int) []*types.ProjectRes {
	projects := []*types.ProjectRes{}
	str := fmt.Sprintf("http://localhost:8000/api/v1/projects/project_urgency/%d", urgency)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return projects
}

func (p *ProjectsHandler) GetAllProjects() []*types.Project {
	projects := []*types.Project{}
	resp, err := http.Get("http://localhost:8000/api/v1/projects/all_projects")
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return projects
}

func (p *ProjectsHandler) CreateProject(project types.ProjectReq) {
	if project.Title == "" {
		return
	}
	if project.ParentID == 0 {
		project.ParentID = -1
	}
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

func (p *ProjectsHandler) GetProjectByID(id string) *types.ProjectRes {
	var project *types.ProjectRes
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

func (p *ProjectsHandler) SearchProjectAssignee(search types.SearchReq) []*types.User {
	users := make([]*types.User, 0)
	str := fmt.Sprintf("http://localhost:8000/api/v1/search_username/%s", search.Text)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil
	}
	return users
}

func (p *ProjectsHandler) EditProject(project types.Project) {
	encodeBody, err := json.Marshal(project)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("PUT", "http://localhost:8000/api/v1/projects/update_project", bytes.NewReader(encodeBody))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func (p *ProjectsHandler) GetChildProjectsByParentID(id int) []*types.Project {
	var projects []*types.Project
	str := fmt.Sprintf("http://localhost:8000/api/v1/projects/child_projects/%v", id)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		fmt.Println(err)
	}
	return projects
}

func (p *ProjectsHandler) DeleteProject(id int) {
	str := fmt.Sprintf("http://localhost:8000/api/v1/projects/delete_project/%d", id)
	req, err := http.NewRequest("DELETE", str, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func (p *ProjectsHandler) SearchProjects(search types.SearchReq) []*types.Project {
	projects := make([]*types.Project, 0)
	str := fmt.Sprintf("http://localhost:8000/api/v1/projects/search_projects/%s", search.Text)

	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil
	}
	return projects
}
