package tokenauth

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	cache "baseGinPro/pkg/models/redis"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// Value Token对应的Value值
type Value struct {
	UserID    int32  `json:"user_id"`    // 用户ID
	SecretKey string `json:"secret_key"` // 秘钥
}

// 与客户端的token验证，token存在redis中，定时刷新
func Token() gin.HandlerFunc {

	return func(c *gin.Context) {
		userTokenPool := cache.UserCache
		// 处理接口验证
		// 1. 判断请求方法，如果不是HTTP POST直接返回错误
		if c.Request.Method != "POST" {
			fmt.Errorf("TokenAuth: RequestMethod!=POST, postArgs:%v\n", c.Request.Form.Encode())
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "错误的请求方法."})
			c.Abort()
			return
		}

		// 1. 获取TOKENID - 无效返回错误
		tokenID := c.PostForm("token_id")
		if tokenID == "" {
			fmt.Errorf("TokenAuth: tokenID empty, postArgs:%v.\n", c.Request.Form.Encode())
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "TokenID为空."})
			c.Abort()
			return
		}

		// 1. 请求Redis服务器获取SecretKey - 失败返回处理
		tokenValue, err := GetTokenValue(tokenID, userTokenPool)
		if err != nil {
			fmt.Errorf("TokenAuth: GetTokenValue, postArgs:%v err:%v\n", c.Request.Form.Encode(), err)
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": err.Error()})
			c.Abort()
			return
		}

		// 1. 获取请求用户ID
		userID := c.PostForm("user_id")
		if userID == "" {
			fmt.Errorf("TokenAuth: UserID empty, postArgs:%v\n", c.Request.Form.Encode())
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "用户账号为空."})
			c.Abort()
			return
		}

		if userID != strconv.Itoa(int(tokenValue.UserID)) {
			fmt.Errorf("TokenAuth: userID!=tokenValue.UserID, postArgs:%v\n", c.Request.Form.Encode())
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "用户账号不匹配."})
			c.Abort()
			return
		}

		// 2. 获取接口参数
		m, err := ToMap(c.Request.Form)
		if err != nil {
			fmt.Errorf("TokenAuth: ToMap postArgs:%v err:%v\n", c.Request.Form.Encode(), err)
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": err.Error()})
			c.Abort()
			return
		}

		// 3. 获取请求签名、计算签名以及验证2者是否合法
		reqSign := c.PostForm("sign")

		rawSignString := MakeSignString(m, tokenValue.SecretKey)
		checkSign := Hexdigest(rawSignString)
		// // 3. 对接口参数进行验证 - 目前限制只能使用HTTP GET或者POST方式
		// checkSign := CalculcateSign(m, tokenValue.SecretKey)
		if reqSign != checkSign {
			fmt.Errorf("TokenAuth: SignNotMatch postArgs:%v checkSign:%v secretKey:%v\n",
				c.Request.Form.Encode(), checkSign, tokenValue.SecretKey)
			c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "鉴权验证请求非法。"})
			c.Abort()
			return
		}

		// Before Request
		c.Next()
	}

}

// Hexdigest md5签名
func Hexdigest(s string) string {
	w := md5.New()
	w.Write([]byte(s))
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}

// MakeSignString 生成签名字符串
func MakeSignString(m map[string]string, secretKey string) string {
	// 对签名字段名进行排序
	mKeys := make([]string, 0)
	for key := range m {
		mKeys = append(mKeys, key)
	}
	sort.Strings(mKeys)

	// 组合签名原始字符串
	// 以k1=value&k2=value&k3=value, 在原始字符串里面最后加key=xxx
	Signs := make([]string, 0)
	for _, key := range mKeys {
		// 如果是sign是，排除签名
		lowerKey := strings.ToLower(key)
		if lowerKey == "sign" {
			continue
		}

		// Value不存在，排除签名
		value, ok := m[key]
		if !ok {
			continue
		}

		// Value 为空，排查签名
		valueStr := fmt.Sprintf("%v", value)
		if valueStr == "" {
			continue
		}

		tmpSign := fmt.Sprintf("%v=%v", lowerKey, value)
		Signs = append(Signs, tmpSign)
	}

	Signs = append(Signs, fmt.Sprintf("key=%v", secretKey))
	rawSign := strings.Join(Signs, "&")
	return rawSign
}

// CalculcateSign 生成签名
func CalculcateSign(m map[string]string, secretKey string) string {
	// 对签名字段名进行排序
	mKeys := make([]string, 0)
	for key := range m {
		mKeys = append(mKeys, key)
	}
	sort.Strings(mKeys)

	// 组合签名原始字符串
	// 以k1=value&k2=value&k3=value, 在原始字符串里面最后加key=xxx
	Signs := make([]string, 0)
	for _, key := range mKeys {
		// 如果是sign是，排除签名
		lowerKey := strings.ToLower(key)
		if lowerKey == "sign" {
			continue
		}

		// Value不存在，排除签名
		value, ok := m[key]
		if !ok {
			continue
		}

		// Value 为空，排查签名
		valueStr := fmt.Sprintf("%v", value)
		if valueStr == "" {
			continue
		}

		tmpSign := fmt.Sprintf("%v=%v", lowerKey, value)
		Signs = append(Signs, tmpSign)
	}

	Signs = append(Signs, fmt.Sprintf("key=%v", secretKey))
	rawSign := strings.Join(Signs, "&")

	w := md5.New()
	w.Write([]byte(rawSign))
	// io.WriteString(w, rawSign)               //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}

// ToMap 转换为map[string]string
func ToMap(values url.Values) (map[string]string, error) {

	m := make(map[string]string)
	for fk, fvs := range values {
		if len(fvs) >= 2 {
			return nil, fmt.Errorf("参数%v存在不确定值", fk)
		}
		m[fk] = fvs[0]
	}
	return m, nil
}

// GetTokenValue 获取Token对应的值
func GetTokenValue(tokenID string, pool *redis.Pool) (*Value, error) {
	conn := pool.Get()
	defer conn.Close()

	tokenKey := fmt.Sprintf("secret:token:%v", tokenID)
	valueBytes, err := redis.Bytes(conn.Do("GET", tokenKey))
	if err != nil {
		fmt.Errorf("GetTokenValue: tokenID:%v err:%v\n", tokenID, err)
		return nil, fmt.Errorf("会话ID不存在或已过期")
	}

	value := &Value{}
	if err := json.Unmarshal(valueBytes, value); err != nil {
		fmt.Errorf("GetTokenValue: tokenID:%v err:%v\n", tokenID, err)
		return nil, fmt.Errorf("数据解码错误。")
	}

	return value, nil
}
