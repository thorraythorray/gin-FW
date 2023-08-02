package schema

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id" binding:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:create_at" json:"create_at" binding:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:update_at" json:"update_at" binding:"-"`
}

type Pagination struct {
	Page    int `form:"page" binding:"required" validate:"gte=1"`
	PerPage int `form:"per_page" binding:"required" validate:"gte=0,lte=50"`
}

/**
* @description:
* @return {offset limit}
 */
func (p *Pagination) PageInfo() (int, int) {
	offset := (p.Page - 1) * p.PerPage
	limit := p.PerPage
	return offset, limit
}

func (p *Pagination) ResponseInfo(items interface{}, total int64) interface{} {
	res := make(map[string]interface{})
	res["page"] = p.Page
	res["per_page"] = p.PerPage
	res["items"] = items
	res["total"] = total
	return res
}
