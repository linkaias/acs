package bootstrap

import (
	router2 "APICallerStats/app/admin/router"
	"APICallerStats/app/grpc"
	"APICallerStats/app/http/router"
	"APICallerStats/app/serve/controller"
	"APICallerStats/common/chart_data_dao_interface"
	"APICallerStats/common/config_dao_interface"
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/crons"
	"APICallerStats/model"
	"APICallerStats/services/mysql_dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"reflect"
	"strconv"
	"time"
)

type Application struct {
	Env            *model.Env
	CfgDao         config_dao_interface.ConfigDao
	RecordDao      record_dao_interface.ACSRecordDao
	RecordGroupDao record_group_dao_interface.ACSRecordGroupDao
	ChartDataDao   chart_data_dao_interface.ChartDataDao
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.ServeInit()
	// 设置运行模式
	if app.Env.Debug {
		log.Println("The ACS is running in development env")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return *app
}

func (a *Application) ServeInit() {
	// 校验配置文件

	// 设置db数据库
	if a.Env.DBType == "mysql" {
		a.CfgDao = mysql_dao.NewMysqlConfigDao(GetMysqlConn(a.Env))
		a.RecordDao = mysql_dao.NewMysqlRecordDao(GetMysqlConn(a.Env))
		a.RecordGroupDao = mysql_dao.NewMysqlRecordGroupDao(GetMysqlConn(a.Env))
		a.ChartDataDao = mysql_dao.NewMysqlChartDataDao(GetMysqlConn(a.Env))
	}

	// 判断是否已安装
	a.autoInstall()

	// 读取用户配置
	a.loadConfigByDb()

	//...
}

func (a *Application) Run() {
	// 记录日志
	controller.RunLog(a.Env)

	// 记录 api调用
	controller.RunRecordTask(a.Env, a.RecordDao)

	// cron
	go crons.RunCron(a.Env, a.RecordDao)

	// 启动http服务 和 admin服务
	go func() {
		r := gin.New()
		// http server
		router.InitFrontRouter(r, a.RecordGroupDao, a.Env)
		//admin
		router2.InitAdminRouter(r, a.ChartDataDao, a.RecordGroupDao, a.Env, a.RecordDao, a.CfgDao)
		err := r.Run(":" + a.Env.HttpServePort)
		if err != nil {
			panic(err)
		}
	}()
	// 启动gRPC服务
	go grpc.RunGRPCServer(a.Env, a.RecordDao, a.RecordGroupDao)

	select {}

}

// 自动安装初始化
func (a *Application) autoInstall() {
	// 待添加重新安装功能...

	if a.CfgDao.Get("is_installed") == "1" {
		return
	}
	// 创建数据表
	if a.Env.DBType == "mysql" {
		// 配置表
		err := GetMysqlConn(a.Env).AutoMigrate(&model.ACSConfigMysqlModel{})
		if err != nil {
			panic(err)
		}
		// api记录表
		err = GetMysqlConn(a.Env).AutoMigrate(&model.RecordMysqlModel{})
		if err != nil {
			panic(err)
		}
		_ = a.CfgDao.Set("install_time", time.Now().Format("2006-01-02 15:04:05"))

		// api记录分组表
		err = GetMysqlConn(a.Env).AutoMigrate(&model.RecordGroupMysqlModel{})
		if err != nil {
			panic(err)
		}
		_ = a.RecordGroupDao.Add(&model.BaseRecordGroupModel{Id: 1, Name: "默认分组", Desc: "默认分组"})
		_ = a.RecordGroupDao.Add(
			&model.BaseRecordGroupModel{
				Id: 2, Name: "Admin访问统计", Desc: "记录admin后台访问记录",
			},
		)
	}
	_ = a.CfgDao.Set("is_installed", "1")
}

func (a *Application) loadConfigByDb() {
	t := reflect.TypeOf(*a.Env)
	v := reflect.ValueOf(*a.Env)
	isReloadDb := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		old := a.CfgDao.Get(field.Name)
		if old != "" && old != fmt.Sprintf("%v", value) {
			switch field.Name {
			case "Debug":
				isReloadDb = true
				val := old == "true"
				a.Env.Debug = val
				break
			case "DBBatchInsertNum":
				val, _ := strconv.Atoi(old)
				a.Env.DBBatchInsertNum = val
				break
			case "DBBatchInsertGap":
				val, _ := strconv.Atoi(old)
				a.Env.DBBatchInsertGap = val
				break
			case "GRPCServePort":
				a.Env.GRPCServePort = old
				break
			case "HttpServePort":
				a.Env.HttpServePort = old
				break
			case "AdminUser":
				a.Env.AdminUser = old
				break
			case "AdminPassword":
				a.Env.AdminPassword = old
				break
			case "SecretKey":
				a.Env.SecretKey = old
				break
			case "HTTPUser":
				a.Env.HTTPUser = old
				break
			case "HTTPPassword":
				a.Env.HTTPPassword = old
				break
			case "GRPCUser":
				a.Env.GRPCUser = old
				break
			case "GRPCPassword":
				a.Env.GRPCPassword = old
				break
			case "JWTExpireTime":
				val, _ := strconv.Atoi(old)
				a.Env.JWTExpireTime = val
				break
			case "JwtAdminExpireTime":
				val, _ := strconv.Atoi(old)
				a.Env.JwtAdminExpireTime = val
				break
			case "RecordSaveDays":
				val, _ := strconv.Atoi(old)
				a.Env.RecordSaveDays = val
				break
			case "LogFilePath":
				a.Env.LogFilePath = old
				break
			case "LogFileMaxAge":
				val, _ := strconv.Atoi(old)
				a.Env.LogFileMaxAge = val
				break
			case "LogFileGap":
				val, _ := strconv.Atoi(old)
				a.Env.LogFileGap = val
				break
			case "OpenAdminRecord":
				val := old == "true"
				a.Env.OpenAdminRecord = val
				break
			}
		}
	}

	if isReloadDb {
		// 设置db数据库
		if a.Env.DBType == "mysql" {
			a.CfgDao = mysql_dao.NewMysqlConfigDao(GetMysqlConn(a.Env))
			a.RecordDao = mysql_dao.NewMysqlRecordDao(GetMysqlConn(a.Env))
			a.RecordGroupDao = mysql_dao.NewMysqlRecordGroupDao(GetMysqlConn(a.Env))
			a.ChartDataDao = mysql_dao.NewMysqlChartDataDao(GetMysqlConn(a.Env))
		}
	}
}

// GetMysqlConn 获取mysql连接
func GetMysqlConn(env *model.Env) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.MysqlUsername,
		env.MysqlPassword,
		env.MysqlHost,
		env.MysqlPort,
		env.MysqlDatabase,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	if env.Debug {
		db.Logger = logger.Default.LogMode(logger.Info)
	} else {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}
	return db
}
