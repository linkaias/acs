package system

import (
	"APICallerStats/common/config_dao_interface"
	"APICallerStats/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"runtime"
)

type AdminSystemCont struct {
	Env       *model.Env
	ConfigDao config_dao_interface.ConfigDao
}

func (c AdminSystemCont) GetInfo(ctx *gin.Context) {
	rt := make(map[string]interface{})
	data := make([]map[string]interface{}, 0)
	// system
	{
		data = append(
			data, map[string]interface{}{
				"System-Debug": c.Env.Debug,
			},
			map[string]interface{}{
				"System-OS": runtime.GOOS,
			},
			map[string]interface{}{
				"System-ARCH": runtime.GOARCH,
			},
			map[string]interface{}{
				"System-NumCPU": runtime.NumCPU(),
			},
		)
	}

	// go
	{
		data = append(
			data, map[string]interface{}{
				"GO-goVersion": runtime.Version(),
			},
			map[string]interface{}{
				"GO-goGoRoot": runtime.GOROOT(),
			},
			map[string]interface{}{
				"GO-NumGoroutines": runtime.NumGoroutine(),
			},
			map[string]interface{}{
				"GO-NumCgoCall": runtime.NumCgoCall(),
			},
			map[string]interface{}{
				"GO-NumGOMAXPROCS": runtime.GOMAXPROCS(0),
			},
			map[string]interface{}{
				"Gin-GinMode": gin.Mode(),
			},
			map[string]interface{}{
				"Gin-GinVersion": gin.Version,
			},
		)
	}
	rt["code"] = 200
	rt["msg"] = "Success！"
	rt["data"] = data
	ctx.JSON(200, rt)
}

func (c AdminSystemCont) GetConfig(ctx *gin.Context) {
	info, _ := json.Marshal(c.Env)
	ctx.JSON(
		200, gin.H{
			"code":    200,
			"msg":     "Success！",
			"data":    string(info),
			"is_save": c.ConfigDao.Get("is_save_config"),
		},
	)
}

func (c AdminSystemCont) SaveConfig(ctx *gin.Context) {
	query := new(model.Env)
	err := ctx.ShouldBind(query)
	if err != nil {
		ctx.JSON(
			200, gin.H{
				"code": 101,
				"msg":  err.Error(),
			},
		)
		return
	}
	c.saveCfgToDb(*query)
	_ = c.ConfigDao.Set("is_save_config", "1")

	ctx.JSON(
		200, gin.H{
			"code": 200,
			"msg":  "Success！",
		},
	)
}

func (c AdminSystemCont) saveCfgToDb(env model.Env) {
	t := reflect.TypeOf(env)
	v := reflect.ValueOf(env)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		_ = c.ConfigDao.Set(field.Name, fmt.Sprintf("%v", value))
	}
}

func (c AdminSystemCont) TestDBConn(ctx *gin.Context) {
	query := new(model.Env)
	err := ctx.ShouldBind(query)
	if err != nil {
		ctx.JSON(
			200, gin.H{
				"code": 101,
				"msg":  err.Error(),
			},
		)
		return
	}
	if query.DBType != "mysql" {
		ctx.JSON(
			200, gin.H{
				"code": 101,
				"msg":  "仅支持mysql",
			},
		)
		return
	}
	conn, err := testMysqlConn(query)
	if err != nil || conn.Error != nil {
		ctx.JSON(
			200, gin.H{
				"code": 101,
				"msg":  "连接失败！",
			},
		)
		return
	}
	ctx.JSON(
		200, gin.H{
			"code": 200,
			"msg":  "Success！",
		},
	)
}

func testMysqlConn(env *model.Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
		env.MysqlUsername,
		env.MysqlPassword,
		env.MysqlHost,
		env.MysqlPort,
		env.MysqlDatabase,
	)
	return gorm.Open(mysql.Open(dsn))
}
