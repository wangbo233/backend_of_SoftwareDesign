/**
    @author:Huchao
    @data:2022/2/12
    @note:
**/
package logic

import (
	"backend/dao/mysql"
	"backend/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetailByID(id uint64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)
}
