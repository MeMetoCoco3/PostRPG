package main

import (
	"PostRPG/internal/database"
	"context"
	"github.com/google/uuid"
)

type Weapon struct {
	ID          uuid.UUID `json:"id"`
	Damage      int       `json:"damage"`
	Role        Role      `json:"Reach"`
	Reach       int       `json:"reach"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewWeapon(name string, description string, damage int, reach int, role Role) *Weapon {
	return &Weapon{
		Name:        name,
		Description: description,
		Damage:      damage,
		Reach:       reach,
		Role:        role,
	}
}

func (w *Weapon) UploadToDb(db *database.Queries) (*Weapon, error) {
	ctx := context.Background()
	data, err := db.CreateNewWeapon(ctx, w.ToParams())
	err = DealWithError(err, "Uploading weapon to database")

	newWeapon := ParamsToWeapon(data)
	return newWeapon, err
}

func (w *Weapon) ToParams() database.CreateNewWeaponParams {
	return database.CreateNewWeaponParams{
		Damage:      int32(w.Damage),
		Reach:       int32(w.Reach),
		Name:        w.Name,
		Description: w.Description,
		Role:        int32(w.Role),
	}
}

func ParamsToWeapon(data database.Weapon) *Weapon {
	return &Weapon{
		ID:          data.ID,
		Damage:      int(data.Damage),
		Reach:       int(data.Reach),
		Name:        data.Name,
		Description: data.Description,
		Role:        Role(data.Role),
	}

}
