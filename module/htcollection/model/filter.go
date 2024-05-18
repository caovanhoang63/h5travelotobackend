package htcollection

import "h5travelotobackend/common"

type Filter struct {
	UserId     int         `json:"-" form:"-"`
	UserFakeId *common.UID `json:"user_id" form:"user_id" binding:"required"`
}

func (f *Filter) Mask() {
	f.UserFakeId = common.NewUIDP(uint32(f.UserId), common.DbTypeUser, 0)
}

func (f *Filter) UnMask() error {
	if f.UserFakeId != nil {
		f.UserId = int(f.UserFakeId.GetLocalID())
	}
	return nil
}
