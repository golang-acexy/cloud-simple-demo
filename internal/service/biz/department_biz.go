package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud/rds"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	repo "github.com/golang-acexy/cloud-simple-demo/internal/service/repo"
	"github.com/golang-acexy/cloud-web/webcloud"
)

var departmentBizService = &DepartmentBizService{
	repo: repo.NewDepartmentRepo(),
}

func NewDepartmentBizService() *DepartmentBizService {
	return departmentBizService
}

type DepartmentBizService struct {
	repo repo.DepartmentRepo
}

func (v *DepartmentBizService) MaxQuerySize() int {
	return 100
}

func (v *DepartmentBizService) DefaultOrderBy() string {
	return "id desc"
}

func (v *DepartmentBizService) DefaultTimeRangeField() string {
	return "created_at"
}

func (v *DepartmentBizService) AllowedTimeRangeFields() []string {
	return []string{
		"created_at",
		"updated_at",
	}
}

func (v *DepartmentBizService) Save(save *model.DepartmentSDTO) (int64, error) {
	if save == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := save.ToT()
	if err != nil {
		return 0, err
	}
	if _, err = v.repo.Save(entity); err != nil {
		return 0, err
	}
	return entity.ID, nil
}

func (v *DepartmentBizService) SaveWithoutZeroFields(save *model.DepartmentSDTO) (int64, error) {
	if save == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := save.ToT()
	if err != nil {
		return 0, err
	}
	if _, err = v.repo.SaveWithoutZeroFields(entity); err != nil {
		return 0, err
	}
	return entity.ID, nil
}

func (v *DepartmentBizService) SaveBatch(saves []*model.DepartmentSDTO) ([]int64, error) {
	if len(saves) == 0 {
		return make([]int64, 0), nil
	}
	entities := make([]*model.Department, 0, len(saves))
	for _, save := range saves {
		if save == nil {
			return nil, webcloud.ErrBadRequestParameters
		}
		entity, err := save.ToT()
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	if _, err := v.repo.SaveBatch(entities); err != nil {
		return nil, err
	}
	ids := make([]int64, len(entities))
	for index, entity := range entities {
		ids[index] = entity.ID
	}
	return ids, nil
}

func (v *DepartmentBizService) BaseQueryByID(condition map[string]any) (*model.DepartmentDTO, error) {
	entity, err := v.repo.QueryOneByMap(rds.NewMapQuery(condition))
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) BaseQueryOne(condition map[string]any) (*model.DepartmentDTO, error) {
	entity, err := v.repo.QueryOneByMap(
		rds.NewMapQuery(condition).OrderBy(v.DefaultOrderBy()),
	)
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) BaseQuery(condition map[string]any) ([]*model.DepartmentDTO, error) {
	entities, err := v.repo.QueryByMap(
		rds.NewMapQuery(condition).
			OrderBy(v.DefaultOrderBy()).
			WithLimit(v.MaxQuerySize()),
	)
	if err != nil {
		return nil, err
	}
	return model.DepartmentSlice(entities).ToDTOs()
}

func (v *DepartmentBizService) BaseQueryPage(query webcloud.PagerDTO[map[string]any]) (webcloud.Pager[model.DepartmentDTO], error) {
	ranges, err := webcloud.ConvertTimeRanges(query.TimeRanges, v.AllowedTimeRangeFields(), v.DefaultTimeRangeField())
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	page, err := v.repo.QueryPageByMap(
		rds.NewMapPageQuery(query.Condition, query.Number, query.Size).
			OrderBy(v.DefaultOrderBy()).
			WithTimeRanges(ranges...),
	)
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	records, err := model.DepartmentSlice(page.Records).ToDTOs()
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	return webcloud.Pager[model.DepartmentDTO]{Records: records, Total: page.Total, Number: page.Number, Size: page.Size}, nil
}

func (v *DepartmentBizService) BaseModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *DepartmentBizService) BaseRemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}

func (v *DepartmentBizService) QueryByID(id int64) (*model.DepartmentDTO, error) {
	entity, err := v.repo.QueryByID(id)
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) QueryByIDs(ids []int64) ([]*model.DepartmentDTO, error) {
	if len(ids) == 0 {
		return make([]*model.DepartmentDTO, 0), nil
	}
	queryIDs := make([]any, len(ids))
	for index, id := range ids {
		queryIDs[index] = id
	}
	entities, err := v.repo.QueryByIDs(queryIDs)
	if err != nil {
		return nil, err
	}
	return model.DepartmentSlice(entities).ToDTOs()
}

func (v *DepartmentBizService) ExistsByID(id int64) (bool, error) {
	return v.repo.ExistsByID(id)
}

func (v *DepartmentBizService) QueryOneByCond(condition model.DepartmentQDTO) (*model.DepartmentDTO, error) {
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	entity, err := v.repo.QueryOneByCond(rds.NewCondQuery(*entityCondition))
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) QueryByCond(condition model.DepartmentQDTO) ([]*model.DepartmentDTO, error) {
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	entities, err := v.repo.QueryByCond(
		rds.NewCondQuery(*entityCondition).
			OrderBy(v.DefaultOrderBy()).
			WithLimit(v.MaxQuerySize()),
	)
	if err != nil {
		return nil, err
	}
	return model.DepartmentSlice(entities).ToDTOs()
}

func (v *DepartmentBizService) CountByCond(condition model.DepartmentQDTO) (int64, error) {
	entityCondition, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.CountByCond(rds.NewCondQuery(*entityCondition))
}

func (v *DepartmentBizService) CountByMap(condition map[string]any) (int64, error) {
	return v.repo.CountByMap(rds.NewMapQuery(condition))
}

func (v *DepartmentBizService) QueryPage(pager webcloud.PagerDTO[model.DepartmentQDTO]) (webcloud.Pager[model.DepartmentDTO], error) {
	condition, err := pager.Condition.ToT()
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	ranges, err := webcloud.ConvertTimeRanges(pager.TimeRanges, v.AllowedTimeRangeFields(), v.DefaultTimeRangeField())
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	page, err := v.repo.QueryPageByCond(
		rds.NewPageQuery(*condition, pager.Number, pager.Size).
			OrderBy(v.DefaultOrderBy()).
			WithTimeRanges(ranges...),
	)
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	records, err := model.DepartmentSlice(page.Records).ToDTOs()
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	return webcloud.Pager[model.DepartmentDTO]{
		Records: records,
		Total:   page.Total,
		Number:  page.Number,
		Size:    page.Size,
	}, nil
}

func (v *DepartmentBizService) ModifyByID(id int64, updated *model.DepartmentMDTO) (int64, error) {
	if updated == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := updated.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.ModifyByID(entity, id)
}

func (v *DepartmentBizService) ModifyByIDWithoutZeroFields(id int64, updated *model.DepartmentMDTO) (int64, error) {
	if updated == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := updated.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.ModifyByIDWithoutZeroFields(entity, id)
}

func (v *DepartmentBizService) ModifyByIDWithMap(id int64, updated map[string]any) (int64, error) {
	if len(updated) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.ModifyByIDWithMap(updated, id)
}

func (v *DepartmentBizService) ModifyByCond(condition model.DepartmentQDTO, updated *model.DepartmentMDTO) (int64, error) {
	if updated == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	conditionEntity, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	updatedEntity, err := updated.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.ModifyByCond(updatedEntity, *conditionEntity)
}

func (v *DepartmentBizService) ModifyByMap(updated, condition map[string]any) (int64, error) {
	if len(updated) == 0 || len(condition) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.ModifyByMap(updated, condition)
}

func (v *DepartmentBizService) RemoveByID(id int64) (int64, error) {
	return v.repo.RemoveByID(id)
}

func (v *DepartmentBizService) RemoveByIDs(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	removeIDs := make([]any, len(ids))
	for index, id := range ids {
		removeIDs[index] = id
	}
	return v.repo.RemoveByIDs(removeIDs)
}

func (v *DepartmentBizService) RemoveByCond(condition model.DepartmentQDTO) (int64, error) {
	entity, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.RemoveByCond(*entity)
}

func (v *DepartmentBizService) RemoveByMap(condition map[string]any) (int64, error) {
	if len(condition) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.RemoveByMap(condition)
}
