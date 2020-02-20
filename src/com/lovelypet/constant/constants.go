package constant

const (
	SignPath = "/lovelypet"
	Signup   = "/signup"
	Signin   = "/signin"
	Signout  = "/signout"

	MoodPath   = "/lovelypet/mood"
	MoodSubmit = "/submit"
	MoodDelete = "/delete"
	MoodQuery  = "/query"
)

const
(
	Token             = "lovelyToken"
	ParamNil          = "参数不正确"
	ParamLost         = "缺少[ %s ]参数"
	SignupSuccess     = "注册成功"
	SignupRepeat      = "该手机号已经被注册"
	SignupNotRecord   = "该手机号未注册"
	SigninSuccess     = "登录成功"
	PwdError          = "密码错误"
	ServerSqlError    = "服务器数据插入失败"
	AuthTokenLost     = "请求未携带token，权限不足,请在[header/cookie/请求参数]任一里面设置,key = lovelyToken"
	AuthTokenResignin = "请重新登录"
)

const
(
	SUCCESS      = 1  //成功
	FATAL        = -1 //失败
	TokenExpired = -2 //token过期
)
