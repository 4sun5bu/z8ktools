package z8kcoff

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Header struct {
	Magic  uint16
	Nscns  uint16
	Timdat uint32
	Symptr uint32
	Nsyms  uint32
	Opthdr uint16
	Flags  uint16
}

type Section struct {
	Name    [8]byte
	Paddr   uint32
	Vaddr   uint32
	Size    uint32
	Scnptr  uint32
	Relptr  uint32
	Lnnoptr uint32
	Nreloc  uint16
	Nlnno   uint16
	Flags   uint32
}

type Relocation struct {
	Vaddr  uint32
	Symndx uint32
	Offset uint32
	Type   uint16
	Stuff  uint16
}

type Symbol struct {
	Name   [8]uint8
	Value  uint32
	Scnum  int16
	Type   uint16
	Sclass uint8
	Numaux uint8
}

func GetFileHeader(fl *os.File) (Header, error) {
	var hdr Header
	err := binary.Read(fl, binary.BigEndian, &hdr)
	if err != nil {
		return hdr, errors.New("COFF header read error")
	}
	return hdr, nil
}

func PrintFileHeader(hdr Header) {
	fmt.Println("[Header]")
	fmt.Printf("  magic : 0x%04x ", hdr.Magic)
	fmt.Printf("  nscns : %2d ", hdr.Nscns)
	fmt.Printf("  symoff : 0x%06x", hdr.Symptr)
	fmt.Printf("  nsyms : %3d ", hdr.Nsyms)
	fmt.Printf("  flags : 0x%04x\n", hdr.Flags)
}
