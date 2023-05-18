package util

import (
	"github.com/gin-gonic/gin"
)

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseLogin struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
	UserId  uint   `json:"userId"`
}

type LoginErrorWrapper struct{}

// ErrorJSON : json error response function
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

// SuccessJSON : json error response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}

func (err LoginErrorWrapper) Error() string {
	return "Já existe cadastro com esses dados"
}

func NewErrorWrapper() error {
	return LoginErrorWrapper{}
}
