package internal

const (
	Active = iota
	Disable
	Deleted
)

const (
	JwtSignKey            = "Can i c u"
	JwtExpireHour         = 2
	JwtExemptRouterString = "/v1/user/register,/v1/login,/v1/token/obtain"
)
