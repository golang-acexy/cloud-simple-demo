package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-simple-demo/internal/service/repo"
	"github.com/golang-acexy/cloud-web/webcloud"
	"gorm.io/gorm"
)

var studentBizService = &StudentBizService[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]{
	repo: repo.NewStudentRepo(),
}

func NewStudentBizService() *StudentBizService[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO] {
	return studentBizService
}

type StudentBizService[ID webcloud.IDType, S, M, Q, D any] struct {
	repo *repo.StudentRepo
}

func (v *StudentBizService[ID, S, M, Q, D]) MaxQueryCount() int {
	return 500
}

func (v *StudentBizService[ID, S, M, Q, D]) DefaultOrderBySQL() string {
	return "id desc"
}

func (v *StudentBizService[ID, S, M, Q, D]) Save(save *model.StudentSDTO) (int64, error) {
	var t = save.ToT()
	_, err := v.repo.SaveExcludeZeroField(t)
	if err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (v *StudentBizService[ID, S, M, Q, D]) QueryByID(condition map[string]any, result *model.StudentDTO) (row int64, err error) {
	var r model.Student
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *StudentBizService[ID, S, M, Q, D]) QueryOne(condition map[string]any, result *model.StudentDTO) (row int64, err error) {
	var r model.Student
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *StudentBizService[ID, S, M, Q, D]) Query(condition map[string]any, result *[]*model.StudentDTO) (row int64, err error) {
	var r []*model.Student
	row, err = v.repo.QueryByGorm(&r, func(db *gorm.DB) {
		db.Where(condition).Order(v.DefaultOrderBySQL()).Limit(v.MaxQueryCount()).Scan(&r)
	})
	if row > 0 {
		model.StudentSlice(r).ParseDTOs(result)
	}
	return
}

func (v *StudentBizService[ID, S, M, Q, D]) QueryByPager(condition map[string]any, pager *webcloud.Pager[model.StudentDTO]) error {
	p := databasecloud.Pager[model.Student]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	err := v.repo.QueryPageByMap(condition, v.DefaultOrderBySQL(), &p)
	if err != nil {
		return err
	}
	pager.Records = model.StudentSlice(p.Records).ToDTOs()
	pager.Total = p.Total
	return nil
}

func (v *StudentBizService[ID, S, M, Q, D]) ModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *StudentBizService[ID, S, M, Q, D]) RemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}
