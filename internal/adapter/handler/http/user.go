package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vitconduck/fun/internal/core/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// getUserRequest represents the request body for getting a user
type getUserRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64			true	"User ID"
//	@Success		200	{object}	userResponse	"User displayed"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (u *UserHandler) GetUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
	}

	user, err := u.svc.GetUser(ctx, uint(req.ID))
	if err != nil {
		handleError(ctx, err)
		return
	}
	rsp := newUserResponse(user)

	handleSuccess(ctx, rsp)

}
