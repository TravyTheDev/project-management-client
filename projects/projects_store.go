package projects

import (
	"database/sql"
)

type PersonalNotes struct {
	projectID int
	notes     string
}

type ProjectsStore struct {
	db *sql.DB
}

func NewProjectsStore(db *sql.DB) *ProjectsStore {
	return &ProjectsStore{
		db: db,
	}
}

func (p *ProjectsStore) CreatePersonalNotes(projectID int, notes string) error {
	stmt := `INSERT INTO personal_notes (project_id, notes) VALUES (?, ?)`

	_, err := p.db.Exec(stmt, projectID, notes)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectsStore) GetPersonalNotesByProjectID(projectID int) (string, error) {
	var notes PersonalNotes
	stmt := `SELECT * FROM personal_notes WHERE project_id = ?`

	rows, err := p.db.Query(stmt, projectID)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err := rows.Scan(
			&notes.projectID,
			&notes.notes,
		)
		if err != nil {
			return "", err
		}
	}
	return notes.notes, nil
}

func (p *ProjectsStore) EditPersonalNotesByProjectID(projectID int, notes string) error {
	stmt := `UPDATE personal_notes SET notes = ? WHERE project_id = ?`

	_, err := p.db.Exec(stmt, notes, projectID)
	if err != nil {
		return err
	}
	return nil
}
