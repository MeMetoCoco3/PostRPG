package main

import (
	"PostRPG/internal/database"
	"context"
	"github.com/google/uuid"
)

type Skill struct {
	ID          uuid.UUID `json:"id"`
	AmountToPay int
	Coin        string
	Damage      int    `json:"damage"`
	Reach       int    `json:"Reach"`
	Role        Role   `json:"role"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewSkill(name string, description string, damage int, reach int, role Role, cost int, coin string) *Skill {
	return &Skill{
		Name:        name,
		Description: description,
		Damage:      damage,
		Reach:       reach,
		AmountToPay: cost,
		Coin:        coin,
		Role:        role,
	}
}

func (s *Skill) UploadToDb(db *database.Queries) (*Skill, error) {
	ctx := context.Background()
	data, err := db.CreateNewSkill(ctx, s.ToParams())
	err = DealWithError(err, "Uploading skill to database")

	newSkill := ParamsToSkill(data)
	return newSkill, err
}

func (s *Skill) ToParams() database.CreateNewSkillParams {
	return database.CreateNewSkillParams{
		Damage:      int32(s.Damage),
		Reach:       int32(s.Reach),
		Coin:        s.Coin,
		AmountToPay: int32(s.AmountToPay),
		Name:        s.Name,
		Description: s.Description,
		Role:        int32(s.Role),
	}
}

func ParamsToSkill(data database.Skill) *Skill {
	return &Skill{
		ID:          data.ID,
		Damage:      int(data.Damage),
		Reach:       int(data.Reach),
		Coin:        data.Coin,
		AmountToPay: int(data.AmountToPay),
		Name:        data.Name,
		Description: data.Description,
		Role:        Role(data.Role),
	}

}
