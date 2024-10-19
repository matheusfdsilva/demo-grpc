package types

import (
	"context"
	pb "demo-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *pb.Order) error
}
