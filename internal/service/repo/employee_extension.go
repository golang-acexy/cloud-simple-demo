package repo

// CountByDepartmentID 统计指定部门的员工数量，演示如何扩展自动生成的 Repository。
func (r EmployeeRepo) CountByDepartmentID(departmentID int64) (int64, error) {
	return r.CountByMap(map[string]any{"department_id": departmentID})
}
