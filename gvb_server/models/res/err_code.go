package res

type ErrorCode int

// 错误状态码
const (
	SettingsError ErrorCode = 1001 // 系统错误
	ArgumentError ErrorCode = 1002 //参数错误
)

//状态码用纯json也行

// 状态码和错误关系映射Map
var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
