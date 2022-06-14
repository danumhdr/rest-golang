package helper

import "database/sql"

func CommitOrRollback(trx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := trx.Rollback()
		PanicIfError(errorRollback)
	} else {
		errorCommit := trx.Commit()
		PanicIfError(errorCommit)
	}
}
