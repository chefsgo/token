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
