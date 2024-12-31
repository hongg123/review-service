// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"review-service/internal/data/model"
)

func newReviewReplyInfo(db *gorm.DB, opts ...gen.DOOption) reviewReplyInfo {
	_reviewReplyInfo := reviewReplyInfo{}

	_reviewReplyInfo.reviewReplyInfoDo.UseDB(db, opts...)
	_reviewReplyInfo.reviewReplyInfoDo.UseModel(&model.ReviewReplyInfo{})

	tableName := _reviewReplyInfo.reviewReplyInfoDo.TableName()
	_reviewReplyInfo.ALL = field.NewAsterisk(tableName)
	_reviewReplyInfo.ID = field.NewInt64(tableName, "id")
	_reviewReplyInfo.CreateBy = field.NewString(tableName, "create_by")
	_reviewReplyInfo.UpdateBy = field.NewString(tableName, "update_by")
	_reviewReplyInfo.CreateAt = field.NewTime(tableName, "create_at")
	_reviewReplyInfo.UpdateAt = field.NewTime(tableName, "update_at")
	_reviewReplyInfo.DeleteAt = field.NewTime(tableName, "delete_at")
	_reviewReplyInfo.Version = field.NewInt32(tableName, "version")
	_reviewReplyInfo.ReplyID = field.NewInt64(tableName, "reply_id")
	_reviewReplyInfo.ReviewID = field.NewInt64(tableName, "review_id")
	_reviewReplyInfo.StoreID = field.NewInt64(tableName, "store_id")
	_reviewReplyInfo.Content = field.NewString(tableName, "content")
	_reviewReplyInfo.PicInfo = field.NewString(tableName, "pic_info")
	_reviewReplyInfo.VideoInfo = field.NewString(tableName, "video_info")
	_reviewReplyInfo.ExtJSON = field.NewString(tableName, "ext_json")
	_reviewReplyInfo.CtrlJSON = field.NewString(tableName, "ctrl_json")

	_reviewReplyInfo.fillFieldMap()

	return _reviewReplyInfo
}

// reviewReplyInfo 评价商家回复表
type reviewReplyInfo struct {
	reviewReplyInfoDo reviewReplyInfoDo

	ALL       field.Asterisk
	ID        field.Int64  // 主键
	CreateBy  field.String // 创建?标识
	UpdateBy  field.String // 更新?标识
	CreateAt  field.Time   // 创建时间
	UpdateAt  field.Time   // 更新时间
	DeleteAt  field.Time   // 逻辑删除标记
	Version   field.Int32  // 乐观锁标记
	ReplyID   field.Int64  // 回复id
	ReviewID  field.Int64  // 评价id
	StoreID   field.Int64  // 店铺id
	Content   field.String // 评价内容
	PicInfo   field.String // 媒体信息：图?
	VideoInfo field.String // 媒体信息：视频
	ExtJSON   field.String // 信息扩展
	CtrlJSON  field.String // 控制扩展

	fieldMap map[string]field.Expr
}

func (r reviewReplyInfo) Table(newTableName string) *reviewReplyInfo {
	r.reviewReplyInfoDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reviewReplyInfo) As(alias string) *reviewReplyInfo {
	r.reviewReplyInfoDo.DO = *(r.reviewReplyInfoDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reviewReplyInfo) updateTableName(table string) *reviewReplyInfo {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreateBy = field.NewString(table, "create_by")
	r.UpdateBy = field.NewString(table, "update_by")
	r.CreateAt = field.NewTime(table, "create_at")
	r.UpdateAt = field.NewTime(table, "update_at")
	r.DeleteAt = field.NewTime(table, "delete_at")
	r.Version = field.NewInt32(table, "version")
	r.ReplyID = field.NewInt64(table, "reply_id")
	r.ReviewID = field.NewInt64(table, "review_id")
	r.StoreID = field.NewInt64(table, "store_id")
	r.Content = field.NewString(table, "content")
	r.PicInfo = field.NewString(table, "pic_info")
	r.VideoInfo = field.NewString(table, "video_info")
	r.ExtJSON = field.NewString(table, "ext_json")
	r.CtrlJSON = field.NewString(table, "ctrl_json")

	r.fillFieldMap()

	return r
}

func (r *reviewReplyInfo) WithContext(ctx context.Context) IReviewReplyInfoDo {
	return r.reviewReplyInfoDo.WithContext(ctx)
}

func (r reviewReplyInfo) TableName() string { return r.reviewReplyInfoDo.TableName() }

func (r reviewReplyInfo) Alias() string { return r.reviewReplyInfoDo.Alias() }

func (r reviewReplyInfo) Columns(cols ...field.Expr) gen.Columns {
	return r.reviewReplyInfoDo.Columns(cols...)
}

func (r *reviewReplyInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reviewReplyInfo) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 15)
	r.fieldMap["id"] = r.ID
	r.fieldMap["create_by"] = r.CreateBy
	r.fieldMap["update_by"] = r.UpdateBy
	r.fieldMap["create_at"] = r.CreateAt
	r.fieldMap["update_at"] = r.UpdateAt
	r.fieldMap["delete_at"] = r.DeleteAt
	r.fieldMap["version"] = r.Version
	r.fieldMap["reply_id"] = r.ReplyID
	r.fieldMap["review_id"] = r.ReviewID
	r.fieldMap["store_id"] = r.StoreID
	r.fieldMap["content"] = r.Content
	r.fieldMap["pic_info"] = r.PicInfo
	r.fieldMap["video_info"] = r.VideoInfo
	r.fieldMap["ext_json"] = r.ExtJSON
	r.fieldMap["ctrl_json"] = r.CtrlJSON
}

func (r reviewReplyInfo) clone(db *gorm.DB) reviewReplyInfo {
	r.reviewReplyInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reviewReplyInfo) replaceDB(db *gorm.DB) reviewReplyInfo {
	r.reviewReplyInfoDo.ReplaceDB(db)
	return r
}

type reviewReplyInfoDo struct{ gen.DO }

type IReviewReplyInfoDo interface {
	gen.SubQuery
	Debug() IReviewReplyInfoDo
	WithContext(ctx context.Context) IReviewReplyInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReviewReplyInfoDo
	WriteDB() IReviewReplyInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReviewReplyInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReviewReplyInfoDo
	Not(conds ...gen.Condition) IReviewReplyInfoDo
	Or(conds ...gen.Condition) IReviewReplyInfoDo
	Select(conds ...field.Expr) IReviewReplyInfoDo
	Where(conds ...gen.Condition) IReviewReplyInfoDo
	Order(conds ...field.Expr) IReviewReplyInfoDo
	Distinct(cols ...field.Expr) IReviewReplyInfoDo
	Omit(cols ...field.Expr) IReviewReplyInfoDo
	Join(table schema.Tabler, on ...field.Expr) IReviewReplyInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReviewReplyInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReviewReplyInfoDo
	Group(cols ...field.Expr) IReviewReplyInfoDo
	Having(conds ...gen.Condition) IReviewReplyInfoDo
	Limit(limit int) IReviewReplyInfoDo
	Offset(offset int) IReviewReplyInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReviewReplyInfoDo
	Unscoped() IReviewReplyInfoDo
	Create(values ...*model.ReviewReplyInfo) error
	CreateInBatches(values []*model.ReviewReplyInfo, batchSize int) error
	Save(values ...*model.ReviewReplyInfo) error
	First() (*model.ReviewReplyInfo, error)
	Take() (*model.ReviewReplyInfo, error)
	Last() (*model.ReviewReplyInfo, error)
	Find() ([]*model.ReviewReplyInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReviewReplyInfo, err error)
	FindInBatches(result *[]*model.ReviewReplyInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ReviewReplyInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReviewReplyInfoDo
	Assign(attrs ...field.AssignExpr) IReviewReplyInfoDo
	Joins(fields ...field.RelationField) IReviewReplyInfoDo
	Preload(fields ...field.RelationField) IReviewReplyInfoDo
	FirstOrInit() (*model.ReviewReplyInfo, error)
	FirstOrCreate() (*model.ReviewReplyInfo, error)
	FindByPage(offset int, limit int) (result []*model.ReviewReplyInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReviewReplyInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reviewReplyInfoDo) Debug() IReviewReplyInfoDo {
	return r.withDO(r.DO.Debug())
}

func (r reviewReplyInfoDo) WithContext(ctx context.Context) IReviewReplyInfoDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reviewReplyInfoDo) ReadDB() IReviewReplyInfoDo {
	return r.Clauses(dbresolver.Read)
}

func (r reviewReplyInfoDo) WriteDB() IReviewReplyInfoDo {
	return r.Clauses(dbresolver.Write)
}

func (r reviewReplyInfoDo) Session(config *gorm.Session) IReviewReplyInfoDo {
	return r.withDO(r.DO.Session(config))
}

func (r reviewReplyInfoDo) Clauses(conds ...clause.Expression) IReviewReplyInfoDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reviewReplyInfoDo) Returning(value interface{}, columns ...string) IReviewReplyInfoDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reviewReplyInfoDo) Not(conds ...gen.Condition) IReviewReplyInfoDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reviewReplyInfoDo) Or(conds ...gen.Condition) IReviewReplyInfoDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reviewReplyInfoDo) Select(conds ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reviewReplyInfoDo) Where(conds ...gen.Condition) IReviewReplyInfoDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reviewReplyInfoDo) Order(conds ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reviewReplyInfoDo) Distinct(cols ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reviewReplyInfoDo) Omit(cols ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reviewReplyInfoDo) Join(table schema.Tabler, on ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reviewReplyInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reviewReplyInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reviewReplyInfoDo) Group(cols ...field.Expr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reviewReplyInfoDo) Having(conds ...gen.Condition) IReviewReplyInfoDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reviewReplyInfoDo) Limit(limit int) IReviewReplyInfoDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reviewReplyInfoDo) Offset(offset int) IReviewReplyInfoDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reviewReplyInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReviewReplyInfoDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reviewReplyInfoDo) Unscoped() IReviewReplyInfoDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reviewReplyInfoDo) Create(values ...*model.ReviewReplyInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reviewReplyInfoDo) CreateInBatches(values []*model.ReviewReplyInfo, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reviewReplyInfoDo) Save(values ...*model.ReviewReplyInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reviewReplyInfoDo) First() (*model.ReviewReplyInfo, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewReplyInfo), nil
	}
}

func (r reviewReplyInfoDo) Take() (*model.ReviewReplyInfo, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewReplyInfo), nil
	}
}

func (r reviewReplyInfoDo) Last() (*model.ReviewReplyInfo, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewReplyInfo), nil
	}
}

func (r reviewReplyInfoDo) Find() ([]*model.ReviewReplyInfo, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReviewReplyInfo), err
}

func (r reviewReplyInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReviewReplyInfo, err error) {
	buf := make([]*model.ReviewReplyInfo, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reviewReplyInfoDo) FindInBatches(result *[]*model.ReviewReplyInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reviewReplyInfoDo) Attrs(attrs ...field.AssignExpr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reviewReplyInfoDo) Assign(attrs ...field.AssignExpr) IReviewReplyInfoDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reviewReplyInfoDo) Joins(fields ...field.RelationField) IReviewReplyInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reviewReplyInfoDo) Preload(fields ...field.RelationField) IReviewReplyInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reviewReplyInfoDo) FirstOrInit() (*model.ReviewReplyInfo, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewReplyInfo), nil
	}
}

func (r reviewReplyInfoDo) FirstOrCreate() (*model.ReviewReplyInfo, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReviewReplyInfo), nil
	}
}

func (r reviewReplyInfoDo) FindByPage(offset int, limit int) (result []*model.ReviewReplyInfo, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reviewReplyInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reviewReplyInfoDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reviewReplyInfoDo) Delete(models ...*model.ReviewReplyInfo) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reviewReplyInfoDo) withDO(do gen.Dao) *reviewReplyInfoDo {
	r.DO = *do.(*gen.DO)
	return r
}
