package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud/rds"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	repo "github.com/golang-acexy/cloud-simple-demo/internal/service/repo"
	"github.com/golang-acexy/cloud-web/webcloud"
)

var employeeBizService = &EmployeeBizService{
	repo: repo.NewEmployeeRepo(),
}

func NewEmployeeBizService() *EmployeeBizService {
	return employeeBizService
}

type EmployeeBizService struct {
	repo repo.EmployeeRepo
}

func (v *EmployeeBizService) MaxQuerySize() int {
	return 100
}

func (v *EmployeeBizService) DefaultOrderBy() string {
	return "id desc"
}

func (v *EmployeeBizService) DefaultTimeRangeField() string {
	return "created_at"
}

func (v *EmployeeBizService) AllowedTimeRangeFields() []string {
	return []string{
		"created_at",
		"updated_at",
	}
}

func (v *EmployeeBizService) Save(save *model.EmployeeSDTO) (int64, error) {
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

func (v *EmployeeBizService) SaveWithoutZeroFields(save *model.EmployeeSDTO) (int64, error) {
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

func (v *EmployeeBizService) SaveBatch(saves []*model.EmployeeSDTO) ([]int64, error) {
	if len(saves) == 0 {
		return make([]int64, 0), nil
	}
	entities := make([]*model.Employee, 0, len(saves))
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

func (v *EmployeeBizService) BaseQueryByID(condition map[string]any) (*model.EmployeeDTO, error) {
	entity, err := v.repo.QueryOneByMap(rds.NewMapQuery(condition))
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) BaseQueryOne(condition map[string]any) (*model.EmployeeDTO, error) {
	entity, err := v.repo.QueryOneByMap(
		rds.NewMapQuery(condition).OrderBy(v.DefaultOrderBy()),
	)
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) BaseQuery(condition map[string]any) ([]*model.EmployeeDTO, error) {
	entities, err := v.repo.QueryByMap(
		rds.NewMapQuery(condition).
			OrderBy(v.DefaultOrderBy()).
			WithLimit(v.MaxQuerySize()),
	)
	if err != nil {
		return nil, err
	}
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) BaseQueryPage(query webcloud.PagerDTO[map[string]any]) (webcloud.Pager[model.EmployeeDTO], error) {
	ranges, err := webcloud.ConvertTimeRanges(query.TimeRanges, v.AllowedTimeRangeFields(), v.DefaultTimeRangeField())
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	page, err := v.repo.QueryPageByMap(
		rds.NewMapPageQuery(query.Condition, query.Number, query.Size).
			OrderBy(v.DefaultOrderBy()).
			WithTimeRanges(ranges...),
	)
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	records, err := model.EmployeeSlice(page.Records).ToDTOs()
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	return webcloud.Pager[model.EmployeeDTO]{Records: records, Total: page.Total, Number: page.Number, Size: page.Size}, nil
}

func (v *EmployeeBizService) BaseModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *EmployeeBizService) BaseRemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}

func (v *EmployeeBizService) QueryByID(id int64) (*model.EmployeeDTO, error) {
	entity, err := v.repo.QueryByID(id)
	if err != nil || entity == nil {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) QueryByIDs(ids []int64) ([]*model.EmployeeDTO, error) {
	if len(ids) == 0 {
		return make([]*model.EmployeeDTO, 0), nil
	}
	queryIDs := make([]any, len(ids))
	for index, id := range ids {
		queryIDs[index] = id
	}
	entities, err := v.repo.QueryByIDs(queryIDs)
	if err != nil {
		return nil, err
	}
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) ExistsByID(id int64) (bool, error) {
	return v.repo.ExistsByID(id)
}

func (v *EmployeeBizService) QueryOneByCond(condition model.EmployeeQDTO) (*model.EmployeeDTO, error) {
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

func (v *EmployeeBizService) QueryByCond(condition model.EmployeeQDTO) ([]*model.EmployeeDTO, error) {
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
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) CountByCond(condition model.EmployeeQDTO) (int64, error) {
	entityCondition, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.CountByCond(rds.NewCondQuery(*entityCondition))
}

func (v *EmployeeBizService) CountByMap(condition map[string]any) (int64, error) {
	return v.repo.CountByMap(rds.NewMapQuery(condition))
}

func (v *EmployeeBizService) QueryPage(pager webcloud.PagerDTO[model.EmployeeQDTO]) (webcloud.Pager[model.EmployeeDTO], error) {
	condition, err := pager.Condition.ToT()
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	ranges, err := webcloud.ConvertTimeRanges(pager.TimeRanges, v.AllowedTimeRangeFields(), v.DefaultTimeRangeField())
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	page, err := v.repo.QueryPageByCond(
		rds.NewPageQuery(*condition, pager.Number, pager.Size).
			OrderBy(v.DefaultOrderBy()).
			WithTimeRanges(ranges...),
	)
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	records, err := model.EmployeeSlice(page.Records).ToDTOs()
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	return webcloud.Pager[model.EmployeeDTO]{
		Records: records,
		Total:   page.Total,
		Number:  page.Number,
		Size:    page.Size,
	}, nil
}

func (v *EmployeeBizService) ModifyByID(id int64, updated *model.EmployeeMDTO) (int64, error) {
	if updated == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := updated.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.ModifyByID(entity, id)
}

func (v *EmployeeBizService) ModifyByIDWithoutZeroFields(id int64, updated *model.EmployeeMDTO) (int64, error) {
	if updated == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := updated.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.ModifyByIDWithoutZeroFields(entity, id)
}

func (v *EmployeeBizService) ModifyByIDWithMap(id int64, updated map[string]any) (int64, error) {
	if len(updated) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.ModifyByIDWithMap(updated, id)
}

func (v *EmployeeBizService) ModifyByCond(condition model.EmployeeQDTO, updated *model.EmployeeMDTO) (int64, error) {
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

func (v *EmployeeBizService) ModifyByMap(updated, condition map[string]any) (int64, error) {
	if len(updated) == 0 || len(condition) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.ModifyByMap(updated, condition)
}

func (v *EmployeeBizService) RemoveByID(id int64) (int64, error) {
	return v.repo.RemoveByID(id)
}

func (v *EmployeeBizService) RemoveByIDs(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	removeIDs := make([]any, len(ids))
	for index, id := range ids {
		removeIDs[index] = id
	}
	return v.repo.RemoveByIDs(removeIDs)
}

func (v *EmployeeBizService) RemoveByCond(condition model.EmployeeQDTO) (int64, error) {
	entity, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.RemoveByCond(*entity)
}

func (v *EmployeeBizService) RemoveByMap(condition map[string]any) (int64, error) {
	if len(condition) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.RemoveByMap(condition)
}
