package workermodel

const EntityName = "Worker"

type Worker struct {
	UserId  int    `json:"user_id" gorm:"column:user_id;"`
	HotelId int    `json:"hotel_id" gorm:"column:hotel_id;"`
	Role    string `json:"role" gorm:"column:role;"`
	Status  int    `json:"status" gorm:"column:status;"`
}

func (Worker) TableName() string {
	return "workers"
}
