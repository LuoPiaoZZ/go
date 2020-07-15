package main
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
//定义超时时间
const TokenExpireDuration =time.Hour*2//2小时

//定义密钥
var MySercret=[]byte("密钥")

//生成一个jwt
func GenToken(username string)(string,error)  {
	//创建一个自己的声明
	c:=MyClaims{
		"username",//自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),//过期时间
			Issuer: "demoproject",//签发人
		},
	}
	//使用定义的签名方法创建签名对象
	token:=jwt.NewWithClaims(jwt.SigningMethodES256,c)
	//使用指定的密钥签名并获得完整的编码后的字符串token
	return token.SignedString(MySercret)
}
//解析jwt
func ParseToken(tokenString string)(*MyClaims,error)  {
	//解析token
	token,err:=jwt.ParseWithClaims(tokenString,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySercret,nil
	})
	if err!=nil {
		return nil, err
	}
	if claims,ok:=token.Claims.(*MyClaims);ok&&token.Valid {//校验
		return claims,nil
	}
	return nil,errors.New("invalid token")
}
