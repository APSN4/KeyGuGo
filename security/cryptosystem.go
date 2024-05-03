package security

import "github.com/APSN4/bencrypt/bc"

func EncryptFile(filePath string, aesKey []byte) (r bool, err error) {
	err = bc.EncryptFile(filePath, filePath, aesKey)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DecryptFile(filePath string, aesKey []byte) (r bool, err error) {
	err = bc.DecryptFile(filePath, filePath, aesKey)
	if err != nil {
		return false, err
	}
	return true, nil
}
