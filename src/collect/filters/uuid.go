package collect

import uuid "github.com/satori/go.uuid"

func Uuid() string {
	u4 := uuid.NewV4()
	return u4.String()
	// return ""
}
