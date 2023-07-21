package request

type Former interface {
	Validate() (bool, error)
}
