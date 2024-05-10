package common

import "time"

type SqlModel struct {
	Id        int        `json:"-" gorm:"column:id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

type SqlModelNoMask struct {
	Id        int        `json:"id" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (m *SqlModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 0)
	m.FakeId = &uid
}
