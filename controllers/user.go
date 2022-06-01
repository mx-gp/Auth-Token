package controllers

import (
	"authtoken/helper"
	"authtoken/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateToken : Generate Token
func GenerateToken(c *gin.Context) {
	// Generate Token
	Token := helper.GenerateRandomString()

	// Get IP Address of local machine
	IpAddress := helper.GetIP()

	// store this token into DB
	user := models.User{
		Token:     Token,
		IPAddress: IpAddress,
	}

	rslt := DB.Create(&user)
	if rslt.Error != nil {
		helper.HandleError(c, rslt.Error, rslt.Error.Error(), 500)
		return
	}

	rslt = DB.Where("id = ?", user.ID).Find(&user)
	if rslt.Error != nil {
		helper.HandleError(c, rslt.Error, rslt.Error.Error(), 500)
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    Token,
		"success": true,
	})
}

// Login With Token
func LoginPOST(c *gin.Context) {
	token := c.Request.Header.Get("X-Auth-Token")

	// find that token exist in flashpoint or not
	existingUser := models.User{}
	rslt := DB.Where("token = ?", token).Find(&existingUser)
	if rslt.Error != nil {
		helper.HandleError(c, rslt.Error, rslt.Error.Error(), 500)
		return
	}

	s := time.Now().Format("2006-01-02 15:04:05")
	currentTime, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		helper.HandleError(nil, err, err.Error(), 400)
		return
	}

	if currentTime.After(existingUser.CreatedAt.Add(168 * time.Hour)) {
		var err error = errors.New("Token Expired")
		helper.HandleError(c, err, "Token is expired", 400)
		return
	}

	if existingUser.Token != token {
		helper.HandleError(nil, nil, "Token is not valid", 400)
		return
	}

	data := &models.User{
		Token:     token,
		ID:        existingUser.ID,
		CreatedAt: existingUser.CreatedAt,
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    data,
		"success": true,
	})
}
