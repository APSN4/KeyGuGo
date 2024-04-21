package database

import (
	"encoding/json"
	"fmt"
	"keygugo/constant"
	"keygugo/database/dao"
	"keygugo/internal"
	"keygugo/security"
	"log"
	"os"
	"path/filepath"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func GetData() []byte {
	dbPathFile := filepath.Join(constant.GetFolderPath(), "storage.db")
	if FileExists(dbPathFile) {
		fmt.Println("Файл существует:", dbPathFile)
	} else {
		fmt.Println("Файл не существует:", dbPathFile)
		err := internal.CreateKeyHash(constant.Key)
		if err != nil {
			fmt.Println(err)
		}
	}
	dbPath, _ := dao.CreateDB()

	fmt.Println(dbPath)
	data, err := dao.GetDataDB()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func GetDataKeys(data map[string]string) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}

func GetDataValue(id int) string {
	dataStruct := dao.DataDB{}
	jsonData := GetData()
	err := json.Unmarshal(jsonData, &dataStruct)
	if err != nil {
		fmt.Println("err", err)
		return "err"
	}
	dataValue := dataStruct.Title
	return dataValue[id]
}

func AddNote(title string, content string) (result bool) {
	r, err := dao.AddNoteDB(title, content)
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func DeleteNoteDBStorage(id int) (result bool) {
	r, err := dao.DeleteNoteDB(id)
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func UpdateNote(id int, title string, content string) (result bool) {
	r, err := dao.EditNoteDB(id, title, content)
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func LoadData() {
	dbPath := filepath.Join(constant.GetFolderPath(), "storage.db")
	bs := make([]byte, 32)
	s := constant.Key
	copy(bs, []byte(s))
	_, err := security.DecryptFile(dbPath, bs)
	if err != nil {
		fmt.Println(err)
	}
	data := GetData()
	fmt.Println(data)
}
