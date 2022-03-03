package storage

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

func UUIDFromTimestamp(timeNano uint64) (string, error) {
	timeBuffer, err := GetBigEndianBytesUint64(timeNano)
	if err != nil {
		return "", err
	}
	defer ReleaseBytesBuffer(timeBuffer)
	return hex.EncodeToString(timeBuffer.Next(8)) + "-" + uuid.New().String(), nil
}

var bytesBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func AcquireBytesBuffer() *bytes.Buffer {
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func ReleaseBytesBuffer(bytes *bytes.Buffer) {
	bytes.Reset()
	bytesBufferPool.Put(bytes)
}

func GetBigEndianBytesUint64(n uint64) (*bytes.Buffer, error) {
	buf := AcquireBytesBuffer()
	err := binary.Write(buf, binary.BigEndian, n)
	return buf, err
}

func MustGetBigEndianBytesUint64(n uint64) *bytes.Buffer {
	buf := AcquireBytesBuffer()
	err := binary.Write(buf, binary.BigEndian, n)
	if err != nil {
		panic(err)
	}
	return buf
}

func GetUnixTimestampNow() uint64 {
	return uint64(time.Now().UnixNano())
}
