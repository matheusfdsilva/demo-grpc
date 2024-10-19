package service

import (
	"context"
	pb "demo-grpc/services/common/genproto/orders"
)

var ordersDb = make([]*pb.Order, 0)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *pb.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*pb.Order {
	return ordersDb
}
