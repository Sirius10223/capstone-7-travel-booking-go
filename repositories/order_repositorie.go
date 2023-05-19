package repositories

import ()

type OrderRepository interface {

}

type orderRepository struct {

}

func NewOrderRepository () orderRepository {
	return orderRepository{}
}