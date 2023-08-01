package form

type Pagination struct {
	Page    int `json:"page" default:"0"`
	PerPage int `json:"per_page" default:"10"`
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
	res["page"] = p.Page
	res["per_page"] = p.PerPage
	res["items"] = items
	res["total"] = total
	return res
}
