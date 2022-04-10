package jike

const AccessToken = "AccessToken"
const RefreshToken = "RefreshToken"

var tokenMap = make(map[string]string)

func getAccessToken(areaCode, phone string) string {
	return tokenMap[AccessToken+"_"+areaCode+"_"+phone]
}

func setAccessToken(areaCode, phone, accessToken string) error {
	tokenMap[AccessToken+"_"+areaCode+"_"+phone] = accessToken
	return nil
}

func getRefreshToken(areaCode, phone string) string {
	return tokenMap[RefreshToken+"_"+areaCode+"_"+phone]
}

func setRefreshToken(areaCode, phone, refreshToken string) error {
	tokenMap[RefreshToken+"_"+areaCode+"_"+phone] = refreshToken
	return nil
}
