package data

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is user model
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	HasToken  *int      `json:"has_token,omitempty"`
	Token     *Token    `json:"token,omitempty"`
}

// GetAll is get all users from user table and order by id
func (u *User) GetAll() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `select id, email, first_name, last_name, password, active, created_at, updated_at,
	case
		when (select count(id) from tokens t where user_id = users.id and t.expiry > NOW()) > 0 then 1
		else 0
	end as has_token
	from users order by id`
	rows, err := db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.HasToken,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, nil
	}

	return users, nil
}

// GetByEmail return one user info by email
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `select id, email, first_name, last_name, password, active, created_at, updated_at from users where email = $1`
	row := db.QueryRowContext(ctx, stmt, email)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByEmail return one user info by userID
func (u *User) GetByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `select id, email, first_name, last_name, password, active, created_at, updated_at from users where id = $1`
	row := db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert inserts a new user into the database, and set the ID in receiver
func (u *User) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return nil
	}

	stmt := `insert into users (email, first_name, last_name, password, created_at, updated_at)
		     values ($1, $2, $3, $4, $5, $6) returning id`

	row := db.QueryRowContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
		hashedPassword,
		time.Now(),
		time.Now(),
	)
	if err := row.Err(); err != nil {
		return err
	}

	err = row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the user info(expected password) by userID
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `update users set
		     email = $1,
			 first_name = $2,
			 last_name = $3,
			 active = $4,
			 updated_at = $5
			 where id = $6`
	_, err := db.ExecContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Active,
		time.Now(),
		u.ID,
	)

	return err
}

// Delete deltes user info from database by sserID
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `delete from users where id = $1`
	_, err := db.ExecContext(ctx, stmt, u.ID)
	return err
}

// PasswordMatched パスワードの検証処理
func (u *User) PasswordMatched(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// 検証失敗
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

// ResetPassword パスワードリセット処理
func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `update users set password = $1 where id = $2`
	_, err = db.ExecContext(ctx, stmt, hashedPassword, u.ID)
	return err
}
