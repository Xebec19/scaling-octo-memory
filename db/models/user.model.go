package db

type User struct {
	UserId     string `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Email      string
	Password   string
	ProfilePic string
	CreatedOn  string
	UpdatedOn  string
	Role       string
}
