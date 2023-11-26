package app

import (
	"medical-vitals-management-system/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock implementation of the database functions
type MockDB struct {
	mock.Mock
}

func (m *MockDB) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) GetUser(username string) (models.User, error) {
	args := m.Called(username)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockDB) UpdateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) DeleteUser(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockDB) UserExists(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func TestCreateUserWithMock(t *testing.T) {
	mockDB := new(MockDB)

	mockDB.On("CreateUser", mock.Anything).Return(nil)

	err := mockDB.CreateUser(models.User{Username: "JohnDoe"})
	assert.NoError(t, err, "Error creating user")

	mockDB.AssertExpectations(t)
}

func TestGetUserWithMock(t *testing.T) {
	mockDB := new(MockDB)

	mockUser := models.User{Username: "JohnDoe"}
	mockDB.On("GetUser", "JohnDoe").Return(mockUser, nil)

	user, err := mockDB.GetUser("JohnDoe")
	assert.NoError(t, err, "Error getting user")
	assert.Equal(t, mockUser, user, "User data mismatch")

	mockDB.AssertExpectations(t)
}

func TestUpdateUserWithMock(t *testing.T) {
	mockDB := new(MockDB)

	mockDB.On("UpdateUser", mock.Anything).Return(nil)

	err := mockDB.UpdateUser(models.User{Username: "JohnDoe"})
	assert.NoError(t, err, "Error updating user")

	mockDB.AssertExpectations(t)
}

func TestDeleteUserWithMock(t *testing.T) {
	mockDB := new(MockDB)

	mockDB.On("DeleteUser", "JohnDoe").Return(nil)

	err := mockDB.DeleteUser("JohnDoe")
	assert.NoError(t, err, "Error deleting user")

	mockDB.AssertExpectations(t)
}

func TestUserExistsWithMock(t *testing.T) {
	mockDB := new(MockDB)

	mockDB.On("UserExists", "JohnDoe").Return(true, nil)

	exists, err := mockDB.UserExists("JohnDoe")
	assert.NoError(t, err, "Error checking user existence")
	assert.True(t, exists, "User should exist")

	mockDB.AssertExpectations(t)
}
