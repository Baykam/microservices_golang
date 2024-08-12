package commands

import (
	"context"
	"project-microservices/dto"
	"project-microservices/mappers"
	"project-microservices/pkg/tracing"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (c *commandsHandler) UpdateUser(ctx context.Context, req *dto.UserUpdateReq) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "apiGateWayService.Commands.updateUser")
	defer span.Finish()

	updateUser := mappers.UserUpdateToGrpc(req)

	dtoBytes, err := proto.Marshal(updateUser)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserUpdated.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
