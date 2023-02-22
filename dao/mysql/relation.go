package mysql

import (
	"quaver/models"
)

func Attention(p *models.Attention) (err error) {
	//可以增加事务，保证sql数据库处理的安全性
	db.Model(&models.Attention{}).Where("token = ? ", p.Token).Updates(models.Attention{ToUserId: p.ToUserId, ActionType: p.ActionType})
	return nil
}

func QuaryList(token string) (commen []int64) {
	//通过token查询关注的人
	attention := models.Attention{}
	db.Where("token=?", token).Find(&attention)
	//fmt.Println(attention.Token, attention.ToUserId)
	commen = append(commen, attention.ToUserId)
	return commen
}

func UserExist(id int64) (token string) {
	CommentList := models.CommentList{}
	db.Where("user_id=?", id).Find(&CommentList)
	if CommentList.UserId == id { //找到对应id
		return CommentList.Token
	}
	return
}

//查询关注用户id
func SearchToUserId(token string) (userId int64) {
	attention := models.Attention{}
	db.Where("token=?", token).Find(&attention)
	return attention.ToUserId
}
