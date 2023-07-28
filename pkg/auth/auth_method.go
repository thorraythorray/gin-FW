package auth

type AuthorizeTokner interface {
	Obtaining() (string, error)
	Authenticating() (int, error)
}

type authorizeImpl struct{}

func (auth *authorizeImpl) Obtain(a AuthorizeTokner) (string, error) {
	return a.Obtaining()
}

func (auth *authorizeImpl) Authenticate(a AuthorizeTokner) (int, error) {
	return a.Authenticating()
}

var AuthorizeImpl = new(authorizeImpl)
