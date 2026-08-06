package repo

import (
	"github.com/golang-acexy/cloud-database/databasecloud/rds"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"gorm.io/gorm"
)

type EmployeeMapper struct {
	gormstarter.BaseMapper[model.Employee]
}

func (m EmployeeMapper) WithTxMapper(tx *gorm.DB) EmployeeMapper {
	return EmployeeMapper{
		BaseMapper: m.GetBaseMapperWithTx(tx),
	}
}

type EmployeeRepo struct {
	rds.Repository[EmployeeRepo, EmployeeMapper, model.Employee]
}

var employeeRepo = newEmployeeRepo()

func newEmployeeRepo() EmployeeRepo {
	repository := rds.NewRepository(
		EmployeeMapper{},
		func(repository rds.Repository[EmployeeRepo, EmployeeMapper, model.Employee]) EmployeeRepo {
			return EmployeeRepo{Repository: repository}
		},
	)
	return EmployeeRepo{Repository: repository}
}

func NewEmployeeRepo() EmployeeRepo {
	return employeeRepo
}

// 在此处扩展自定义 Mapper 能力。

// 在此处扩展自定义 Repository 能力。
