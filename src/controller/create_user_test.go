package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserController_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_got_error_invalid_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@",
			Password: "Teste@@2020",
			Name:     "Test",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_undefined_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			// Email: gomock.Nil().String(),
			Password: "Teste@@2020",
			Name:     "Test",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_undefined_password", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email: "teste@teste.com",
			// Password: "T",
			Name: "Test",
			Age:  20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_too_short_password", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "T",
			Name:     "Test",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_undefined_name", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			// Name: "Test",
			Age: 20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_too_short_name", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "Tes",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_too_long_name", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "TesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteX",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_undefined_age", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "Teste",
			// Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_too_young", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "TesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteTesteX",
			Age:      17,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_too_old", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "Teste",
			Age:      101,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	
}
