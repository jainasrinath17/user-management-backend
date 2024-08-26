package repository_test

import (
	"errors"
	"testing"
	"user-management-backend/internal/model"
	"user-management-backend/internal/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
	*gorm.DB
}

var _ = Describe("UserRepository", func() {
	var (
		db       *gorm.DB
		userRepo repository.UserRepository
		testUser *model.User
	)

	BeforeEach(func() {
		db = initTestDB()
		userRepo = repository.NewUserRepository(db)
		testUser = &model.User{
			UserName:   "testuser",
			FirstName:  "Test",
			LastName:   "User",
			Email:      "testuser@example.com",
			UserStatus: "A",
			Department: "IT",
		}
	})

	AfterEach(func() {
		db.Migrator().DropTable(&model.User{})
	})

	Describe("Create", func() {
		It("should create a new user", func() {
			err := userRepo.Create(testUser)
			Expect(err).NotTo(HaveOccurred())

			var user model.User
			db.First(&user, "user_name = ?", testUser.UserName)
			Expect(user.UserName).To(Equal(testUser.UserName))
		})
	})

	Describe("FindByUserName", func() {
		Context("when user exists", func() {
			It("should return the user", func() {
				db.Create(testUser)
				user, err := userRepo.GetByUserName(testUser.UserName)
				Expect(err).NotTo(HaveOccurred())
				Expect(user.UserName).To(Equal(testUser.UserName))
			})
		})

		Context("when user does not exist", func() {
			It("should return an error", func() {
				user, err := userRepo.GetByUserName("nonexistent")
				Expect(user).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete", func() {
		It("should delete a user", func() {
			db.Create(testUser)
			err := userRepo.Delete(testUser.ID)
			Expect(err).NotTo(HaveOccurred())

			var user model.User
			result := db.First(&user, testUser.ID)
			Expect(errors.Is(result.Error, gorm.ErrRecordNotFound)).To(BeTrue())
		})
	})
})

func initTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	return db
}

func TestUserRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserRepository Suite")
}
