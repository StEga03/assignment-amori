package user

import (
	"context"
	"database/sql"

	"github.com/assignment-amori/internal/entity"
	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
	timeutils "github.com/assignment-amori/pkg/time_utils"
)

func (r *Repository) CreateUser(ctx context.Context, param entity.NewUserParams, cel *consistency.ConsistencyElement) (uint64, error) {
	var (
		userId uint64
		tx     *pgx.Tx
		err    error

		now = timeutils.Now()
	)

	if cel != nil {
		tx = cel.Txn
	}

	userId, err = r.sf.NextID()
	if err != nil {
		return userId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToGenerateID)
	}

	err = r.db.ExecuteInTx(ctx, tx, func(tx *pgx.Tx) error {
		// Insert User.
		user := userTable{
			ID:        userId,
			FirstName: param.FirstName,
			LastName: sql.NullString{
				String: param.LastName,
				Valid:  true,
			},
			BirthDate: sql.NullTime{
				Time:  param.BirthDate,
				Valid: true,
			},
			Gender: sql.NullString{
				String: param.Gender,
				Valid:  true,
			},
			GenderInterest: sql.NullString{
				String: param.GenderInterest,
				Valid:  true,
			},
			PhoneNumber: sql.NullString{
				String: param.PhoneNumber,
				Valid:  true,
			},
			RelationshipStatus: sql.NullString{
				String: param.RelationshipStatus,
				Valid:  true,
			},
			RelationshipGoal: sql.NullString{
				String: param.RelationshipGoal,
				Valid:  true,
			},
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := r.insertUser(ctx, tx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return userId, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedCreateFromRepoUser)
	}

	return userId, nil
}

/*
================================================ UNEXPORTED FUNCTIONS =====================================================
*/

func (r *Repository) insertUser(ctx context.Context, tx txResource, user userTable) (err error) {
	_, err = tx.Insert(ctx, &user)

	return errorwrapper.Wrap(err, errorwrapper.ErrIDInsertDB,
		errorwrapper.WithMetaData(
			errorwrapper.MetaKV{
				"stmt": insertUser,
			}))
}
