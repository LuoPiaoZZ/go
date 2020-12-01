package define

const (
	YMD ="2006-01-02"
	LPDB="lpdb"
	LPCache="lpcache"
	TableUserPropEveryDay="user_prop_everyday"
	// redis缓存key
	MiddlewareUserCache = "usercache"
)

const (
	ErrorCode   = -1
	SuccessCode = 0
)

const (
	ErrParam   = "参数错误"
	ErrDB      = "数据库查询出错"
	StatusOk   = "successful"
	ErrRequest = "请求次数太频繁"
)
