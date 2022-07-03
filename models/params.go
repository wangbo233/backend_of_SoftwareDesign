/**
    @author:huchao
    @data:2022/2/15
    @note:
**/
package models

/**
 * @Author huchao
 * @Description //TODO magic string
 * @Date 21:56 2022/2/15
 **/
const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamPostList struct {
	CommunityID uint64 `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页数量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
