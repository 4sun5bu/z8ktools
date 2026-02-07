package z8kcoff

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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

func GetSections(fl *os.File, n uint16) ([]Section, error) {
	var scns []Section
	for i := 0; i < int(n); i++ {
		var scn Section
		err := binary.Read(fl, binary.BigEndian, &scn)
		if err != nil {
			return nil, errors.New("section header read error")
		}
		scns = append(scns, scn)
	}
	return scns, nil
}

func printScn(scn Section) {
	fmt.Printf("  name : [")
	for _, c := range scn.Name {
		fmt.Printf("0x%02x ", c)
	}
	fmt.Printf("\b]")
	data := scn.Name[:]
	name := string(data)
	fmt.Printf(" %s\n", name)
	fmt.Printf("  paddress : 0x%06x", scn.Paddr)
	fmt.Printf("  vaddr : 0x%06x", scn.Vaddr)
	fmt.Printf("  size : %2d\n", scn.Size)
	fmt.Printf("  scnoff : 0x%06x", scn.Scnptr)
	fmt.Printf("  reloff : 0x%06x", scn.Relptr)
	fmt.Printf("  nreloc : %2d", scn.Nreloc)
	fmt.Printf("  flags : 0x%06x\n", scn.Flags)
}

func PrintSections(scns []Section) {
	fmt.Println("[Section]")
	for _, scn := range scns {
		printScn(scn)
	}
}

func GetRelocations(fl *os.File, scns []Section) ([][]Relocation, error) {
	var rels [][]Relocation
	for _, scn := range scns {
		if scn.Nreloc == 0 {
			continue
		}
		_, err := fl.Seek(int64(scn.Relptr), io.SeekStart)
		if err != nil {
			return nil, errors.New("relocation info read error")
		}
		var scnRels []Relocation
		for i := 0; i < int(scn.Nreloc); i++ {
			var rel Relocation
			err := binary.Read(fl, binary.BigEndian, &rel)
			if err != nil {
				return nil, errors.New("relocation info read error")
			}
			scnRels = append(scnRels, rel)
		}
		rels = append(rels, scnRels)
	}
	return rels, nil
}

func printReloc(rel Relocation) {
	fmt.Printf("  reloc vaddr : 0x%06x", rel.Vaddr)
	fmt.Printf("    symndx : %2d", rel.Symndx)
	fmt.Printf("    type : 0x%04x\n", rel.Type)
}

func PrintRelocations(rels [][]Relocation) {
	fmt.Printf("[Relocatio]\n")
	for n, scn := range rels {
		for _, rel := range scn {
			fmt.Printf("  scnum : %2d", n+1)
			printReloc(rel)
		}
	}
}

func GetSectionData(fl *os.File, scn Section) ([]byte, error) {
	fl.Seek(int64(scn.Scnptr), io.SeekStart)
	var buf []byte
	var data byte
	for i := 0; i < int(scn.Size); i++ {
		err := binary.Read(fl, binary.BigEndian, &data)
		if err != nil {
			return nil, errors.New("section data read error")
		}
		buf = append(buf, data)
	}
	return buf, nil
}

func DumpSectionData(data []byte) {
	for addr, d := range data {
		if (addr & 0x0f) == 0 {
			fmt.Printf("  %04x  ", addr)
		}
		if (addr & 0x0f) == 8 {
			fmt.Printf(" ")
		}
		fmt.Printf("%02x ", d)
		if (addr & 0x0f) == 0x0f {
			fmt.Print("\b\n")
		}
	}
	fmt.Print("\n")
}
