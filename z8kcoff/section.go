package z8kcoff

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

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

func PrintSections(scns []Section) {
	fmt.Println("[Section]")
	for n, scn := range scns {
		fmt.Printf("  %d: name: [", n+1)
		for _, c := range scn.Name {
			fmt.Printf("0x%02x ", c)
		}
		fmt.Printf("\b]")
		data := scn.Name[:]
		name := string(data)
		fmt.Printf(" %s\n", name)
		fmt.Printf("     paddress: 0x%06x", scn.Paddr)
		fmt.Printf("  vaddr: 0x%06x", scn.Vaddr)
		fmt.Printf("  size: %2d\n", scn.Size)
		fmt.Printf("     scnoff: 0x%06x", scn.Scnptr)
		fmt.Printf("  reloff: 0x%06x", scn.Relptr)
		fmt.Printf("  nreloc: %2d", scn.Nreloc)
		fmt.Printf("  flags: 0x%06x\n", scn.Flags)
	}
}

func GetSectionData(fl *os.File, scn Section) ([]byte, error) {
	_, err := fl.Seek(int64(scn.Scnptr), io.SeekStart)
	if err != nil {
		return nil, errors.New("section data read error")
	}
	var buf []byte
	var data byte
	for i := 0; i < int(scn.Size); i++ {
		err = binary.Read(fl, binary.BigEndian, &data)
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
			fmt.Printf("  %04x:  ", addr)
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
