package api_record_group

import (
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type apiRecordGroupCont struct {
	dao record_group_dao_interface.ACSRecordGroupDao
}

func NewApiRecordGroupCont(dao record_group_dao_interface.ACSRecordGroupDao) *apiRecordGroupCont {
	return &apiRecordGroupCont{dao: dao}
}

func (c apiRecordGroupCont) List(ctx *gin.Context) {
	rt := map[string]interface{}{}
	list, err := c.dao.List()
	if err != nil {
		rt["msg"] = "获取失败！"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}
	rt["data"] = list
	rt["msg"] = "Success！"
	rt["code"] = 200
	ctx.JSON(http.StatusOK, rt)
}
func (c apiRecordGroupCont) Update(ctx *gin.Context) {
	rt := map[string]interface{}{}
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id <= 0 {
		rt["msg"] = "id 不能为空！"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}
	if err := VerifyCanDesc(id); err != nil {
		rt["msg"] = err.Error()
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}
	name := ctx.PostForm("name")
	desc := ctx.PostForm("desc")
	if name == "" {
		rt["msg"] = "name 不能为空!"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}

	err := c.dao.UpByGroupId(
		&model.BaseRecordGroupModel{
			Id:   id,
			Name: name,
			Desc: desc,
		},
	)
	if err != nil {
		rt["msg"] = "更新失败！"
		rt["code"] = 102
		ctx.JSON(http.StatusOK, rt)
		return
	}

	rt["msg"] = "Success！"
	rt["code"] = 200
	ctx.JSON(http.StatusOK, rt)
}
func (c apiRecordGroupCont) Delete(ctx *gin.Context) {
	rt := map[string]interface{}{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id <= 0 {
		rt["msg"] = "id 不能为空！"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}
	if err := VerifyCanDesc(id); err != nil {
		rt["msg"] = err.Error()
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}

	err := c.dao.DelByGroupId(id)
	if err != nil {
		rt["msg"] = "删除失败！"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}
	rt["msg"] = "Success！"
	rt["code"] = 200
	ctx.JSON(http.StatusOK, rt)
}
func (c apiRecordGroupCont) Add(ctx *gin.Context) {
	rt := map[string]interface{}{}
	name := ctx.PostForm("name")
	desc := ctx.PostForm("desc")
	if name == "" {
		rt["msg"] = "name 不能为空!"
		rt["code"] = 101
		ctx.JSON(http.StatusOK, rt)
		return
	}

	err := c.dao.Add(
		&model.BaseRecordGroupModel{
			Name: name,
			Desc: desc,
		},
	)
	if err != nil {
		rt["msg"] = "添加失败！请检查是否重复添加！"
		rt["code"] = 102
		ctx.JSON(http.StatusOK, rt)
		return
	}

	rt["msg"] = "Success！"
	rt["code"] = 200
	ctx.JSON(http.StatusOK, rt)
}

func VerifyCanDesc(id int) error {
	if id == 1 || id == 2 {
		return errors.New("系统预设分组不支持修改删除")
	}
	return nil
}
