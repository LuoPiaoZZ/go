package main

import "github.com/BurntSushi/toml"

type Postgresql struct {
	UserName string `toml:"userName"`
	PassWord string `toml:passWord`
}
type A struct {
	A1 string `toml:"a1"`
}

type Config struct {
	Testkey int      `toml:"testkey"`
	TestArray []int `toml:"testarray"`
	Postgresql Postgresql `toml:"postgresql"`
	A       map[string]A `toml:"a"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		return c, err
	}
	return c, nil
}

//获取testkey
func (c Config)GetTestkey() int{
	return c.Testkey
}
//获取testArray
func (c Config)GetTestArray() []int  {
	return c.TestArray
}
//获取PostgreSQL
func (c Config)GetPostgresql() Postgresql{
	return c.Postgresql
}
//根据参数返回a下面的子表
func (c Config)getA(key string)A  {
	return c.A[key]
}
