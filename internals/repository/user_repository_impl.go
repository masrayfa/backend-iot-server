package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

// UserRepositoryImpl implements UserRepository interface
type UserRepositoryImpl struct {
}

// NewUserRepositoryImpl returns a new instance of UserRepositoryImpl
func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

// FindAll returns all users
func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx pgx.Tx) ([]domain.User, error) {
	return nil, nil
}

// FindById returns a user by id
func (r *UserRepositoryImpl) FindById(ctx context.Context, tx pgx.Tx,id int64) (domain.User, error) {
	var emptyUser domain.User
	return emptyUser, nil
}

// FindByEmail returns a user by email
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, tx pgx.Tx,email string) (domain.User, error) {

	var emptyUser domain.User
	return emptyUser, nil
}

// FindByUsername returns a user by username
func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, tx pgx.Tx,username string) (domain.User, error) {
	var emptyUser domain.User
	return emptyUser, nil
}

// FindByToken returns a user by token
func (r *UserRepositoryImpl) FindByToken(ctx context.Context, tx pgx.Tx,token string) (domain.User, error) {

	var emptyUser domain.User
	return emptyUser, nil
}

// Save saves a user
func (r *UserRepositoryImpl) Save(ctx context.Context, tx pgx.Tx, user domain.User) (domain.User, error) {
	token := uuid.New().String()
	status := false
	isAdmin := false


	h := sha256.New()
	_, err := h.Write([]byte(user.Password))
	if err != nil {
		return user, err
	}

	hashedBytes := h.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	script := "INSERT INTO users (username, email, password, token, status, is_admin) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := tx.QueryRow(ctx, script, user.Username, user.Email, hashedString, token, status, isAdmin)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		return user, err
	}

	var emptyUser domain.User
	return emptyUser, nil
}

// Update updates a user
func (r *UserRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, user domain.User) (domain.User, error) {
	var emptyUser domain.User
	return emptyUser, nil
}

// Delete deletes a user
func (r *UserRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx,id int64) error {
	return nil
}

// UpdateStatus updates a user status
func (r *UserRepositoryImpl) UpdateStatus(ctx context.Context, tx pgx.Tx, id int64, status bool) error {
	return nil
}

// UpdatePassword updates a user password
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx pgx.Tx, id int64, password string) error {
	return nil
}