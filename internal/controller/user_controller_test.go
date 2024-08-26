package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"user-management-backend/internal/controller"
	"user-management-backend/internal/model"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) UpdateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) GetUserByID(id uint) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = Describe("UserController", func() {
	var (
		e              *echo.Echo
		req            *http.Request
		rec            *httptest.ResponseRecorder
		mockService    *MockUserService
		userController *controller.UserController
	)

	BeforeEach(func() {
		e = echo.New()
		rec = httptest.NewRecorder()
		mockService = new(MockUserService)
		userController = &controller.UserController{Service: mockService}
	})

	Describe("CreateUser", func() {
		Context("when creating a new user", func() {
			It("should return status created and user data", func() {
				user := &model.User{
					UserName:   "testuser",
					FirstName:  "Test",
					LastName:   "User",
					Email:      "testuser@example.com",
					UserStatus: "A",
					Department: "IT",
				}

				mockService.On("CreateUser", user).Return(nil)

				body, _ := json.Marshal(user)
				req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

				c := e.NewContext(req, rec)
				err := userController.CreateUser(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusCreated))
				Expect(rec.Body.String()).To(ContainSubstring("testuser"))
			})
		})

		Context("when user_name already exists", func() {
			It("should return status bad request", func() {
				user := &model.User{
					UserName:   "testuser",
					FirstName:  "Test",
					LastName:   "User",
					Email:      "testuser@example.com",
					UserStatus: "A",
					Department: "IT",
				}

				mockService.On("CreateUser", user).Return(errors.New("user_name already exists"))

				body, _ := json.Marshal(user)
				req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

				c := e.NewContext(req, rec)
				err := userController.CreateUser(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(rec.Body.String()).To(ContainSubstring("user_name already exists"))
			})
		})
	})

	Describe("DeleteUser", func() {
		Context("when deleting a user", func() {
			It("should return status no content", func() {
				userID := 1
				mockService.On("DeleteUser", uint(userID)).Return(nil)

				req = httptest.NewRequest(http.MethodDelete, "/users/"+strconv.Itoa(userID), nil)
				c := e.NewContext(req, rec)
				c.SetParamNames("id")
				c.SetParamValues(strconv.Itoa(userID))

				err := userController.DeleteUser(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusNoContent))
			})
		})
	})
})
