package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type UserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func newUserRequest(ctx *fasthttp.RequestCtx) (*UserRequest, error) {
	ur := new(UserRequest)
	if err := json.Unmarshal(ctx.Request.Body(), &ur); err != nil {
		return nil, err
	}
	return ur, nil
}

func (dr *UserRequest) Validate() error {
	if dr.Email == "" {
		return fmt.Errorf("Email incorrect or null")
	}

	return nil
}

func (h *Handler) UsersRegistration(ctx *fasthttp.RequestCtx) {
	ur, err := newUserRequest(ctx)
	if err != nil {
		h.replyStatusBadRequest(err, ctx)
		return
	}

	if err := ur.Validate(); err != nil {
		h.replyStatusBadRequest(err, ctx)
		return
	}

	token := h.storage.NewUser(ur.Email)
	reply, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		token,
	})

	if err != nil {
		h.replyInternalServerError(ctx)
		return
	}

	ctx.SetBody(reply)
}
