package util

import (
	"go-web-test/pkg/setting" // 导入项目设置包
	"time"                    // 导入时间包

	"github.com/golang-jwt/jwt" // 导入JWT库
)

var jwtSecret = []byte(setting.JwtSecret) // 从设置中读取JWT密钥并转换为字节切片

// Claims 定义了JWT的payload中包含的字段
type Claims struct {
	Username           string `json:"username"` // 用户名
	Password           string `json:"password"` // 用户密码 (注意：在生产环境中，密码不应该直接存储在JWT中，而是应该使用更安全的方式，例如只存储用户ID)
	jwt.StandardClaims        // JWT标准声明，包含ExpiresAt, Issuer等
}

// GenerateToken 生成JWT token
// 实际生产环境中，密码不应该直接存储在JWT中，而是应该使用更安全的方式，例如只存储用户ID和用户名，jwt只能验证途中是否被篡改
// 以验证是否可以信任这个用户
// 真实业务中 jwt里确实不应该放 密码等隐私问题
// 煎鱼打算 的golang demo是带你入门为目标的，真实 场景下 还需要考虑很多，比如 POST PUT 的参数应该放 body中，authToken应该放在请求头中 ，jwt签名的时候需要把ini文件中的，盐🧂 加进去
// 对于多用户登录的系统 jwt应该是放到redis中去维护的
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()                    // 获取当前时间
	expireTime := nowTime.Add(3 * time.Hour) // 设置过期时间为当前时间加上3小时

	claims := Claims{
		username, // 设置用户名
		password, // 设置密码 (注意：在生产环境中，密码不应该直接存储在JWT中)
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 设置过期时间戳
			Issuer:    "gin-blog",        // 设置签发者
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 创建一个新的JWT token，使用HS256加密算法
	token, err := tokenClaims.SignedString(jwtSecret)                // 使用密钥签名token

	return token, err // 返回token和错误
}

// ParseToken 解析JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil // 提供密钥
	})

	if tokenClaims != nil { // 验证token是否有效
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // 类型断言并检查token是否有效
			return claims, nil // 返回claims和nil错误
		}
	}

	return nil, err // 返回nil和错误
}
