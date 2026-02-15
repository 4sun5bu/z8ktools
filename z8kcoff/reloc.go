package z8kcoff

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

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

func PrintRelocations(rels [][]Relocation) {
	fmt.Printf("[Relocatio]\n")
	for n, scn := range rels {
		for _, rel := range scn {
			fmt.Printf("  scnum: %d", n+1)
			fmt.Printf("  reloc vaddr: 0x%06x", rel.Vaddr)
			fmt.Printf("  symndx: %2d", rel.Symndx)
			fmt.Printf("  type  0x%04x - ", rel.Type)
			printRelType(rel.Type)
		}
	}
}

func printRelType(typ uint16) {
	var typstr string
	switch typ {
	case RelWAbs:
		typstr = "16bit Abs"
	case RelJrDsp:
		typstr = "Jr Disp"
	case RelLAbs:
		typstr = "32bit Abs"
	case RelBAbs:
		typstr = "8bit Abs"
	case RelLwNbl:
		typstr = "Lower Nibble"
	case RelWPCRel:
		typstr = "16bit PC rel"
	case RelCallr:
		typstr = "Callr Disp"
	case RelSegNum:
		typstr = "Segment"
	case RelUpNbl:
		typstr = "Upper Nibble"
	case RelDjnzDsp:
		typstr = "Djnz Disp"
		typstr = "Unknown"
	}
	fmt.Println(typstr)
}
