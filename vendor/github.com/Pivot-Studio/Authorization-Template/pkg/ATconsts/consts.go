package ATconsts

const (
	//Set-Cookie name prosperity
	COOKIE_NAME = "token"

	//Token is invalid in one month
	EXPIRE_TIME_TOKEN = 30 * 24 * 60 * 60

	TOKEN_SCRECT_KEY = "tokenemail"

	//Verify code is vaild in 20min
	VERIFYCODE_VALID_TIME = 1200

	VERIFYCODE_LENGTH = 4

	RATA_NUMBER_LIMITER = 2

	//Redis prefix string
	BLACK_LIST_PREFIX = "blacklist:"

	CONFIG_FILE_NAME = "Config.json"

	USER  = "user"
	ADMIN = "admin"
	ROOT  = "root"

	REDIS_VERIFY_CODE_SUFFIX = "verify_code:"
)
