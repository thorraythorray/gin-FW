package request

type User struct {
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	Status   int
}

func (u *User) Validate() {

}
