package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Token struct {
	ID        int       `json:"-"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// GetByToken トークン文字列でTokenデータを取得
func (t *Token) GetByToken(plainText string) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `select id, user_id, email, token, token_hash, expiry, created_at, updated_at from tokens where token = $1`
	row := db.QueryRowContext(ctx, stmt, plainText)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var token Token
	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Email,
		&token.Token,
		&token.TokenHash,
		&token.Expiry,
		&token.CreatedAt,
		&token.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// GetUserByToken トータルでユーザー情報を取得
func (t *Token) GetUserByToken(plainText string) (*User, error) {
	// get token info from database
	tkn, err := t.GetByToken(plainText)
	if err != nil {
		return nil, err
	}

	if tkn.Expiry.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	var user User
	userData, err := user.GetByID(tkn.UserID)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

// GenerateToken ユーザーIDで26桁のトークン情報を生成
func (t *Token) GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Token))
	token.TokenHash = hash[:]

	return token, nil
}

// AuthenticateToken HTTPリクエストトークン情報の認証
func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("cannot get authorization header")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("the authorization header format is not correct")
	}

	token := headerParts[1]
	if len(token) != 26 {
		return nil, errors.New("token format is not correct")
	}

	return t.GetUserByToken(token)
}

// ValidToken トークンの有効性をチェック
func (t *Token) ValidToken(plainText string) (bool, error) {
	_, err := t.GetUserByToken(plainText)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Insert inserts a new token into the database, and set the ID in receiver
func (t *Token) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	// 新しいトータルを登録する前に、DBに既存のトータル情報を全て削除
	err := t.DeleteByUserID(t.UserID)
	if err != nil {
		return err
	}

	stmt := `insert into tokens (user_id, email, token, token_hash, expiry, created_at, updated_at)
	         values ($1, $2, $3, $4, $5, $6, $7) returning id`
	row := db.QueryRowContext(ctx, stmt,
		t.UserID,
		t.Email,
		t.Token,
		t.TokenHash,
		t.Expiry,
		time.Now(),
		time.Now(),
	)
	if err := row.Err(); err != nil {
		return err
	}

	err = row.Scan(&t.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserID ユーザーIDでトータル情報を作成
func (t *Token) DeleteByUserID(userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `delete from tokens where user_id = $1`
	_, err := db.ExecContext(ctx, stmt, userID)
	return err
}

// DeleteByToken トークン文字列でトータル情報を作成
func (t *Token) DeleteByToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIME_OUT)
	defer cancel()

	stmt := `delete from tokens where token = $1`
	_, err := db.ExecContext(ctx, stmt, token)
	return err
}
