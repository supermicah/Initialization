package orm

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
