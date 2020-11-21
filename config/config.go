package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	//"github.com/json-iterator/go"
	"github.com/bitly/go-simplejson"
)

type Config struct {
	DBConfigs        map[string]DBConfig `json:"dbs"`
	Env              string              `json:"env"`
	RedisClusterName string              `json:"redis_cluster_name"`
	RedisHosts       []string            `json:"redis_hosts"`
}

type DBConfig struct {
	Database string          `json:"database"`
	Settings string          `json:"settings"`
	WriteDB  DBConnectInfo   `json:"write"`
	ReadDB   []DBConnectInfo `json:"read"` //attention
}

type DBConnectInfo struct {
	AuthKey         string `json:"auth_key"` //拓展使用，一般作为动态密钥
	Consul          string `json:"consul"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	DefaultHostPort string `json:"default_host_port"`
}

var (
	ConfInstance *Config
	DBSettings   *simplejson.Json
)

func isProduct() bool {
	return ConfInstance.Env == "prod"
}

func NewConfig(file string) (*Config, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}
	DBSettings, err = simplejson.NewJson(content)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func Init(file string) error {
	if ConfInstance != nil {
		return nil
	}
	conf, err := NewConfig(file)
	if err != nil {
		return err
	}
	if len(conf.Env) == 0 {
		conf.Env = "dev"
	}
	ConfInstance = conf
	return nil
}

func CheckEnv() string {
	if os.Getenv("PRODUCT_ENV") != "" {
		return "prod"
	}
	return "dev"
}

var Input_ConfDir string

func FlagInit() {
	flag.StringVar(&Input_ConfDir, "confdir", "", "配置文件路径")
	flag.Parse()
	if Input_ConfDir == "" {
		panic("flaginit error")
	}
}
