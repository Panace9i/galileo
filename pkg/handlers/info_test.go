package handlers

import (
	"encoding/json"
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.Info(&ctx)

	if ctx.Response.StatusCode() != fasthttp.StatusOK {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusOK, ctx.Response.StatusCode())
	}

	host, _ := os.Hostname()
	status := Status{}
	json.Unmarshal(ctx.Response.Body(), &status)

	if status.Host != host {
		t.Errorf("Expected host %s, got %s", host, status.Host)
	}
}
