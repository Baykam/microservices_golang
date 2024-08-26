package userKafkaConn

import (
	"context"
	userKafkaProto "project-microservices/kafka_protos/user"
	"project-microservices/pkg/tracing"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (u *userMessageProcessor) updateUser(ctx context.Context, reader *kafka.Reader, message kafka.Message) {
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, message.Headers, "userService.Kafka.userUpdated")
	defer span.Finish()

	msg := &userKafkaProto.PostUser{}
	if err := proto.Unmarshal(message.Value, msg); err != nil {
		u.log.WarnMsg("proto.Unmarshal", err)
		return
	}

	err := u.us.Commands.UpdateUser(ctx)
	if err != nil {
		return
	}

	u.commitMessage(ctx, reader, message)
}
