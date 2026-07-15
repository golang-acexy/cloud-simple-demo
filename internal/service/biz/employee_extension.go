package biz

// CountByDepartmentID 统计指定部门的员工数量。
func (v *EmployeeBizService) CountByDepartmentID(departmentID int64) (int64, error) {
	return v.repo.CountByDepartmentID(departmentID)
}
