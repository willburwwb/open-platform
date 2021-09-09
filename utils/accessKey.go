package utils

import (
	"encoding/json"
	"log"

	"github.com/UniqueStudio/open-platform/pkg"
	"github.com/pkg/errors"
)

// GenAccessKey is a function to gen AccessKey
func GenAccessKey(UserID string, IsAdmin bool) string {
	user := pkg.AccessUser{
		UserID:  UserID,
		IsAdmin: IsAdmin,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	AccessKey, err := Encrypt(string(rs))
	if err != nil {
		log.Fatal(err)
	}
	return AccessKey
}

// LoadAccessKey is a func to load AccessKey
func LoadAccessKey(AccessKey string) (userID string, isAdmin bool, err error) {
	if AccessKey == "" {
		return "", false, errors.New("Empyt AccessKey")
	}
	u := new(pkg.AccessUser)
	data, err := Decrypt(AccessKey)
	if err != nil {
		return "", false, err
	}
	if err := json.Unmarshal([]byte(data), &u); err != nil {
		return "", false, err
	}
	return u.UserID, u.IsAdmin, nil
}
