package middleware

import (
	"encoding/json"
	"src/golang_mssql/dbconfig"
	"src/golang_mssql/dto"
	"src/golang_mssql/models"
	utils "src/golang_mssql/util"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	var userDto dto.ProfileData
	err := json.NewDecoder(c.Request.Body).Decode(&userDto)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to decode the request body."})
	}

	user, err := utils.GetByUserId(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if len(user) > 0 {
		db := dbconfig.Connection()
		db.Model(&models.User{}).Where("id = ?", id).Updates(userDto)
		db.Commit()
		c.JSON(200, gin.H{"message": "Your Profile has been successfully changed."})
	} else {
		c.JSON(400, gin.H{"message": "User ID not found."})
	}

}
