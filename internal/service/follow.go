package service

type FollowActionRequest struct {
	//UserId uint   `form:"user_id"  binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   uint   `form:"to_user_id" binding:"required"`
	ActionType int64  `form:"action_type" binding:"required"`
}

type FollowListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type FollowListResponse struct {
	ResponseCommon
	UserList []UserInfo `json:"user_list" binding:"required"`
}

// IsFollowRequest 判断A是否关注B
type IsFollowRequest struct {
	A uint
	B uint
}


func (svc *Service) FollowList(userId uint) (res FollowListResponse, err error) {
	follows, err := svc.dao.FollowList(userId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowedId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		userInfo.IsFollow = true
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}

func (svc *Service) FollowerList(userId uint) (res FollowListResponse, err error) {
	follows, err := svc.dao.FollowerList(userId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowerId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}
