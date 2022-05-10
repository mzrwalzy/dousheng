package backend

import "bytedance-douyin/service"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:56
 */
type ApiGroup struct {
	UserApi
	CommentApi
	LikeApi
	FollowApi
	VideoApi
}

// service
var (
	userService = service.ServiceGroupApp.BackendServiceGroup.UserService
)