package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	Conf       *Config
)

// Config 对应配置文件结构
type Config struct {
	Listen                  string                  `toml:"listen"`
	GRPCListen              string                  `toml:"grpc_listen"`           // grpc服务地址
	DBServers               map[string]DBServer     `toml:"dbservers"`
	RedisServers            map[string]RedisServer  `toml:"redisservers"`
}

//RedisServer 表示 redis 服务器配置
type RedisServer struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"middlewares"`
}

// DBServer 表示DB服务器配置
type DBServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	if _, err := toml.DecodeFile(tomlfile, &Conf); err != nil {
		return Conf, err
	}
	return Conf, nil
}

// DBServerConf 获取数据库配置
func (c Config) DBServerConf(key string) (DBServer, bool) {
	s, ok := c.DBServers[key]
	return s, ok
}

//RedisServerConf 获取数据库配置
func (c Config) RedisServerConf(key string) (RedisServer, bool) {
	s, ok := c.RedisServers[key]
	return s, ok
}

// GetListenAddr 监听地址
func (c Config) GetListenAddr() string {
	return c.Listen
}

// ConnectString 表示连接数据库的字符串
func (m DBServer) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Host, m.Port, m.User, m.Password, m.DBName)
}