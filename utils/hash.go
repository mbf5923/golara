package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// generate random string and return sha256
func GenerateRandomString() string {
	randomString := make([]rune, 32)
	for i := range randomString {
		randomString[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	//merge b with timestamp
	timeString := time.Now().Unix()
	randomString = append(randomString, []rune(strconv.FormatInt(timeString, 10))...)

	//make sha256
	shaString := sha256.New()
	shaString.Write([]byte(string(randomString)))
	apiToken := hex.EncodeToString(shaString.Sum(nil))

	return apiToken
}

func Md5Hash(password string) string {
	//make md5 hash
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
