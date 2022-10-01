package db

type User struct {
	UserId         string `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
	CreatedOn  string `json:"created_on"`
	UpdatedOn  string `json:"updated_on"`
	Role       string `json:"role"`
}
