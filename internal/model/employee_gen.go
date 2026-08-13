package model

import (
	"github.com/acexy/golang-toolkit/util/json"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"github.com/jinzhu/copier"
)

const TableNameEmployee = "demo_employee"

// Employee 框架演示员工表
type Employee struct {
	ID           int64                 `gorm:"primaryKey;<-:false" json:"id"` // 主键
	DepartmentID int64                 `json:"departmentId"`                  // 所属部门ID
	UserID       int64                 `json:"userId"`                        // 数据所属用户ID
	EmployeeNo   string                `json:"employeeNo"`                    // 员工编号
	Name         string                `json:"name"`                          // 员工姓名
	Email        string                `json:"email"`                         // 邮箱
	Status       int32                 `json:"status"`                        // 状态：1在职，0离职
	CreatedAt    gormstarter.Timestamp `gorm:"<-:false" json:"createdAt"`     // 创建时间
	UpdatedAt    gormstarter.Timestamp `gorm:"<-:false" json:"updatedAt"`     // 更新时间
}

// EmployeeColumns 定义 Employee Mapper 可使用的类型化字段。
type EmployeeColumns struct {
	ID           gormstarter.Column[Employee, int64]
	DepartmentID gormstarter.Column[Employee, int64]
	UserID       gormstarter.Column[Employee, int64]
	EmployeeNo   gormstarter.Column[Employee, string]
	Name         gormstarter.Column[Employee, string]
	Email        gormstarter.Column[Employee, string]
	Status       gormstarter.Column[Employee, int32]
	CreatedAt    gormstarter.Column[Employee, gormstarter.Timestamp]
	UpdatedAt    gormstarter.Column[Employee, gormstarter.Timestamp]
}

var employeeColumns = EmployeeColumns{
	ID:           gormstarter.NewColumn(func(entity *Employee) *int64 { return &entity.ID }),
	DepartmentID: gormstarter.NewColumn(func(entity *Employee) *int64 { return &entity.DepartmentID }),
	UserID:       gormstarter.NewColumn(func(entity *Employee) *int64 { return &entity.UserID }),
	EmployeeNo:   gormstarter.NewColumn(func(entity *Employee) *string { return &entity.EmployeeNo }),
	Name:         gormstarter.NewColumn(func(entity *Employee) *string { return &entity.Name }),
	Email:        gormstarter.NewColumn(func(entity *Employee) *string { return &entity.Email }),
	Status:       gormstarter.NewColumn(func(entity *Employee) *int32 { return &entity.Status }),
	CreatedAt:    gormstarter.NewColumn(func(entity *Employee) *gormstarter.Timestamp { return &entity.CreatedAt }),
	UpdatedAt:    gormstarter.NewColumn(func(entity *Employee) *gormstarter.Timestamp { return &entity.UpdatedAt }),
}

// EmployeeColumnSet 返回 Employee 的只读字段元数据。
func EmployeeColumnSet() *EmployeeColumns {
	return &employeeColumns
}

func (Employee) TableName() string {
	return TableNameEmployee
}
func (Employee) DBType() gormstarter.DBType {
	return gormstarter.DBTypeMySQL
}

// EmployeeSDTO 定义保存操作允许接收的字段。
type EmployeeSDTO struct {
	DepartmentID int64  `json:"departmentId"`
	UserID       int64  `json:"userId"`
	EmployeeNo   string `json:"employeeNo"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       int32  `json:"status"`
}

// EmployeeMDTO 定义修改操作允许接收的字段。
type EmployeeMDTO struct {
	DepartmentID int64  `json:"departmentId"`
	UserID       int64  `json:"userId"`
	EmployeeNo   string `json:"employeeNo"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       int32  `json:"status"`
}

// EmployeeQDTO 定义查询操作允许接收的字段。
type EmployeeQDTO struct {
	ID           int64          `json:"id"`
	DepartmentID int64          `json:"departmentId"`
	UserID       int64          `json:"userId"`
	EmployeeNo   string         `json:"employeeNo"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Status       int32          `json:"status"`
	CreatedAt    json.Timestamp `json:"createdAt"`
	UpdatedAt    json.Timestamp `json:"updatedAt"`
}

// EmployeeDTO 定义允许返回的字段。
type EmployeeDTO struct {
	ID           int64          `json:"id"`
	DepartmentID int64          `json:"departmentId"`
	UserID       int64          `json:"userId"`
	EmployeeNo   string         `json:"employeeNo"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Status       int32          `json:"status"`
	CreatedAt    json.Timestamp `json:"createdAt"`
	UpdatedAt    json.Timestamp `json:"updatedAt"`
}

func (v Employee) ToDTO() (*EmployeeDTO, error) {
	var dto EmployeeDTO
	if err := copier.Copy(&dto, v); err != nil {
		return nil, err
	}
	return &dto, nil
}

func (v Employee) ParseDTO(dto *EmployeeDTO) error {
	return copier.Copy(dto, v)
}

type EmployeeSlice []*Employee

func (v EmployeeSlice) ToDTOs() ([]*EmployeeDTO, error) {
	dtos := make([]*EmployeeDTO, 0, len(v))
	if err := copier.Copy(&dtos, v); err != nil {
		return nil, err
	}
	return dtos, nil
}

func (v EmployeeSlice) ParseDTOs(dtos *[]*EmployeeDTO) error {
	return copier.Copy(dtos, v)
}

func (v EmployeeSDTO) ToT() (*Employee, error) {
	var entity Employee
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v EmployeeSDTO) ParseT(entity *Employee) error {
	return copier.Copy(entity, v)
}

func (v EmployeeMDTO) ToT() (*Employee, error) {
	var entity Employee
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v EmployeeMDTO) ParseT(entity *Employee) error {
	return copier.Copy(entity, v)
}

func (v EmployeeQDTO) ToT() (*Employee, error) {
	var entity Employee
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v EmployeeQDTO) ParseT(entity *Employee) error {
	return copier.Copy(entity, v)
}

func (v EmployeeDTO) ToT() (*Employee, error) {
	var entity Employee
	if err := copier.Copy(&entity, v); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (v EmployeeDTO) ParseT(entity *Employee) error {
	return copier.Copy(entity, v)
}
