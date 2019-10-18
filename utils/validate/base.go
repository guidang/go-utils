package validate

// Mode 生成校验码的方式
type Mode int

const (
	// ByteXOR 前后字节异或方式
	ByteXOR Mode = iota
	// ByteSuperposition 字节累计叠加方式
	ByteSuperposition
)
