package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr);
	if(err !=nil) {
       return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)


	duration, err := time.ParseDuration(maxIdleTime)

	if(err != nil) {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration);

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	er := db.PingContext(ctx)

	if(er != nil) {
		return nil, er
	}

	return db, nil
}