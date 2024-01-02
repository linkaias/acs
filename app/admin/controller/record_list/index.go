package record_list

import (
	"APICallerStats/app/admin/model"
	"APICallerStats/common/record_dao_interface"
	"github.com/gin-gonic/gin"
)

type AdminRecordCont struct {
	Dao record_dao_interface.ACSRecordDao
}

func (d AdminRecordCont) List(ctx *gin.Context) {
	query := new(model.AdminRecordListRequestModel)
	if err := ctx.ShouldBind(query); err != nil {
		ctx.JSON(
			200, gin.H{
				"code": 101,
				"msg":  err.Error(),
			},
		)
		return
	}
	query.Verify()
	count, list, err := d.Dao.List(&query.BaseRecordModel, &query.BaseListRequestModel)
	if err != nil {
		ctx.JSON(
			200, gin.H{
				"code": 102,
				"msg":  err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		200, gin.H{
			"code":  200,
			"msg":   "Success！",
			"data":  list,
			"total": count,
			"page":  query.Page,
		},
	)
}
