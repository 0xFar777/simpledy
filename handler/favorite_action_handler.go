package handler

import (
	"errors"
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandleFavoriteActionPost(token string, video_id int64, action_type string) (model.FavoriteActionResponse, error) {
	var statusCode = -1
	var statusMsg = ""
	var err error

	//将Token转化为UserId
	claims, err := utils.ParseToken(token)
	var userId int64
	if err != nil {
		log.Print(err)
	}
	userId = int64(claims["userId"].(float64))
	//检测Favorite表中是否存在此点赞消息
	_, exist := repository.IfFavoriteActionYes(userId, video_id, action_type)
	if exist > 0 {
		//如果不存在，就新建该点赞信息进到Favorite表
		newFavoriteInformation := model.Favorite{
			UserId:  userId,
			VideoId: video_id,
		}
		repository.CreateFavoriteInformation(newFavoriteInformation)
	} else {
		//如果存在，就从Favorite表中移除该数据
		repository.DeleteFavoriteInformation(userId, video_id)
	}
	change := repository.UpdateFavoriteCount(token, video_id, action_type)
	// 点赞状态未改变，行为操作失败
	if change != true {
		statusCode = -1
		if action_type == "1" {
			err = errors.New("点赞失败")
			statusMsg = err.Error()
		} else {
			err = errors.New("取消点赞失败")
			statusMsg = err.Error()
		}
	} else {
		//点赞状态改变，设置响应消息
		statusCode = 0
		if action_type == "1" {
			statusMsg = "点赞成功"
		} else {
			statusMsg = "取消点赞成功"
		}
	}
	//更新并返回响应
	resp := model.GenerateFavoriteActionResponse(int32(statusCode), statusMsg)
	return resp, err
}
