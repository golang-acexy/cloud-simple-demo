package model

import (
	"github.com/acexy/golang-toolkit/logger"
	"github.com/acexy/golang-toolkit/util/json"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"github.com/jinzhu/copier"
)

const TableNameTeacher = "demo_teacher"

// Teacher 教师表
type Teacher struct {
	ID         int64                 `gorm:"primaryKey;<-:false" json:"id"`
	CreateTime gormstarter.Timestamp `gorm:"<-:false" json:"createTime"` // 创建时间
	UpdateTime gormstarter.Timestamp `gorm:"<-:false" json:"updateTime"` // 更新时间
	Name       string                `json:"name"`                       // 姓名
	Sex        string                `json:"sex"`                        // 性别
	Age        int32                 `json:"age"`                        // 年龄
	ClassNo    string                `json:"classNo"`                    // 班级编号
}

func (Teacher) TableName() string {
	return TableNameTeacher
}
func (Teacher) DBType() gormstarter.DBType {
	return gormstarter.DBTypeMySQL
}

// TeacherSDTO 保存时结构体 限定只允许操作的字段
type TeacherSDTO struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int32  `json:"age"`
	ClassNo string `json:"classNo"`
}

// TeacherMDTO 修改时结构体 限定只允许操作的字段
type TeacherMDTO struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int32  `json:"age"`
	ClassNo string `json:"classNo"`
}

// TeacherQDTO 查询时结构体 限定只允许使用的字段
type TeacherQDTO struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int32  `json:"age"`
	ClassNo string `json:"classNo"`
}

// TeacherDTO 结果结构体 限定只允许展示的字段
type TeacherDTO struct {
	ID         int64          `json:"id"`
	CreateTime json.Timestamp `json:"createTime"`
	UpdateTime json.Timestamp `json:"updateTime"`
	Name       string         `json:"name"`
	Sex        string         `json:"sex"`
	Age        int32          `json:"age"`
	ClassNo    string         `json:"classNo"`
}

func (v Teacher) ToDTO() *TeacherDTO {
	var dto TeacherDTO
	err := copier.Copy(&dto, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
		return nil
	}
	return &dto
}

func (v Teacher) ParseDTO(dto *TeacherDTO) {
	err := copier.Copy(dto, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
	}
}

type TeacherSlice []*Teacher

func (v TeacherSlice) ToDTOs() []*TeacherDTO {
	var dtos []*TeacherDTO
	err := copier.Copy(&dtos, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
		return nil
	}
	return dtos
}

func (v TeacherSlice) ParseDTOs(dtos *[]*TeacherDTO) {
	err := copier.Copy(dtos, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
	}
}

func (v TeacherSDTO) ToT() *Teacher {
	var t Teacher
	err := copier.Copy(&t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
		return nil
	}
	return &t
}

func (v TeacherSDTO) ParseT(t *Teacher) {
	err := copier.Copy(t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
	}
}

func (v TeacherMDTO) ToT() *Teacher {
	var t Teacher
	err := copier.Copy(&t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
		return nil
	}
	return &t
}

func (v TeacherMDTO) ParseT(t *Teacher) {
	err := copier.Copy(t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
	}
}

func (v TeacherQDTO) ToT() *Teacher {
	var t Teacher
	err := copier.Copy(&t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
		return nil
	}
	return &t
}

func (v TeacherQDTO) ParseT(t *Teacher) {
	err := copier.Copy(t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
	}
}

func (v TeacherDTO) ToT() *Teacher {
	var t Teacher
	err := copier.Copy(&t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
		return nil
	}
	return &t
}

func (v TeacherDTO) ParseT(t *Teacher) {
	err := copier.Copy(t, v)
	if err != nil {
		logger.Logrus().Errorln("copier.Copy error:", err)
	}
}
