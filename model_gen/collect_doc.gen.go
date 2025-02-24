// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/SelfDown/collect/model"
)

func newCollectDoc(db *gorm.DB, opts ...gen.DOOption) collectDoc {
	_collectDoc := collectDoc{}

	_collectDoc.collectDocDo.UseDB(db, opts...)
	_collectDoc.collectDocDo.UseModel(&model.CollectDoc{})

	tableName := _collectDoc.collectDocDo.TableName()
	_collectDoc.ALL = field.NewAsterisk(tableName)
	_collectDoc.CollectDocID = field.NewString(tableName, "collect_doc_id")
	_collectDoc.Title = field.NewString(tableName, "title")
	_collectDoc.SubTitle = field.NewString(tableName, "sub_title")
	_collectDoc.Type = field.NewString(tableName, "type")
	_collectDoc.ParentDir = field.NewString(tableName, "parent_dir")
	_collectDoc.Code = field.NewString(tableName, "code")
	_collectDoc.CodeDesc = field.NewString(tableName, "code_desc")
	_collectDoc.OrderIndex = field.NewInt32(tableName, "order_index")
	_collectDoc.CreateTime = field.NewString(tableName, "create_time")
	_collectDoc.CreateUser = field.NewString(tableName, "create_user")
	_collectDoc.IsDelete = field.NewString(tableName, "is_delete")
	_collectDoc.CodeResult = field.NewString(tableName, "code_result")

	_collectDoc.fillFieldMap()

	return _collectDoc
}

type collectDoc struct {
	collectDocDo

	ALL          field.Asterisk
	CollectDocID field.String
	Title        field.String
	SubTitle     field.String
	Type         field.String
	ParentDir    field.String
	Code         field.String
	CodeDesc     field.String
	OrderIndex   field.Int32
	CreateTime   field.String
	CreateUser   field.String
	IsDelete     field.String
	CodeResult   field.String

	fieldMap map[string]field.Expr
}

func (c collectDoc) Table(newTableName string) *collectDoc {
	c.collectDocDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c collectDoc) As(alias string) *collectDoc {
	c.collectDocDo.DO = *(c.collectDocDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *collectDoc) updateTableName(table string) *collectDoc {
	c.ALL = field.NewAsterisk(table)
	c.CollectDocID = field.NewString(table, "collect_doc_id")
	c.Title = field.NewString(table, "title")
	c.SubTitle = field.NewString(table, "sub_title")
	c.Type = field.NewString(table, "type")
	c.ParentDir = field.NewString(table, "parent_dir")
	c.Code = field.NewString(table, "code")
	c.CodeDesc = field.NewString(table, "code_desc")
	c.OrderIndex = field.NewInt32(table, "order_index")
	c.CreateTime = field.NewString(table, "create_time")
	c.CreateUser = field.NewString(table, "create_user")
	c.IsDelete = field.NewString(table, "is_delete")
	c.CodeResult = field.NewString(table, "code_result")

	c.fillFieldMap()

	return c
}

func (c *collectDoc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *collectDoc) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["collect_doc_id"] = c.CollectDocID
	c.fieldMap["title"] = c.Title
	c.fieldMap["sub_title"] = c.SubTitle
	c.fieldMap["type"] = c.Type
	c.fieldMap["parent_dir"] = c.ParentDir
	c.fieldMap["code"] = c.Code
	c.fieldMap["code_desc"] = c.CodeDesc
	c.fieldMap["order_index"] = c.OrderIndex
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["create_user"] = c.CreateUser
	c.fieldMap["is_delete"] = c.IsDelete
	c.fieldMap["code_result"] = c.CodeResult
}

func (c collectDoc) clone(db *gorm.DB) collectDoc {
	c.collectDocDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c collectDoc) replaceDB(db *gorm.DB) collectDoc {
	c.collectDocDo.ReplaceDB(db)
	return c
}

type collectDocDo struct{ gen.DO }

type ICollectDocDo interface {
	gen.SubQuery
	Debug() ICollectDocDo
	WithContext(ctx context.Context) ICollectDocDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICollectDocDo
	WriteDB() ICollectDocDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICollectDocDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICollectDocDo
	Not(conds ...gen.Condition) ICollectDocDo
	Or(conds ...gen.Condition) ICollectDocDo
	Select(conds ...field.Expr) ICollectDocDo
	Where(conds ...gen.Condition) ICollectDocDo
	Order(conds ...field.Expr) ICollectDocDo
	Distinct(cols ...field.Expr) ICollectDocDo
	Omit(cols ...field.Expr) ICollectDocDo
	Join(table schema.Tabler, on ...field.Expr) ICollectDocDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICollectDocDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICollectDocDo
	Group(cols ...field.Expr) ICollectDocDo
	Having(conds ...gen.Condition) ICollectDocDo
	Limit(limit int) ICollectDocDo
	Offset(offset int) ICollectDocDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectDocDo
	Unscoped() ICollectDocDo
	Create(values ...*model.CollectDoc) error
	CreateInBatches(values []*model.CollectDoc, batchSize int) error
	Save(values ...*model.CollectDoc) error
	First() (*model.CollectDoc, error)
	Take() (*model.CollectDoc, error)
	Last() (*model.CollectDoc, error)
	Find() ([]*model.CollectDoc, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectDoc, err error)
	FindInBatches(result *[]*model.CollectDoc, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CollectDoc) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICollectDocDo
	Assign(attrs ...field.AssignExpr) ICollectDocDo
	Joins(fields ...field.RelationField) ICollectDocDo
	Preload(fields ...field.RelationField) ICollectDocDo
	FirstOrInit() (*model.CollectDoc, error)
	FirstOrCreate() (*model.CollectDoc, error)
	FindByPage(offset int, limit int) (result []*model.CollectDoc, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICollectDocDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c collectDocDo) Debug() ICollectDocDo {
	return c.withDO(c.DO.Debug())
}

func (c collectDocDo) WithContext(ctx context.Context) ICollectDocDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c collectDocDo) ReadDB() ICollectDocDo {
	return c.Clauses(dbresolver.Read)
}

func (c collectDocDo) WriteDB() ICollectDocDo {
	return c.Clauses(dbresolver.Write)
}

func (c collectDocDo) Session(config *gorm.Session) ICollectDocDo {
	return c.withDO(c.DO.Session(config))
}

func (c collectDocDo) Clauses(conds ...clause.Expression) ICollectDocDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c collectDocDo) Returning(value interface{}, columns ...string) ICollectDocDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c collectDocDo) Not(conds ...gen.Condition) ICollectDocDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c collectDocDo) Or(conds ...gen.Condition) ICollectDocDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c collectDocDo) Select(conds ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c collectDocDo) Where(conds ...gen.Condition) ICollectDocDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c collectDocDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICollectDocDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c collectDocDo) Order(conds ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c collectDocDo) Distinct(cols ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c collectDocDo) Omit(cols ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c collectDocDo) Join(table schema.Tabler, on ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c collectDocDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c collectDocDo) RightJoin(table schema.Tabler, on ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c collectDocDo) Group(cols ...field.Expr) ICollectDocDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c collectDocDo) Having(conds ...gen.Condition) ICollectDocDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c collectDocDo) Limit(limit int) ICollectDocDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c collectDocDo) Offset(offset int) ICollectDocDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c collectDocDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectDocDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c collectDocDo) Unscoped() ICollectDocDo {
	return c.withDO(c.DO.Unscoped())
}

func (c collectDocDo) Create(values ...*model.CollectDoc) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c collectDocDo) CreateInBatches(values []*model.CollectDoc, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c collectDocDo) Save(values ...*model.CollectDoc) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c collectDocDo) First() (*model.CollectDoc, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDoc), nil
	}
}

func (c collectDocDo) Take() (*model.CollectDoc, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDoc), nil
	}
}

func (c collectDocDo) Last() (*model.CollectDoc, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDoc), nil
	}
}

func (c collectDocDo) Find() ([]*model.CollectDoc, error) {
	result, err := c.DO.Find()
	return result.([]*model.CollectDoc), err
}

func (c collectDocDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectDoc, err error) {
	buf := make([]*model.CollectDoc, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c collectDocDo) FindInBatches(result *[]*model.CollectDoc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c collectDocDo) Attrs(attrs ...field.AssignExpr) ICollectDocDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c collectDocDo) Assign(attrs ...field.AssignExpr) ICollectDocDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c collectDocDo) Joins(fields ...field.RelationField) ICollectDocDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c collectDocDo) Preload(fields ...field.RelationField) ICollectDocDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c collectDocDo) FirstOrInit() (*model.CollectDoc, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDoc), nil
	}
}

func (c collectDocDo) FirstOrCreate() (*model.CollectDoc, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDoc), nil
	}
}

func (c collectDocDo) FindByPage(offset int, limit int) (result []*model.CollectDoc, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c collectDocDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c collectDocDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c collectDocDo) Delete(models ...*model.CollectDoc) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *collectDocDo) withDO(do gen.Dao) *collectDocDo {
	c.DO = *do.(*gen.DO)
	return c
}
