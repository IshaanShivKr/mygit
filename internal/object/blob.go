package object

import "fmt"

type Blob struct {
	Data []byte
}

func NewBlob(data []byte) *Blob {
	return &Blob{
		Data: data,
	}
}

func (b *Blob) Type() string {
	return "blob"
}

func (b *Blob) Serialize() ([]byte, error) {
	header := fmt.Sprintf("blob %d\x00", len(b.Data))
	serialized := append([]byte(header), b.Data...)
	return serialized, nil
}