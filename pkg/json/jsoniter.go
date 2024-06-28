package json

import (
	"io"

	"github.com/goccy/go-json"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}
