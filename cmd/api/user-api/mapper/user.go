package mapper

import (
	"ktrain/cmd/api/user-api/dto"
	"ktrain/cmd/model"
	"time"
)

func ToUserResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Username:  user.Username,
		Gender:    user.Gender,
		Birthday:  user.Birthday.Format("02/01/2006"),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func ToUserModel(user *dto.UserRequest) *model.User {
	birthday, _ := time.Parse("02/01/2006", user.Birthday)
	pReq := &model.User{
		Fullname:   user.Fullname,
		Username:   user.Username,
		Gender:     user.Gender,
		Birthday:   birthday,
		AuthTokens: []model.AuthToken{},
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	return pReq
}
func ToListUsersResponse(users []*model.User) []*dto.UserResponse {
	listUsersResponse := []*dto.UserResponse{}
	for _, user := range users {
		userResponse := &dto.UserResponse{
			ID:        user.ID,
			Fullname:  user.Fullname,
			Username:  user.Username,
			Gender:    user.Gender,
			Birthday:  user.Birthday.Format("02/01/2006"),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		listUsersResponse = append(listUsersResponse, userResponse)
	}
	return listUsersResponse
}
func ToActionResponse(actions []*dto.ActionRequest) *dto.ActionResponse {
	listAction := dto.ActionResponse{
		Action: make([]string, 0),
	}
	for _, action := range actions {
		listAction.Action = append(listAction.Action, action.Action)
	}
	return &listAction
}

func ToJWTTokenResponse(token string)*dto.JWTResponse{
	return &dto.JWTResponse{
		Token: token,
	}
}