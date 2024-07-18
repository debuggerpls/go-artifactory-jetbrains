package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type JetbrainsProduct struct {
	Code                string           // RR
	Name                string           // RustRover
	Releases            []ProductRelease // array
	ForSale             bool
	ProductFamilyName   string          // RustRover
	AdditionalLinks     json.RawMessage // array
	IntellijProductCode string          // RR
	AlternativeCodes    json.RawMessage // array
	SalesCode           string          // RR
	Link                string          // https://www.jetbrains.com/rust/
	Description         string          // A powerful IDE for Rust
	Tags                json.RawMessage // array
	Types               json.RawMessage // array
	Categories          json.RawMessage // array - could be interesting
	Distributions       json.RawMessage // object
}

type ProductRelease struct {
	Date                   string          // 2024-07-16
	Type                   string          // eap | release
	Version                string          // 2024.2
	MajorVersion           string          // 2024.2
	Patches                json.RawMessage // object
	Downloads              json.RawMessage // object
	LicenseRequired        bool
	Build                  string // 242.19890.39
	Whatsnew               string
	uninstallFeedbackLinks json.RawMessage // object
}

/*
Downloads object:

"linux: object"
"linuxARM64: object"
"macM1: object"
"mac: object"
"windows: object"
"windowsARM64: object"

*/

type ReleaseDownload struct {
	Link         string // https://download.jetbrains.com/rustrover/RustRover-242.19890.39.tar.gz
	Size         int    // 859881903
	ChecksumLink string // https://download.jetbrains.com/rustrover/RustRover-242.19890.39.tar.gz.sha256
}

func (d *ReleaseDownload) String() string {
	return fmt.Sprintf(d.Link)
}

func (r *ProductRelease) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "\ttype: %s\n", r.Type)
	fmt.Fprintf(&b, "\tversion: %s\n", r.Version)
	fmt.Fprintf(&b, "\tdownloads:\n")
	return b.String()
}

func (p *JetbrainsProduct) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "name: %s\ncode: %s\n", p.Name, p.Code)
	for _, r := range p.Releases {
		if r.Type == "release" {
			fmt.Fprintf(&b, "%s", r.String())
		}
	}

	return b.String()
}

func main() {
	resp, err := http.Get("https://data.services.jetbrains.com/products")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(string(body))

	var products []JetbrainsProduct
	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range products {
		fmt.Println(p.String())
	}

}
