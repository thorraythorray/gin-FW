package schema

type Former interface {
	Validate() (bool, error)
}
