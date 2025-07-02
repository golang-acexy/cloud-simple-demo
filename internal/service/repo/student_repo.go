package repo

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"gorm.io/gorm"
)

var studentRepo = &StudentRepo{
	GormRepository: databasecloud.GormRepository[
		gormstarter.IBaseMapper[gormstarter.BaseMapper[model.Student], model.Student],
		gormstarter.BaseMapper[model.Student],
		model.Student,
	]{
		Mapper: StudentMapper{},
	},
}

func NewStudentRepo() *StudentRepo {
	return studentRepo
}

func (r StudentRepo) RawMapper() StudentMapper {
	return StudentMapper{r.GormRepository.RawMapper()}
}

func (r StudentRepo) WithTxRepo(tx *gorm.DB) StudentRepo {
	return StudentRepo{
		GormRepository: r.GormRepository.WithTxRepo(tx),
	}
}

func (r StudentRepo) NewTxRepo() StudentRepo {
	return StudentRepo{
		GormRepository: r.GormRepository.NewTxRepo(),
	}
}

type StudentMapper struct {
	gormstarter.BaseMapper[model.Student]
}
type StudentRepo struct {
	databasecloud.GormRepository[gormstarter.IBaseMapper[gormstarter.BaseMapper[model.Student],
		model.Student], gormstarter.BaseMapper[model.Student], model.Student]
}

// 在此处拓展自定义Mapper的业务功能

//func (m StudentMapper) MyMapperMethod() {
//
//}

// ---------------------------------------

// 在此处拓展自定义Repo的业务功能

//func (r StudentRepo) MyRepoMethod()  {
//
//}
