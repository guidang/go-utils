package validate

import (
	"bytes"
	"errors"
)

// Validate Validate
type Validate struct {
	data,
	code []byte
	flag Mode
}

// NewValidate Validate init
func NewValidate(data []byte, code []byte, flag Mode) *Validate {
	t := &Validate{}
	t.data = data // 数据
	t.code = code // 校验码
	t.flag = flag // 校验码生成方式
	return t
}

// ValidateCode 校验码检测一致性
func (t *Validate) ValidateCode() error {
	if t.code == nil {
		return errors.New("校验码不存在")
	}

	oldCode := t.code
	if err := t.MakeCode(); err != nil {
		return errors.New("生成校验码失败")
	}

	if bytes.Equal(oldCode, t.code) {
		return nil
	}

	return errors.New("校验码不匹配")
}

// MakeCode 生成校验码
func (t *Validate) MakeCode() error {
	switch t.flag {
	case ByteXOR: // 前后字节异或方式
		return t.makeByteXOR()

	case ByteSuperposition: // 字节累计叠加方式
		return t.makeByteSuperposition()
	}

	return errors.New("不支持此校验码生成方式")
}

// GetCode 获取校验码
func (t *Validate) GetCode() []byte {
	return t.code
}

// makeByteXOR 前后字节异或方式
func (t *Validate) makeByteXOR() error {
	dataLen := len(t.data)

	if dataLen == 0 {
		return errors.New("空数据无法生成校验码")
	}

	code := t.data[0]
	for i := 1; i < dataLen; i++ {
		code ^= t.data[i]
	}

	t.code = append([]byte{}, code)
	return nil
}

// makeByteSuperposition 字节累计叠加方式
func (t *Validate) makeByteSuperposition() error {
	dataLen := len(t.data)
	if dataLen == 0 {
		return errors.New("空数据无法生成校验码")
	}

	var value byte

	for i := 0; i < dataLen; i++ {
		value += t.data[i]
	}

	t.code = append([]byte{}, value)
	return nil
}
