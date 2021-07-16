package app

import (
	"time"

	"github.com/go-programming-tour-book/blog-service/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-programming-tour-book/blog-service/global"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	// 它是 jwt-go 库中预定义的，也是 JWT 的规范，
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}
// 它承担了整个流程中比较重要的职责，也就是生成 JWT Token 的行 为，主体的函数流程逻辑是根据客户端传入的
// AppKey 和 AppSecret 以及在项目配置中所设置的签 发者（Issuer）和过期时间（ExpiresAt），根据指定的算法生成签名后的 Token。
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
    // 体创建 Token 实例，
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	 // 生成签名字符串
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// 解析和校验 Token
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返 回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		// Valid：验证基于时间的声明
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
