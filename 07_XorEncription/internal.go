package main

import (
	"errors"
	"io"
)

func Encrypt(src io.Reader, dst io.Writer, key []byte) error {
	return xor(src, dst, key)
}

func Decrypt(src io.Reader, dst io.Writer, key []byte) error {
	return xor(src, dst, key)
}

func xor(src io.Reader, dst io.Writer, key []byte) error {
	if len(key) == 0 {
		return errors.New("key must not be empty")
	}
	buf := make([]byte, 1024)
	keyIdx := 0
	for {
		n, err := src.Read(buf)
		if n > 0 {
			for i := range n {
				buf[i] = buf[i] ^ key[keyIdx]
				keyIdx = (keyIdx + 1) % len(key)
			}
			if _, werr := dst.Write(buf[:n]); werr != nil {
				return werr
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
