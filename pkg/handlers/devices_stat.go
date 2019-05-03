package handlers

import (
	"encoding/json"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/valyala/fasthttp"
)

func (h *Handler) DevicesStats(ctx *fasthttp.RequestCtx, token string) {
	if !h.storage.DeviceExists(token) {
		h.replyUnprocessableEntity(ctx)
		return
	}

	var collection []storage.DeviceStat
	if err := json.Unmarshal(ctx.Request.Body(), &collection); err != nil {
		h.replyUnprocessableEntity(ctx)
		return
	}

	h.storage.AddStat(token, collection)

	ctx.SetStatusCode(fasthttp.StatusCreated)
}
