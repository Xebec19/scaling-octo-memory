package auth

type userRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"email"`
	Role      string `json:"role" binding:"oneof=teacher student"`
}
