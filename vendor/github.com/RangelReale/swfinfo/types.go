package swfinfo

import (
	"io"

	"github.com/32bitkid/bitreader"
)

type Compression uint8

const (
	COMPRESS_NONE Compression = iota
	COMPRESS_ZLIB
	COMPRESS_LZMA
)

func (c Compression) String() string {
	switch c {
	case COMPRESS_NONE:
		return "No compression"
	case COMPRESS_ZLIB:
		return "ZLIB compression"
	case COMPRESS_LZMA:
		return "LZMA compression"
	}
	return "Unknown compression type"
}

type Twips int32

func (t Twips) Pixels() float32 {
	return float32(t) / 20.0
}

type Rect struct {
	Xmin, Xmax, Ymin, Ymax Twips
}

func (r *Rect) ReadFrom(f io.Reader) error {
	br := bitreader.NewBitReader(f)
	// read size of each property in bits
	v, err := br.Read32(5)
	if err != nil {
		return err
	}

	tsize := uint(v)
	var t uint32

	// Xmin
	t, err = br.Read32(tsize)
	if err != nil {
		return err
	}
	r.Xmin = Twips(t)

	// Xmax
	t, err = br.Read32(tsize)
	if err != nil {
		return err
	}
	r.Xmax = Twips(t)

	// Ymin
	t, err = br.Read32(tsize)
	if err != nil {
		return err
	}
	r.Ymin = Twips(t)

	// Ymax
	t, err = br.Read32(tsize)
	if err != nil {
		return err
	}
	r.Ymax = Twips(t)

	_, err = br.ByteAlign()
	if err != nil {
		return err
	}

	return nil
}

func (r *Rect) Width() Twips {
	return r.Xmax - r.Xmin
}

func (r *Rect) Height() Twips {
	return r.Ymax - r.Ymin
}
