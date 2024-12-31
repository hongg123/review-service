package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"
)

type ReviewRepo interface {
	SaveReview(ctx context.Context, info *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error)
	SaveReply(ctx context.Context, replyInfo *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
}

type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ReviewUsecase) SaveReview(ctx context.Context, reviewInfo *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] CreateReview, req: %v", reviewInfo)
	// 1. 数据校验
	// 1.1 参数基础校验：正常来说不应该在这一层，在service层或者框架层处理（validate参数校验）

	// 1.2 参数业务校验：带业务逻辑的参数校验，比如已经评价过的订单不能再创建评价
	reviews, err := uc.repo.GetReviewByOrderID(ctx, reviewInfo.OrderID)
	if err != nil {
		return nil, v1.ErrorDbFailed("查询数据库失败")
	}
	if len(reviews) > 0 {
		return nil, v1.ErrorOrderReviewed("订单：%d已经评价过", reviewInfo.OrderID)
	}
	// 2. 生成review ID(雪花算法，很多公司是有公共的生成分布式ID的服务，直接接入即可)
	reviewInfo.ReviewID = snowflake.GenID()
	// 3. 查询订单和商品快照信息
	/* 实际业务场景下需要查询订单服务和商品信息服务（RPC调用） */
	// 4. 拼装数据入库
	return uc.repo.SaveReview(ctx, reviewInfo)
}

// CreateReply 创建评价回复
func (uc *ReviewUsecase) CreateReply(ctx context.Context, param *ReplyParam) (*model.ReviewReplyInfo, error) {
	// 调用data层创建一个评价的回复
	uc.log.WithContext(ctx).Debugf("[biz] CreateReply param:%v", param)
	reply := &model.ReviewReplyInfo{
		ReplyID:   snowflake.GenID(),
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	}
	return uc.repo.SaveReply(ctx, reply)
}
