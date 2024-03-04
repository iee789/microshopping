package handler

import (
	"bytes"
	"context"
	"log"
	"math/rand"

	pb "recommendationservice/proto"
)

// 日志
var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

// 推荐服务结构体
type RecommendationService struct {
	ProductCatalogService pb.ProductCatalogServiceClient
}

// 列出推荐
func (s *RecommendationService) ListRecommendations(ctx context.Context, in *pb.ListRecommendationsRequest) (out *pb.ListRecommendationsResponse, e error) {
	maxResponsesCount := 5
	out = new(pb.ListRecommendationsResponse)
	
	// 查询商品类别
	catalog, err := s.ProductCatalogService.ListProducts(ctx, &pb.Empty{})
	if err != nil {
		return out, err
	}
	// 接收一个推荐请求，从产品目录中获取所有商品，然后根据请求中排除的商品ID列表，随机选择一定数量的商品作为推荐
	filteredProductsIDs := make([]string, 0, len(catalog.Products))
	for _, p := range catalog.Products {
		// 过滤掉请求中已存在的商品id，只保留不再请求列表的商品id
		if contains(p.Id, in.ProductIds) {
			continue
		}
		filteredProductsIDs = append(filteredProductsIDs, p.Id)
	}

	productIDs := sample(filteredProductsIDs, maxResponsesCount)
	logger.Printf("[Recv ListRecommendations] product_ids=%v", productIDs)
	
	out.ProductIds = productIDs
	return out, nil
}

// 判断是否包含
func contains(target string, source []string) bool {
	for _, s := range source {
		if target == s {
			return true
		}
	}
	return false
}

// 随机抽样，source切片中选择的c个元素是随机且不重复的
func sample(source []string, c int) []string {
	n := len(source)
	if n <= c {
		return source
	}
	
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}
	
	result := make([]string, 0, c)
	for i := 0; i < c; i++ {
		result = append(result, source[indices[i]])
	}
	
	return result
}
