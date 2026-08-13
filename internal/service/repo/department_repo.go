package repo

import (
	"github.com/golang-acexy/cloud-database/databasecloud/rds"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"gorm.io/gorm"
)

type DepartmentMapper struct {
	gormstarter.BaseMapper[model.Department]
}

func (DepartmentMapper) Columns() *model.DepartmentColumns {
	return model.DepartmentColumnSet()
}

func (m DepartmentMapper) WithTxMapper(tx *gorm.DB) DepartmentMapper {
	return DepartmentMapper{
		BaseMapper: m.GetBaseMapperWithTx(tx),
	}
}

type DepartmentRepo struct {
	rds.Repository[DepartmentRepo, DepartmentMapper, model.Department]
}

// Columns 返回 Department 对应的类型安全字段集合。
func (r DepartmentRepo) Columns() *model.DepartmentColumns {
	return r.RawMapper().Columns()
}

var departmentRepo = newDepartmentRepo()

func newDepartmentRepo() DepartmentRepo {
	repository := rds.NewRepository(
		DepartmentMapper{},
		func(repository rds.Repository[DepartmentRepo, DepartmentMapper, model.Department]) DepartmentRepo {
			return DepartmentRepo{Repository: repository}
		},
	)
	return DepartmentRepo{Repository: repository}
}

func NewDepartmentRepo() DepartmentRepo {
	return departmentRepo
}

// 在此处扩展自定义 Mapper 能力。

// 在此处扩展自定义 Repository 能力。
