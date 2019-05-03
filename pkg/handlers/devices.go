package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (h *Handler) Devices(ctx *fasthttp.RequestCtx, token string) {
	if !h.storage.UserExists(token) {
		h.replyNotFound("User not found", ctx)
		return
	}

	reply, err := json.Marshal(h.storage.GetDeviceStat(token))
	if err != nil {
		h.replyInternalServerError(ctx)
		return
	}

	ctx.SetBody(reply)
}
