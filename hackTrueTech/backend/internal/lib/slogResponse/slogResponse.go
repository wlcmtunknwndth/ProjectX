package slogResponse

import "log/slog"

func SlogErr(err error) slog.Attr {
	return slog.Attr{Key: "error", Value: slog.StringValue(err.Error())}
}

func SlogOp(op string) slog.Attr {
	return slog.Attr{Key: "op", Value: slog.StringValue(op)}
}
