package service

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/skycodes912/grpc-order-backend/proto"
)

type OrderServiceServer struct {
	proto.OrderServiceServer
	mu     sync.Mutex
	orders map[string]*proto.OrderResponse
}

func NewOrderServiceServer() *OrderServiceServer {
	return &OrderServiceServer{
		orders: make(map[string]*proto.OrderResponse),
	}
}

// handle order service create order
func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Println("Recieved order request: %+v\n", req)
	order := proto.OrderResponse{
		Id:       req.Id,
		Item:     req.Item,
		Quantity: req.Quantity,
		Price:    req.Price,
		Status:   "Order Created",
	}
	s.orders[order.Id] = &order
	return &order, nil
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, req *proto.OrderID) (*proto.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received GetOrder request: ID=%s\n", req.Id)
	order, exists := s.orders[req.Id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received UpdateOrder request: %+v\n", req)
	order, exists := s.orders[req.Id]
	if !exists {
		return nil, errors.New("order not found")
	}

	order.Item = req.Item
	order.Quantity = req.Quantity
	order.Price = req.Price
	order.Status = "Order Updated"

	s.orders[req.Id] = order
	return order, nil
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, req *proto.OrderID) (*proto.OrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received DeleteOrder request: %+v\n", req)
	order, exists := s.orders[req.Id]
	if !exists {
		return nil, errors.New("order not found")
	}

	delete(s.orders, req.Id)
	return order, nil
}
