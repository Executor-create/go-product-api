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

	"golang-api/internal/model"
)

func newGeneratedReview(db *gorm.DB, opts ...gen.DOOption) generatedReview {
	_generatedReview := generatedReview{}

	_generatedReview.generatedReviewDo.UseDB(db, opts...)
	_generatedReview.generatedReviewDo.UseModel(&model.GeneratedReview{})

	tableName := _generatedReview.generatedReviewDo.TableName()
	_generatedReview.ALL = field.NewAsterisk(tableName)
	_generatedReview.ID = field.NewString(tableName, "id")
	_generatedReview.ProductID = field.NewString(tableName, "product_id")
	_generatedReview.Rating = field.NewString(tableName, "rating")
	_generatedReview.Pros = field.NewString(tableName, "pros")
	_generatedReview.Cons = field.NewString(tableName, "cons")
	_generatedReview.CreatedAt = field.NewTime(tableName, "created_at")
	_generatedReview.Overview = field.NewString(tableName, "overview")

	_generatedReview.fillFieldMap()

	return _generatedReview
}

type generatedReview struct {
	generatedReviewDo

	ALL       field.Asterisk
	ID        field.String
	ProductID field.String
	Rating    field.String
	Pros      field.String
	Cons      field.String
	CreatedAt field.Time
	Overview  field.String

	fieldMap map[string]field.Expr
}

func (g generatedReview) Table(newTableName string) *generatedReview {
	g.generatedReviewDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g generatedReview) As(alias string) *generatedReview {
	g.generatedReviewDo.DO = *(g.generatedReviewDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *generatedReview) updateTableName(table string) *generatedReview {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewString(table, "id")
	g.ProductID = field.NewString(table, "product_id")
	g.Rating = field.NewString(table, "rating")
	g.Pros = field.NewString(table, "pros")
	g.Cons = field.NewString(table, "cons")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.Overview = field.NewString(table, "overview")

	g.fillFieldMap()

	return g
}

func (g *generatedReview) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *generatedReview) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["product_id"] = g.ProductID
	g.fieldMap["rating"] = g.Rating
	g.fieldMap["pros"] = g.Pros
	g.fieldMap["cons"] = g.Cons
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["overview"] = g.Overview
}

func (g generatedReview) clone(db *gorm.DB) generatedReview {
	g.generatedReviewDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g generatedReview) replaceDB(db *gorm.DB) generatedReview {
	g.generatedReviewDo.ReplaceDB(db)
	return g
}

type generatedReviewDo struct{ gen.DO }

type IGeneratedReviewDo interface {
	gen.SubQuery
	Debug() IGeneratedReviewDo
	WithContext(ctx context.Context) IGeneratedReviewDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGeneratedReviewDo
	WriteDB() IGeneratedReviewDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGeneratedReviewDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGeneratedReviewDo
	Not(conds ...gen.Condition) IGeneratedReviewDo
	Or(conds ...gen.Condition) IGeneratedReviewDo
	Select(conds ...field.Expr) IGeneratedReviewDo
	Where(conds ...gen.Condition) IGeneratedReviewDo
	Order(conds ...field.Expr) IGeneratedReviewDo
	Distinct(cols ...field.Expr) IGeneratedReviewDo
	Omit(cols ...field.Expr) IGeneratedReviewDo
	Join(table schema.Tabler, on ...field.Expr) IGeneratedReviewDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGeneratedReviewDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGeneratedReviewDo
	Group(cols ...field.Expr) IGeneratedReviewDo
	Having(conds ...gen.Condition) IGeneratedReviewDo
	Limit(limit int) IGeneratedReviewDo
	Offset(offset int) IGeneratedReviewDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGeneratedReviewDo
	Unscoped() IGeneratedReviewDo
	Create(values ...*model.GeneratedReview) error
	CreateInBatches(values []*model.GeneratedReview, batchSize int) error
	Save(values ...*model.GeneratedReview) error
	First() (*model.GeneratedReview, error)
	Take() (*model.GeneratedReview, error)
	Last() (*model.GeneratedReview, error)
	Find() ([]*model.GeneratedReview, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GeneratedReview, err error)
	FindInBatches(result *[]*model.GeneratedReview, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GeneratedReview) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGeneratedReviewDo
	Assign(attrs ...field.AssignExpr) IGeneratedReviewDo
	Joins(fields ...field.RelationField) IGeneratedReviewDo
	Preload(fields ...field.RelationField) IGeneratedReviewDo
	FirstOrInit() (*model.GeneratedReview, error)
	FirstOrCreate() (*model.GeneratedReview, error)
	FindByPage(offset int, limit int) (result []*model.GeneratedReview, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGeneratedReviewDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g generatedReviewDo) Debug() IGeneratedReviewDo {
	return g.withDO(g.DO.Debug())
}

func (g generatedReviewDo) WithContext(ctx context.Context) IGeneratedReviewDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g generatedReviewDo) ReadDB() IGeneratedReviewDo {
	return g.Clauses(dbresolver.Read)
}

func (g generatedReviewDo) WriteDB() IGeneratedReviewDo {
	return g.Clauses(dbresolver.Write)
}

func (g generatedReviewDo) Session(config *gorm.Session) IGeneratedReviewDo {
	return g.withDO(g.DO.Session(config))
}

func (g generatedReviewDo) Clauses(conds ...clause.Expression) IGeneratedReviewDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g generatedReviewDo) Returning(value interface{}, columns ...string) IGeneratedReviewDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g generatedReviewDo) Not(conds ...gen.Condition) IGeneratedReviewDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g generatedReviewDo) Or(conds ...gen.Condition) IGeneratedReviewDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g generatedReviewDo) Select(conds ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g generatedReviewDo) Where(conds ...gen.Condition) IGeneratedReviewDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g generatedReviewDo) Order(conds ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g generatedReviewDo) Distinct(cols ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g generatedReviewDo) Omit(cols ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g generatedReviewDo) Join(table schema.Tabler, on ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g generatedReviewDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g generatedReviewDo) RightJoin(table schema.Tabler, on ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g generatedReviewDo) Group(cols ...field.Expr) IGeneratedReviewDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g generatedReviewDo) Having(conds ...gen.Condition) IGeneratedReviewDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g generatedReviewDo) Limit(limit int) IGeneratedReviewDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g generatedReviewDo) Offset(offset int) IGeneratedReviewDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g generatedReviewDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGeneratedReviewDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g generatedReviewDo) Unscoped() IGeneratedReviewDo {
	return g.withDO(g.DO.Unscoped())
}

func (g generatedReviewDo) Create(values ...*model.GeneratedReview) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g generatedReviewDo) CreateInBatches(values []*model.GeneratedReview, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g generatedReviewDo) Save(values ...*model.GeneratedReview) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g generatedReviewDo) First() (*model.GeneratedReview, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GeneratedReview), nil
	}
}

func (g generatedReviewDo) Take() (*model.GeneratedReview, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GeneratedReview), nil
	}
}

func (g generatedReviewDo) Last() (*model.GeneratedReview, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GeneratedReview), nil
	}
}

func (g generatedReviewDo) Find() ([]*model.GeneratedReview, error) {
	result, err := g.DO.Find()
	return result.([]*model.GeneratedReview), err
}

func (g generatedReviewDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GeneratedReview, err error) {
	buf := make([]*model.GeneratedReview, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g generatedReviewDo) FindInBatches(result *[]*model.GeneratedReview, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g generatedReviewDo) Attrs(attrs ...field.AssignExpr) IGeneratedReviewDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g generatedReviewDo) Assign(attrs ...field.AssignExpr) IGeneratedReviewDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g generatedReviewDo) Joins(fields ...field.RelationField) IGeneratedReviewDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g generatedReviewDo) Preload(fields ...field.RelationField) IGeneratedReviewDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g generatedReviewDo) FirstOrInit() (*model.GeneratedReview, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GeneratedReview), nil
	}
}

func (g generatedReviewDo) FirstOrCreate() (*model.GeneratedReview, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GeneratedReview), nil
	}
}

func (g generatedReviewDo) FindByPage(offset int, limit int) (result []*model.GeneratedReview, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g generatedReviewDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g generatedReviewDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g generatedReviewDo) Delete(models ...*model.GeneratedReview) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *generatedReviewDo) withDO(do gen.Dao) *generatedReviewDo {
	g.DO = *do.(*gen.DO)
	return g
}
