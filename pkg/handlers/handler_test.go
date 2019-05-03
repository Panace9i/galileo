package handlers

import (
	"fmt"
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestNew(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	if h.logger == nil {
		t.Error("Expected new logger, got nil")
	}
	if h.config == nil {
		t.Error("Expected new config, got nil")
	}
	if h.storage == nil {
		t.Error("Expected new storage, got nil")
	}
	if h.stats == nil {
		t.Error("Expected new stats, got nil")
	}
}

func TestNotFound(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.NotFound(&ctx)

	if ctx.Response.StatusCode() != fasthttp.StatusNotFound {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusNotFound, ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusNotFound) {
		t.Errorf("Expected response %s, got %s", fasthttp.StatusMessage(fasthttp.StatusNotFound), ctx.Response.String())
	}
}

func TestMethodNotAllowed(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.MethodNotAllowed(&ctx)

	if ctx.Response.StatusCode() != fasthttp.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusMethodNotAllowed, ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed) {
		t.Errorf("Expected response %s, got %s", fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed), ctx.Response.String())
	}
}

func TestPanicHandler(t *testing.T) {
	h := New(new(logrus.Logger), new(config.Config), new(storage.Storage))

	ctx := fasthttp.RequestCtx{}
	h.PanicHandler(&ctx, fmt.Errorf("test"))

	if ctx.Response.StatusCode() != fasthttp.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", fasthttp.StatusInternalServerError, ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusInternalServerError) {
		t.Errorf("Expected response %s, got %s", fasthttp.StatusMessage(fasthttp.StatusInternalServerError), ctx.Response.String())
	}
}
