package service

import "github.com/gin-gonic/gin"

// 合约地址
var Contract_address string

// 启动服务
func Start(contract_address string) (err error) {
	r := gin.Default()

	r.SetTrustedProxies([]string{"169.254.0.42"})

	// 初始化合约地址
	Contract_address = contract_address

	// 注册
	register := r.Group("/register")
	{
		register.POST("/sendEmail", user_register_sendEmailHandler)
		register.POST("/verifyEmail", user_register_verifyEmailHandler)
		register.POST("/user", user_registerHandler)
		register.POST("/verifyBizlicense", gainer_register_verifyBizlicense)
	}

	// 登录
	login := r.Group("/login")
	{
		login.POST("/user", user_loginHandler)
	}

	// 用户（分享者）
	user := r.Group("/user")
	{
		user.POST("/verifyIDCard", user_verifyIDCardHandler)
		user.POST("/readMedicalInformation", user_readMedicalInformation)
	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		gainer.POST("/")
	}

	err = r.Run(":80")
	return err
}
