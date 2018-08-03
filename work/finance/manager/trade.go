package manager

/*
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	//	"sub_account_service/finance/config"
	//	"sub_account_service/finance/db"
	"sub_account_service/finance/lib"
	_ "sub_account_service/finance/memory" //引用memory包初始化函数
	//	"sub_account_service/finance/models"
	//	"sub_account_service/finance/protocol"
	//	"sub_account_service/finance/session"
	third "sub_account_service/finance/third_part_pay"
	//	"sub_account_service/finance/utils"
)

func RefundTrade(ctx *gin.Context) {
	tradeNo := ctx.Query("trade_no")
	if tradeNo == "" {
		ctx.JSON(http.StatusOK, lib.Result.Fail(-1, "trade number is empty"))
	}

	if err := third.CustomerRefundTrade(tradeNo); err != nil {
		glog.Errorln("customer refund trade error", err)
		ctx.JSON(http.StatusOK, lib.Result.Fail(-1, fmt.Sprintf("customer refund trade error: %v", err)))
	}

	ctx.JSON(http.StatusOK, lib.Result.Success("success", nil))
}
*/
