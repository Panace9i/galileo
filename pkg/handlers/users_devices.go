package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (h *Handler) UsersDevices(ctx *fasthttp.RequestCtx, token string) {
	if !h.storage.UserExists(token) {
		h.replyNotFound("Device not found", ctx)
		return
	}

	reply, err := json.Marshal(h.storage.GetUserDevices(token))
	if err != nil {
		h.replyInternalServerError(ctx)
		return
	}

	ctx.SetBody(reply)
}
