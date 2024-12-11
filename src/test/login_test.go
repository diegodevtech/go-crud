package test

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("email_and_password_valid_returns_token", func(t *testing.T) {
		recorderCreateuser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateuser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      50,
		}

		bCreate, _ := json.Marshal(userRequest)
		stringReaderCreateuser := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreateuser)
		UserController.CreateUser(ctxCreateUser)

		recorderLogin := httptest.NewRecorder()
		ctxLogin := GetTestGinContext(recorderLogin)

		loginRequest := request.LoginRequest{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(loginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLogin, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.Login(ctxLogin)

		assert.EqualValues(t, http.StatusOK, recorderLogin.Code)
		assert.NotEmpty(t, recorderLogin.Result().Header.Get("Authorization"))
	})

	t.Run("invalid_password_returns_401", func(t *testing.T) {
		recorderCreateuser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateuser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: "Teste!@!@!@11221212",
			Name:     "Test User",
			Age:      50,
		}

		bCreate, _ := json.Marshal(userRequest)
		stringReaderCreateuser := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreateuser)
		UserController.CreateUser(ctxCreateUser)

		recorderLogin := httptest.NewRecorder()
		ctxLogin := GetTestGinContext(recorderLogin)

		loginRequest := request.LoginRequest{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(loginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLogin, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.Login(ctxLogin)

		assert.EqualValues(t, http.StatusUnauthorized, recorderLogin.Code)
		assert.Empty(t, recorderLogin.Result().Header.Get("Authorization"))
	})

	t.Run("invalid_email_returns_404", func(t *testing.T) {
		recorderCreateuser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateuser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@", rand.Int())

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: password,
			Name:     "Test User",
			Age:      50,
		}

		bCreate, _ := json.Marshal(userRequest)
		stringReaderCreateuser := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreateuser)
		UserController.CreateUser(ctxCreateUser)

		recorderLogin := httptest.NewRecorder()
		ctxLogin := GetTestGinContext(recorderLogin)

		loginRequest := request.LoginRequest{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(loginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLogin, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.Login(ctxLogin)

		assert.EqualValues(t, http.StatusUnauthorized, recorderLogin.Code)
		assert.Empty(t, recorderLogin.Result().Header.Get("Authorization"))
	})
}
