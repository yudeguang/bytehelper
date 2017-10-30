//用于根据一定的规则对切片进行读取,用于补充"encoding/binary"包对于读取数据的一些便之处

package bytehelper

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

type ByteHelper struct {
	data []byte
}

//实例化切片
func New(data []byte) *ByteHelper {
	var b ByteHelper
	b.data = data
	return &b
}

//截取nSize长度切片
func (b *ByteHelper) ReadByte(nSize int) []byte {
	s := b.data[:nSize]
	b.data = b.data[nSize:]
	return s
}

//截取nSize长度文本
func (b *ByteHelper) ReadString(nSize int) string {
	return string(b.ReadByte(nSize))
}

//截取Uint8
func (b *ByteHelper) ReadUint8() uint8 {
	s := uint8(b.data[0:1][0])
	b.data = b.data[1:]
	return s
}

//截取Uint16
func (b *ByteHelper) ReadUint16() uint16 {
	n := binary.LittleEndian.Uint16(b.data)
	b.data = b.data[2:]
	return n
}

//截取Uint16 BigEndian
func (b *ByteHelper) ReadUint16BigEndian() uint16 {
	n := binary.BigEndian.Uint16(b.data)
	b.data = b.data[2:]
	return n
}

//截取Uint32
func (b *ByteHelper) ReadUint32() uint32 {
	n := binary.LittleEndian.Uint32(b.data)
	b.data = b.data[4:]
	return n
}

//截取Uint32 BigEndian
func (b *ByteHelper) ReadUint32BigEndian() uint32 {
	n := binary.BigEndian.Uint32(b.data)
	b.data = b.data[4:]
	return n
}

//截取Uint64
func (b *ByteHelper) ReadUint64() uint64 {
	n := binary.BigEndian.Uint64(b.data)
	b.data = b.data[8:]
	return n
}

//截取Uint64 BigEndian
func (b *ByteHelper) ReadUint64BigEndian() uint64 {
	n := binary.BigEndian.Uint64(b.data)
	b.data = b.data[8:]
	return n
}

//截取数据,其中第1个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint8() []byte {
	nSize := b.ReadUint8()
	return b.ReadByte(int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint16() []byte {
	nSize := b.ReadUint16()
	return b.ReadByte(int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint16BigEndian() []byte {
	nSize := b.ReadUint16BigEndian()
	return b.ReadByte(int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint32() []byte {
	nSize := b.ReadUint32()
	return b.ReadByte(int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint32BigEndian() []byte {
	nSize := b.ReadUint32BigEndian()
	return b.ReadByte(int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint64() []byte {
	nSize := b.ReadUint64()
	return b.ReadByte(int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度
func (b *ByteHelper) ReadByteUint64BigEndian() []byte {
	nSize := b.ReadUint64BigEndian()
	return b.ReadByte(int(nSize))
}

//截取数据,其中第1个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUint8() string {
	nSize := b.ReadUint8()
	return b.ReadString(int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUIN16() string {
	nSize := b.ReadUint16()
	return b.ReadString(int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUIN16BigEndian() string {
	nSize := b.ReadUint16BigEndian()
	return b.ReadString(int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUIN32() string {
	nSize := b.ReadUint32()
	return b.ReadString(int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUIN32BigEndian() string {
	nSize := b.ReadUint32BigEndian()
	return b.ReadString(int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUIN64() string {
	nSize := b.ReadUint64()
	return b.ReadString(int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度,截取完成后转化为文本
func (b *ByteHelper) ReadStringUIN64BigEndian() string {
	nSize := b.ReadUint64BigEndian()
	return b.ReadString(int(nSize))
}

//把16进制转化为文本形式
func (b *ByteHelper) ReadHexToString(nSize int) string {
	s := strings.ToUpper(hex.EncodeToString(b.data[:nSize]))
	b.data = b.data[nSize:]
	return s
}

//截取剩余切片长度
func (b *ByteHelper) Len() int {
	return len(b.data)
}

//截取nSize长度的切片
func ReadByte(data []byte, nSize int) ([]byte, []byte) {
	return data[:nSize], data[nSize:]
}

//截取nSize长度的切片并转化为文本
func ReadString(data []byte, nSize int) (string, []byte) {
	return string(data[:nSize]), data[nSize:]
}

//截取Uint8
func ReadUint8(data []byte) (uint8, []byte) {
	return uint8(data[0:1][0]), data[1:]
}

//截取Uint16
func ReadUint16(data []byte) (uint16, []byte) {
	n := binary.LittleEndian.Uint16(data)
	return n, data[2:]
}

//截取Uint16 BigEndian
func ReadUint16BigEndian(data []byte) (uint16, []byte) {
	n := binary.BigEndian.Uint16(data)
	return n, data[2:]
}

//截取Uint32
func ReadUint32(data []byte) (uint32, []byte) {
	n := binary.LittleEndian.Uint32(data)
	return n, data[4:]
}

//截取Uint32 BigEndian
func ReadUint32BigEndian(data []byte) (uint32, []byte) {
	n := binary.BigEndian.Uint32(data)
	return n, data[4:]
}

//截取Uint64
func ReadUint64(data []byte) (uint64, []byte) {
	n := binary.LittleEndian.Uint64(data)
	return n, data[8:]
}

//截取Uint64 BigEndian
func ReadUint64BigEndian(data []byte) (uint64, []byte) {
	n := binary.BigEndian.Uint64(data)
	return n, data[8:]
}

//截取数据,其中第1个字节表示后续待截取切片长度
func ReadByteUint8(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint8(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度
func ReadByteUint16(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint16(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度
func ReadByteUint16BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint16BigEndian(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度
func ReadByteUint32(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint32(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度
func ReadByteUint32BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint32BigEndian(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度
func ReadByteUint64(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint64(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度
func ReadByteUint64BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := ReadUint64BigEndian(data)
	return ReadByte(data, int(nSize))
}

//截取数据,其中第1个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUint8(data []byte) (string, []byte) {
	nSize, data := ReadUint8(data)
	return ReadString(data, int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUIN16(data []byte) (string, []byte) {
	nSize, data := ReadUint16(data)
	return ReadString(data, int(nSize))
}

//截取数据,其中前2个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUIN16BigEndian(data []byte) (string, []byte) {
	nSize, data := ReadUint16BigEndian(data)
	return ReadString(data, int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUIN32(data []byte) (string, []byte) {
	nSize, data := ReadUint32(data)
	return ReadString(data, int(nSize))
}

//截取数据,其中前4个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUIN32BigEndian(data []byte) (string, []byte) {
	nSize, data := ReadUint32BigEndian(data)
	return ReadString(data, int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUIN64(data []byte) (string, []byte) {
	nSize, data := ReadUint64(data)
	return ReadString(data, int(nSize))
}

//截取数据,其中前8个字节表示后续待截取切片长度,截取完成后转化为文本
func ReadStringUIN64BigEndian(data []byte) (string, []byte) {
	nSize, data := ReadUint64BigEndian(data)
	return ReadString(data, int(nSize))
}

//把16进制转化为文本形式
func ReadHexToString(data []byte, nSize int) (string, []byte) {
	s := strings.ToUpper(hex.EncodeToString(data[:nSize]))
	return s, data[nSize:]
}
