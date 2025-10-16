package pkg

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"runtime"

	"github.com/klauspost/compress/zstd"
)

func CompressWind[T any](windItems WindowItem[*T]) ([]byte, error) {
	var buf bytes.Buffer
	encoder, err := zstd.NewWriter(&buf, zstd.WithEncoderLevel(zstd.SpeedFastest), zstd.WithEncoderConcurrency(runtime.NumCPU()))
	if err != nil {
		return nil, fmt.Errorf("failed to create writer: %w", err)
	}
	if err := gob.NewEncoder(encoder).Encode(windItems); err != nil {
		return nil, fmt.Errorf("failed to encode window items: %w", err)
	}
	return buf.Bytes(), nil
}

func DecompressWind[T any](data []byte) (WindowItem[*T], error) {
	decoder, err := zstd.NewReader(bytes.NewReader(data), zstd.WithDecoderConcurrency(runtime.NumCPU()))
	if err != nil {
		return WindowItem[*T]{}, fmt.Errorf("failed to create reader: %w", err)
	}
	var windItems WindowItem[*T]
	if err := gob.NewDecoder(decoder).Decode(&windItems); err != nil {
		return WindowItem[*T]{}, fmt.Errorf("failed to decode window items: %w", err)
	}
	return windItems, nil
}
