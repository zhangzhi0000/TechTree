package general

const (
	DBName = "tech"

	RespKeyStatus = "status"
	RespKeyType   = "type"
	RespKeyData   = "data"

	UserActive   = 0x10
	UserInactive = 0x11

	ErrSucceed       = 0x0
	ErrInvalidParams = 0x1
	ErrMysql       = 0x3
	ErrNotFound    = 0x4
)