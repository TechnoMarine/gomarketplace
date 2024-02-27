package db

import (
	"context"
	"fmt"
	"gomarketplace/common"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePostTx(t *testing.T) {
	store := NewStore(testConn)

	n := 10

	arg := CreateUserParams{
		Username: "Test",
		Role:     "Testerer",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		log.Fatal("User doesnt create")
	}

	errs := make(chan error)
	results := make(chan Post)

	for i := 0; i < n; i++ {
		go func() {
			args := CreatePostTxParams{
				Title:  common.GetRandomArticleName(),
				Body:   common.GetTextMock(),
				UserId: user.ID,
				Status: common.GetRandomStatuses(),
			}
			result, err := store.CreatePostTx(context.Background(), args)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)

		fmt.Print(result)
	}

}
