package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hast(p *PasswordProtector)
}

// creamos el constructor de PasswordProtector
func NewPassrordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

// vamos hacer que HashAlgorithm sea intercambiable
func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash

}

func (p *PasswordProtector) Hast() {
	p.hashAlgorithm.Hast(p)
}

type SHA struct{}

func (SHA) Hast(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hast(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 fot %s\n", p.passwordName)
}

func main() {

	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPassrordProtector("Nestor", "gmail pasword", sha)
	passwordProtector.Hast()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hast()
}
