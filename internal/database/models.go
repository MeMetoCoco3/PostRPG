// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/google/uuid"
)

type Character struct {
	ID       uuid.UUID
	Health   int32
	Mana     int32
	Stamina  int32
	Strength int32
	Job      int32
	Name     string
	SkillID  uuid.NullUUID
	WeaponID uuid.NullUUID
	Icon     string
}

type Skill struct {
	ID          uuid.UUID
	Coin        string
	AmountToPay int32
	Damage      int32
	Role        int32
	Reach       int32
	Name        string
	Description string
}

type Terrain struct {
	ID     int32
	Matrix [][]int32
}

type Weapon struct {
	ID          uuid.UUID
	Name        string
	Description string
	Damage      int32
	Reach       int32
	Role        int32
}
