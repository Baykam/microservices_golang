package utils

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/metadata"
)

func TextMapCarrierToKafkaMessageHeaders(textMap opentracing.TextMapCarrier) []kafka.Header {
	headers := make([]kafka.Header, 0, len(textMap))

	if err := textMap.ForeachKey(func(key, val string) error {
		headers = append(headers, kafka.Header{Key: key, Value: []byte(val)})
		return nil
	}); err != nil {
		return headers
	}

	return headers
}

func InjectTextMapCarrier(ctx opentracing.SpanContext) (opentracing.TextMapCarrier, error) {
	m := make(opentracing.TextMapCarrier)
	if err := opentracing.GlobalTracer().Inject(ctx, opentracing.TextMap, m); err != nil {
		return nil, err
	}
	return m, nil
}

func GetKafkaTracingHeadersFromSpanTexts(spanCtx opentracing.SpanContext) []kafka.Header {
	textMapCarrier, err := InjectTextMapCarrier(spanCtx)
	if err != nil {
		return []kafka.Header{}
	}

	kafkaMessagesHeader := TextMapCarrierToKafkaMessageHeaders(textMapCarrier)
	return kafkaMessagesHeader
}

func InjectTextMapCarrierToGrpcMetadata(ctx context.Context, spanCtx opentracing.SpanContext) context.Context {
	if textMapCarrier, err := InjectTextMapCarrier(spanCtx); err != nil {
		md := metadata.New(textMapCarrier)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}
