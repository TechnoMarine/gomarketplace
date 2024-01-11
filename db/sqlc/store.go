package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db, Queries: New(db)}
}

func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreatePostTxParams struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int32  `json:"user_id"`
	Status string `json:"status"`
}

type CreateUserResult struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
}

func (store *Store) CreatePostTx(ctx context.Context, arg CreatePostTxParams) (Post, error) {
	var err error
	var post Post

	err = store.execTx(ctx, func(queries *Queries) error {

		post, err = store.CreatePost(ctx, CreatePostParams{Title: arg.Title, Body: arg.Body, Status: arg.Status, UserID: arg.UserId})
		if err != nil {
			fmt.Errorf("creating post error %v", err)
		}
		return nil
	})

	return post, err
}
