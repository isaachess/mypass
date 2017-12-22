package data

import "encoding/base64"

type EncodableBytes struct {
	encoder *base64.Encoding
	b       []byte
}

func NewEncodableBytes(b []byte, enc *base64.Encoding) *EncodableBytes {
	return &EncodableBytes{encoder: enc, b: b}
}

func (e *EncodableBytes) String() string {
	return e.encoder.EncodeToString(e.b)
}

func (e *EncodableBytes) Bytes() []byte {
	return e.b
}
