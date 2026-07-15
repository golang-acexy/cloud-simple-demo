package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
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

func (v *EmployeeBizService) Save(save *model.EmployeeSDTO) (int64, error) {
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
	db, err := v.repo.TableGORMDB()
	if err != nil {
		return nil, err
	}
	err = db.Where(condition).Order(v.DefaultOrderBy()).Limit(v.MaxQuerySize()).Scan(&entities).Error
	if err != nil {
		return nil, err
	}
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) BaseQueryPage(condition map[string]any, pager *webcloud.Pager[model.EmployeeDTO]) error {
	countDB, err := v.repo.TableGORMDB()
	if err != nil {
		return err
	}
	if err = countDB.Where(condition).Count(&pager.Total).Error; err != nil || pager.Total == 0 {
		return err
	}
	entities := make([]*model.Employee, 0, pager.Size)
	pageDB, err := v.repo.TableGORMDB()
	if err != nil {
		return err
	}
	if err = pageDB.Where(condition).
		Order(v.DefaultOrderBy()).
		Limit(pager.Size).
		Offset((pager.Number - 1) * pager.Size).
		Scan(&entities).Error; err != nil {
		return err
	}
	records, err := model.EmployeeSlice(entities).ToDTOs()
	if err != nil {
		return err
	}
	pager.Records = records
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

func (v *EmployeeBizService) QueryOneByCond(condition *model.EmployeeQDTO) (*model.EmployeeDTO, error) {
	if condition == nil {
		return nil, webcloud.ErrBadRequestParameters
	}
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	var entity model.Employee
	row, err := v.repo.QueryOneByCond(entityCondition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *EmployeeBizService) QueryByCond(condition *model.EmployeeQDTO) ([]*model.EmployeeDTO, error) {
	if condition == nil {
		return nil, webcloud.ErrBadRequestParameters
	}
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	entities := make([]*model.Employee, 0)
	if _, err = v.repo.QueryByCond(entityCondition, v.DefaultOrderBy(), &entities); err != nil {
		return nil, err
	}
	return model.EmployeeSlice(entities).ToDTOs()
}

func (v *EmployeeBizService) QueryPage(pager webcloud.PagerDTO[model.EmployeeQDTO]) (webcloud.Pager[model.EmployeeDTO], error) {
	condition, err := pager.Condition.ToT()
	if err != nil {
		return webcloud.Pager[model.EmployeeDTO]{}, err
	}
	page := databasecloud.Pager[model.Employee]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	if err = v.repo.QueryPageByCond(condition, rds.PageQuery{OrderBySQL: v.DefaultOrderBy()}, &page); err != nil {
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

func (v *EmployeeBizService) RemoveByCond(condition *model.EmployeeQDTO) (int64, error) {
	if condition == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.RemoveByCond(entity)
}

func (v *EmployeeBizService) RemoveByMap(condition map[string]any) (int64, error) {
	if len(condition) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.RemoveByMap(condition)
}
