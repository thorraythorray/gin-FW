package internal

const (
	Active = iota
	Disable
	Deleted
)

const (
	SignKey    = "Can i c u"
	ExpireHour = 2
)

const JwtExemptRouterString = "/v1/user/register,/v1/login,/v1/token/obtain"
