package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-simple-demo/internal/service/repo"
	"github.com/golang-acexy/cloud-web/webcloud"
	"gorm.io/gorm"
)

var teacherBizService = &TeacherBizService[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO]{
	repo: repo.NewTeacherRepo(),
}

func NewTeacherBizService() *TeacherBizService[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO] {
	return teacherBizService
}

type TeacherBizService[ID webcloud.IDType, S, M, Q, D any] struct {
	repo *repo.TeacherRepo
}

func (v *TeacherBizService[ID, S, M, Q, D]) MaxQueryCount() int {
	return 500
}

func (v *TeacherBizService[ID, S, M, Q, D]) DefaultOrderBySQL() string {
	return "id desc"
}

func (v *TeacherBizService[ID, S, M, Q, D]) Save(save *model.TeacherSDTO) (int64, error) {
	var t = save.ToT()
	_, err := v.repo.SaveExcludeZeroField(t)
	if err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (v *TeacherBizService[ID, S, M, Q, D]) QueryByID(condition map[string]any, result *model.TeacherDTO) (row int64, err error) {
	var r model.Teacher
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *TeacherBizService[ID, S, M, Q, D]) QueryOne(condition map[string]any, result *model.TeacherDTO) (row int64, err error) {
	var r model.Teacher
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *TeacherBizService[ID, S, M, Q, D]) Query(condition map[string]any, result *[]*model.TeacherDTO) (row int64, err error) {
	var r []*model.Teacher
	row, err = v.repo.QueryByGorm(&r, func(db *gorm.DB) {
		db.Where(condition).Order(v.DefaultOrderBySQL()).Limit(v.MaxQueryCount()).Scan(&r)
	})
	if row > 0 {
		model.TeacherSlice(r).ParseDTOs(result)
	}
	return
}

func (v *TeacherBizService[ID, S, M, Q, D]) QueryByPager(condition map[string]any, pager *webcloud.Pager[model.TeacherDTO]) error {
	p := databasecloud.Pager[model.Teacher]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	err := v.repo.QueryPageByMap(condition, v.DefaultOrderBySQL(), &p)
	if err != nil {
		return err
	}
	pager.Records = model.TeacherSlice(p.Records).ToDTOs()
	pager.Total = p.Total
	return nil
}

func (v *TeacherBizService[ID, S, M, Q, D]) ModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *TeacherBizService[ID, S, M, Q, D]) RemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}
