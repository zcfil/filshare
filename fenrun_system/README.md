# TCO-admin

English | 简体中文

基于Gin + Vue + Element UI的前后端分离权限管理系统

系统初始化极度简单，只需要配置文件中，修改数据库连接，系统启动后会自动初始化数据库信息以及必须的基础数据

## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库

- 配置文件简单的模型映射，快速能够得到想要的配置

- 代码生成工具（即将push，demo已经发布可以体验了）

- 表单构建工具（即将push，demo已经发布可以体验了）

- TODO: 单元测试

## 🎁 内置功能

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 参数管理：对系统动态配置常用参数。
8. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 系统接口：根据业务代码自动生成相关的api接口文档。
11. 代码生成：根据数据表结构生成对应的增删改查相对应业务，全部可视化编程。
12. 表单构建：自定义页面样式，拖拉拽实现页面布局。

## 🎁 项目目录结构

```
xAdmin : 主业目录
  apis : 后台服务API接口
    |-- apis.go: 游戏下注总数,彩票投注详情,统计当月游戏玩次数,连胜统计,游戏抽水收益查询,充提币统计
    |-- bet.go : 彩票管理
    |-- betrelease.go : 通兑管理
    |-- binary.go : 二元期权管理
    |-- bonus.go : 庄家池数据
    |-- captcha.go : 图片验证码
    |-- coincharge.go : 财务管理
    |-- connfig.go : 配置数据
    |-- dept.go : 部门管理
    |-- diceRecord.go : Dice游戏管理
    |-- dictData.go : 字典数据管理
    |-- dictType.go : 字典类型管理
    |-- exception.go :充值异常处理
    |-- exchange.go : 通兑管理
    |-- fomoRecord.go :Fomo 管理
    |-- info.go : 系统基础信息
    |-- loginLog.go : 系统登录日志
    |-- lottery.go : 彩票管理
    |-- lotteryset.go : 彩票设置
    |-- menu.go : 菜单管理
    |-- operlog.go : 系统操作日志
    |-- pledge.go : 质押管理
    |-- post.go : 职位管理
    |-- role.go : 角色管理
    |-- rolemenu.go : 角色菜单
    |-- sysuser.go : 系统管理员
    |-- UserMember : 会员管理
    |-- usermoneylog.go : 财务管理
    |-- wallets.go : 节点钱包管理
    |-- withdraw.go : 提币管理
  config : 项目配置文件
    |-- application.go: 项目配置信息
    |-- cxcblock.go: CXC 公链配置
    |-- database.go: db初始化配置
    |-- db.sql: 初始化数据库
    |-- redis.go: redis 配置
    |-- settings.yml: 配置文件
  database : 数据库 
    |-- mysql.go: Mysql数据库操作
  docs  : swagger :接口文档
    |-- docs.go :文档操作
    |-- swagger.json : swagger json 数据
    |-- swagger.yaml : swagger yaml文件
  handler : 公共处理目录
  	|-- cornjob	: 定时任务
  	|-- sd : 服务器检查
  	|-- auth.go : 权限校验【登录，退出，授权】
  	|-- nofound.go : 无权限访问
  	|-- redisconn : Redis 链接
  models :  数据库操作件目录
  	|-- accountRecord.go : 账户记录
	|-- apismodel.go
	|-- block.go  : 同步节点区块信息
	|-- bonusConfig.go : 平台庄家池数据
	|--casbinrule.go
	|-- comm.go		:models层的一些通用方法
	|--coinCharging.go : 用户充币
	|--coinWithdraw.go : 用户提币
	|--collectWallet.go : 节点资金归集
	|--dept.go : 部门
	|--dictType.go : 字典类型
	|--gooleAuth.go  : 谷歌验证
	|--initdb.go : 数据库脚本初始化
	|--load.go : 获取总数
	|--pair.go : 交易对
	|--payOrder.go : 支付订单
	|--response.go : 请求响应
	|--sysAdmin.go : 系统管理员
	|--sysConfig.go : 系统配置
	|--sysDataScope.go : 数据权限
	|--sysDictData.go  : 字典数据
	|--sysLogin.go : 登录
	|--sysLoginlog.go : 登录日志
	|--sysMenu.go	: 系统菜单
	|--sysOperlog.go : 系统操作日志
	|--sysPost.go : 岗位
	|--sysRole.go : 角色
	|--sysRoleDept.go : 系统角色部门
	|--sysRoleMenu.go : 角色菜单
	|--userKyc.go : 用户KYC审核
	|--userAssets.go : 用户资产
	|--userLevel.go : 会员等级
	|--userMember.go : 会员表
	|--userMoneyLog.go : 用户资金记录
  pkg: 资源包
  	|-- auth : 定时自动工作
		|--ollect.go : CXC节点资金自动归集（未用）
		|--autoEther.go : ETH节点资金自动归集
	|-- casbin : 授权库，支持访问控制模型
	|-- cxcblock : CXC 公链
	|-- export : EXECEL 导入导出
	|-- file :  文件操作
	|-- googleAuthenticator :  谷歌验证码操作
	|-- jwtauth : JWT权限验证
	|-- setting : 系统设置 
	|-- utils.go : 工具类
  router：API路由	
  	|-- middleware :  中间件
	|-- auth.go :  权限初始化
	|-- customerror.go : 自定义错误
	|-- demo.go : 
	|-- header.go : 请求头
	|-- logger.go : 日志记录到文件
	|-- permission.go : 权限检查中间件
	|-- requestid.go : 请求ID
	|-- router.go : API路由初始化
  runtime : 项目运行目录主要是数据缓存
  static : 服务区静态目录、上传的图片，及资源
  utils : 工具类
  	|-- string.go : 字符串工具
	|-- user.go : 用户操作
  go.mod : go mod 包管理
  main.go :服务启动入口
  README.md: 工程描述文件
```

## 🧡配置详情

1. 配置文件说明

```yml
settings:
  application:  
    # 项目启动环境            
    env: dev  
    # 当 env:demo 时，GET以外的请求操作提示
    envmsg: "谢谢您的参与，但为了大家更好的体验，所以本次提交就算了吧！" 
    # 主机ip 或者域名，默认0.0.0.0
    host: 0.0.0.0 
    # 是否需要初始化数据库结构以及基本数据；true：需要；false：不需要 
    isinit: false  
    # JWT加密字符串
    jwtsecret: 123abc  
    # log存放路径
    logpath: temp/logs/log.log   
    # 服务名称
    name: go-admin   
    # 服务端口
    port: 8000   
    readtimeout: 1   
    writertimeout: 2 
  database:
    # 数据库名称
    database: dbname 
    # 数据库类型
    dbtype: mysql    
    # 数据库地址
    host: 127.0.0.1  
    # 数据库密码
    password: password  
    # 数据库端口
    port: 3306       
    # 数据库用户名
    username: root   
  # redis 可忽略
  redis: 
    # redis链接地址
    addr: 0.0.0.0:6379 
    # db 
    db: 0   
    # 密码            
    password: password  
    # 读超时时长
    readtimeout: 50   
  # CXC 公链链接
  cxcblock:
    # 服务器地址
    serverHost: 47.57.65.221 
    # 服务器端口
    serverPort: 7318
    # 链接用户
    user: cxcsrpc
    # 密码
    passwd:  5rFDt7XjbhP9ge5vSaEuA68XeFo49Sd3C23wE4fQ3pzu
    # 使用SSL加密
    useSSL: false
```

2. 文件路径 go-admin/config/settings.yml

## 📦 本地开发

首次启动说明

```bash
# 获取代码
git clone 

# 进入工作路径
cd ./toc-admin

# 编译项目
go build

# 修改配置
vi ./config/setting.yml (更改isinit和数据库连接)

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据)
# 2. 确认数据库初始化参数 
# 注意: settings.application.isinit 如果是首次启动，请把当前值设置成true，系统会自动初始化数据库结构以及基本的数据信息；
# 3. 确认log路径


# 启动项目，也可以用IDE进行调试
./toc-admin

# 也可以在WIKI中查看说明
```

## 📦 数据库详情

文档生成

```bash
swag init  
```

如果没有swag命令 go get安装一下即可

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

交叉编译

```bash
env GOOS=windows GOARCH=amd64 go build main.go

# or

env GOOS=linux GOARCH=amd64 go build main.go
```

## 🤝 使用的开源项目

[gin](https://github.com/gin-gonic/gin)

[casbin](https://github.com/casbin/casbin)

[spf13/viper](https://github.com/spf13/viper)

[gorm](https://github.com/jinzhu/gorm)

[gin-swagger](https://github.com/swaggo/gin-swagger)

[jwt-go](https://github.com/dgrijalva/jwt-go)

[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)

[ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue)

## 版本

### 2020-04 新功能

## ss