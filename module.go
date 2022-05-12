package token

import (
	"sync"
	"time"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
)

func init() {
	chef.Register(NAME, module)
}

var (
	module = &Module{
		config: Config{
			Driver: chef.DEFAULT,
		},
		drivers: map[string]Driver{},
	}
)

type (
	// Level 日志级别，从小到大，数字越小越严重
	Level = int

	// 日志模块定义
	Module struct {
		//mutex 锁
		mutex sync.Mutex

		// 几项运行状态
		connected, initialized, launched bool

		//config 令牌配置
		config Config

		//drivers 驱动注册表
		drivers map[string]Driver

		// connect 令牌连接
		connect Connect
	}

	// LogConfig 日志模块配置
	Config struct {
		// Driver 令牌驱动，默认为 default
		Driver string

		// Secret 密钥
		Secret string

		// Expiry 默认过期时间
		// 0 表示不过期
		Expiry time.Duration

		// Setting 是为不同驱动准备的自定义参数
		// 具体参数表，请参考各不同的驱动
		Setting Map
	}

	Token struct {
		ActId      string `json:"d,omitempty"`
		Authorized bool   `json:"a,omitempty"`
		Identity   string `json:"i,omitempty"`
		Payload    Map    `json:"l,omitempty"`
		Expiry     int64  `json:"e,omitempty"`
	}
)

//Driver 为log模块注册驱动
func (this *Module) Driver(name string, driver Driver, override bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if driver == nil {
		panic("Invalid log driver: " + name)
	}

	if override {
		this.drivers[name] = driver
	} else {
		if this.drivers[name] == nil {
			this.drivers[name] = driver
		}
	}
}

func (this *Module) Config(config Config, override bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.config = config
}
