package repos

import "github.com/gin-gonic/gin"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserInfo() map[string]interface{} {
	return gin.H{
		"username": "john_doe",
		"email":    "john@gmail.com",
		"age":      30,
		"address":  "123 Main St, Springfield, USA",
	}
}
