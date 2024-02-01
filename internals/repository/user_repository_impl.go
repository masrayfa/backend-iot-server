package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
)

// UserRepositoryImpl implements UserRepository interface
type UserRepositoryImpl struct {
}

// NewUserRepositoryImpl returns a new instance of UserRepositoryImpl
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// FindAll returns all users
func (r *UserRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.User, error) {
	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)

	script := "SELECT id_user, email, username, status, isAdmin from user_person"

	rows, err := tx.Query(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.IsAdmin)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users, nil
}

// FindById returns a user by id
func (r *UserRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int64) (domain.UserRead, error) {
	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)

	script := "SELECT id_user, email, username, status, isAdmin FROM user_person WHERE id_user = $1"

	var user domain.UserRead 
	err = tx.QueryRow(ctx, script, id).Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.IsAdmin)
	helper.PanicIfError(err)

	return user, nil
}

// FindByEmail returns a user by email
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, dbpool *pgxpool.Pool, email string) (domain.User, error) {

	var emptyUser domain.User
	return emptyUser, nil
}

// FindByUsername returns a user by username
func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, dbpool *pgxpool.Pool, username string) (domain.User, error) {
	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)

	script := "SELECT id_user, email, username, status, isAdmin from user_person WHERE username = $1"

	rows := tx.QueryRow(ctx, script, username)

	var user domain.User
	err = rows.Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.IsAdmin)
	helper.PanicIfError(err)

	return user, nil
}

// Save saves a user
func (r *UserRepositoryImpl) Save(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	status := false
	isAdmin := false

	hashedPassword, err := helper.HashPassword(user.Password)

	tx, err := dbpool.Begin(ctx)

	defer func() {
		err := recover()
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				return
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				return
			}
		}
	}()

	script := "INSERT INTO user_person (username, email, password, status, isadmin) VALUES ($1, $2, $3, $4, $5) RETURNING id_user"
	row := tx.QueryRow(ctx, script, user.Username, user.Email, hashedPassword, status, isAdmin)

	var idUser int64
	err = row.Scan(&idUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update updates a user
func (r *UserRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// Delete deletes a user
func (r *UserRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int64) error {
	tx, err := dbpool.Begin(ctx)
	defer helper.CommitOrRollback(ctx, tx)

	script := "DELETE FROM user_person WHERE id_user = $1"

	res, err := tx.Exec(ctx, script, id)
	helper.PanicIfError(err)

	if res.RowsAffected() != 1 {
		http.Error(nil, fmt.Sprintf("No row affected on delete user with id: %d", id), http.StatusBadRequest)
		return err
	}

	return nil
}

// UpdateStatus updates a user status
func (r *UserRepositoryImpl) UpdateStatus(ctx context.Context, dbpool *pgxpool.Pool, id int64, status bool) error {
	return nil
}

// UpdatePassword updates a user password
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, dbpool *pgxpool.Pool, id int64, password string) error {
	tx, err := dbpool.Begin(ctx)
	defer helper.CommitOrRollback(ctx, tx)

	script := "UPDATE user_person SET password = $1 WHERE id_user = $2"

	hashedPassword, err := helper.HashPassword(password)

	res, err := tx.Exec(ctx, script, hashedPassword, id)
	helper.PanicIfError(err)

	if res.RowsAffected() != 1 {
		http.Error(nil, fmt.Sprintf("No row affected on update user password with id: %d", id), http.StatusBadRequest)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) MatchPassword(ctx context.Context, dbpool *pgxpool.Pool, id int64, password string) error {
	var userPassword string

	script := "SELECT password FROM user_person WHERE id_user = $1"

	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)

	err = tx.QueryRow(ctx, script, id).Scan(&userPassword)
	helper.PanicIfError(err)

	hashedPassword, err := helper.HashPassword(password)
	helper.PanicIfError(err)

	// compare password
	if userPassword != hashedPassword {
		return err
	}

	return nil
}