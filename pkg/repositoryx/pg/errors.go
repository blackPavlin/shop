package pg

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

// IsUniqueViolationErr
func IsUniqueViolationErr(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation
}

// IsForeignKeyViolationErr
func IsForeignKeyViolationErr(err error, fkName string) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) &&
		pgErr.Code == pgerrcode.ForeignKeyViolation &&
		pgErr.ConstraintName == fkName
}
