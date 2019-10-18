// 粘包处理

package pack

import (
	"errors"
)

// Pack Pack
type Pack struct {
	data  []byte
	start []byte
	end   []byte
}

// NewPack Pack init
func NewPack(data, start, end []byte) *Pack {
	t := &Pack{}
	t.data = data
	t.start = start
	t.end = end
	return t
}

// Pack 封装包数据流 (添加包头、包尾)
func (t *Pack) Pack(data []byte) []byte {
	data = append(t.start, data...)
	data = append(data, t.end...)
	return data
}

// Unpack 拆解包数据流 (根据消息头、消息尾拆解粘包)
func (t *Pack) Unpack() (resp [][]byte, err error) {
	data := t.data
	dataLen := len(data)                   // 数据长度
	sLen, eLen := len(t.start), len(t.end) // 包头包层长度

	if dataLen <= 0 {
		return nil, errors.New("无法拆解空的包数据流")
	}

	if t.start == nil && t.end == nil {
		return append(resp, data), errors.New("数据包的消息头标志位和尾标志位未设置")
	}

	s1, s2 := -1, -1 // 开始、结束位置
	x1, x2 := 0, 0   // 开始、结束索引
	offset := 0

	for {
		// 索引超出
		if offset >= dataLen {
			break
		}

		// 没有开始位置, 处理开始位置
		if s1 == -1 {
			// 匹配包头位置
			if data[offset] == t.start[x1] {
				// 完全匹配
				if x1 == sLen-1 {
					// 开始位置
					s1 = offset - x1
					s2, x1 = -1, 0
				} else {
					x1++
				}
			} else {
				x1 = 0
			}

		} else {
			// 匹配包尾位置
			if data[offset] == t.end[x2] {
				// 完全匹配
				if x2 == eLen-1 {
					s2 = offset + 1
					resp = append(resp, data[s1:s2])
					// 起始位置重置
					s1, s2, x2 = -1, -1, 0
				} else {
					x2++
				}
			} else {
				x2 = 0
			}
		}
		offset++
	}

	return
}
