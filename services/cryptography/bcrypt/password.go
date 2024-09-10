package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//generate hash from password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	//check password with hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
