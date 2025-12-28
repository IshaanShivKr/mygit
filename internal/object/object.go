package object

type Object interface {
	Type() string
	Serialize() ([]byte, error)
}

type Header struct {
    Type string
    Size int
}
