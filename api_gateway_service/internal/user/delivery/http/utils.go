package userHttp

import "github.com/opentracing/opentracing-go"

func (u *userHandlers) traceError(span opentracing.Span, err error) {
	span.SetTag("userHandlers", true)
	span.LogKV(err)
	u.metrics.ErrorHttpRequests.Inc()
}
