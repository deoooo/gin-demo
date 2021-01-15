package models

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckPassword(username string, password string) bool {
	var u User
	db.Find(&u).Where("username = ?", username)
	return u.Password == password
}
