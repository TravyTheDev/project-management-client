package cookie

import (
	"database/sql"
	"net/http"
)

type Cookies []*http.Cookie

type CookieStore struct {
	db *sql.DB
}

func NewCookieStore(db *sql.DB) *CookieStore {
	return &CookieStore{
		db: db,
	}
}

func (c *CookieStore) SaveCookies(cookieSlice []*http.Cookie) error {
	stmt := `INSERT INTO cookies (name, value) VALUES (?, ?)`

	for _, co := range cookieSlice {
		_, err := c.db.Exec(stmt, co.Name, co.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CookieStore) GetCookies() (Cookies, error) {
	cookies := make(Cookies, 0)
	stmt := `SELECT * FROM cookies`
	rows, err := c.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := new(http.Cookie)
		err := rows.Scan(
			&c.Name,
			&c.Value,
		)
		if err != nil {
			return nil, err
		}
		cookies = append(cookies, c)
	}
	return cookies, nil
}

func (c *CookieStore) DeleteCookies() error {
	stmt := `DELETE FROM cookies`
	_, err := c.db.Exec(stmt)
	return err
}

func (c *CookieStore) DeleteAuthCookie() error {
	var name = "authentication"
	stmt := `DELETE FROM cookies where name = ?`
	_, err := c.db.Exec(stmt, name)
	return err
}
