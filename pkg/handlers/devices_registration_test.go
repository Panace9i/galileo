package handlers

import (
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestDevicesRegistration(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.DevicesRegistration(&ctx)

	if ctx.Response.StatusCode() != fasthttp.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusBadRequest, ctx.Response.StatusCode())
	}
}
