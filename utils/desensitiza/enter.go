package desensitiza

import "strings"

// DesensitizationTel 手机号脱敏
func DesensitizationTel(tel string) string {
	// 18300001111
	// 183****1111
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}

// DesensitizationEmail 邮箱脱敏
func DesensitizationEmail(email string) string {
	// 不同的邮箱，邮箱账号不一样
	eList := strings.Split(email, "@") // 从@开始分割
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}
