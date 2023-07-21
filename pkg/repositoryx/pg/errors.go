package pg

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

// IsUniqueViolationErr tells if given error is a postgres unique violation error "23505".
func IsUniqueViolationErr(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation
}

// IsForeignKeyViolationErr tells if given error is a postgres foreign key violation error "23503" of a given fk.
func IsForeignKeyViolationErr(err error, fkName string) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) &&
		pgErr.Code == pgerrcode.ForeignKeyViolation &&
		pgErr.ConstraintName == fkName
}
