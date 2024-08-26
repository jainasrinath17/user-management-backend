package service_test

import (
	"testing"
	"user-management-backend/internal/model"
	"user-management-backend/internal/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) GetByUserName(userName string) (*model.User, error) {
	args := m.Called(userName)
	if result := args.Get(0); result != nil {
		return result.(*model.User), nil
	  }
	  return nil, args.Error(1)
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id uint) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = Describe("UserService", func() {
	var (
		mockRepo    *MockUserRepository
		userService service.UserService
		testUser    *model.User
	)

	BeforeEach(func() {
		mockRepo = new(MockUserRepository)
		userService = service.NewUserService(mockRepo)
		testUser = &model.User{
			UserName:   "testuser",
			FirstName:  "Test",
			LastName:   "User",
			Email:      "testuser@example.com",
			UserStatus: "A",
			Department: "CSE",
		}
	})

	Describe("CreateUser", func() {
		Context("when user_name already exists", func() {
			It("should return an error", func() {
				mockRepo.On("GetByUserName", testUser.UserName).Return(testUser, nil)
				err := userService.CreateUser(testUser)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("user already exists"))
			})
		})

		Context("when user_name does not exist", func() {
			It("should create a new user", func() {
				mockRepo.On("GetByUserName", testUser.UserName).Return(nil, nil)
				mockRepo.On("Create", testUser).Return(nil)

				err := userService.CreateUser(testUser)
				Expect(err).NotTo(HaveOccurred())
				mockRepo.AssertExpectations(GinkgoT())
			})
		})
	})
})
