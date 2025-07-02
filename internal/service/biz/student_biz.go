package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
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

func (v *StudentBizService[ID, S, M, Q, D]) BaseQueryByID(condition map[string]any, result *model.StudentDTO) (row int64, err error) {
	var r model.Student
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *StudentBizService[ID, S, M, Q, D]) BaseQueryOne(condition map[string]any, result *model.StudentDTO) (row int64, err error) {
	var r model.Student
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *StudentBizService[ID, S, M, Q, D]) BaseQuery(condition map[string]any, result *[]*model.StudentDTO) (row int64, err error) {
	var r []*model.Student
	row, err = v.repo.QueryByGorm(&r, func(db *gorm.DB) {
		db.Where(condition).Order(v.DefaultOrderBySQL()).Limit(v.MaxQueryCount()).Scan(&r)
	})
	if row > 0 {
		model.StudentSlice(r).ParseDTOs(result)
	}
	return
}

func (v *StudentBizService[ID, S, M, Q, D]) BaseQueryByPager(condition map[string]any, pager webcloud.Pager[model.StudentDTO]) error {
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

func (v *StudentBizService[ID, S, M, Q, D]) BaseModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *StudentBizService[ID, S, M, Q, D]) BaseRemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}

// QueryByID 通过主键查询
func (v *StudentBizService[ID, S, M, Q, D]) QueryByID(id ID) *model.StudentDTO {
	var r model.Student
	row, err := v.repo.QueryByID(id, &r)
	if row > 0 && err == nil {
		return r.ToDTO()
	}
	return nil
}

// QueryOneByCond 通过条件查询一条数据
func (v *StudentBizService[ID, S, M, Q, D]) QueryOneByCond(condition *model.StudentQDTO) *model.StudentDTO {
	var r model.Student
	row, err := v.repo.QueryOneByCond(condition.ToT(), &r)
	if row > 0 && err == nil {
		return r.ToDTO()
	}
	return nil
}

// QueryByCond 通过条件查询多条数据
func (v *StudentBizService[ID, S, M, Q, D]) QueryByCond(condition *model.StudentQDTO) []*model.StudentDTO {
	var rs []*model.Student
	row, err := v.repo.QueryByCond(condition.ToT(), v.DefaultOrderBySQL(), &rs)
	if row > 0 && err == nil {
		return model.StudentSlice(rs).ToDTOs()
	}
	return nil
}

// QueryByPager 分页查询
func (v *StudentBizService[ID, S, M, Q, D]) QueryByPager(pager webcloud.PagerDTO[model.StudentQDTO]) webcloud.Pager[model.StudentDTO] {
	p := databasecloud.Pager[model.Student]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	r := webcloud.Pager[model.StudentDTO]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	err := v.repo.QueryPageByCond(pager.Condition.ToT(), v.DefaultOrderBySQL(), &p)
	if err == nil {
		r.Total = p.Total
		r.Records = model.StudentSlice(p.Records).ToDTOs()
	}
	return r
}

// ModifyByID 根据主键修改数据
func (v *StudentBizService[ID, S, M, Q, D]) ModifyByID(updated *model.StudentMDTO) bool {
	row, err := v.repo.ModifyByID(updated.ToT())
	return row > 0 && err == nil
}

// ModifyByIDExcludeZeroField 根据主键修改数据 不包括零值数据
func (v *StudentBizService[ID, S, M, Q, D]) ModifyByIDExcludeZeroField(updated *model.StudentMDTO) bool {
	row, err := v.repo.ModifyByIDExcludeZeroField(updated.ToT())
	return row > 0 && err == nil
}

// ModifyByIdUseMap 根据主键修改数据 使用map
func (v *StudentBizService[ID, S, M, Q, D]) ModifyByIdUseMap(updated map[string]any, id ID) bool {
	row, err := v.repo.ModifyByIdUseMap(updated, id)
	return row > 0 && err == nil
}

// RemoveByID 根据主键删除数据
func (v *StudentBizService[ID, S, M, Q, D]) RemoveByID(id ID) bool {
	row, err := v.repo.RemoveByID(id)
	return row > 0 && err == nil
}
