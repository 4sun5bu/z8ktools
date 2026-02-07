package z8kcoff

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

func GetSymbols(fl *os.File, hdr Header) ([]Symbol, error) {
	_, err := fl.Seek(int64(hdr.Symptr), io.SeekStart)
	var syms []Symbol
	var symb Symbol
	for n := 0; n < int(hdr.Nsyms); n++ {
		if err != nil {
			return nil, errors.New("symbol not read")
		}
		err = binary.Read(fl, binary.BigEndian, &symb)
		if err != nil {
			return nil, errors.New("symbol not read")
		}
		syms = append(syms, symb)
	}
	return syms, nil
}

func PrintSymbols(syms []Symbol) {
	fmt.Println("[Symbol]")
	for n, sym := range syms {
		fmt.Printf(" %2d: name: [", n)
		for _, c := range sym.Name {
			fmt.Printf("0x%02x ", c)
		}
		fmt.Printf("\b] ")
		if sym.Name[0] == 0x00 {
			p := int64(sym.Name[1])<<16 + int64(sym.Name[2])<<8 + int64(sym.Name[3])
			fmt.Printf("0x%06x\n", p)

		} else {
			data := sym.Name[:]
			name := string(data)
			fmt.Printf("%s\n", name)
		}
		fmt.Printf("     value: 0x%08x", sym.Value)
		fmt.Printf(" scnum: %d ", sym.Scnum)
		fmt.Printf(" type: 0x%04x\n", sym.Type)
		fmt.Printf("     sclass: 0x%04x ", sym.Sclass)
		fmt.Printf(" numaux: %d\n", sym.Numaux)
	}
}
