package util

// CreateToken 创建token
// @uid 用户id
func CreateToken(uid int64) string {

	claims := Claims{
		Uid: uid,
	}

	token, err := JwtEncode(claims, 7200, nil)
	if err != nil {
		return ""
	}
	return token
}

// VerifyToken 验证token
func VerifyToken(token string) (int64, error) {

	claims, err := JwtDecode(token, nil)
	if err != nil {
		return 0, err
	}
	return claims.Uid, nil
}
