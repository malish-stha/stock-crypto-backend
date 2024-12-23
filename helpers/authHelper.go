package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, userType string) (err error){
	userType := c.GetString("user_type")
	err = nil
	if userType != role{
		err = errors.New("You are not authorized to access this route")
		return err
	}
}

func MatchUserTypeToUid(c *gin.Context, userId string)(err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && userId != userId{
		err = errors.New("You are not authorized to access this route")
		return err
	}
	err = CheckUserType(c, userType)
	return err

}