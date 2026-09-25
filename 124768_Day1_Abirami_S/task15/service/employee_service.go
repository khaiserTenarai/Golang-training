package service

import "task15/task15/dao"

func AddEmployee(id int, name string, salary int) {
	dao.AddEmployee(id, name, salary)
}
func SearchEmployee(id int) bool {
	return dao.SearchEmployee(id)
}
func DisplayEmployees() {
	dao.DisplayEmployees()
}
func DeleteEmployee(id int) bool {
	return dao.DeleteEmployee(id)
}
