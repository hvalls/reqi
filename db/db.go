package db

import (
	"database/sql"
	"errors"
	"os"
	"os/user"

	_ "github.com/mattn/go-sqlite3"

	"reqi/requesttpl"
)

func SaveRequestTpl(tpl *requesttpl.RequestTpl) error {
	DeleteRequestTpl(tpl.Name)
	db, err := getDB()
	if err != nil {
		return err
	}
	insertTpl, err := db.Prepare("INSERT INTO request_templates(name, definition) values(?, ?)")
	if err != nil {
		return err
	}
	str, err := tpl.String()
	if err != nil {
		return err
	}
	_, err = insertTpl.Exec(tpl.Name, str)
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
	row := db.QueryRow("SELECT name, definition FROM request_templates WHERE name=?", tplName)
	var name, definition string
	switch err = row.Scan(&name, &definition); err {
	case nil:
	case sql.ErrNoRows:
		return nil, errors.New("template not found")
	default:
		return nil, err
	}
	return requesttpl.NewYaml([]byte(definition))
}

func GetRequestTpls() ([]*requesttpl.RequestTpl, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT name, definition FROM request_templates")
	if err != nil {
		return nil, err
	}
	var tpls []*requesttpl.RequestTpl
	var name, definition string
	for rows.Next() {
		err := rows.Scan(&name, &definition)
		if err != nil {
			return nil, err
		}
		tpl, err := requesttpl.NewYaml([]byte(definition))
		if err != nil {
			return nil, err
		}
		tpls = append(tpls, tpl)
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
		"name"			TEXT NOT NULL,
		"definition"	TEXT NOT NULL,
		PRIMARY KEY("name")
	);`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
