package user

import (
	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/internal/entity/generic"
)

func (u *userTable) ToEntity() entity.User {
	return entity.User{
		ID:                 u.ID,
		FirstName:          u.FirstName,
		LastName:           u.LastName.String,
		BirthDate:          u.BirthDate.Time,
		Gender:             u.Gender.String,
		GenderInterest:     u.GenderInterest.String,
		PhoneNumber:        u.PhoneNumber.String,
		RelationshipStatus: u.RelationshipStatus.String,
		RelationshipGoal:   u.RelationshipGoal.String,
		MetaInfo: generic.MetaInfo{
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	}
}
