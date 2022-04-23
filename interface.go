package token

import (
	"fmt"
	"time"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/util"
)

func (this *Module) Register(name string, value Any, override bool) {
	switch obj := value.(type) {
	case Driver:
		this.Driver(name, obj, override)
	}
}
func (this *Module) Configure(value Any) {
	if cfg, ok := value.(Config); ok {
		this.config = cfg
		return
	}

	var global Map
	if cfg, ok := value.(Map); ok {
		global = cfg
	} else {
		return
	}

	var config Map
	if vvv, ok := global["token"].(Map); ok {
		config = vvv
	}

	//设置驱动
	if driver, ok := config["driver"].(string); ok {
		this.config.Driver = driver
	}

	// secret 密钥
	if secret, ok := config["secret"].(string); ok {
		this.config.Secret = secret
	}

	//默认过期时间，单位秒
	if expiry, ok := config["expiry"].(string); ok {
		dur, err := util.ParseDuration(expiry)
		if err == nil {
			this.config.Expiry = dur
		}
	}
	if expiry, ok := config["expiry"].(int); ok {
		this.config.Expiry = time.Second * time.Duration(expiry)
	}
	if expiry, ok := config["expiry"].(float64); ok {
		this.config.Expiry = time.Second * time.Duration(expiry)
	}

	if setting, ok := config["setting"].(Map); ok {
		this.config.Setting = setting
	}
}
func (this *Module) Initialize() {
	if this.initialized {
		return
	}

	this.initialized = true
}
func (this *Module) Connect() {
	if this.connected {
		return
	}

	driver, ok := this.drivers[this.config.Driver]
	if ok == false {
		panic("Invalid token driver: " + this.config.Driver)
	}

	// 建立连接
	connect, err := driver.Connect(this.config)
	if err != nil {
		panic("Failed to connect to token: " + err.Error())
	}

	// 打开连接
	err = connect.Open()
	if err != nil {
		panic("Failed to open token connect: " + err.Error())
	}

	// 保存连接
	this.connect = connect

	this.connected = true
}
func (this *Module) Launch() {
	if this.launched {
		return
	}

	fmt.Println("token Launch")

	this.launched = true
}
func (this *Module) Terminate() {
	if this.connect != nil {
		this.connect.Close()
	}

	this.launched = false
	this.connected = false
	this.initialized = false
}
