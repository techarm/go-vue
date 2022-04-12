package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/techarm/go-vue/vue-api/internal/data"
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

	if !user.Active {
		app.errorJSON(w, errors.New("ユーザーが無効になっています。管理者に連絡してください。"), http.StatusOK)
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
	user.Token = token
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

// ValidateToken トークン有効性チェック
func (app *application) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("json format is not corret"))
		return
	}

	valid, _ := app.models.Token.ValidToken(payload.Token)
	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data:    valid,
	})
}

// GetUsers ユーザー一覧を取得する
func (app *application) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.models.User.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("not found users"))
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

// GetUser ユーザー情報を取得する
func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("パラメータ不正です"))
		return
	}
	user, err := app.models.User.GetByID(userID)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("not found user"))
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data: warpper{
			"user": user,
		},
	})
}

// CreateUser　ユーザー作成する
func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("パラメータ不正"))
		return
	}
	err = user.Insert()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザー情報登録失敗しました。"), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "ユーザー情報を登録しました。",
		Data: warpper{
			"user": user,
		},
	})
}

// UpdateUser ユーザー更新する
func (app *application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("パラメータ不正です"))
		return
	}

	var payload data.User
	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("パラメータ不正"))
		return
	}

	user, err := app.models.User.GetByID(userID)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("user not found"), http.StatusNotFound)
		return
	}

	user.Email = payload.Email
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Active = payload.Active

	err = user.Update()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザー情報更新失敗しました。"), http.StatusInternalServerError)
		return
	}

	// reset password if password is set
	if payload.Password != "" {
		err := user.ResetPassword(payload.Password)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, errors.New("パスワード更新失敗しました。"), http.StatusInternalServerError)
			return
		}
		app.infoLog.Printf("ユーザーパスワード変更しました。 userID=%d\n", user.ID)
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "ユーザー情報を更新しました。",
	})
}

// DeleteUser ユーザー削除する
func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("パラメータ不正です"))
		return
	}

	var user = data.User{
		ID: userID,
	}

	err = user.Delete()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザー削除失敗しました。"), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "ユーザーを削除しました。",
	})
}

// LogUserOutAndSetInactive 対象ユーザーをログアウトさせ、状態を無効にする
func (app *application) LogUserOutAndSetInactive(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("パラメータ不正です"))
		return
	}

	user, err := app.models.User.GetByID(userID)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("user not found"), http.StatusNotFound)
		return
	}

	user.Active = false
	err = user.Update()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("ユーザー更新処理が失敗しました。"), http.StatusInternalServerError)
		return
	}

	// 対象ユーザーのログイントークンを削除
	err = app.models.Token.DeleteByUserID(userID)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("トークン削除処理が失敗しました。"), http.StatusInternalServerError)
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "対象ユーザーを無効にし、ログアウトさせました。",
	})
}

// GetAllBooks 全部の本リストを取得
func (app *application) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := app.models.Book.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("books not found"), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data: warpper{
			"books": books,
		},
	})
}

// GetBook slugで本情報を取得
func (app *application) GetBook(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	book, err := app.models.Book.GetOneBySlug(slug)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("book not found"), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
		Data: warpper{
			"book": book,
		},
	})
}
