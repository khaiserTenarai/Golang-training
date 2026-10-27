package repository

// UserRepository is a mock repository for the demo. Task 6; replace with auth_db later (Task 7).
type UserRepository struct{}

func NewUserRepository() *UserRepository { return &UserRepository{} }
func (r *UserRepository) IsValidDemoUser(username, password string) bool {
	return username == "demo" && password == "demo123"
}
