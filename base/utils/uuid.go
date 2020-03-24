package utils

import uuid "github.com/satori/go.uuid"

func Uuid() uuid.UUID {
	id, _ := uuid.NewV4()
	return id
}
