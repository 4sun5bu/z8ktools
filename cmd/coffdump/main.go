package main

import (
	"fmt"
	"os"
	"z8ktools/z8kcoff"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}
	defer file.Close()

	var hdr z8kcoff.Header
	hdr, err = z8kcoff.GetFileHeader(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	z8kcoff.PrintFileHeader(hdr)

	var scns []z8kcoff.Section
	scns, err = z8kcoff.GetSections(file, hdr.Nscns)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	z8kcoff.PrintSections(scns)

	var relocs [][]z8kcoff.Relocation
	relocs, err = z8kcoff.GetRelocations(file, scns)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	z8kcoff.PrintRelocations(relocs)

	for _, scn := range scns {
		if scn.Scnptr == 0 {
			continue
		}
		fmt.Printf("[%s]\n", string(scn.Name[:]))
		data, err := z8kcoff.GetSectionData(file, scn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		z8kcoff.DumpSectionData(data)
	}
}
