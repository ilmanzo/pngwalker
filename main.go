package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type chunk struct {
	Length    uint32
	ChunkType [4]byte
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename.png\n", os.Args[0])
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	header := make([]byte, 8)
	_, err = f.Read(header)
	fmt.Printf("header: %v\n", header)
	if err != nil {
		panic(err)
	}
	var data chunk
	for {
		err = binary.Read(f, binary.BigEndian, &data)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("chunk len=%d, type: %s\n", data.Length, string(data.ChunkType[:4]))
		f.Seek(int64(data.Length+4), 1)
	}
}
