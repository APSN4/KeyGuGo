package internal

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"keygugo/constant"
	"os"
	"path/filepath"
)

func computeFileHash(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Преобразовать хеш-сумму в строку
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CreateKeyHash(key string) error {
	// Получить путь к файлу .hs
	filepathHS := filepath.Join(constant.GetFolderPath(), "hs_check.hs")

	// Вычислить хеш ключа
	hash := md5.New()
	hash.Write([]byte(key))
	keyHash := hex.EncodeToString(hash.Sum(nil))

	// Записать хеш ключа в файл
	if err := os.WriteFile(filepathHS, []byte(keyHash), 0644); err != nil {
		return err
	}

	return nil
}

func CheckKeyHash(key string) error {
	// Получить путь к файлу .hs
	filepathHS := filepath.Join(constant.GetFolderPath(), "hs_check.hs")

	// Прочитать ожидаемый хеш из файла .hs
	expectedHash, err := ioutil.ReadFile(filepathHS)
	if err != nil {
		return err
	}

	// Вычислить хеш ключа
	hash := md5.New()
	hash.Write([]byte(key))
	keyHash := hex.EncodeToString(hash.Sum(nil))

	// Сравнить хеш ключа с ожидаемым хешем из файла .hs
	if keyHash != string(expectedHash) {
		return errors.New("хеш ключа не совпадает с ожидаемым хешем из файла .hs")
	}

	return nil
}

func ReadFileHash() (string, error) {
	filepathHS := filepath.Join(constant.GetFolderPath(), "hs_check.hs")

	data, err := os.ReadFile(filepathHS)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
