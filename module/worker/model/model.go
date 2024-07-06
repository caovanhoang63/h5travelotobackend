package workermodel

import (
	"errors"
	"h5travelotobackend/common"
	"time"
)

const EntityName = "Worker"

type Worker struct {
	UserId    int                `json:"-" gorm:"column:user_id;"`
	User      *common.SimpleUser `json:"user" gorm:"foreignKey:UserId;preload=false"`
	HotelId   int                `json:"hotel_id" gorm:"column:hotel_id;"`
	Status    int                `json:"status" gorm:"column:status;"`
	UpdatedAt *time.Time         `json:"updated_at" gorm:"column:updated_at;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
}

func (Worker) TableName() string {
	return "workers"
}
func (w *Worker) GetUserId() int {
	return w.UserId
}
func (w *Worker) GetHotelId() int {
	return w.HotelId
}

type WorkerCreate struct {
	UserFakeId common.UID `json:"user_id" gorm:"-"`
	UserID     int        `json:"-" gorm:"column:user_id;"`
	HotelId    int        `json:"-" gorm:"column:hotel_id;"`
}

func (WorkerCreate) TableName() string {
	return Worker{}.TableName()
}

func (w *WorkerCreate) UnMask() {
	w.UserID = int(w.UserFakeId.GetLocalID())
}

var (
	ErrWorkerAlreadyExist = common.NewCustomError(
		errors.New("worker already existed"),
		"worker already exist",
		"ERR_WORKER_ALREADY_EXIST")
	ErrNotHotelWorker = common.NewCustomError(
		errors.New("user is not hotel worker"),
		"user is not hotel worker",
		"ERR_NOT_HOTEL_WORKER")

	ErrWorkerNotExist = common.NewCustomError(
		errors.New("worker not existed"),
		"worker not exist",
		"ERR_WORKER_NOT_EXIST")
)
