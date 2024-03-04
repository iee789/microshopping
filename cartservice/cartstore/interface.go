package cartstore

import (
	"context"

	pb "cartservice/proto"
)

// 购物车接口
type CartStore interface {
	AddItem(ctx context.Context, userID, productID string, quantity int32, out *pb.Empty) (r *pb.Empty, err error)
	EmptyCart(ctx context.Context, userID string) (*pb.Empty, error)
	GetCart(ctx context.Context, userID string) (*pb.Cart, error)
}

// 购物车使用map来保存数据，包括：添加商品、删除商品、清空购物车

// 实例化CartStore
func NewMemoryCartStore() CartStore {
	return &memoryCartStore{
		carts: make(map[string]map[string]int32),
	}
}
