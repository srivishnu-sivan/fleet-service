package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/srivishnu-sivan/fleet-service/internal/domain"
	"github.com/srivishnu-sivan/fleet-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}


func NewUserHandler (service *service.UserService) *UserHandler{
	return &UserHandler{
		service : service,
	}
}


func (u *UserHandler) CreateUser(c *gin.Context){
	var req domain.RegisterUser

	if err := c.ShouldBindBodyWithJSON(&req); err != nil{
		c.JSON(400, gin.H{
			"error" : "Invalid request feild" + err.Error(),
		})
		return
	}


user := domain.RegisterUser{
	Email: req.Email,
	Password: req.Password,
	RoleID: req.RoleID,
}


id, err := u.service.CreateUser(c.Request.Context(), user)

if err != nil{
	c.JSON(500, gin.H{"error" : err.Error()})
	return
}

c.JSON(201, gin.H{"id": id})















}