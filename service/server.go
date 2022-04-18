package service

import "github.com/gin-gonic/gin"

// 合约地址
var Contract_address string

// 启动服务
func Start(contract_address string) {
	r := gin.Default()

	// 初始化合约地址
	Contract_address = contract_address

	// 注册
	register := r.Group("/register")
	{
		register.POST("/sendEmail", user_register_sendEmailHandler)
		register.POST("/verifyEmail", user_register_verifyEmailHandler)

	}

	// 用户（分享者）
	user := r.Group("/user")
	{
		register = user.Group("/register")
		{
			register.POST("/", user_registerHandler)
		}

		login := user.Group("/login")
		{
			login.POST("/", user_loginHandler)
		}

		user.POST("/verifyIDCard", user_verifyIDCardHandler)
		user.POST("/readMedicalInformation", user_readMedicalInformation)

	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		register = gainer.Group("/register")
		{
			register.POST("/verifyBizlicense", gainer_register_verifyBizlicense)
		}

		gainer.POST("/")
	}

	r.Run(":80")

}
