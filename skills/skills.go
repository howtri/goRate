package skills

import (
	"errors"
	"log"

	"github.com/rs/xid"
)

type Skill struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Rankings []int  `json:"rankings"`
}

type Ranking struct {
	ID      string `json:"id"`
	Ranking int    `json:"ranking"`
}

type db struct {
	skills []*Skill
}

var overall db

func GetSkill(s Skill) *Skill {
	for _, v := range overall.skills {
		if v.ID == s.ID {
			log.Println("Found Skill")
			return v
		}
	}
	log.Println("No skill with ID found")
	return &Skill{}

}

func AddSkill(s Skill) int {
	s.ID = xid.New().String()
	overall.skills = append(overall.skills, &s)
	return 1
}

func RankSkill(r Ranking) error {
	for _, v := range overall.skills {
		if v.ID == r.ID {
			log.Println("Adding ranking")
			v.Rankings = append(v.Rankings, r.Ranking)
			return nil
		}
	}
	log.Println("No skill with ID found")
	return errors.New("no skill with that ID")
}

func GetAll() []*Skill {
	log.Println("Returning all skills")
	return overall.skills
}
