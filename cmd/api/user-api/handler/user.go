package handler

import (
	"ktrain/cmd/api/user-api/dto"
	"ktrain/cmd/api/user-api/mapper"
	"ktrain/cmd/model"
	"ktrain/cmd/repository"
	"ktrain/pkg/errors"
	"ktrain/pkg/httputil"
	"ktrain/rambbitmq"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type userHandler struct {
	userRepository        repository.IUserRepository
	activityLogRepository repository.ActivityLogRepository
	rabbitmq              *rambbitmq.RabbitMqManager
	*validator.Validate
}

func NewUserHandler(rabbitmq *rambbitmq.RabbitMqManager, userRepository repository.IUserRepository, activityLogRepository repository.ActivityLogRepository) *userHandler {
	return &userHandler{
		userRepository:        userRepository,
		activityLogRepository: activityLogRepository,
		rabbitmq:              rabbitmq,
		Validate:              validator.New(),
	}
}
func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	req := dto.UserRequest{}
	var binder httputil.JsonBinder
	if err := binder.BindRequest(&req, r); err != nil {
		if err.Error() == "Error reading body request" {
			httputil.RespondError(w, http.StatusInternalServerError, "Error reading body request")
			return
		} else {
			httputil.RespondError(w, http.StatusInternalServerError, "Error unmarshal body request")
		}
		return
	}
	err := h.Validate.Struct(req)
	if err != nil {
		httputil.RespondError(w, http.StatusBadRequest, "Error when validate request")
		return
	}
	ctx := r.Context()
	body := dto.UserActivityLogMessage{
		ID:  ctx.Value("userID").(int64),
		Log: "Update user",
	}
	err = h.rabbitmq.Publish(body)
	if err != nil {
		httputil.FailOnError(err, err.Error())
	}
	_, err = h.userRepository.GetUserByID(int64(id))
	if err != nil {
		if errors.IsDataNotFound(err) {
			httputil.RespondError(w, http.StatusNotFound, "User not found in database")
			return
		}
		httputil.RespondError(w, http.StatusInternalServerError, "Error when getting user ")
		return
	}
	user := mapper.ToUserModel(&req)
	user.ID = int64(id)
	resp, err := h.userRepository.UpdateUser(user)
	if err != nil {
		httputil.RespondError(w, http.StatusInternalServerError, "Error when update user")
		return
	}
	httputil.RespondSuccessWithData(w, http.StatusOK, mapper.ToUserResponse(resp))
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := dto.UserActivityLogMessage{
		ID:  ctx.Value("userID").(int64),
		Log: "Update user",
	}
	err := h.rabbitmq.Publish(body)
	if err != nil {
		httputil.FailOnError(err, err.Error())
	}
	ID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err = h.userRepository.DeleteUser(int64(ID))
	if err != nil {
		httputil.RespondError(w, http.StatusInternalServerError, "Error when delete user")
		return
	}
}

func (h *userHandler) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := dto.UserActivityLogMessage{
		ID:  ctx.Value("userID").(int64),
		Log: "Update user",
	}
	err := h.rabbitmq.Publish(body)
	if err != nil {
		httputil.FailOnError(err, err.Error())
	}
	user, err := h.userRepository.GetUserByID(ctx.Value("userID").(int64))
	if err != nil {
		if errors.IsDataNotFound(err) {
			httputil.RespondError(w, http.StatusNotFound, "Your profile not found")
			return
		}
		httputil.RespondError(w, http.StatusInternalServerError, "Error when getting user profile")
		return
	}
	httputil.RespondSuccessWithData(w, http.StatusOK, mapper.ToUserResponse(user))
}

func (h *userHandler) GetListUsers(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var ids []int64
	if values["ids"] != nil {
		req := dto.UserQuery{}
		var binder httputil.QueryURLBinder
		if err := binder.BindRequest(&req, r); err != nil {
			httputil.RespondError(w, http.StatusInternalServerError, "Error when query users list")
			return
		}
		for _, v := range req.Ids {
			id, _ := strconv.Atoi(v)
			ids = append(ids, int64(id))
		}
	}
	ctx := r.Context()
	body := dto.UserActivityLogMessage{
		ID:  ctx.Value("userID").(int64),
		Log: "Update user",
	}
	err := h.rabbitmq.Publish(body)
	if err != nil {
		httputil.FailOnError(err, err.Error())
	}
	users, err := h.userRepository.GetListUser(ids)
	if err != nil {
		httputil.RespondError(w, http.StatusInternalServerError, "Error when getting users list")
		return
	}
	httputil.RespondSuccessWithData(w, http.StatusOK, mapper.ToListUsersResponse(users))
}

func (h *userHandler) GetInformationUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		httputil.RespondError(w, http.StatusBadRequest, "Request invalid")
		return
	}
	ctx := r.Context()
	body := dto.UserActivityLogMessage{
		ID:  ctx.Value("userID").(int64),
		Log: "Update user",
	}
	err = h.rabbitmq.Publish(body)
	if err != nil {
		httputil.FailOnError(err, err.Error())
	}

	user, err := h.userRepository.GetUserByID(int64(userID))
	if err != nil {
		if errors.IsDataNotFound(err) {
			httputil.RespondError(w, http.StatusNotFound, "User not found")
			return
		}
		httputil.RespondError(w, http.StatusInternalServerError, "Error when getting user profile")
		return
	}

	httputil.RespondSuccessWithData(w, http.StatusOK, mapper.ToUserResponse(user))
}

func (h *userHandler) PostNewUser(w http.ResponseWriter, r *http.Request) {
	u := dto.CreateUserRequest{}
	var binder httputil.JsonBinder
	if err := binder.BindRequest(&u, r); err != nil {
		if err.Error() == "Error reading body request" {
			httputil.RespondError(w, http.StatusInternalServerError, "Error reading body request")
			return
		} else {
			httputil.RespondError(w, http.StatusInternalServerError, "Error unmarshal body request")
		}
		return
	}
	err := h.Validate.Struct(u)
	if err != nil {
		httputil.RespondError(w, http.StatusBadRequest, "Error when validate request")
		return
	}
	birthday, _ := time.Parse("2006-01-02", u.Birthday)
	User := &model.User{
		Fullname: u.Fullname,
		Username: u.Username,
		Gender:   u.Gender,
		Birthday: birthday,
	}
	ctx := r.Context()
	body := dto.UserActivityLogMessage{
		ID:  ctx.Value("userID").(int64),
		Log: "Update user",
	}
	err = h.rabbitmq.Publish(body)
	if err != nil {
		httputil.FailOnError(err, err.Error())
	}

	newUser, err := h.userRepository.CreateUser(User)
	if err != nil {
		httputil.RespondError(w, http.StatusInternalServerError, "Error when creating new user")
		return
	}
	httputil.RespondSuccessWithData(w, http.StatusOK, mapper.ToUserResponse(newUser))
}
