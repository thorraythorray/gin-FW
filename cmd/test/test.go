// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gillesdemey/go-dicom"
// )

// func main() {

// 	bytes, _ := os.ReadFile("C:\\Users\\leidong\\Downloads\\1662948907500-0d5940f1.dcm")

// 	parser, _ := dicom.NewParser()
// 	data, _ := parser.Parse(bytes)

// 	for _, elem := range data.Elements {
// 		fmt.Printf("%+v\n", &elem)
// 	}

// }

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/googlecloudplatform/go-dicom-parser/dicom"
)

func main() {
	r, err := os.Open("C:\\Users\\leidong\\Downloads\\1662948907500-0d5940f1.dcm")
	if err != nil {
		log.Fatalf("os.Open(_) => %v", err)
	}
	dataSet, err := dicom.Parse(r)
	if err != nil {
		log.Fatalf("dicom.Parse(_) => %v", err)
	}

	for tag, element := range dataSet.Elements {
		// fmt.Println(tag, element.VR, element.ValueField)
		fmt.Println(tag, element.VR, element.VR.Name)
	}
}
