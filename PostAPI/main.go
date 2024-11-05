package main

import (
	"fmt"
)

type Role int

const (
	WARRIOR Role = iota
	WIZZARD
	ARCHER
)

type Skill struct {
	payment     []int
	damage      int
	distance    int
	canBeLearnt Role
	name        string
	description string
	payWith     []string
}

type Weapon struct {
	damage      int
	reach       int
	name        string
	description string
}

type Character struct {
	id        int
	health    int
	mana      int
	stamina   int
	strength  int
	job       Role
	jobString string
	name      string
	skill     Skill
	weapon    Weapon
	icon      rune
}

func NewCharacter(name string, job Role, icon rune) *Character {
	var role string
	var health, mana, stamina, strength int
	switch job {
	case 0:
		role = "WARRIOR"
		health = 100
		mana = 0
		stamina = 4
		strength = 10
	case 1:
		role = "WIZZARD"
		health = 30
		mana = 30
		stamina = 2
		strength = 2
	case 2:
		role = "ARCHER"
		health = 50
		mana = 0
		stamina = 4
		strength = 6
	default:
		panic("Not correct job")
	}

	c := &Character{
		health:    health,
		mana:      mana,
		stamina:   stamina,
		strength:  strength,
		job:       job,
		jobString: role,
		name:      name,
		icon:      icon,
	}
	return c
}

func (c *Character) PrintStats() {
	fmt.Println("🌟 Character Sheet 🌟")
	fmt.Println("--------------------")
	fmt.Printf("Name:           %s\n", c.name)
	fmt.Printf("Icon:           %c\n", c.icon)
	fmt.Printf("Job:            %s\n", c.jobString)
	fmt.Println()
	fmt.Println("    📊 Stats 📊")
	fmt.Printf("Health:             %d\n", c.health)
	fmt.Printf("Mana:               %d\n", c.mana)
	fmt.Printf("Stamina:            %d\n", c.stamina)
	fmt.Printf("Strength:           %d\n", c.strength)
	fmt.Println()
	/*
		fmt.Println("🔮 Skill:")
		fmt.Printf("Name:       %s\n", skill.Name)
		fmt.Printf("Power:      %d\n", skill.Power)
		fmt.Println()
		fmt.Println("🗡️ Weapon:")
		fmt.Printf("Name:       %s\n", weapon.Name)
		fmt.Printf("Damage:     %d\n", weapon.Damage)
		fmt.Printf("Range:      %d\n", weapon.Range)
		fmt.Println("--------------------")
	*/
}

func main() {
	c := NewCharacter("Katusha", WARRIOR, '♡')
	c.PrintStats()
}
