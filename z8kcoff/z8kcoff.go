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

const TypZ8002 uint16 = 0x2000
const TypZ8001 uint16 = 0x1000
const TypBigEndian uint16 = 0x0200
const TypNoLno = 0x0004
const TypNoSym = 0x0008
const TypNoRel = 0x0001

const SecText uint32 = 0x0020
const SecData uint32 = 0x0040
const SecBSS uint32 = 0x0080

const RelWAbs uint16 = 0x01
const RelJrDsp uint16 = 0x02
const RelLAbs uint16 = 0x11
const RelBAbs uint16 = 0x22
const RelLwNbl uint16 = 0x23
const RelWPCRel uint16 = 0x04
const RelCallr uint16 = 0x05
const RelSegNum uint16 = 0x10
const RelUpNbl uint16 = 0x24
const RelDjnzDsp uint16 = 0x25

const ClsNull uint8 = 0
const ClsExtSym uint8 = 2
const ClsStatSym uint8 = 3
const ClsLabel uint8 = 6

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
	switch hdr.Flags & 0x3000 {
	case TypZ8001:
		fmt.Printf("  Z8001,")
	case TypZ8002:
		fmt.Printf("  Z8002,")
	default:
		fmt.Printf("  No segment type,")
	}
	if (hdr.Flags & TypNoLno) != 0 {
		fmt.Printf(" line number stripped,")
	}
	if (hdr.Flags & TypNoSym) != 0 {
		fmt.Printf(" symbol table stripped,")
	}
	if (hdr.Flags & TypNoRel) != 0 {
		fmt.Printf(" relocation info stripped,")
	}
	fmt.Printf("\b\n")
}
