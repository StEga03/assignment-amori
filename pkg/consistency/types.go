package consistency

import (
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/sql/pgx"
)

var (
	//ErrTxIsEmpty is error when got nil in Element.Txn
	ErrTxIsEmpty = errorwrapper.New(errorwrapper.ErrIDTxElementIsEmpty)
)

// ConsistencyElement is a structure consist of variable that supporting consistency process
type ConsistencyElement struct {
	//Txn is sequel of transaction
	Txn *pgx.Tx
	//PostCmd is command that will execute after transaction wrapped up
	PostCmd []func() error

	// PostCmdUndo is command that will execute after transaction rolled back.
	PostCmdUndo []func() error
}
