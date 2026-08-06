package biz

import (
	"time"

	"github.com/acexy/golang-toolkit/util/str"
	"github.com/golang-acexy/cloud-database/databasecloud"
	"github.com/golang-acexy/cloud-database/databasecloud/rds"
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	repo "github.com/golang-acexy/cloud-simple-demo/internal/service/repo"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gorm/gormstarter"
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

// convertTimeRanges 校验时间字段白名单，并转换为 GORM 分页查询参数。
func (v *EmployeeBizService) convertTimeRanges(timeRanges []webcloud.TimeRange) ([]gormstarter.TimeRange, error) {
	allowedFields := v.AllowedTimeRangeFields()
	result := make([]gormstarter.TimeRange, 0, len(timeRanges))
	defaultFieldUsed := false
	for _, timeRange := range timeRanges {
		field := str.CamelToSnake(timeRange.Field)
		if field == "" {
			if defaultFieldUsed {
				return nil, webcloud.ErrBadRequestParameters
			}
			defaultFieldUsed = true
			field = v.DefaultTimeRangeField()
		}
		allowed := false
		for _, allowedField := range allowedFields {
			if field == allowedField {
				allowed = true
				break
			}
		}
		if !allowed || timeRange.Start.IsZero() && timeRange.End.IsZero() {
			return nil, webcloud.ErrBadRequestParameters
		}
		if !timeRange.Start.IsZero() && !timeRange.End.IsZero() && !timeRange.Start.Before(timeRange.End.Time) {
			return nil, webcloud.ErrBadRequestParameters
		}
		var startTime, endTime *time.Time
		if !timeRange.Start.IsZero() {
			value := timeRange.Start.Time
			startTime = &value
		}
		if !timeRange.End.IsZero() {
			value := timeRange.End.Time
			endTime = &value
		}
		result = append(result, gormstarter.TimeRange{
			Field:     field,
			StartTime: startTime,
			EndTime:   endTime,
		})
	}
	return result, nil
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
	var entity model.Employee
	row, err := v.repo.QueryOneByMap(condition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) BaseQueryOne(condition map[string]any) (*model.EmployeeDTO, error) {
	var entity model.Employee
	row, err := v.repo.QueryOneByMap(condition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) BaseQuery(condition map[string]any) ([]*model.EmployeeDTO, error) {
	entities := make([]*model.Employee, 0)
	db := v.repo.TableGormDB()
	err := db.Where(condition).Order(v.DefaultOrderBy()).Limit(v.MaxQuerySize()).Scan(&entities).Error
	if err != nil {
		return nil, err
	}
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) BaseQueryPage(condition map[string]any, timeRanges []webcloud.TimeRange, pager *webcloud.Pager[model.EmployeeDTO]) error {
	ranges, err := v.convertTimeRanges(timeRanges)
	if err != nil {
		return err
	}
	page := databasecloud.Pager[model.Employee]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	query := rds.PageQuery{
		OrderBySQL: v.DefaultOrderBy(),
		TimeRanges: ranges,
	}
	if len(condition) == 0 {
		err = v.repo.QueryPageByCond(model.Employee{}, query, &page)
	} else {
		err = v.repo.QueryPageByMap(condition, query, &page)
	}
	if err != nil {
		return err
	}
	records, err := model.EmployeeSlice(page.Records).ToDTOs()
	if err != nil {
		return err
	}
	pager.Records = records
	pager.Total = page.Total
	return nil
}

func (v *EmployeeBizService) BaseModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *EmployeeBizService) BaseRemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}

func (v *EmployeeBizService) QueryByID(id int64) (*model.EmployeeDTO, error) {
	var entity model.Employee
	row, err := v.repo.QueryByID(id, &entity)
	if err != nil || row == 0 {
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
	entities := make([]*model.Employee, 0, len(ids))
	if _, err := v.repo.QueryByIDs(queryIDs, &entities); err != nil {
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
	var entity model.Employee
	row, err := v.repo.QueryOneByCond(*entityCondition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) QueryByCond(condition model.EmployeeQDTO) ([]*model.EmployeeDTO, error) {
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	entities := make([]*model.Employee, 0)
	if _, err = v.repo.QueryByCond(*entityCondition, v.DefaultOrderBy(), &entities); err != nil {
		return nil, err
	}
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) CountByCond(condition model.EmployeeQDTO) (int64, error) {
	entityCondition, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.CountByCond(*entityCondition)
}

func (v *EmployeeBizService) QueryPage(pager webcloud.PagerDTO[model.EmployeeQDTO]) (webcloud.Pager[model.EmployeeDTO], error) {
	condition, err := pager.Condition.ToT()
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	ranges, err := v.convertTimeRanges(pager.TimeRanges)
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	page := databasecloud.Pager[model.Employee]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	if err = v.repo.QueryPageByCond(*condition, rds.PageQuery{
		OrderBySQL: v.DefaultOrderBy(),
		TimeRanges: ranges,
	}, &page); err != nil {
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
	entity.ID = id
	return v.repo.ModifyByID(entity)
}

func (v *EmployeeBizService) ModifyByIDWithoutZeroFields(id int64, updated *model.EmployeeMDTO) (int64, error) {
	if updated == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := updated.ToT()
	if err != nil {
		return 0, err
	}
	entity.ID = id
	return v.repo.ModifyByIDWithoutZeroFields(entity)
}

func (v *EmployeeBizService) ModifyByIDWithMap(id int64, updated map[string]any) (int64, error) {
	if len(updated) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.ModifyByIDWithMap(updated, id)
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
