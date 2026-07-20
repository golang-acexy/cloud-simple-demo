package biz

import (
	"github.com/golang-acexy/cloud-database/databasecloud"
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
	var entity model.Department
	row, err := v.repo.QueryOneByMap(condition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) BaseQueryOne(condition map[string]any) (*model.DepartmentDTO, error) {
	var entity model.Department
	row, err := v.repo.QueryOneByMap(condition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) BaseQuery(condition map[string]any) ([]*model.DepartmentDTO, error) {
	entities := make([]*model.Department, 0)
	db := v.repo.TableGormDB()
	err := db.Where(condition).Order(v.DefaultOrderBy()).Limit(v.MaxQuerySize()).Scan(&entities).Error
	if err != nil {
		return nil, err
	}
	return model.DepartmentSlice(entities).ToDTOs()
}

func (v *DepartmentBizService) BaseQueryPage(condition map[string]any, pager *webcloud.Pager[model.DepartmentDTO]) error {
	countDB := v.repo.TableGormDB()
	if err := countDB.Where(condition).Count(&pager.Total).Error; err != nil || pager.Total == 0 {
		return err
	}
	entities := make([]*model.Department, 0, pager.Size)
	pageDB := v.repo.TableGormDB()
	if err := pageDB.Where(condition).
		Order(v.DefaultOrderBy()).
		Limit(pager.Size).
		Offset((pager.Number - 1) * pager.Size).
		Scan(&entities).Error; err != nil {
		return err
	}
	records, err := model.DepartmentSlice(entities).ToDTOs()
	if err != nil {
		return err
	}
	pager.Records = records
	return nil
}

func (v *DepartmentBizService) BaseModifyByID(update, condition map[string]any) (int64, error) {
	return v.repo.ModifyByMap(update, condition)
}

func (v *DepartmentBizService) BaseRemoveByID(condition map[string]any) (int64, error) {
	return v.repo.RemoveByMap(condition)
}

func (v *DepartmentBizService) QueryByID(id int64) (*model.DepartmentDTO, error) {
	var entity model.Department
	row, err := v.repo.QueryByID(id, &entity)
	if err != nil || row == 0 {
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
	entities := make([]*model.Department, 0, len(ids))
	if _, err := v.repo.QueryByIDs(queryIDs, &entities); err != nil {
		return nil, err
	}
	return model.DepartmentSlice(entities).ToDTOs()
}

func (v *DepartmentBizService) ExistsByID(id int64) (bool, error) {
	return v.repo.ExistsByID(id)
}

func (v *DepartmentBizService) QueryOneByCond(condition *model.DepartmentQDTO) (*model.DepartmentDTO, error) {
	if condition == nil {
		return nil, webcloud.ErrBadRequestParameters
	}
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	var entity model.Department
	row, err := v.repo.QueryOneByCond(entityCondition, &entity)
	if err != nil || row == 0 {
		return nil, err
	}
	return entity.ToDTO()
}

func (v *DepartmentBizService) QueryByCond(condition *model.DepartmentQDTO) ([]*model.DepartmentDTO, error) {
	if condition == nil {
		return nil, webcloud.ErrBadRequestParameters
	}
	entityCondition, err := condition.ToT()
	if err != nil {
		return nil, err
	}
	entities := make([]*model.Department, 0)
	if _, err = v.repo.QueryByCond(entityCondition, v.DefaultOrderBy(), &entities); err != nil {
		return nil, err
	}
	return model.DepartmentSlice(entities).ToDTOs()
}

func (v *DepartmentBizService) CountByCond(condition *model.DepartmentQDTO) (int64, error) {
	if condition == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entityCondition, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.CountByCond(entityCondition)
}

func (v *DepartmentBizService) QueryPage(pager webcloud.PagerDTO[model.DepartmentQDTO]) (webcloud.Pager[model.DepartmentDTO], error) {
	condition, err := pager.Condition.ToT()
	if err != nil {
		return webcloud.Pager[model.DepartmentDTO]{}, err
	}
	page := databasecloud.Pager[model.Department]{
		Number: pager.Number,
		Size:   pager.Size,
	}
	if err = v.repo.QueryPageByCond(condition, rds.PageQuery{OrderBySQL: v.DefaultOrderBy()}, &page); err != nil {
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
	entity.ID = id
	return v.repo.ModifyByID(entity)
}

func (v *DepartmentBizService) ModifyByIDWithoutZeroFields(id int64, updated *model.DepartmentMDTO) (int64, error) {
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

func (v *DepartmentBizService) ModifyByIDWithMap(id int64, updated map[string]any) (int64, error) {
	if len(updated) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.ModifyByIDWithMap(updated, id)
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

func (v *DepartmentBizService) RemoveByCond(condition *model.DepartmentQDTO) (int64, error) {
	if condition == nil {
		return 0, webcloud.ErrBadRequestParameters
	}
	entity, err := condition.ToT()
	if err != nil {
		return 0, err
	}
	return v.repo.RemoveByCond(entity)
}

func (v *DepartmentBizService) RemoveByMap(condition map[string]any) (int64, error) {
	if len(condition) == 0 {
		return 0, webcloud.ErrBadRequestParameters
	}
	return v.repo.RemoveByMap(condition)
}
