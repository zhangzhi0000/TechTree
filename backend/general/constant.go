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
	ErrUserExists    = 0x2
	ErrMysql         = 0x3
	ErrNotFound      = 0x4

	NodeNoActive = 0x10
	NodeActive   = 0x11
	NodePassed   = 0x12
	NoseUserAdd  = 0x13

	Agree = 0x20
	Disagree = 0x21

	SessionUserID = "userid"
)