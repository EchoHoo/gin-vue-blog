package pwd

import "golang.org/x/crypto/bcrypt"

//加密密码
func HashPwd(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

//验证密码
func CheckPwd(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	} else {
		return true
	}

}
