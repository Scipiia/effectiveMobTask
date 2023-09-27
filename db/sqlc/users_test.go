package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/scipiia/effectivemobiletask/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:       util.RandomName(),
		Surname:    util.RandomName(),
		Patronymic: util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Name, arg.Name)
	require.Equal(t, user.Surname, arg.Surname)
	require.Equal(t, user.Patronymic, arg.Patronymic)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Name, user1.Name)
	require.Equal(t, user2.Surname, user1.Surname)
	require.Equal(t, user2.Patronymic, user1.Patronymic)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:      user1.ID,
		Name:    user1.Name,
		Surname: user1.Surname,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Name, arg.Name)
	require.Equal(t, user2.Surname, arg.Surname)
	require.Equal(t, user2.Patronymic, user1.Patronymic)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
