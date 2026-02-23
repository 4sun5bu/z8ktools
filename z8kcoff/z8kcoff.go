package z8kcoff

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Header struct {
	Magic  uint16
	NScns  uint16
	Timdat uint32
	Symptr uint32
	NSyms  uint32
	Opthdr uint16
	Flags  uint16
}

const (
	TypeZ8002 = 0x2000
	TypeZ8001 = 0x1000
	BigEndian = 0x0200
	NoLineNo  = 0x0004
	NoSymbol  = 0x0008
	NoReloc   = 0x0001
)

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

const (
	SecText = 0x0020
	SecData = 0x0040
	SecBSS  = 0x0080
)

type Relocation struct {
	Vaddr  uint32
	Symndx uint32
	Offset uint32
	Type   uint16
	Stuff  uint16
}

const (
	RelWAbs    = 0x01
	RelJrDsp   = 0x02
	RelLAbs    = 0x11
	RelBAbs    = 0x22
	RelLwNbl   = 0x23
	RelWPCRel  = 0x04
	RelCallr   = 0x05
	RelSegNum  = 0x10
	RelUpNbl   = 0x24
	RelDjnzDsp = 0x25
)

type Symbol struct {
	Name   [8]byte
	Value  uint32
	Scnum  int16
	Type   uint16
	Sclass uint8
	Numaux uint8
}

const (
	ClassNull    = 0
	ClassExtSym  = 2
	ClassStatSym = 3
	ClassLabel   = 6
)

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
	fmt.Printf("  nscns : %2d ", hdr.NScns)
	fmt.Printf("  symoff : 0x%06x", hdr.Symptr)
	fmt.Printf("  nsyms : %3d ", hdr.NSyms)
	fmt.Printf("  flags : 0x%04x\n", hdr.Flags)
	switch hdr.Flags & 0x3000 {
	case TypeZ8001:
		fmt.Printf("  Z8001,")
	case TypeZ8002:
		fmt.Printf("  Z8002,")
	default:
		fmt.Printf("  No segment type,")
	}
	if (hdr.Flags & NoLineNo) != 0 {
		fmt.Printf(" line number stripped,")
	}
	if (hdr.Flags & NoSymbol) != 0 {
		fmt.Printf(" symbol table stripped,")
	}
	if (hdr.Flags & NoReloc) != 0 {
		fmt.Printf(" relocation info stripped,")
	}
	fmt.Printf("\b\n")
}
