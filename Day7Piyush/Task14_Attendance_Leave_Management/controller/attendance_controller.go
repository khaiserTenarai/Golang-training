package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task14_attendance_leave_management/models"
	"task14_attendance_leave_management/repository"

	"github.com/gorilla/mux"
)

type AttendanceController struct {
	attRepo   *repository.AttendanceRepository
	leaveRepo *repository.LeaveRepository
}

func NewAttendanceController(ar *repository.AttendanceRepository, lr *repository.LeaveRepository) *AttendanceController {
	return &AttendanceController{attRepo: ar, leaveRepo: lr}
}

func (c *AttendanceController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	created, err := c.attRepo.CreateEmployee(emp)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "Employee created", Data: created})
}

func (c *AttendanceController) ListEmployees(w http.ResponseWriter, r *http.Request) {
	emps, err := c.attRepo.GetAllEmployees()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error fetching employees"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Success", Data: emps})
}

func (c *AttendanceController) CheckIn(w http.ResponseWriter, r *http.Request) {
	empID, _ := strconv.Atoi(mux.Vars(r)["id"])
	att, err := c.attRepo.CheckIn(empID)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Check-in failed: " + err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Checked in", Data: att})
}

func (c *AttendanceController) CheckOut(w http.ResponseWriter, r *http.Request) {
	empID, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.attRepo.CheckOut(empID); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Check-out failed: " + err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Checked out"})
}

func (c *AttendanceController) AttendanceReport(w http.ResponseWriter, r *http.Request) {
	empID, _ := strconv.Atoi(mux.Vars(r)["id"])
	records, err := c.attRepo.GetAttendanceReport(empID)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Attendance report", Data: records})
}

func (c *AttendanceController) ApplyLeave(w http.ResponseWriter, r *http.Request) {
	var leave models.LeaveRequest
	json.NewDecoder(r.Body).Decode(&leave)
	created, err := c.leaveRepo.ApplyLeave(leave)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Leave application failed: " + err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "Leave applied", Data: created})
}

func (c *AttendanceController) MyLeaves(w http.ResponseWriter, r *http.Request) {
	empID, _ := strconv.Atoi(mux.Vars(r)["id"])
	leaves, err := c.leaveRepo.GetLeavesByEmployee(empID)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Leave records", Data: leaves})
}

func (c *AttendanceController) PendingLeaves(w http.ResponseWriter, r *http.Request) {
	leaves, err := c.leaveRepo.GetPendingLeaves()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Pending leaves", Data: leaves})
}

func (c *AttendanceController) ApproveLeave(w http.ResponseWriter, r *http.Request) {
	leaveID, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.leaveRepo.ApproveLeave(leaveID, "Admin"); err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Approval failed"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Leave approved"})
}

func (c *AttendanceController) RejectLeave(w http.ResponseWriter, r *http.Request) {
	leaveID, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.leaveRepo.RejectLeave(leaveID, "Admin"); err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Rejection failed"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Leave rejected"})
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
