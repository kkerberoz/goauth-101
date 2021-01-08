package controllers

import (
	"ginauth101/models"
	"ginauth101/services"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//AuthController is for auth logic
type AuthController struct{}

//Login is to process login request
func (auth *AuthController) Login(c *gin.Context) {

	var loginInfo models.User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	userservice := services.Userservice{}
	user, errf := userservice.Find(&loginInfo)
	if errf != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}

	token, err := user.GetJwtToken()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

//Profile is to provide current user info
func (auth *AuthController) Profile(c *gin.Context) {
	user := c.MustGet("user").(*(models.User))

	c.JSON(200, gin.H{
		"username": user.Name,
		"email":    user.Email,
	})
}

//Register is for user signup
func (auth *AuthController) Register(c *gin.Context) {

	type registerInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
	}
	var info registerInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	user := models.User{}
	user.Email = info.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)
	user.Name = info.Name
	userservice := services.Userservice{}
	err = userservice.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Register successfully!"})
	}
	return
}
