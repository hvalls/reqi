package db

import (
	"database/sql"
	"os"
	"os/user"

	_ "github.com/mattn/go-sqlite3"

	"reqi/requesttpl"
)

func SaveRequestTpl(tpl requesttpl.RequestTpl) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	insertTpl, _ := db.Prepare("INSERT INTO request_templates(name, description, url, method, body) values(?, ?, ?, ?, ?)")
	insertTpl.Exec(tpl.Name, tpl.Description, tpl.URL, tpl.Method, tpl.Body)
	return nil
}

func GetRequestTpl(tplName string) (*requesttpl.RequestTpl, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	rows, _ := db.Query("SELECT name, description, url, method, body FROM request_templates WHERE name=?", tplName)
	rows.Next()
	var name, description, url, method, body string
	rows.Scan(&name, &description, &url, &method, &body)
	return requesttpl.New(name, description, url, method, body), nil
}

func getDB() (*sql.DB, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	reqiPath := usr.HomeDir + "/.reqi"
	dbPath := reqiPath + "/reqi.db"
	var db *sql.DB
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		db, err = initDB(reqiPath, dbPath)
		if err != nil {
			return nil, err
		}
	} else {
		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func initDB(reqiPath, dbPath string) (*sql.DB, error) {
	if _, err := os.Stat(reqiPath); os.IsNotExist(err) {
		os.Mkdir(reqiPath, 0700)
	}
	var db *sql.DB
	os.Create(dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE "request_templates" (
		"name"	TEXT NOT NULL,
		"description"	TEXT NOT NULL,
		"url"	TEXT NOT NULL,
		"method"	TEXT NOT NULL,
		"body"	TEXT NOT NULL,
		PRIMARY KEY("name")
	);`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
