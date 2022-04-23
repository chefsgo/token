package token

type (
	Driver interface {
		Connect(config Config) (Connect, error)
	}
	Connect interface {
		// Open 打开连接
		Open() error
		// Close 关闭结束
		Close() error

		// Sign 签名
		Sign(*Token) (string, error)

		// Validate 验签
		Validate(token string) (*Token, error)
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
