package data

import (
	"context"
	"time"
)

// User is user model
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_Name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
}

// GetAll is get all users from user table and order by id
func (u *User) GetAll() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `select id, email, first_name, last_name, password, created_at, updated_at from users order by id`
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
			&user.CreatedAt,
			&user.UpdatedAt,
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

	stmt := `select id, email, first_name, last_name, password, created_at, updated_at from users where email = $1`
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

	stmt := `select id, email, first_name, last_name, password, created_at, updated_at from users where email = $1`
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
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert inserts a new user into the database, and returns the ID
func (u *User) Insert() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `insert into users (email, first_name, last_name, password, created_at, updated_at)
		     values ($1, $2, $3, $4, $5, $6) returning id`

	row := db.QueryRowContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Password,
		time.Now(),
		time.Now(),
	)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates the user info(expected password) by userID
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `update users set(email, first_name, last_name, updated_at) values ($1, $2, $3, $4) where id = $5`
	_, err := db.ExecContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
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
