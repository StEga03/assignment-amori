package enforcer

import (
	"context"
	"log/slog"

	"github.com/assignment-amori/pkg/consistency"
	"github.com/assignment-amori/pkg/errorwrapper"
)

// New is initialization of struct rule Module
func New(db dbResource) *Module {
	return &Module{
		db: db,
	}
}

func (s *Module) RunAsUnit(ctx context.Context, action func(celTemp *consistency.ConsistencyElement) error) error {
	var (
		err error
		cel *consistency.ConsistencyElement
	)

	// Start consistency unit
	cel, err = s.start(ctx)
	if err != nil {
		return err
	}

	// do action
	err = action(cel)

	if err != nil {
		// rollback if failed
		errRollBack := s.undo(ctx, cel)
		if errRollBack != nil {
			slog.Error("[RunAsUnit] failed rollback", errRollBack)
		}

		// Run executable function assigned by other repo layer after transaction rolled back.
		for _, f := range cel.PostCmdUndo {
			if postCmdErr := f(); postCmdErr != nil {
				slog.Error("[RunAsUnit] failed run post cmd undo", postCmdErr)
			}
		}

		return err
	}

	// commit in complete
	err = s.complete(ctx, cel)
	return err
}

// Start is func to start the unit of work process should be call only once in every process
func (s *Module) start(ctx context.Context) (*consistency.ConsistencyElement, error) {
	el := consistency.ConsistencyElement{}

	if tx, err := s.db.Beginx(ctx); err != nil {
		return nil, errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToBegin)
	} else {
		el.Txn = tx
	}

	return &el, nil
}

// Complete is func to end the unit of work process should be call only once in every process
func (s *Module) complete(ctx context.Context, element *consistency.ConsistencyElement) error {

	if element == nil || element.Txn == nil {
		return consistency.ErrTxIsEmpty
	}

	if err := element.Txn.Commit(ctx); err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToCommit)
	}

	//Run executable function assigned by other repo layer after transaction wrapped up
	for _, f := range element.PostCmd {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

// Undo is func to backward the condition the unit of work process it happen because of one of process is failed and can call multiple in every process
func (s *Module) undo(ctx context.Context, element *consistency.ConsistencyElement) error {

	if element == nil || element.Txn == nil {
		return consistency.ErrTxIsEmpty
	}

	if err := element.Txn.Rollback(ctx); err != nil {
		return errorwrapper.Wrap(err, errorwrapper.ErrIDFailedToRollback)
	}

	return nil
}
