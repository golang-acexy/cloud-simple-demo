package adm

import (
	"strconv"

	"github.com/golang-acexy/starter-gin/ginstarter"
)

// CountByDepartment 返回指定部门的员工数量，演示如何扩展自动生成的 Router。
func (r *EmployeeRouter) CountByDepartment() ginstarter.HandlerWrapper {
	return func(request *ginstarter.Request) (ginstarter.Response, error) {
		departmentID, err := strconv.ParseInt(request.GetPathParam("departmentId"), 10, 64)
		if err != nil || departmentID <= 0 {
			return ginstarter.RespRestBadParameters(), nil
		}
		count, err := r.bizService.CountByDepartmentID(departmentID)
		if err != nil {
			return nil, err
		}
		return ginstarter.RespRestSuccess(map[string]int64{"count": count}), nil
	}
}
