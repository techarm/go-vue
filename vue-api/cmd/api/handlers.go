package main

import (
	"errors"
	"net/http"
	"time"
)

type credentials struct {
	UserId   string `json:"email"`
	Password string `json:"password"`
}

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type warpper map[string]interface{}

// Login ログイン処理
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	err := app.readJSON(w, r, &creds)
	if err != nil {
		message := "json format is not corret"
		app.errorLog.Printf("%s. error: %s", message, err)
		app.errorJSON(w, errors.New(message))
		return
	}

	// TODO authenticate
	app.infoLog.Println(creds.UserId, creds.Password)

	// ユーザーID(email)でユーザー情報を取得する
	user, err := app.models.User.GetByEmail(creds.UserId)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザまたはパスワードが正しくありません。"))
		return
	}

	// パスワードの検証
	valid, err := user.PasswordMatched(creds.Password)
	if err != nil || !valid {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザまたはパスワードが正しくありません。"))
		return
	}

	// Tokenを生成する(有効時間を24時間設定)
	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("トークン情報が生成失敗しました。"))
		return
	}

	// Token情報をDBに保存する
	token.UserID = user.ID
	token.Email = user.Email
	err = token.Insert()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("トークン情報が生成失敗しました。"))
		return
	}
	app.infoLog.Printf("Create token info and tokenID is %d\n", token.ID)

	// send back a response
	err = app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data:    warpper{"token": token},
	})

	if err != nil {
		app.errorLog.Println(err)
	}
}
