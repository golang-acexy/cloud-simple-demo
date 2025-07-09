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
	if save == nil {
		return 0, nil
	}
	var t = save.ToT()
	_, err := v.repo.SaveExcludeZeroField(t)
	if err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (v *TeacherBizService[ID, S, M, Q, D]) BaseQueryByID(condition map[string]any, result *model.TeacherDTO) (row int64, err error) {
	var r model.Teacher
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *TeacherBizService[ID, S, M, Q, D]) BaseQueryOne(condition map[string]any, result *model.TeacherDTO) (row int64, err error) {
	var r model.Teacher
	row, err = v.repo.QueryOneByMap(condition, &r)
	if row > 0 {
		r.ParseDTO(result)
	}
	return
}

func (v *TeacherBizService[ID, S, M, Q, D]) BaseQuery(condition map[string]any, result *[]*model.TeacherDTO) (row int64, err error) {
	var r []*model.Teacher
	row, err = v.repo.QueryByGorm(&r, func(db *gorm.DB) {
		db.Where(condition).Order(v.DefaultOrderBySQL()).Limit(v.MaxQueryCount()).Scan(&r)
	})
	if row > 0 {
		model.TeacherSlice(r).ParseDTOs(result)
	}
	return
}

func (v *TeacherBizService[ID, S, M, Q, D]) BaseQueryByPager(condition map[string]any, pager *webcloud.Pager[model.TeacherDTO]) error {
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

func (v *TeacherBizService[ID, S, M, Q, D]) BaseModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *TeacherBizService[ID, S, M, Q, D]) BaseRemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}

// QueryByID 通过主键查询
func (v *TeacherBizService[ID, S, M, Q, D]) QueryByID(id ID) *model.TeacherDTO {
	var r model.Teacher
	row, err := v.repo.QueryByID(id, &r)
	if row > 0 && err == nil {
		return r.ToDTO()
	}
	return nil
}

// QueryOneByCond 通过条件查询一条数据
func (v *TeacherBizService[ID, S, M, Q, D]) QueryOneByCond(condition *model.TeacherQDTO) *model.TeacherDTO {
	if condition == nil {
		return nil
	}
	var r model.Teacher
	row, err := v.repo.QueryOneByCond(condition.ToT(), &r)
	if row > 0 && err == nil {
		return r.ToDTO()
	}
	return nil
}

// QueryByCond 通过条件查询多条数据
func (v *TeacherBizService[ID, S, M, Q, D]) QueryByCond(condition *model.TeacherQDTO) []*model.TeacherDTO {
	if condition == nil {
		return nil
	}
	var rs []*model.Teacher
	row, err := v.repo.QueryByCond(condition.ToT(), v.DefaultOrderBySQL(), &rs)
	if row > 0 && err == nil {
		return model.TeacherSlice(rs).ToDTOs()
	}
	return nil
}

// QueryByPager 分页查询
func (v *TeacherBizService[ID, S, M, Q, D]) QueryByPager(pager webcloud.PagerDTO[model.TeacherQDTO]) webcloud.Pager[model.TeacherDTO] {
	p := databasecloud.Pager[model.Teacher]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	r := webcloud.Pager[model.TeacherDTO]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	err := v.repo.QueryPageByCond(pager.Condition.ToT(), v.DefaultOrderBySQL(), &p)
	if err == nil {
		r.Total = p.Total
		r.Records = model.TeacherSlice(p.Records).ToDTOs()
	}
	return r
}

// ModifyByID 根据主键修改数据
func (v *TeacherBizService[ID, S, M, Q, D]) ModifyByID(updated *model.TeacherMDTO) bool {
	if updated == nil {
		return false
	}
	row, err := v.repo.ModifyByID(updated.ToT())
	return row > 0 && err == nil
}

// ModifyByIDExcludeZeroField 根据主键修改数据 不包括零值数据
func (v *TeacherBizService[ID, S, M, Q, D]) ModifyByIDExcludeZeroField(updated *model.TeacherMDTO) bool {
	if updated == nil {
		return false
	}
	row, err := v.repo.ModifyByIDExcludeZeroField(updated.ToT())
	return row > 0 && err == nil
}

// ModifyByIdUseMap 根据主键修改数据 使用map
func (v *TeacherBizService[ID, S, M, Q, D]) ModifyByIdUseMap(updated map[string]any, id ID) bool {
	row, err := v.repo.ModifyByIdUseMap(updated, id)
	return row > 0 && err == nil
}

// RemoveByID 根据主键删除数据
func (v *TeacherBizService[ID, S, M, Q, D]) RemoveByID(id ID) bool {
	row, err := v.repo.RemoveByID(id)
	return row > 0 && err == nil
}

// RemoveByCond 根据条件删除数据
func (v *TeacherBizService[ID, S, M, Q, D]) RemoveByCond(condition *model.TeacherDTO) bool {
	if condition == nil {
		return false
	}
	row, err := v.repo.RemoveByCond(condition.ToT())
	return row > 0 && err == nil
}

// RemoveByMap 根据条件删除数据
func (v *TeacherBizService[ID, S, M, Q, D]) RemoveByMap(condition map[string]any) bool {
	row, err := v.repo.RemoveByMap(condition)
	return row > 0 && err == nil
}
