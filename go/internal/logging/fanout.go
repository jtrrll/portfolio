package logging

import (
	"context"
	"log/slog"
)

// FanoutHandler is a [slog.Handler] that delegates to multiple underlying handlers.
type FanoutHandler struct {
	handlers []slog.Handler
}

// NewFanoutHandler creates a [FanoutHandler] that writes to all provided handlers.
// Panics if fewer than two handlers are provided.
func NewFanoutHandler(handlers ...slog.Handler) *FanoutHandler {
	if len(handlers) < 2 {
		panic("FanoutHandler requires at least two handlers")
	}
	return &FanoutHandler{handlers: handlers}
}

func (h *FanoutHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *FanoutHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, record.Level) {
			if err := handler.Handle(ctx, record.Clone()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *FanoutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	cloned := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		cloned[i] = handler.WithAttrs(attrs)
	}
	return &FanoutHandler{handlers: cloned}
}

func (h *FanoutHandler) WithGroup(name string) slog.Handler {
	cloned := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		cloned[i] = handler.WithGroup(name)
	}
	return &FanoutHandler{handlers: cloned}
}
