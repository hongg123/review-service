package data

import (
	"context"
	"errors"
	"review-service/internal/data/model"
	"review-service/internal/data/query"

	"review-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, reviewInfo *model.ReviewInfo) (*model.ReviewInfo, error) {
	err := r.data.query.ReviewInfo.
		WithContext(ctx).
		Save(reviewInfo)
	return reviewInfo, err
}

// 根据订单ID查询评价
func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error) {
	review, err := r.data.query.ReviewInfo.
		WithContext(ctx).
		Where(r.data.query.ReviewInfo.OrderID.Eq(orderID)).
		Find()
	return review, err
}

func (r *reviewRepo) SaveReply(ctx context.Context, replyInfo *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	/* 1.数据校验 */
	// 1.1 数据业务合法性校验（已回复的评价不允许商家再次回复）
	review, err := r.data.query.ReviewInfo. // 现根据回复的reviewID查询评价表里HasReply字段
						WithContext(ctx).
						Where(r.data.query.ReviewInfo.ReviewID.Eq(replyInfo.ReviewID)).
						First()
	if err != nil {
		return nil, err
	}
	if review.HasReply == 1 {
		return nil, errors.New("该评价已回复")
	}
	// 1.2 水平越权校验（A商家不能给B商家的评价）
	if review.StoreID != replyInfo.StoreID {
		return nil, errors.New("水平越权")
	}
	/* 2.更新数据（评价回复表和评价表同时更新 ---> 涉及事务操作） */
	// 事务操作
	err = r.data.query.Transaction(func(tx *query.Query) error {
		// 回复表插入一条数据
		if err := tx.ReviewReplyInfo.
			WithContext(ctx).
			Save(replyInfo); err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply create reply fail, err:%v\n", err)
			return err
		}
		// 评价表更新hasReply字段
		if _, err := tx.ReviewInfo.
			WithContext(ctx).
			Where(tx.ReviewInfo.ReviewID.Eq(replyInfo.ReviewID)).
			Update(tx.ReviewInfo.HasReply, 1); err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply update review fail, err:%v", err)
			return err
		}
		return nil
	})
	/* 3.返回结果 */
	return replyInfo, err
}

func (r *reviewRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *reviewRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *reviewRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *reviewRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
