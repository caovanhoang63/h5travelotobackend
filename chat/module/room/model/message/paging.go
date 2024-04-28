package chatmessage

type Paging struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total" form:"total"`

	// Support cursor with UID
	Cursor     int `json:"cursor" form:"cursor"`
	NextCursor int `json:"next_cursor"`
}

func (p *Paging) FullFill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 30
	}
}

func (p *Paging) GetOffSet() int {
	return (p.Page - 1) * p.Limit
}
