package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type Role int

const (
	WARRIOR Role = iota
	WIZZARD
	ARCHER
)

type Action int

const (
	ATTACK Action = iota
	WEAPON
	SKILL
)

func DealWithError(err error, message string) error {
	if err != nil {
		// Log the error with additional context
		log.Printf("%s: %v", message, err)
		return fmt.Errorf("%s: %w", message, err)
	}
	return nil
}
func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// TODO: Valid may be u!=uuid.nil
func ToNullUUID(u uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  u,
		Valid: true,
	}
}

func ToNullInt32(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: i != 0,
	}
}
