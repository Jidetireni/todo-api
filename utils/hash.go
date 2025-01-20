package utils

import "golang.org/x/crypto/bcrypt"

func HashPasswd(passwd string) (string, error) {
	byt, err := bcrypt.GenerateFromPassword([]byte(passwd), 14)
	return string(byt), err
}

func CheckPasswdHash(passwd, hashedPasswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(passwd))
	return err == nil
}
