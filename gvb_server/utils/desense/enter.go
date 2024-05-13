package desense

//脱敏工具包
import "strings"

// DesensitizationTel 脱敏手机号
func DesensitizationTel(tel string) string {
	//脱敏手机号
	if len(tel) == 11 {
		return tel[:3] + "****" + tel[7:]
	}
	return tel
}

// 脱敏邮箱
func DesensitizationEmail(email string) string {
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}
