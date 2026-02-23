package z8kcoff

import (
	"bufio"
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
	for n := 0; n < int(hdr.NSyms); n++ {
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
	var naux uint8 = 0
	for n, sym := range syms {
		fmt.Printf(" %2d: name: [", n)
		for _, c := range sym.Name {
			fmt.Printf("0x%02x ", c)
		}
		fmt.Printf("\b] ")
		if sym.Name[0] == 0x00 {
			if naux != 0 {
				// p := int64(sym.Name[1])<<16 + int64(sym.Name[2])<<8 + int64(sym.Name[3])
				// fmt.Printf("0x%06x\n", p)
				fmt.Printf("AUX\n")
			} else {
				p := int64(sym.Name[5])<<16 + int64(sym.Name[6])<<8 + int64(sym.Name[7])
				fmt.Printf("Strings at 0x%06x\n", p)
			}
		} else {
			data := sym.Name[:]
			name := string(data)
			fmt.Printf("%s\n", name)
		}
		fmt.Printf("     value: 0x%08x", sym.Value)
		fmt.Printf(" scnum: %d ", sym.Scnum)
		fmt.Printf(" type: 0x%04x\n", sym.Type)
		fmt.Printf("     sclass: 0x%04x ", sym.Sclass)
		fmt.Printf(" numaux: %d - ", sym.Numaux)
		printSymClass(sym.Sclass)
		naux = sym.Numaux
	}
}

func printSymClass(clas uint8) {
	var clsstr string
	switch clas {
	case ClassNull:
		clsstr = "No entry"
	case ClassExtSym:
		clsstr = "Ext symb"
	case ClassStatSym:
		clsstr = "Static symb"
	case ClassLabel:
		clsstr = "Label"
	default:
		clsstr = "Unknown"
	}
	fmt.Println(clsstr)
}

func GetStrings(fl *os.File, hdr Header) ([]string, error) {
	_, err := fl.Seek(int64(hdr.Symptr)+int64(hdr.NSyms)*18+4, io.SeekStart)
	if err != nil {
		return nil, errors.New("no strings")
	}
	rd := bufio.NewReader(fl)
	var strs []string
	var str string
	for {
		str, err = rd.ReadString(0x00)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.New("symbol not read")
		}
		strs = append(strs, str)
	}
	return strs, nil
}

func PrintStrings(strs []string) {
	fmt.Println("[Strings]")
	if len(strs) == 0 {
		fmt.Println("  no strings")
	} else {
		addr := 0x0004
		for _, str := range strs {
			fmt.Printf("  0x%04x: %s\n", addr, str)
			addr += len(str)
		}
	}
}
