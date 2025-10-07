package test

import (
	"errors"
	"testing"

	"golang-microservice/internal/entity"
	"golang-microservice/internal/usecase"
)

// mockUserRepository is a mock implementation of UserRepository interface
type mockUserRepository struct {
	users  map[int64]entity.User
	nextID int64
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users:  make(map[int64]entity.User),
		nextID: 1,
	}
}

func (m *mockUserRepository) Create(user *entity.User) error {
	user.ID = m.nextID
	m.users[m.nextID] = *user
	m.nextID++
	return nil
}

func (m *mockUserRepository) FindAll() ([]entity.User, error) {
	result := []entity.User{}
	for _, u := range m.users {
		result = append(result, u)
	}
	return result, nil
}

func (m *mockUserRepository) FindByID(id int64) (*entity.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (m *mockUserRepository) Update(user *entity.User) error {
	if _, exists := m.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	m.users[user.ID] = *user
	return nil
}

func (m *mockUserRepository) Delete(id int64) error {
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

func setupUsecase() usecase.UserUsecase {
	repo := newMockUserRepository()
	return usecase.NewUserUsecase(repo)
}

func TestCreateUser(t *testing.T) {
	uc := setupUsecase()

	user := &entity.User{Name: "Eko", Email: "eko@example.com"}
	err := uc.Create(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID == 0 {
		t.Fatalf("expected user ID to be set, got %v", user.ID)
	}
}

func TestGetAllUsers(t *testing.T) {
	uc := setupUsecase()

	_ = uc.Create(&entity.User{Name: "Eko", Email: "eko@example.com"})
	_ = uc.Create(&entity.User{Name: "Budi", Email: "budi@example.com"})

	users, err := uc.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	uc := setupUsecase()
	user := &entity.User{Name: "Eko", Email: "eko@example.com"}
	_ = uc.Create(user)

	fetched, err := uc.GetByID(user.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fetched.Name != user.Name {
		t.Fatalf("expected name %s, got %s", user.Name, fetched.Name)
	}
}

func TestUpdateUser(t *testing.T) {
	uc := setupUsecase()
	user := &entity.User{Name: "Eko", Email: "eko@example.com"}
	_ = uc.Create(user)

	user.Name = "Eko Updated"
	err := uc.Update(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fetched, _ := uc.GetByID(user.ID)
	if fetched.Name != "Eko Updated" {
		t.Fatalf("expected updated name, got %s", fetched.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	uc := setupUsecase()
	user := &entity.User{Name: "Eko", Email: "eko@example.com"}
	_ = uc.Create(user)

	err := uc.Delete(user.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = uc.GetByID(user.ID)
	if err == nil {
		t.Fatalf("expected error, got nil (user should be deleted)")
	}
}
