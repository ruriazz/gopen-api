package constants

type EncryptioMode int

const (
	AES_CBC EncryptioMode = iota
	AES_GCM
)

func (c EncryptioMode) String() string {
	return [...]string{"AES-CBC", "AES-GCM"}[c]
}
