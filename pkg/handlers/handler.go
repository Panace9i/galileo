package handlers

import (
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/logger"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/valyala/fasthttp"
	"time"
)

type Handler struct {
	logger  logger.Logger
	config  *config.Config
	stats   *stats
	storage *storage.Storage
}

type stats struct {
	requests        *Requests
	averageDuration time.Duration
	maxDuration     time.Duration
	totalDuration   time.Duration
	requestsCount   time.Duration
	startTime       time.Time
}

func New(logger logger.Logger, config *config.Config, storage *storage.Storage) *Handler {
	return &Handler{
		logger: logger,
		config: config,
		stats: &stats{
			requests:  new(Requests),
			startTime: time.Now(),
		},
		storage: storage,
	}
}

func (h *Handler) Wrap(handle func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		timer := time.Now()

		if len(ctx.Request.Body()) == 0 {
			ctx.Request.SetBody([]byte(ctx.URI().String()))
		}

		handle(ctx)
		h.countDuration(timer)
		h.collectCodes(ctx)
	}
}

func (h *Handler) WrapAuth(handle func(ctx *fasthttp.RequestCtx, token string)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		timer := time.Now()
		auth := ctx.Request.Header.Peek("Authorization")
		token := string(auth)

		if len(ctx.Request.Body()) == 0 {
			ctx.Request.SetBody([]byte(ctx.URI().String()))
		}

		if token == "" {
			h.replyUnauthorized(ctx)
			return
		}

		handle(ctx, token)
		h.countDuration(timer)
		h.collectCodes(ctx)
	}
}

func (h *Handler) NotFound(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusNotFound))
}

func (h *Handler) MethodNotAllowed(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed))
}

func (h *Handler) PanicHandler(ctx *fasthttp.RequestCtx, i interface{}) {
	timer := time.Now()

	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusInternalServerError))

	h.logger.Error(i)

	h.countDuration(timer)
	h.collectCodes(ctx)
}

func (h *Handler) countDuration(timer time.Time) {
	if !timer.IsZero() {
		h.stats.requestsCount++
		took := time.Now()
		duration := took.Sub(timer)
		h.stats.totalDuration += duration
		if duration > h.stats.maxDuration {
			h.stats.maxDuration = duration
		}
		h.stats.averageDuration = h.stats.totalDuration / h.stats.requestsCount
		h.stats.requests.Duration.Max = h.stats.maxDuration.String()
		h.stats.requests.Duration.Average = h.stats.averageDuration.String()
	}
}

func (h *Handler) collectCodes(ctx *fasthttp.RequestCtx) {
	sc := ctx.Response.Header.StatusCode()
	if sc >= 500 {
		h.stats.requests.Codes.C5xx++
	} else if sc >= 400 {
		h.stats.requests.Codes.C4xx++
	} else if sc >= 200 && sc < 300 {
		h.stats.requests.Codes.C2xx++
	}
}

func (h *Handler) replyStatusBadRequest(err error, ctx *fasthttp.RequestCtx) {
	h.reply(err.Error(), fasthttp.StatusBadRequest, ctx)
}

func (h *Handler) replyUnprocessableEntity(ctx *fasthttp.RequestCtx) {
	h.reply(fasthttp.StatusMessage(fasthttp.StatusUnprocessableEntity), fasthttp.StatusUnprocessableEntity, ctx)
}

func (h *Handler) replyNotFound(msg string, ctx *fasthttp.RequestCtx) {
	h.reply(msg, fasthttp.StatusNotFound, ctx)
}

func (h *Handler) replyUnauthorized(ctx *fasthttp.RequestCtx) {
	h.reply(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized, ctx)
}

func (h *Handler) replyInternalServerError(ctx *fasthttp.RequestCtx) {
	h.reply(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError, ctx)
}

func (h *Handler) reply(msg string, statusCode int, ctx *fasthttp.RequestCtx) {
	ctx.Error(msg, statusCode)

	h.logger.Errorf("%s: %s", ctx.Request.Body(), ctx.Response.Body())
}
