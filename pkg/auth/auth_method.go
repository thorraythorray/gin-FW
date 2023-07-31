package auth

type AuthorizeTokner interface {
	Obtaining(string) (string, error)
	Authenticating(string) (int, error)
}

type authorizeImpl struct{}

func (auth *authorizeImpl) Obtain(a AuthorizeTokner, u string) (string, error) {
	return a.Obtaining(u)
}

func (auth *authorizeImpl) Authenticate(a AuthorizeTokner, j string) (int, error) {
	return a.Authenticating(j)
}

var AuthorizeImpl = new(authorizeImpl)
