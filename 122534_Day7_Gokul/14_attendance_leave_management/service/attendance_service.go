package service

import "context"

type AttendanceService interface {
	AddEmployee(context.Context, string) error
	CheckIn(context.Context, int64) error
	CheckOut(context.Context, int64) error
	ShowAttendance(context.Context) error
	ApplyLeave(context.Context, int64, string, string) error
	UpdateLeave(context.Context, int64, string) error
	ShowLeaves(context.Context) error
}
