package handler

import (
	"simpledy/model"
	"simpledy/repository"
)

func HandleFavoriteListGet(token string, user_id int64) model.FavoriteListResponse {
	var statusCode = -1
	var statusMsg = ""
	// claims, err := utils.ParseToken(token)
	// var userId int
	// if err != nil {
	// 	log.Print(err)
	// }
	// userId = int(claims["userId"].(float64))

	//获取用户喜欢的所有Video的Id
	videoIds, _ := repository.FindVideoIdByUserId(user_id)

	//根据视频Id获取到视频对应的信息
	videos, _ := repository.FindVideoInfoByVideoId(videoIds)

	//根据视频信息获取到对应的作者信息
	authors, _ := repository.FindAuthorInfoByVideoInfo(videos)

	statusCode = 0
	statusMsg = "喜欢的视频列表返回成功"
	resp := model.GenerateFavoriteListResponse(int64(statusCode), statusMsg, videos, authors)
	return resp
}

// func HandlePublishListGet(token string) model.PublishListResponse {
// 	var statusCode = -1
// 	var statusMsg = ""
// 	claims, err := utils.ParseToken(token)
// 	var userId int
// 	//fmt.Println("userid =", claims["userId"])
// 	//fmt.Println("useridInt =", claims["userId"].(int))
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	userId = int(claims["userId"].(float64))
// 	videos, _ := repository.FindVideosByUserId(userId)
// 	author := repository.FindUserInfoByUserId(userId)
// 	statusCode = 0
// 	statusMsg = "视频列表返回成功"
// 	resp := model.GeneratePulishListResponse(int64(statusCode), statusMsg, videos, author)
// 	return resp
// }
