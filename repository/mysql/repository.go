package mysql

import (
	"dependency-injection/domain/user"
	"dependency-injection/usecase/register"
	"errors"
	"fmt"
)

type dsn string

func DataSource(user string, password string, host string, port uint, db string) dsn {
	return dsn(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			user, password, host, port, db,
		),
	)
}

func NewRepository(dsn dsn) register.Repository {
	return &repository{string(dsn)}
}

type repository struct {
	dsn string
}

func (r *repository) Add(user.User) error {
	// db, err := sql.Open("mysql", string(r.dsn))
	return errors.New("unimplemented")
}
