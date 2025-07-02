package repo

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"gorm.io/gorm"
)

var teacherRepo = &TeacherRepo{
	GormRepository: databasecloud.GormRepository[
		gormstarter.IBaseMapper[gormstarter.BaseMapper[model.Teacher], model.Teacher],
		gormstarter.BaseMapper[model.Teacher],
		model.Teacher,
	]{
		Mapper: TeacherMapper{},
	},
}

func NewTeacherRepo() *TeacherRepo {
	return teacherRepo
}

func (r TeacherRepo) RawMapper() TeacherMapper {
	return TeacherMapper{r.GormRepository.RawMapper()}
}

func (r TeacherRepo) WithTxRepo(tx *gorm.DB) TeacherRepo {
	return TeacherRepo{
		GormRepository: r.GormRepository.WithTxRepo(tx),
	}
}

func (r TeacherRepo) NewTxRepo() TeacherRepo {
	return TeacherRepo{
		GormRepository: r.GormRepository.NewTxRepo(),
	}
}

type TeacherMapper struct {
	gormstarter.BaseMapper[model.Teacher]
}
type TeacherRepo struct {
	databasecloud.GormRepository[gormstarter.IBaseMapper[gormstarter.BaseMapper[model.Teacher],
		model.Teacher], gormstarter.BaseMapper[model.Teacher], model.Teacher]
}

// 在此处拓展自定义Mapper的业务功能

//func (m TeacherMapper) MyMapperMethod() {
//
//}

// ---------------------------------------

// 在此处拓展自定义Repo的业务功能

//func (r TeacherRepo) MyRepoMethod()  {
//
//}
