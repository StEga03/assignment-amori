package entity

import (
	"time"

	"github.com/assignment-amori/internal/entity/generic"
)

type User struct {
	ID                 uint64    `json:"id"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	BirthDate          time.Time `json:"birthDate"`
	Gender             string    `json:"gender"`
	GenderInterest     string    `json:"genderInterest"`
	PhoneNumber        string    `json:"phoneNumber"`
	RelationshipStatus string    `json:"relationshipStatus"`
	RelationshipGoal   string    `json:"relationshipGoal"`
	generic.MetaInfo
}

type NewUserParams struct {
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	BirthDate          time.Time `json:"birthDate"`
	Gender             string    `json:"gender"`
	GenderInterest     string    `json:"genderInterest"`
	PhoneNumber        string    `json:"phoneNumber"`
	RelationshipStatus string    `json:"relationshipStatus"`
	RelationshipGoal   string    `json:"relationshipGoal"`
}
