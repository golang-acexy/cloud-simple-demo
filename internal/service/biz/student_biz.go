package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-simple-demo/internal/service/repo"
	"github.com/golang-acexy/cloud-web/webcloud"
	"gorm.io/gorm"
)

var studentBizService = &StudentBizService[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]{
	reop:          repo.NewStudentRepo(),
	maxQueryCount: 500,
	baseOrderBy:   "id desc",
}

func NewStudentBizService() *StudentBizService[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO] {
	return studentBizService
}

type StudentBizService[ID webcloud.IDType, S, M, Q, D any] struct {
	reop          *repo.StudentRepo
	maxQueryCount int
	baseOrderBy   string
}

func (s *StudentBizService[ID, S, M, Q, D]) Save(save *model.StudentSDTO) (int64, error) {
	var t = save.ToT()
	_, err := s.reop.Save(t)
	if err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (s *StudentBizService[ID, S, M, Q, D]) QueryByID(condition map[string]any, result *model.StudentDTO) (int64, error) {
	var r model.Student
	defer r.ParseDTO(result)
	return s.reop.QueryOneByMap(condition, &r)
}

func (s *StudentBizService[ID, S, M, Q, D]) QueryOne(condition map[string]any, result *model.StudentDTO) (int64, error) {
	var r model.Student
	defer r.ParseDTO(result)
	return s.reop.QueryOneByMap(condition, &r)
}

func (s *StudentBizService[ID, S, M, Q, D]) Query(condition map[string]any, result *[]*model.StudentDTO) (int64, error) {
	var r []*model.Student
	defer model.StudentSlice(r).ParseDTOs(result)
	return s.reop.QueryByGorm(&r, func(db *gorm.DB) {
		db.Where(condition).Order(s.baseOrderBy).Limit(s.maxQueryCount).Scan(&r)
	})
}

func (s *StudentBizService[ID, S, M, Q, D]) QueryByPager(condition map[string]any, pager *webcloud.Pager[model.StudentDTO]) error {
	p := databasecloud.Pager[model.Student]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	defer func() {
		pager.Records = model.StudentSlice(p.Records).ToDTOs()
		pager.Total = p.Total
	}()
	return s.reop.QueryPageByMap(condition, s.baseOrderBy, &p)
}

func (s *StudentBizService[ID, S, M, Q, D]) ModifyByID(update, condition map[string]any) (int64, error) {
	return s.reop.ModifyByMap(update, condition)
}

func (s *StudentBizService[ID, S, M, Q, D]) RemoveByID(condition map[string]any) (int64, error) {
	return s.reop.RemoveByMap(condition)
}
