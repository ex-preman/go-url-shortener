package mysql

import (
	"database/sql"
)

type MysqlDB struct {
	db *sql.DB
}

func (m *MysqlDB) Init() error {
	var err error
	m.db, err = sql.Open("mysql", "root:toor@tcp(localhost:3306)/url_db?charset=utf8") //change this conn
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDB) Save(URL, code string) error {
	// perform a db.Query insert
	_, err := m.db.Query("INSERT INTO url (id, url, code, created_at) VALUES (NULL, ?, ?, CURRENT_TIMESTAMP)", URL, code)
	if err != nil {
		return err
	}

	return nil
}

func (m *MysqlDB) Load(code string) (string, error) {
	type Result struct {
		URL string `json:"url"`
	}

	// Execute the query
	var rs Result
	err := m.db.QueryRow("SELECT url FROM url where code = ?", code).Scan(&rs.URL)
	if err != nil {
		return "", err
	}

	return string(rs.URL), nil
}
