package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/configs"
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
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, dbpool *pgxpool.Pool, email string) (user domain.User, err error) {
	log.Println("Find by email")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return user, err 
	}

	script := "SELECT id_user, email, username, status, isAdmin from user_person WHERE email = $1"

	rows := tx.QueryRow(ctx, script, email)

	err = rows.Scan(&user.IdUser, &user.Email, &user.Username, &user.Status, &user.IsAdmin)
	defer helper.CommitOrRollback(ctx, tx)
	if err != nil {
		log.Println("Tidak ada email user")
		return user, err
	}

	return user, nil
}

// FindByUsername returns a user by username
func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, dbpool *pgxpool.Pool, username string) (user domain.User, err error) {
	log.Println("Find by username", username)

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return user, err
	}

	sqlStatement := `SELECT id_user, email, username,  status, isadmin FROM user_person WHERE username= $1 `
	err = tx.QueryRow(ctx, sqlStatement, username).Scan(
		&user.IdUser,
		&user.Email,
		&user.Username,
		&user.Status,
		&user.IsAdmin,
	)
	if err != nil {
		log.Println("Tidak ada username")
		return user, err
	}

	log.Println("user di repository: ", user)

	fmt.Println("findby username success")
	return user, nil
}

// Save saves a user
func (r *UserRepositoryImpl) Save(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	status := false
	isAdmin := false

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)

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
	log.Println("id user dari save repository: ", idUser)

	user.IdUser = idUser

	return user, nil
}

// Update updates a user
func (r *UserRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// Delete deletes a user
func (r *UserRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int64) error {
	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)
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
	log.Println("update status repository")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(ctx, tx)

	script := "UPDATE user_person SET status = $1 WHERE id_user = $2"

	res, err := tx.Exec(ctx, script, status, id)
	if err != nil {
		return err
	}
	count := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("no row affected on update user status with id: %d", id)
	}

	log.Println("status berhasil diupdate dari repository")

	return nil
}

// UpdatePassword updates a user password
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, dbpool *pgxpool.Pool, id int64, password string) error {
	tx, err := dbpool.Begin(ctx)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx, tx)

	script := "UPDATE user_person SET password = $1 WHERE id_user = $2"

	hashedPassword, err := helper.HashPassword(password)
	helper.PanicIfError(err)

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
	if err != nil {
		return err
	}

	err = tx.QueryRow(ctx, script, id).Scan(&userPassword)
	if err != nil {
		return err
	}

	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	// compare password
	log.Println("user password: ", userPassword)
	log.Println("hashed password: ", hashedPassword)
	if userPassword != hashedPassword {
		log.Println("password tidak sama")
		return errors.New("password tidak sama")
	} 

	log.Println("password berhasil di match sama")
	return nil
}

func (r *UserRepositoryImpl) SendEmailForgotPassword(ctx context.Context, dbpool *pgxpool.Pool, user domain.User, password string) error {
	subject := "Forgot Password"
	body := fmt.Sprintf(`<html>
	 	<head></head>
		<body>
		<h3>Hi %s</h3>
		<p>Your forgot password request is already received. Here is your new password: %s</p>
		</body>

		</html>`, user.Username, password)

	err := helper.SendEmail(user.Email, subject, body)
	helper.PanicIfError(err)

	return nil
}

func (r *UserRepositoryImpl) SendEmailActivation(ctx context.Context, dbpool *pgxpool.Pool, user domain.UserRead) error {
	config := configs.GetConfig()

	jwtToken, err := helper.SignUserToken(user)
	helper.PanicIfError(err)

	urlCode := fmt.Sprintf("%s/api/v1/user/activate?token=%s", config.Server.Domain, jwtToken)
	subject := "Activation Account"
	body := fmt.Sprintf(`<html>
	 	<head>
			<title>Activation Account</title>
		</head>

		<body>
		<h3>Hi %s</h3>
		<p>Your account has been activated. You can now login to your account</p>
		<p>Click this link to activate your account: <a href="%s">Activate</a></p>
		</body>

		</html>`, user.Username, urlCode)

	err = helper.SendEmail(user.Email, subject, body)
	helper.PanicIfError(err)

	return nil
}