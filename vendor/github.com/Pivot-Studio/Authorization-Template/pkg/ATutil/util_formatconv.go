package ATutil

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

func ConvAtoUi(str string) (ui uint) {
	int_str, err := strconv.Atoi(str)
	if err != nil {
		log.Panic(err)
	}
	ui = uint(int_str)
	return
}

// A Hash function using salt with bcrypt libriry to hash password
func HashWithSalt(plainText string) (HashText string) {

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	CheckError(err)
	HashText = string(hash)
	return
}
func ParseTimeStampToReadableTime(timestamp int64) time.Time {
	i, err := strconv.ParseInt(strconv.Itoa(int(timestamp)), 10, 64)
	if err != nil {
		panic(err)
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tm := time.Unix(i, 0).In(loc)
	return tm
}
