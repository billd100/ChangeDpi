package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	exif2 "github.com/dsoprea/go-exif/v2"
	exifcommon "github.com/dsoprea/go-exif/v2/common"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
)

var (
	dpi int = 105
)

func SetExifData(filepath string) error {
	jmp := jpegstructure.NewJpegMediaParser()

	intfc, err := jmp.ParseFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	sl := intfc.(*jpegstructure.SegmentList)

	_, err = sl.DropExif()
	if err != nil {
		log.Fatal(err)
	}

	im := exif2.NewIfdMapping()

	err = exif2.LoadStandardIfds(im)
	if err != nil {
		log.Fatal(err)
	}
	ti := exif2.NewTagIndex()
	rootIb := exif2.NewIfdBuilder(im, ti, exifcommon.IfdPathStandard, exifcommon.EncodeDefaultByteOrder)

	err = rootIb.AddStandardWithName("XResolution", []exifcommon.Rational{{Numerator: uint32(dpi), Denominator: uint32(1)}})
	if err != nil {
		log.Fatal(err)
	}
	err = rootIb.AddStandardWithName("YResolution", []exifcommon.Rational{{Numerator: uint32(dpi), Denominator: uint32(1)}})
	if err != nil {
		log.Fatal(err)
	}
	err = sl.SetExif(rootIb)
	if err != nil {
		log.Fatal(err)
	}
	b := new(bytes.Buffer)

	err = sl.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("/Users/billdavis/Downloads/new-label.jpg", b.Bytes(), 0644); err != nil {
		fmt.Printf("write file err: %v", err)
	}
	return nil
}

func main() {
	SetExifData("/Users/billdavis/Downloads/label.jpg")
}
