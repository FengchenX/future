//author xinbing
//time 2018/9/11 17:00
//
package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ibs_service/saas_background/dbs"
	"ibs_service/saas_background/utils"
)
type companyDTO struct {
	ID	int
	Name string
	BranchNo string
	BranchName string
}

func GetCompaniesByUserId(userId uint) *utils.Resp {
	rows, err := dbs.SaasGormDB.Client.Table("saas_user_company_rels").Select("company_id").
		Where("user_id = ?", userId).Rows()
	if err != nil {
		logrus.WithError(err).Errorln("GetCompaniesByUserId error!")
		return utils.Resp{}.Failed("查找公司失败！")
	}
	var companyIds []int
	for rows.Next() {
		id := 0
		fmt.Println(rows.Scan(&id))
		companyIds = append(companyIds , id)
	}
	var companies []*companyDTO
	err = dbs.OrderServerGormDB.Client.Table("companies").Where("id IN (?)",companyIds).Find(&companies).Error
	if err != nil {
		logrus.WithError(err).Errorln("GetCompaniesByUserId error!")
		return utils.Resp{}.Failed("查找公司失败！")
	}
	return utils.Resp{}.Success("查找成功！", companies)
}
