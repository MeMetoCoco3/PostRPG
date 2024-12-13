package main

import (
	"database/sql"

	"github.com/google/uuid"
)

type Role int

const (
	WARRIOR Role = iota
	WIZZARD
	ARCHER
)

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func ToNullUUID(u uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  u,
		Valid: u != uuid.Nil,
	}
}

func ToNullInt32(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: i != 0,
	}
}
