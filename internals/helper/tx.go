package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// CommitOrRollback is a helper function to commit or rollback a transaction
func CommitOrRollback(ctx context.Context, tx pgx.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback(ctx)
		PanicIfError(err)
	} else {
		err := tx.Commit(ctx)
		PanicIfError(err)
	}
}