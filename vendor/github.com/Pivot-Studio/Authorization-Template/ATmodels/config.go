package ATmodels

type Config struct {
	Key                 string              `json:"Key"`
	DbSettings          DbSettings          `json:"DbSettings"`
	EmailSenderSettings EmailSenderSettings `json:"EmailSenderSettings"`
	RedisSettings       RedisSettings       `json:"RedisSettings"`
	HostName            string              `json:"HostName"`
}
type RedisSettings struct {
	Address  string `json:"Address"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}
type DbSettings struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}

//Case Sensitive! Need a uppercase suffix
type EmailSenderSettings struct {
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	Servername string `json:"Servername"`
}
