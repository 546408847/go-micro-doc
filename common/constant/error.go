package constant

const (
	/*错误码 2位系统 + 2位模块 + 3位业务*/
	ErrorOk = 0 //正确

	/*系统 - 00*/
	ErrorSystemError     = 100000 //系统异常
	ErrorSystemRateLimit = 100001 //请求过于频繁
	ErrorSystemUpgrade   = 100002 //系统正在升级维护
	ErrorDb              = 100003 //数据库错误
	ErrorParams          = 100004 //参数错误
	ErrorHttpRequest     = 100005 //http请求第三方错误
	ErrorRpcRequest      = 100006 //rpc请求第三方错误

	/*用户 - 01*/
	ErrorNoLogin                = 100100 //用户未登录
	ErrorUserMobile             = 100101 //手机号码错误
	ErrorUserAlreadyRegister    = 100102 //用户已注册
	ErrorPhoneOrPasswordInvalid = 100103 //手机号码或密码不正确
	ErrorUserFroze              = 100104 //用户被冻结,请联系客服

	/*订单 - 02*/
	ErrorApplyNotFound = 100200 //订单不存在

	/*支付 - 03*/
	ErrorNoDoublePayment = 100300 //不能重复支付

)

var errorMap = map[int]string{
	ErrorOk:                     "成功",
	ErrorSystemError:            "系统异常",
	ErrorSystemRateLimit:        "请求过于频繁",
	ErrorSystemUpgrade:          "系统正在升级维护",
	ErrorDb:                     "保存错误",
	ErrorParams:                 "参数错误",
	ErrorHttpRequest:            "请求第三方错误",
	ErrorRpcRequest:             "请求第三方错误",
	ErrorNoLogin:                "用户未登录",
	ErrorUserMobile:             "手机号码错误",
	ErrorUserAlreadyRegister:    "用户已注册",
	ErrorPhoneOrPasswordInvalid: "手机号码或密码不正确",
	ErrorUserFroze:              "尊敬的用户，此手机号不可使用，请更换手机号进行操作",
	ErrorApplyNotFound:          "订单不存在",
	ErrorNoDoublePayment:        "不能重复支付",
}

/**
 * 获取错误信息
 */
func ErrorMessage(code int) string {
	if message, ok := errorMap[code]; ok {
		return message
	}
	return "系统异常"
}
