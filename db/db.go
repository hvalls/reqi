package db

import (
	"database/sql"
	"os"
	"os/user"

	_ "github.com/mattn/go-sqlite3"

	"reqi/requesttpl"
)

func SaveRequestTpl(tpl requesttpl.RequestTpl) error {
	DeleteRequestTpl(tpl.Name)
	db, err := getDB()
	if err != nil {
		return err
	}
	insertTpl, err := db.Prepare("INSERT INTO request_templates(name, description, url, method, body) values(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = insertTpl.Exec(tpl.Name, tpl.Description, tpl.URL, tpl.Method, tpl.Body)
	return err
}

func DeleteRequestTpl(tplName string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	deleteTpl, err := db.Prepare("DELETE FROM request_templates WHERE name=?")
	if err != nil {
		return err
	}
	_, err = deleteTpl.Exec(tplName)
	return err
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
	rows.Close()
	return requesttpl.New(name, description, url, method, body), nil
}

func GetRequestTpls() ([]*requesttpl.RequestTpl, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT name, description, url, method, body FROM request_templates")
	if err != nil {
		return nil, err
	}
	var tpls []*requesttpl.RequestTpl
	var name, description, url, method, body string
	for rows.Next() {
		err := rows.Scan(&name, &description, &url, &method, &body)
		if err != nil {
			return nil, err
		}
		tpls = append(tpls, requesttpl.New(name, description, url, method, body))
	}
	rows.Close()
	return tpls, nil
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
