package auth

type Authorizer interface {
	Creating() (string, error)
	Parsing() (int, error)
}

type authorizeImpl struct{}

func (auth *authorizeImpl) Obtain(a Authorizer) (string, error) {
	return a.Creating()
}

func (auth *authorizeImpl) Authenticate(a Authorizer) (int, error) {
	return a.Parsing()
}

var AuthorizeImpl = new(authorizeImpl)
