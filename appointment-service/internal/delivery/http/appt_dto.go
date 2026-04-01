package delivery

import "appointment-service/internal/domain"

type CreateAppointmentRequest struct {
	Title       string `json:"title" binding:"required"`
	DoctorID    string `json:"doctor_id" binding:"required"`
	Description string `json:"description"`
}

type UpdateAppointmentRequest struct {
	Status string `json:"status" binding:"required"`
}

type AppointmentResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	DoctorID    string `json:"doctor_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func ToDomain(req CreateAppointmentRequest) *domain.Appointment {
	return &domain.Appointment{
		Title:       req.Title,
		DoctorID:    req.DoctorID,
		Description: req.Description,
	}
}

func FromDomain(d *domain.Appointment) AppointmentResponse {
	return AppointmentResponse{
		ID:          d.ID,
		Title:       d.Title,
		DoctorID:    d.DoctorID,
		Description: d.Description,
		Status:      string(d.Status),
	}
}
