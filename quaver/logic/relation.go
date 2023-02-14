package logic

import (
	"quaver/dao/mysql"
	"quaver/models"
)

// RelationAction 关注操作
func RelationAction(currentUserID int64, p *models.ParamRelationAction) (err error) {
	follow := new(models.Follow)
	follow.UserID = currentUserID
	follow.FollowID = p.ToUserId
	follow.IsFollow = p.ActionType

	return mysql.RelationAction(follow)
}
