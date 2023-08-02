package schema

type User struct {
	Username string `json:"username" binding:"required" validate:"min=6,max=20"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}
type UserProfile struct {
	User
	Phone string `json:"phone" binding:"required" validate:"phone"`
	Email string `json:"email" binding:"required" validate:"email"`
}

type UserModel struct {
	BaseModel
	UserProfile
	Status   uint8  `binding:"-"`
	Identity string `binding:"-"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserRegister struct {
	UserProfile
}

func (form *UserRegister) Validate() error {
	return nil
}

type UserLogin struct {
	User
}

func (form *UserLogin) Validate() error {
	return nil
}
