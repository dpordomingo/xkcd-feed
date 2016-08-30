package interfaces

type Connectable interface {
	Retrieve(id interface{}) (string, error)
	Hydrate(input string) (Postable, error)
}
