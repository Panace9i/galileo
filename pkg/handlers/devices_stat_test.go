package handlers

import (
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestDevicesStats(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.DevicesStats(&ctx, "123")

	if ctx.Response.StatusCode() != fasthttp.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusUnprocessableEntity, ctx.Response.StatusCode())
	}
}
