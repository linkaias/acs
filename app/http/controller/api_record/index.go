package api_record

import (
	model2 "APICallerStats/app/http/model"
	"APICallerStats/app/serve/controller"
	"APICallerStats/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type curlApiRecordCont struct{}

func NewCurlApiRecordCont() *curlApiRecordCont {
	return new(curlApiRecordCont)
}

func (c curlApiRecordCont) Add(ctx *gin.Context) {
	rt := map[string]interface{}{}
	val := new(model2.ApiRecordModel)
	err := ctx.ShouldBind(val)
	if err != nil {
		rt["msg"] = "参数解析错误！"
		rt["code"] = 100
		ctx.JSON(http.StatusOK, rt)
		return
	}
	if val.GroupId == 0 || val.Name == "" || val.Url == "" || val.Method == "" {
		rt["msg"] = "参数错误!"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}
	ip := ctx.ClientIP()
	if val.Ip != "" {
		ip = val.Ip
	}
	controller.AddACSRecord(
		&model.BaseRecordModel{
			GroupId: val.GroupId,
			Name:    val.Name,
			Url:     val.Url,
			Method:  val.Method,
			IP:      ip,
			AddTime: time.Now(),
		},
	)
	rt["msg"] = "Success！"
	rt["code"] = 200
	ctx.JSON(http.StatusOK, rt)
}
