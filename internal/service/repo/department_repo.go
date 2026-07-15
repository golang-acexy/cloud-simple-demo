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

func (m DepartmentMapper) WithTxMapper(tx *gorm.DB) DepartmentMapper {
	return DepartmentMapper{
		BaseMapper: m.GetBaseMapperWithTx(tx),
	}
}

type DepartmentRepo struct {
	rds.Repository[DepartmentRepo, DepartmentMapper, model.Department]
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
