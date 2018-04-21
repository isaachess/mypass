package data

// Secret is to prevent secret bytes (such as an unencryped password)
// from being accidentally printed
type Secret []byte

func NewSecret(b []byte) Secret {
	return Secret(b)
}

func (s Secret) String() string {
	return "< unprintable >"
}

func (s Secret) DangerousString() string {
	return string(s)
}
