package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	microConfig "github.com/micro/go-micro/config"
	"go-micro-doc/common/constant"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Project  string             `yaml:"project"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Domain   map[string]string  `yaml:"domain"`
	Log      log                `yaml:"log"`
	Registry registry           `yaml:"registry"`
	Jaeger   jaeger             `yaml:"jaeger"`
}

// 数据库连接配置
type DatabaseConfig struct {
	Dialect        string
	Database       string
	User           string
	Password       string
	Host           string
	Port           int
	Charset        string
	URL            string
	MaxIdleConnNum int
	MaxOpenConnNum int
}

// redis连接配置
type RedisConfig struct {
	Host      string
	Port      int
	Password  string
	URL       string
	MaxIdle   int
	MaxActive int
}


type log struct {
	Path          string
	MaxAge        int
	RotationTime  int
	RotationCount int
}

type registry struct {
	Address  string
	Client   map[string]string
	Ttl      int
	Interval int
	Version  string
	MetaData map[string]string
}

type jaeger struct {
	Host        string
	Port        int
	ServiceName string
}

var (
	application string
	runEnv      string
	AppConfig   *Config
	once        sync.Once // go 的并发包 sync 里的工具类，保证某个操作只能执行一次
)

/**
 * 初始化配置
 */
func InitConfig(app, env, path string) {
	once.Do(func() {
		application = app
		runEnv = env
		gin.SetMode(constant.GetGinMode(env))
		if path == "" {
			pwd, _ := os.Getwd()
			path = filepath.Join(pwd, "config", env) + ".yml"
		}
		AppConfig = loadConfig(path)
		//todo getLogHooks() 日志组件
	})
}

/**
 * 获取应用名称
 */
func GetApplication() string {
	return application
}

/**
 * 获取运行时环境
 */
func GetEnv() string {
	return runEnv
}

/**
 * 加载配置
 */
func loadConfig(path string) *Config {
	config := new(Config)
	if err := microConfig.LoadFile(path); err != nil {
		panic(fmt.Errorf("fatal error: load config file: %s\n", err))
	}
	if err := microConfig.Scan(config); err != nil {
		panic(fmt.Errorf("fatal error: scan config value: %s\n", err))
	}
	return config
}
