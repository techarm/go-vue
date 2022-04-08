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

	app.infoLog.Printf("user[%s] was try to login", creds.UserId)

	// ユーザーID(email)でユーザー情報を取得する
	user, err := app.models.User.GetByEmail(creds.UserId)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザまたはパスワードが正しくありません。"), http.StatusOK)
		return
	}

	// パスワードの検証
	valid, err := user.PasswordMatched(creds.Password)
	if err != nil || !valid {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザまたはパスワードが正しくありません。"), http.StatusOK)
		return
	}

	// Tokenを生成する(有効時間を24時間設定)
	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("トークン情報が生成失敗しました。"), http.StatusInternalServerError)
		return
	}

	// Token情報をDBに保存する
	token.UserID = user.ID
	token.Email = user.Email
	err = token.Insert()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("トークン情報が生成失敗しました。"), http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Create token info and tokenID is %d\n", token.ID)

	// send back a response
	user.Token = *token
	err = app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data: warpper{
			"user": user,
		},
	})

	if err != nil {
		app.errorLog.Println(err)
	}
}

// Logout ログアウト処理、トークン情報を削除する
func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &request)
	if err != nil {
		message := "json format is not corret"
		app.errorLog.Printf("%s. error: %s", message, err)
		app.errorJSON(w, errors.New(message))
		return
	}

	// 処理失敗しても、正常の応答データを返す
	err = app.models.Token.DeleteByToken(request.Token)
	if err != nil {
		app.errorLog.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
	})

	if err != nil {
		app.errorLog.Println(err)
	}
}

// GetUsers ユーザー一覧を取得する
func (app *application) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.models.User.GetAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data: warpper{
			"users": users,
		},
	})

	if err != nil {
		app.errorLog.Println(err)
	}
}
