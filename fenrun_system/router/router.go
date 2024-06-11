package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	. "xAdmin/apis"
	_ "xAdmin/docs"
	"xAdmin/handler"
	"xAdmin/handler/sd"
	_ "xAdmin/pkg/jwtauth"
	"xAdmin/router/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(middleware.LoggerToFile())
	r.Use(middleware.CustomError)
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestId())
	r.Use(middleware.DemoEvn())
	r.Use(gin.Recovery())
	r.Static("/static", "./static")
	//导出
	r.Static("/export", "./runtime/export")
	r.GET("/info", Ping)
	r.GET("/heath", Heath)
	//r.MaxMultipartMemory = config.UpdateFile.MaxSize // 上传文件大小限制
	// 监控信息
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/os", sd.OSCheck)
	}

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()

	if err != nil {
		_ = fmt.Errorf("JWT Error", err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler) //登录

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	//r.GET("/dashboard", Dashboard)
	r.GET("/routes", Dashboard)
	//r.GET("/getpost", PostTest)

	//无权限限制接口
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/getCaptcha", GenerateCaptchaHandler) //获取图片验证码

		apiv1.GET("/rolemenu", GetRoleMenu)
		apiv1.POST("/rolemenu", InsertRoleMenu)
		apiv1.DELETE("/rolemenu/:id", DeleteRoleMenu)
		apiv1.GET("/dict/databytype/:dictType", GetDictDataByDictType)

	}

	//权限限制接口
	auth := r.Group("/api/v1")

	auth.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		auth.Static("/export", "./runtime/export")
		auth.POST("/logout", handler.LogOut)            //退出系统
		apiv1.GET("/menuTreeselect", GetMenuTreeSelect) //获取菜单树

		auth.GET("/deptList", GetDeptList)   //获取部门列表
		auth.GET("/deptTree", GetDeptTree)   //获取部门树形图
		auth.GET("/dept/:deptId", GetDept)   //根据部门ID 获取部门
		auth.POST("/dept", InsertDept)       //添加部门
		auth.PUT("/dept", UpdateDept)        //修改部门
		auth.DELETE("/dept/:id", DeleteDept) //删除部门

		auth.GET("/dict/datalist", GetDictDataList)         //获取字典列表
		auth.GET("/dict/data/:dictCode", GetDictData)       //根据字典代码获取字典
		auth.POST("/dict/data", InsertDictData)             //添加字典
		auth.PUT("/dict/data/", UpdateDictData)             // 修改字典
		auth.DELETE("/dict/data/:dictCode", DeleteDictData) //删除字典数据

		auth.GET("/dict/typelist", GetDictTypeList)       //字典类型列表数据
		auth.GET("/dict/type/:dictId", GetDictType)       //通过字典id获取字典类型
		auth.POST("/dict/type", InsertDictType)           //添加字典类型
		auth.PUT("/dict/type", UpdateDictType)            //修改字典类型
		auth.DELETE("/dict/type/:dictId", DeleteDictType) //删除字典类型

		auth.GET("/sysUserList", GetSysUserList)               // 系统用户列表
		auth.GET("/sysUser/:userId", GetSysUser)               //根据用户ID 获取用户
		auth.GET("/sysUser/", GetSysUserInit)                  //获取用户角色和职位
		auth.POST("/sysUser", InsertSysUser)                   //添加管理员
		auth.PUT("/sysUser", UpdateSysUser)                    //修改用户数据
		auth.DELETE("/sysUser/:userId", DeleteSysUser)         //删除用户数据
		auth.GET("/passingSysUserList", GetPassingSysUserList) // 获取审核中的列表
		auth.GET("/userAuditRecord", GetUserAuditRecord)       // 获取审核记录
		auth.POST("/submitNewUser", SubmitNewUser)             // 提交新的用户
		auth.POST("/allowNewUserPass", AllowNewUserPass)       // 允许新用户通过

		auth.GET("/rolelist", GetRoleList)              //角色列表
		auth.GET("/role/:roleId", GetRole)              //角色
		auth.PUT("/roleAdd", InsertRole)                //添加角色
		auth.PUT("/role", UpdateRole)                   //修改角色
		auth.PUT("/roledatascope", UpdateRoleDataScope) //修改角色数据
		auth.DELETE("/role/:roleId", DeleteRole)        //删除角色
		auth.PUT("/roleDel", DeleteRole)                //删除角色

		//参数设置
		auth.GET("/configList", GetConfigList)    //配置列表
		auth.GET("/config/:id", GetConfig)        //获取配置
		auth.POST("/config", InsertConfig)        //添加配置
		auth.PUT("/config", UpdateConfig)         //修改配置
		auth.DELETE("/config/:id", DeleteConfig)  //删除配置
		auth.POST("/config/export", ExportConfig) //导出配置

		auth.GET("/roleMenuTreeselect/:roleId", GetMenuTreeRoleselect) //获取角色菜单树
		auth.GET("/roleDeptTreeselect/:roleId", GetDeptTreeRoleselect) //获取部门菜单树

		auth.GET("/getinfo", GetInfo)                        //获取管理员信息
		auth.GET("/user/profile", GetSysUserProfile)         //获取当前登录用户
		auth.POST("/user/profileAvatar", InsetSysUserAvatar) //修改用户头像
		auth.PUT("/user/pwd", SysUserUpdatePwd)              //修改管理员密码
		auth.GET("/user/getVerificationCode", CreateSecret)  //创建谷歌验证码
		auth.POST("/user/bindVerificationCode", BindCode)    //绑定谷歌验证

		auth.GET("/postlist", GetPostList)       //职位列表
		auth.GET("/post/:postId", GetPost)       //职位列表数据
		auth.POST("/post", InsertPost)           //添加职位
		auth.PUT("/post", UpdatePost)            //修改职位
		auth.DELETE("/post/:postId", DeletePost) //删除职位

		auth.GET("/menulist", GetMenuList)   //菜单列表
		auth.GET("/menu/:id", GetMenu)       //菜单数据
		auth.POST("/menu", InsertMenu)       //添加菜单
		auth.PUT("/menu", UpdateMenu)        //修改菜单
		auth.DELETE("/menu/:id", DeleteMenu) //删除菜单
		auth.GET("/menurole", GetMenuRole)   //获取角色权限
		auth.GET("/menuids", GetMenuIDS)     //获取角色对应的菜单id数组

		auth.GET("/loginloglist", GetLoginLogList)       //登录日志
		auth.GET("/loginlog/:infoId", GetLoginLog)       //通过编码获取登录日志
		auth.POST("/loginlog", InsertLoginLog)           //添加日志
		auth.PUT("/loginlog", UpdateLoginLog)            //修改日志
		auth.DELETE("/loginlog/:infoId", DeleteLoginLog) //删除日志

		auth.GET("/operloglist", GetOperLogList)       //操作日志
		auth.GET("/operlog/:operId", GetOperLog)       //获取日志
		auth.DELETE("/operlog/:operId", DeleteOperLog) //删除操作日志

		auth.GET("/configKey/:configKey", GetConfigByConfigKey) //通过配置Key获取配置信息

		//会员管理
		//auth.GET("/memberList", GetUsersList)                   //会员管理列表
		//auth.POST("/member/memberDelete", UsersDelete)          //删除会员
		//auth.POST("/member/memberEdit", UsersEdit)              //编辑会员
		//
		//
		//auth.GET("/memberExport", GetUsersListExport)              //会员列表 导出

		// 上传文件
		auth.POST("/uploadfile", UploadFile)          // 用户上传文件
		auth.GET("/getuploadfileList", GetUploadFile) // 当前用户上传文件列表
		auth.GET("/delfile", DeleteFile)              // 删除上传文件
		//矿多多
		auth.POST("/updateConfig", UpdateConfig1)
		auth.GET("/getConfigKey", GetConfigKey)

		auth.POST("/customerProfitEdit", CustomerProfitEdit) //编辑客户
		auth.GET("/getCustomerByid", GetCustomerByid)        //获取客户信息

		auth.GET("/individualPerformance", IndividualPerformance) //个人业绩

		auth.GET("/getInvestmentByid", GetInvestmentByid) //获取投资信息
		auth.POST("/investmentRevoke", InvestmentRevoke)  //撤销客户投资

		auth.GET("/profitconfigList", ProfitconfigList)              //分润配置列表
		auth.POST("/profitconfigOnce", ProfitconfigOnce)             //添加分润配置（一次性分配）
		auth.POST("/updateProfitConfigOnce", UpdateProfitconfigOnce) //修改分润配置（一次性分配）
		auth.POST("/profitconfigEdit", ProfitconfigEdit)             //业务员分润配置(提交修改)
		auth.POST("/delProfitconfigOnce", DelProfitconfigOnce)       //添加分润配置（一次性分配）

		auth.GET("/statementSalesman", StatementSalesman)             //业务员报表
		auth.GET("/statementSalesmanExport", StatementSalesmanExport) //业务员报表导出
		auth.GET("/statementCustomer", StatementCustomer)             //顾客报表
		auth.GET("/statementCustomerExport", StatementCustomerExport) //顾客报表导出

		auth.GET("/statementConfigOnce", StatementConfigOnce)     //一次性报表
		auth.POST("/statementSettlement", StatementSettlement)    //结算
		auth.GET("/statementNoSettlement", StatementNoSettlement) //未结算订单列表
		auth.GET("/statementIsSettlement", StatementIsSettlement) //已结算订单列表
		//auth.GET("/statementAddOnce", StatementAddOnce)		//结算一次性报表
		auth.GET("/statementConfigOnceExport", StatementConfigOnceExport) //一次性报表导出

		auth.GET("/statementSummary", StatementSummary)                //报表汇总
		auth.GET("/statementSummaryExport", StatementSummaryExport)    // 报表汇总导出
		auth.POST("/statementCustomerSettle", StatementCustomerSettle) // 客户结算
		auth.GET("/statementSettleHistory", StatementSettleHistory)    // 获取结算历史

		auth.GET("/getUserProfits", GetUserProfits) //收益

		auth.GET("/customerAuditList", CustomerAuditList) //审核列表
		auth.GET("/customerLogList", CustomerLogList)     //审核记录
		auth.GET("/customerLog", CustomerLog)             //审核记录
		auth.POST("/customerLogAudit", CustomerLogAudit)  //审核

		auth.POST("/editlevellock", Editlevellock) //等级锁定

		auth.GET("/getUserVipLevel", GetUserVipLevel)    // 获取用户的vip等级
		auth.POST("/editUserVipLevel", EditUserVipLevel) // 编辑玩家的vip等级

		auth.GET("/getUserPerformance", GetUserPerformance) // 获取主页显示内容

		auth.GET("/getUserReferrer", GetUserReferrer)   //上级
		auth.GET("/getUserReferrals", GetUserReferrals) //下级

		//auth.GET("/investmentExport", InvestmentExport)		//投资列表导出
		//auth.POST("/investmentImport", InvestmentImport)		//导入
		//auth.POST("/emptyInvestment", EmptyInvestment)		//清空订单信息

		//++++++++++分润++++++++++
		auth.GET("/financeConfigList", FinanceConfigList)  //配置列表
		auth.POST("/financeConfigEdit", FinanceConfigEdit) //修改配置

		auth.GET("/customerList", GetcustomerList)   //客户列表
		auth.POST("/customerAdd", CustomerAdd)       //添加客户
		auth.POST("/customerEdit", CustomerEdit)     //编辑客户
		auth.POST("/customerDelete", CustomerDelete) //删除客户

		auth.GET("/investmentList", InvestmentList)      //客户投资列表
		auth.POST("/investmentAdd", InvestmentAdd)       //添加客户投资
		auth.POST("/investmentEdit", InvestmentEdit)     //编辑客户投资
		auth.POST("/investmentBreak", InvestmentBreak)   //终止客户投资
		auth.POST("/investmentDelete", InvestmentDelete) //删除客户投资

		auth.GET("/userList", GetUserList)

		auth.GET("/getWeekList", GetWeekList) // 获取以周而结算算收益列表列表
		//auth.GET("/getWeekCustomerList", GetWeekCustomerList)                     // 获取周的用户列表
		auth.GET("/getWeekCustomerList", GetWeekCustomerList1)                    // 获取周的用户列表
		auth.GET("/getWeekCustomerListExport", GetWeekCustomerListExport)         // 获取周的用户列表导出
		auth.GET("/getWeekCustomerInvestmentList", GetWeekCustomerInvestmentList) // 获取周用户投资列表

		auth.POST("/transferWeek", TransferWeek)                 // 周转账
		auth.POST("/transferWeekCustomer", TransferWeekCustomer) // 用户周转账

		auth.GET("/getTransferList", GetTransferList) // 转账记录
		auth.GET("/getSettleList", GetSettleList)     // 获取结算记录

		auth.GET("/getwalletdefault", GetWallet)

		auth.POST("/addmigrate", AddMigrate)       // 添加迁移用户数据
		auth.GET("/getmigrate", GetMigrate)        // 查询所有订单
		auth.POST("/deletemigrate", DeleteMigrate) // 删除订单
		auth.POST("/editmigrate", EditMigrate)     //修改编辑订单
		//auth.POST("/confirmmigrate", ConfirmMigrate) //确认订单是否生效
		auth.POST("/breakmigreate", BreakMigreate)       // 终止订单
		auth.GET("/impressionslist", BillingImpressions) // 迁移用户数据展示
	}

	// 客户接口
	customer := r.Group("/customer")
	{
		customer.POST("/login", CustomerLogin)
		customer.GET("/investmentList", CustomerInvestmentList)
		customer.POST("/changePassword", CustomerChangePassword)
		customer.GET("/settlementList", CustomerSettlementList)
		customer.GET("/customerTransferList", CustomerTransferList)
		customer.GET("/impressionsList", CustomerSettlementListMigrate) // 迁移用户数据展示
		customer.GET("/homepage", CustomerHomepage)

	}
	//批量导入地址
	r.POST("/api/v1/uploadAddress", UploadAddress) //批量导入账号地址
	//接口文档
	r.GET("/admin/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func Ping(c *gin.Context) {
	log.Println(viper.GetString(`settings.database.type`))
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func Heath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}

func Dashboard(c *gin.Context) {

	var user = make(map[string]interface{})
	user["login_name"] = "admin"
	user["user_id"] = 1
	user["user_name"] = "管理员"
	user["dept_id"] = 1

	var cmenuList = make(map[string]interface{})
	cmenuList["children"] = nil
	cmenuList["parent_id"] = 1
	cmenuList["title"] = "用户管理"
	cmenuList["name"] = "Sysuser"
	cmenuList["icon"] = "user"
	cmenuList["order_num"] = 1
	cmenuList["id"] = 4
	cmenuList["path"] = "sysuser"
	cmenuList["component"] = "sysuser/index"

	var lista = make([]interface{}, 1)
	lista[0] = cmenuList

	var menuList = make(map[string]interface{})
	menuList["children"] = lista
	menuList["parent_id"] = 1
	menuList["name"] = "Upms"
	menuList["title"] = "权限管理"
	menuList["icon"] = "example"
	menuList["order_num"] = 1
	menuList["id"] = 4
	menuList["path"] = "/upms"
	menuList["component"] = "Layout"

	var list = make([]interface{}, 1)
	list[0] = menuList
	var data = make(map[string]interface{})
	data["user"] = user
	data["menuList"] = list

	var r = make(map[string]interface{})
	r["code"] = 200
	r["data"] = data

	c.JSON(200, r)
}
