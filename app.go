package main

import (
	"context"
	"encoding/json"
	"fmt"
	"keygugo/constant"
	"keygugo/database"
	"keygugo/database/dao"
	"keygugo/internal"
	"keygugo/security"
	"os"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")
	bs := make([]byte, 32)
	s := constant.Key
	copy(bs, []byte(s))
	_, err := security.EncryptFile(dbPath, bs)
	if err != nil {
		fmt.Println(err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetDB() string {
	return string(database.GetData())
}

func (a *App) AddNoteToDB(jsonData string) {
	type localData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	data := localData{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(data)
	database.AddNote(data.Title, data.Content)
}

func (a *App) DeleteNoteFromDB(id int) {
	dao.DeleteNoteDB(id)
}

func (a *App) EditNoteToDB(id int, title string, content string) {
	database.UpdateNote(id, title, content)
}

func (a *App) GetKey(key string) bool {
	constant.Key = key
	constant.HashSum, _ = internal.ReadFileHash()
	fmt.Println(constant.Key, constant.HashSum)
	err := internal.CheckKeyHash(constant.Key)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			internal.CreateKeyHash(constant.Key)
		} else {
			return false
		}
	}
	database.LoadData()
	return true
}
