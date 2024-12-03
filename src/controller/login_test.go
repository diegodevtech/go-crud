package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_Login(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_got_error_invalid_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.LoginRequest{
			Email:    "teste@",
			Password: "Teste@@2020",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.Login(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_undefined_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.LoginRequest{
			// Email:    "teste@",
			Password: "Teste@@2020",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.Login(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_undefined_password", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.LoginRequest{
			Email:    "teste@teste.com",
			// Password: "Teste@@2020",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.Login(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_got_error_too_short_password", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.LoginRequest{
			Email:    "teste@teste.com",
			Password: "Te",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))
		

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.Login(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.LoginRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
		}

		domain := model.NewLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginService(domain).Return(nil, "", rest_err.NewInternalServerError("Error Test"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.Login(context)
		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		mockedToken := primitive.NewObjectID().Hex()

		userRequest := request.LoginRequest{
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
		}

		domain := model.NewLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginService(domain).Return(domain, mockedToken, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.Login(context)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], mockedToken)
	})
}