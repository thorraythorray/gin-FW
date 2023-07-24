package form

type Pagination struct {
	Page    int `form:"page" binding:"required" json:"page" default:"0"`
	PerPage int `form:"per_page" binding:"required" json:"per_page" default:"10"`
}

/**
* @description:
* @return {offset limit}
 */
func (p *Pagination) PageInfo() (int, int) {
	offset := p.Page * p.PerPage
	limit := p.PerPage
	return offset, limit
}

func (p *Pagination) ResponseInfo(items interface{}, total uint64) interface{} {
	res := make(map[string]interface{})
	res["Page"] = p.Page
	res["PerPage"] = p.PerPage
	res["Items"] = items
	res["Total"] = total
	return res
}
