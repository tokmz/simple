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

	"simple/internal/types/entity"
)

func newPosition(db *gorm.DB, opts ...gen.DOOption) position {
	_position := position{}

	_position.positionDo.UseDB(db, opts...)
	_position.positionDo.UseModel(&entity.Position{})

	tableName := _position.positionDo.TableName()
	_position.ALL = field.NewAsterisk(tableName)
	_position.ID = field.NewInt64(tableName, "id")
	_position.DepartmentID = field.NewInt64(tableName, "department_id")
	_position.Name = field.NewString(tableName, "name")
	_position.Code = field.NewString(tableName, "code")
	_position.Sort = field.NewInt64(tableName, "sort")
	_position.Status = field.NewInt64(tableName, "status")
	_position.Remark = field.NewString(tableName, "remark")
	_position.CreatedAt = field.NewTime(tableName, "created_at")
	_position.UpdatedAt = field.NewTime(tableName, "updated_at")
	_position.DeletedAt = field.NewField(tableName, "deleted_at")

	_position.fillFieldMap()

	return _position
}

// position 系统岗位表
type position struct {
	positionDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键ID|Primary key
	DepartmentID field.Int64  // 部门ID|Department ID
	Name         field.String // 岗位名称|Position name
	Code         field.String // 岗位编码|Position code
	Sort         field.Int64  // 排序|Sort
	Status       field.Int64  // 状态 1:启用 2:禁用|Status 1:Enable 2:Disable
	Remark       field.String // 备注|Remark
	CreatedAt    field.Time   // 创建时间|Created Time
	UpdatedAt    field.Time   // 更新时间|Updated Time
	DeletedAt    field.Field  // 删除时间|Deleted Time

	fieldMap map[string]field.Expr
}

func (p position) Table(newTableName string) *position {
	p.positionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p position) As(alias string) *position {
	p.positionDo.DO = *(p.positionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *position) updateTableName(table string) *position {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.DepartmentID = field.NewInt64(table, "department_id")
	p.Name = field.NewString(table, "name")
	p.Code = field.NewString(table, "code")
	p.Sort = field.NewInt64(table, "sort")
	p.Status = field.NewInt64(table, "status")
	p.Remark = field.NewString(table, "remark")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *position) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *position) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["id"] = p.ID
	p.fieldMap["department_id"] = p.DepartmentID
	p.fieldMap["name"] = p.Name
	p.fieldMap["code"] = p.Code
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["status"] = p.Status
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt

}

func (p position) clone(db *gorm.DB) position {
	p.positionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p position) replaceDB(db *gorm.DB) position {
	p.positionDo.ReplaceDB(db)
	return p
}

type positionDo struct{ gen.DO }

func (p positionDo) Debug() *positionDo {
	return p.withDO(p.DO.Debug())
}

func (p positionDo) WithContext(ctx context.Context) *positionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p positionDo) ReadDB() *positionDo {
	return p.Clauses(dbresolver.Read)
}

func (p positionDo) WriteDB() *positionDo {
	return p.Clauses(dbresolver.Write)
}

func (p positionDo) Session(config *gorm.Session) *positionDo {
	return p.withDO(p.DO.Session(config))
}

func (p positionDo) Clauses(conds ...clause.Expression) *positionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p positionDo) Returning(value interface{}, columns ...string) *positionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p positionDo) Not(conds ...gen.Condition) *positionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p positionDo) Or(conds ...gen.Condition) *positionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p positionDo) Select(conds ...field.Expr) *positionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p positionDo) Where(conds ...gen.Condition) *positionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p positionDo) Order(conds ...field.Expr) *positionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p positionDo) Distinct(cols ...field.Expr) *positionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p positionDo) Omit(cols ...field.Expr) *positionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p positionDo) Join(table schema.Tabler, on ...field.Expr) *positionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p positionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *positionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p positionDo) RightJoin(table schema.Tabler, on ...field.Expr) *positionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p positionDo) Group(cols ...field.Expr) *positionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p positionDo) Having(conds ...gen.Condition) *positionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p positionDo) Limit(limit int) *positionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p positionDo) Offset(offset int) *positionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p positionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *positionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p positionDo) Unscoped() *positionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p positionDo) Create(values ...*entity.Position) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p positionDo) CreateInBatches(values []*entity.Position, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p positionDo) Save(values ...*entity.Position) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p positionDo) First() (*entity.Position, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Position), nil
	}
}

func (p positionDo) Take() (*entity.Position, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Position), nil
	}
}

func (p positionDo) Last() (*entity.Position, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Position), nil
	}
}

func (p positionDo) Find() ([]*entity.Position, error) {
	result, err := p.DO.Find()
	return result.([]*entity.Position), err
}

func (p positionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Position, err error) {
	buf := make([]*entity.Position, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p positionDo) FindInBatches(result *[]*entity.Position, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p positionDo) Attrs(attrs ...field.AssignExpr) *positionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p positionDo) Assign(attrs ...field.AssignExpr) *positionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p positionDo) Joins(fields ...field.RelationField) *positionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p positionDo) Preload(fields ...field.RelationField) *positionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p positionDo) FirstOrInit() (*entity.Position, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Position), nil
	}
}

func (p positionDo) FirstOrCreate() (*entity.Position, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Position), nil
	}
}

func (p positionDo) FindByPage(offset int, limit int) (result []*entity.Position, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p positionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p positionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p positionDo) Delete(models ...*entity.Position) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *positionDo) withDO(do gen.Dao) *positionDo {
	p.DO = *do.(*gen.DO)
	return p
}
