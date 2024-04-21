package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"keygugo/constant"
	"log"
	"path/filepath"
)

type DataDB struct {
	ID      []int    `json:"id"`
	Title   []string `json:"title"`
	Content []string `json:"content"`
}

func CreateDB() (string, error) {

	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer db.Close()
	db.Exec("CREATE TABLE notes(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, content TEXT);")
	return dbPath, nil
}

func GetDataDB() (data []byte, err error) {
	localData := DataDB{}
	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := db.Query("SELECT id, title, content FROM notes")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name, note string
		var id int
		if err := rows.Scan(&id, &name, &note); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		localData.ID = append(localData.ID, id)
		localData.Title = append(localData.Title, name)
		localData.Content = append(localData.Content, note)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return data, err
	}
	jsonData, err := json.Marshal(localData)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
	}
	defer db.Close()
	return jsonData, nil
}

func AddNoteDB(title string, content string) (result bool, err error) {
	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO notes (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		log.Fatal(err)
	}
	result = true
	return
}

func DeleteNoteDB(id int) (result bool, err error) {
	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	result = true
	return
}

func EditNoteDB(id int, title string, content string) (result bool, err error) {
	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE notes SET title = ?, content = ? WHERE id = ?", title, content, id)
	if err != nil {
		log.Fatal(err)
	}
	result = true
	return
}
