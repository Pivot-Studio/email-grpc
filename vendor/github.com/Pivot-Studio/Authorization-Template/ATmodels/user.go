package ATmodels

type AuthUser struct {
	Email                string `gorm:"primary_key"`
	Username             string
	Password             string
	Is_email_activated   bool
	Register_timestamp   int64
	Last_login_timestamp int64
	Register_ip          string
	Last_login_ip        string
	Role                 string
}
