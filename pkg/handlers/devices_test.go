package handlers

import (
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestDevices(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.Devices(&ctx, "123")

	if ctx.Response.StatusCode() != fasthttp.StatusNotFound {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusNotFound, ctx.Response.StatusCode())
	}
}
