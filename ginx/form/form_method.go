package form

type FormMethod interface {
	Validate() error
}
