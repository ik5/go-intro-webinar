package crypto

// type of encryption to use
const (
	SCrypt = iota + 1
	BCrypt
	Argon2
	PBKDF2
)

// The default values for generation
const (
	DefaultSaltSize = 16
)
