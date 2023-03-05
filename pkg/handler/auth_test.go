package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/service"
	mock_service "github.com/Phaseant/MusicAPI/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
)

func TestHandler_register(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAutorization, user entity.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           entity.User
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "test",
				"password": "pwd"
			}`,
			inputUser: entity.User{
				Username: "test",
				Password: "pwd",
			},
			mockBehaviour: func(s *mock_service.MockAutorization, user entity.User) {
				s.EXPECT().NewUser(user).Return("1", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":"1"}`,
		},
		{
			name: "Missing fields",
			inputBody: `{
				"username": "test"
			}`,
			mockBehaviour:       func(s *mock_service.MockAutorization, user entity.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Error":"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name: "Server error",
			inputBody: `{
				"username": "test",
				"password": "pwd"
			}`,
			inputUser: entity.User{
				Username: "test",
				Password: "pwd",
			},
			mockBehaviour: func(s *mock_service.MockAutorization, user entity.User) {
				s.EXPECT().NewUser(user).Return("", errors.New("unable to insert into mongodb"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"Error":"internal error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAutorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Autorization: auth}

			handler := NewHandler(services)

			//test server
			r := gin.New()

			r.POST("/register", handler.register)

			//test http request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(testCase.inputBody))

			//perform request
			r.ServeHTTP(w, req)

			//asserts
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_login(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAutorization, user entity.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           entity.User
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "test",
				"password": "pwd"
			}`,
			inputUser: entity.User{
				Username: "test",
				Password: "pwd",
			},
			mockBehaviour: func(s *mock_service.MockAutorization, user entity.User) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("a.b.c.d", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"token":"a.b.c.d"}`,
		},
		{
			name: "Missing fields",
			inputBody: `{
				"username": "test"
			}`,
			mockBehaviour:       func(s *mock_service.MockAutorization, user entity.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Error":"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name: "Server error",
			inputBody: `{
				"username": "test",
				"password": "pwd"
			}`,
			inputUser: entity.User{
				Username: "test",
				Password: "pwd",
			},
			mockBehaviour: func(s *mock_service.MockAutorization, user entity.User) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("", errors.New("couldn't generate token"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"Error":"couldn't generate token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAutorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Autorization: auth}

			handler := NewHandler(services)

			//test server
			r := gin.New()

			r.POST("/login", handler.login)

			//test http request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(testCase.inputBody))

			//perform request
			r.ServeHTTP(w, req)

			//asserts
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
