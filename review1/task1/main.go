package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getIntMaxValue(12, 12, 12, 12))
	fmt.Println(getUintMaxValue(12, 12, 12, 12))
	fmt.Println(getBits("string"))
}
func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	var (
		maxInt8Value  int8
		maxInt16Value int16
		maxInt32Value int32
		maxInt64Value int64
	)
	if in8 != 0 {
		maxInt8Value = int8(1<<(8-1) - 1)
	}
	if in16 != 0 {
		maxInt16Value = int16(1<<(16-1) - 1)
	}
	if in32 != 0 {
		maxInt32Value = int32(1<<(32-1) - 1)
	}
	if in64 != 0 {
		maxInt64Value = int64(1<<(64-1) - 1)
	}
	return maxInt8Value, maxInt16Value, maxInt32Value, maxInt64Value
}
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	var (
		maxUint8Value  uint8
		maxUint16Value uint16
		maxUint32Value uint32
		maxUint64Value uint64
	)
	if uin8 != 0 {
		maxUint8Value = uint8(1<<(8) - 1)
	}
	if uin16 != 0 {
		maxUint16Value = uint16(1<<(16) - 1)
	}
	if uin32 != 0 {
		maxUint32Value = uint32(1<<(32) - 1)
	}
	if uin64 != 0 {
		maxUint64Value = uint64(1<<(64) - 1)
	}
	return maxUint8Value, maxUint16Value, maxUint32Value, maxUint64Value
}
func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)

	return bits
}
