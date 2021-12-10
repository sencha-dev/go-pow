// ethash: C/C++ implementation of Ethash, the Ethereum Proof of Work algorithm.
// Copyright 2018-2019 Pawel Bylica.
// Licensed under the Apache License, Version 2.0.

package progpow

import (
	"encoding/binary"
	"unsafe"
)

func isLittleEndian() bool {
	n := uint32(0x01020304)
	return *(*byte)(unsafe.Pointer(&n)) == 0x04
}

/* Array utils */

func uint32ArrayToBytes(c []uint32) []byte {
	buf := make([]byte, len(c)*4)
	if isLittleEndian() {
		for i, v := range c {
			binary.LittleEndian.PutUint32(buf[i*4:], v)
		}
	} else {
		for i, v := range c {
			binary.BigEndian.PutUint32(buf[i*4:], v)
		}
	}
	return buf
}

func bytesToUint32Array(arr []byte) []uint32 {
	length := len(arr) / 4
	data := make([]uint32, length)

	for i := 0; i < length; i++ {
		data[i] = binary.LittleEndian.Uint32(arr[i*4 : (i+1)*4])
	}

	return data
}
