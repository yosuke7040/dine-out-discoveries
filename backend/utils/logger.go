package utils

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"log"

// 	"golang.org/x/exp/slog"
// )

// type SimpleHandler struct {
// 	slog.Handler
// 	logger *log.Logger
// }

// func NewSimpleHander(out io.Writer, level slog.Level) *SimpleHandler {
// 	prefix := ""
// 	h := &SimpleHandler{
// 		Handler: slog.NewJSONHandler(out, &slog.HandlerOptions{
// 			Level: level,
// 		}),
// 		logger: log.New(out, prefix, 0),
// 	}
// 	return h
// }

// func (h *SimpleHandler) Handle(_ context.Context, record slog.Record) error {
// 	ts := record.Time.Format("[2006-01-02 15:04:05.000]")
// 	level := fmt.Sprintf("[%5s]", record.Level.String())
// 	h.logger.Println(ts, level, record.Message)
// 	return nil
// }
