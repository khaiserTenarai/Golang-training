package service

import (
	"context"
	"errors"
	"example.com/q14-attendance-leave/repository"
)

type AttendanceServiceImpl struct {
	repo repository.AttendanceRepository
}

func NewAttendanceService(r repository.AttendanceRepository) AttendanceService {
	return &AttendanceServiceImpl{repo: r}
}
func (s *AttendanceServiceImpl) AddEmployee(c context.Context, n string) error {
	if n == "" {
		return errors.New("employee name is required")
	}
	return s.repo.AddEmployee(c, n)
}
func (s *AttendanceServiceImpl) CheckIn(c context.Context, id int64) error {
	return s.repo.CheckIn(c, id)
}
func (s *AttendanceServiceImpl) CheckOut(c context.Context, id int64) error {
	return s.repo.CheckOut(c, id)
}
func (s *AttendanceServiceImpl) ShowAttendance(c context.Context) error {
	return s.repo.ShowAttendance(c)
}
func (s *AttendanceServiceImpl) ApplyLeave(c context.Context, id int64, d, r string) error {
	if d == "" || r == "" {
		return errors.New("date and reason are required")
	}
	return s.repo.ApplyLeave(c, id, d, r)
}
func (s *AttendanceServiceImpl) UpdateLeave(c context.Context, id int64, status string) error {
	return s.repo.UpdateLeave(c, id, status)
}
func (s *AttendanceServiceImpl) ShowLeaves(c context.Context) error { return s.repo.ShowLeaves(c) }
