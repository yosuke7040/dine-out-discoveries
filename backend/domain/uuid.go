package domain

import ulidv2 "github.com/oklog/ulid/v2"

func NewULID() string {
	return ulidv2.MustNew(ulidv2.Now(), nil).String()
}

func IsValidULID(ulid string) bool {
	_, err := ulidv2.Parse(ulid)
	return err == nil
}
