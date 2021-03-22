package main

import "fmt"

func main() {

	sha := &SHA{}
	md5 := &MD5{}

	password := NewPasswordProtector("user", "password", sha)
	password.Hash()
	password.SetHashAlogrithm(md5)
	password.Hash()

}

/***************************************************************
*                          INTERFACES                          *
****************************************************************/
type HashAlogrithm interface {
	Hash(p *PasswordProtector)
}

/***************************************************************
*                          STRUCTS                             *
****************************************************************/
type PasswordProtector struct {
	user          string
	password      string
	hashAlogrithm HashAlogrithm
}

type SHA struct{}

type MD5 struct{}

/***************************************************************
*                         FUNCTIONS                            *
****************************************************************/
func NewPasswordProtector(user, password string, hash HashAlogrithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		password:      password,
		hashAlogrithm: hash,
	}
}
func (p *PasswordProtector) SetHashAlogrithm(hash HashAlogrithm) {
	p.hashAlogrithm = hash
}
func (p *PasswordProtector) Hash() {
	p.hashAlogrithm.Hash(p)
}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for: %s\n", p.password)
}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for: %s\n", p.password)
}
