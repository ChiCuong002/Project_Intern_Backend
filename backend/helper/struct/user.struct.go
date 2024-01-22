package helper

type UpdateData struct {
	UserID    uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}
type UserResponse struct {
	UserID      uint   `json:"user_id"`
	RoleID      uint   `json:"role_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsActive    bool   `json:"is_active"`
}

func (UserResponse) TableName() string {
	return "users"
}
