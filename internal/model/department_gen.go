package model

import (
	"github.com/acexy/golang-toolkit/util/json"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"github.com/jinzhu/copier"
)

const TableNameDepartment = "demo_department"

// Department 框架演示部门表
type Department struct {
	ID          int64                 `gorm:"primaryKey;<-:false" json:"id"` // 主键
	Code        string                `json:"code"`                          // 部门编码
	Name        string                `json:"name"`                          // 部门名称
	Description string                `json:"description"`                   // 部门描述
	Status      int32                 `json:"status"`                        // 状态：1启用，0停用
	CreatedAt   gormstarter.Timestamp `gorm:"<-:false" json:"createdAt"`     // 创建时间
	UpdatedAt   gormstarter.Timestamp `gorm:"<-:false" json:"updatedAt"`     // 更新时间
}

// DepartmentColumns 定义 Department Mapper 可使用的类型化字段。
type DepartmentColumns struct {
	ID          gormstarter.Column[Department, int64]
	Code        gormstarter.Column[Department, string]
	Name        gormstarter.Column[Department, string]
	Description gormstarter.Column[Department, string]
	Status      gormstarter.Column[Department, int32]
	CreatedAt   gormstarter.Column[Department, gormstarter.Timestamp]
	UpdatedAt   gormstarter.Column[Department, gormstarter.Timestamp]
}

var departmentColumns = DepartmentColumns{
	ID:          gormstarter.NewColumn(func(entity *Department) *int64 { return &entity.ID }),
	Code:        gormstarter.NewColumn(func(entity *Department) *string { return &entity.Code }),
	Name:        gormstarter.NewColumn(func(entity *Department) *string { return &entity.Name }),
	Description: gormstarter.NewColumn(func(entity *Department) *string { return &entity.Description }),
	Status:      gormstarter.NewColumn(func(entity *Department) *int32 { return &entity.Status }),
	CreatedAt:   gormstarter.NewColumn(func(entity *Department) *gormstarter.Timestamp { return &entity.CreatedAt }),
	UpdatedAt:   gormstarter.NewColumn(func(entity *Department) *gormstarter.Timestamp { return &entity.UpdatedAt }),
}

// DepartmentColumnSet 返回 Department 的只读字段元数据。
func DepartmentColumnSet() *DepartmentColumns {
	return &departmentColumns
}

func (Department) TableName() string {
	return TableNameDepartment
}
func (Department) DBType() gormstarter.DBType {
	return gormstarter.DBTypeMySQL
}

// DepartmentSDTO 定义保存操作允许接收的字段。
type DepartmentSDTO struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
}

// DepartmentMDTO 定义修改操作允许接收的字段。
type DepartmentMDTO struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
}

// DepartmentQDTO 定义查询操作允许接收的字段。
type DepartmentQDTO struct {
	ID          int64          `json:"id"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      int32          `json:"status"`
	CreatedAt   json.Timestamp `json:"createdAt"`
	UpdatedAt   json.Timestamp `json:"updatedAt"`
}

// DepartmentDTO 定义允许返回的字段。
type DepartmentDTO struct {
	ID          int64          `json:"id"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      int32          `json:"status"`
	CreatedAt   json.Timestamp `json:"createdAt"`
	UpdatedAt   json.Timestamp `json:"updatedAt"`
}

func (v Department) ToDTO() (*DepartmentDTO, error) {
	var dto DepartmentDTO
	if err := copier.Copy(&dto, v); err != nil {
		return nil, err
	}
	return &dto, nil
}

func (v Department) ParseDTO(dto *DepartmentDTO) error {
	return copier.Copy(dto, v)
}

type DepartmentSlice []*Department

func (v DepartmentSlice) ToDTOs() ([]*DepartmentDTO, error) {
	dtos := make([]*DepartmentDTO, 0, len(v))
	if err := copier.Copy(&dtos, v); err != nil {
		return nil, err
	}
	return dtos, nil
}

func (v DepartmentSlice) ParseDTOs(dtos *[]*DepartmentDTO) error {
	return copier.Copy(dtos, v)
}

func (v DepartmentSDTO) ToT() (*Department, error) {
	var entity Department
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v DepartmentSDTO) ParseT(entity *Department) error {
	return copier.Copy(entity, v)
}

func (v DepartmentMDTO) ToT() (*Department, error) {
	var entity Department
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v DepartmentMDTO) ParseT(entity *Department) error {
	return copier.Copy(entity, v)
}

func (v DepartmentQDTO) ToT() (*Department, error) {
	var entity Department
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v DepartmentQDTO) ParseT(entity *Department) error {
	return copier.Copy(entity, v)
}

func (v DepartmentDTO) ToT() (*Department, error) {
	var entity Department
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v DepartmentDTO) ParseT(entity *Department) error {
	return copier.Copy(entity, v)
}
