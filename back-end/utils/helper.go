package utils

import "github.com/google/uuid"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateUUID() string {
	return uuid.New().String()
}
