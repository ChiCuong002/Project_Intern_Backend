package helper

type UpdateData struct {
	UserID    uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}
