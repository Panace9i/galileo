package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type DeviceRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func newDeviceRequest(ctx *fasthttp.RequestCtx) (*DeviceRequest, error) {
	dr := new(DeviceRequest)
	if err := json.Unmarshal(ctx.Request.Body(), &dr); err != nil {
		return nil, err
	}
	return dr, nil
}

func (dr *DeviceRequest) Validate() error {
	if dr.Email == "" {
		return fmt.Errorf("Email incorrect or null")
	}
	if dr.Name == "" {
		return fmt.Errorf("Name incorrect or null")
	}
	return nil
}

func (h *Handler) DevicesRegistration(ctx *fasthttp.RequestCtx) {
	dr, err := newDeviceRequest(ctx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	if err := dr.Validate(); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	if !h.storage.UserExistsByEmail(dr.Email) {
		ctx.Error("User not found", fasthttp.StatusNotFound)
		return
	}

	token := h.storage.NewDevice(dr.Email, dr.Name)
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
