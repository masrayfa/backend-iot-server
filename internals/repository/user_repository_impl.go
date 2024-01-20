package repository

import (
	"context"

	"github.com/google/uuid"
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

	script := "SELECT id_user, email, username, status, token, isAdmin from user_person"

	rows, err := tx.Query(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.Token, &user.IsAdmin)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users, nil
}

// FindById returns a user by id
func (r *UserRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int64) (domain.User, error) {
	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)

	script := "SELECT id_user, email, username, status, token, isAdmin FROM user_person WHERE id_user = $1"

	var user domain.User
	err = tx.QueryRow(ctx, script, id).Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.Token, &user.IsAdmin)
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

	script := "SELECT id_user, email, username, status, token, isAdmin from user_person WHERE username = $1"

	rows := tx.QueryRow(ctx, script, username)

	var user domain.User
	err = rows.Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.Token, &user.IsAdmin)
	helper.PanicIfError(err)

	return user, nil
}

// FindByToken returns a user by token
func (r *UserRepositoryImpl) FindByToken(ctx context.Context, dbpool *pgxpool.Pool, token string) (domain.User, error) {
	var emptyUser domain.User
	return emptyUser, nil
}

// Save saves a user
func (r *UserRepositoryImpl) Save(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	token := uuid.New().String()
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

	script := "INSERT INTO user_person (username, email, password, token, status, isadmin) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_user"
	row := tx.QueryRow(ctx, script, user.Username, user.Email, hashedPassword, token, status, isAdmin)

	var idUser int64
	err = row.Scan(&idUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update updates a user
func (r *UserRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	var emptyUser domain.User
	return emptyUser, nil
}

// Delete deletes a user
func (r *UserRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int64) error {
	return nil
}

// UpdateStatus updates a user status
func (r *UserRepositoryImpl) UpdateStatus(ctx context.Context, dbpool *pgxpool.Pool, id int64, status bool) error {
	return nil
}

// UpdatePassword updates a user password
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, dbpool *pgxpool.Pool, id int64, password string) error {
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