package constant

const (
	SignPath  = "/lovelypet"
	Signup = "/signup"
	Signin = "/signin"
	Signout = "/signout"

	MoodPath = "/lovelypet/mood"
	MoodSubmit = "/submit"
	MoodDelete = "/delete"
)

const
(
	PARAM_NIL = "参数不正确"
	PARAM_LOST = "缺少[ %s ]参数"
	SIGNUP_SUCCESS = "注册成功"
	SIGNUP_FAIL = "注册失败"
	SIGNIN_SUCCESS = "登录成功"
	SIGNIN_FAIL = "登录失败"

	AUTH_TOKEN_LOST = "请求未携带token，权限不足,请在[header/cookie/请求参数]任一里面设置,key = lovelyToken"
	AUTH_TOKEN_INVALID = "token无效，请重新登录"
)

const
(
	SUCCESS = 1//成功
	FATAL = -1//失败
)

