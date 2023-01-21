package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"io/ioutil"
)

func Tables(pool *pgxpool.Pool) (err error) {

	sqlfile := "db/model.sql"
	c, err := ioutil.ReadFile(sqlfile)
	if err != nil {
		return fmt.Errorf("Can not find sql file: %v", err)
	}
	connect, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("Can not return connection from the Pool: %v", err)
	}
	defer connect.Release()
	sql := string(c)
	_, err = connect.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("Can not run init script: %v", err)
	}
	return nil
}
