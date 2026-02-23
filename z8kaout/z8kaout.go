package z8kaout

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type Header struct {
	Fmagic uint16
	TSize  uint16
	DSize  uint16
	BSize  uint16
	SSize  uint16
	Entry  uint16
	pad    uint16
	Relflg uint16
}

const (
	NMagic1  = 0xe707 // Nonsegmented executable
	NMagic2  = 0xe711 // Nonsegmented separate I/D
	RelStrip = 0x0001 // Reloc info stripped
)

type Symbol struct {
	Name  [8]byte
	SType uint8
	Spare uint8
	Value uint16
}

const ()

func GetFileHeader(fl *os.File) (Header, error) {
	var hdr Header
	err := binary.Read(fl, binary.BigEndian, &hdr)
	if err != nil {
		return hdr, errors.New("a.out header read error")
	}
	return hdr, nil
}

func SetFileHeader(fl *os.File, hdr Header) error {
	_, err := fl.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New("file Seek Error")
	}
	err = binary.Write(fl, binary.BigEndian, &hdr)
	if err != nil {
		return errors.New("header Write Error")
	}
	return nil
}

func PrintFileHeader(hdr Header) {
	fmt.Println("[Header]")
	fmt.Printf("  fmagic : 0x%04x ", hdr.Fmagic)
	fmt.Printf("  tsize : %5d ", hdr.TSize)
	fmt.Printf("  dsize : %5d ", hdr.DSize)
	fmt.Printf("  bsize : %5d ", hdr.BSize)
	fmt.Printf("  ssize : %5d", hdr.SSize)
	fmt.Printf("  entry : 0x%04x ", hdr.Entry)
	fmt.Printf("  relflg : 0x%04x\n", hdr.Relflg)
}
