package z8kaout

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Header struct {
	Fmagic uint16
	Tsize  uint16
	Dsize  uint16
	Bsize  uint16
	Ssize  uint16
	Entry  uint16
	pad    uint16
	Relflg uint16
}

type Symbol struct {
	Name  [8]byte
	Stype uint8
	Spare uint8
	Value uint16
}

func GetFileHeader(fl *os.File) (Header, error) {
	var hdr Header
	err := binary.Read(fl, binary.BigEndian, &hdr)
	if err != nil {
		return hdr, errors.New("a.out header read error")
	}
	return hdr, nil
}

func PrintFileHeader(hdr Header) {
	fmt.Println("[Header]")
	fmt.Printf("  fmagic : 0x%04x ", hdr.Fmagic)
	fmt.Printf("  tsize : %5d ", hdr.Tsize)
	fmt.Printf("  dsize : %5d ", hdr.Dsize)
	fmt.Printf("  bsize : %5d ", hdr.Bsize)
	fmt.Printf("  ssize : %5d", hdr.Ssize)
	fmt.Printf("  entry : 0x%04x ", hdr.Entry)
	fmt.Printf("  relflg : 0x%04x\n", hdr.Relflg)
}
