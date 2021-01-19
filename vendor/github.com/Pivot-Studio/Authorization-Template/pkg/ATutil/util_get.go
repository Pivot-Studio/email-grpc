package ATutil

import (
	"errors"
	"github.com/Pivot-Studio/Authorization-Template/pkg/ATconsts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func GetEmailFromCookie(c *gin.Context) (email string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("您未登录，请登陆后查看")
		}
	}()
	cookie, err := c.Cookie(ATconsts.COOKIE_NAME)
	if err != nil {
		err = errors.New("您未登录，请登陆后查看")
		return
	}
	claim, _ := GetClaimFromToken(cookie)
	email = claim.(jwt.MapClaims)["email"].(string)
	return
}
func GetClaimFromToken(tokenString string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(ATconsts.TOKEN_SCRECT_KEY), err
	})
	if err != nil {
		return nil, err
	} else {
		claims = token.Claims.(jwt.MapClaims)
		return claims, nil
	}
}
func GetTimeStamp() (t int64) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t = time.Now().In(loc).Unix()
	return
}
func GetTodayStartTimeStamp() (timestamp int64) {

	t := time.Now()
	year, month, day := t.Date()
	timestamp = time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Unix()
	return
}
