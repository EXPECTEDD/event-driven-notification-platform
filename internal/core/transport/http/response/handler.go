package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	"go.uber.org/zap"
)

var (
	contentTypeStr = "Content-Type"
	jsonStr        = "application/json"
)

type HTTPResponseHandler struct {
	rw  http.ResponseWriter
	log *core_logger.Logger
}

func NewHTTPResponseHandler(
	rw http.ResponseWriter,
	log *core_logger.Logger,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		rw:  rw,
		log: log,
	}
}

func (h *HTTPResponseHandler) SendResponse(resp any, statusCode int) {
	h.rw.Header().Set(contentTypeStr, jsonStr)

	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(resp); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErrorResponse(msg string, err error) {
	var statusCode int
	logFunc := func(string, ...zap.Field) {}

	switch {
	case errors.Is(err, core_error.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_error.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Info
	case errors.Is(err, core_error.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Info
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(msg, err, statusCode)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(msg, err, statusCode)
}

func (h *HTTPResponseHandler) errorResponse(
	msg string,
	err error,
	statusCode int,
) {
	h.rw.Header().Set(contentTypeStr, jsonStr)

	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.rw.WriteHeader(http.StatusNoContent)
}
