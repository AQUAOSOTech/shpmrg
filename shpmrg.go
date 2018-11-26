package main

import (
    "flag"
    "fmt"
    "github.com/jonas-p/go-shp"
    "os"
    "path/filepath"
)

var inPath = flag.String("i", "", "Input file glob path to shapefiles")
var outPath = flag.String("o", "", "Output file location")
var shapeType = flag.Int("t", 25, "Default is polygon; Shape type from https://godoc.org/github.com/jonas-p/go-shp#ShapeType")

func main() {
    flag.Parse()
    if *inPath == "" {
        fmt.Println("Missing -i input file(s)")
        flag.PrintDefaults()
        os.Exit(1)
    }
    if *outPath == "" {
        fmt.Println("Missing -o output location")
        flag.PrintDefaults()
        os.Exit(1)
    }

    fileMatches, err := filepath.Glob(*inPath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    if len(fileMatches) < 1 {
        fmt.Println("No matches for input pattern")
        os.Exit(1)
    }
    outputFile, err := shp.Create(*outPath, shp.ShapeType(*shapeType))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    wroteFields := false

    for i, shapePath := range fileMatches {
        shapefile, err := shp.Open(shapePath)
        if err != nil {
            fmt.Println("Problem reading", shapePath, ", skipping. ", err)
            continue
        }

        // fields from the attribute table (DBF)
        fields := shapefile.Fields()
        if !wroteFields {
            wroteFields = true
            err = outputFile.SetFields(fields)
            if err != nil {
                fmt.Println("Failed setting output shapefile fields, aborting!", err, fields)
                os.Exit(1)
            }
        }

        // loop through all features in the shapefile
        fmt.Println("Adding shapefile", shapePath, "(", i+1, "of", len(fileMatches), ")")
        for shapefile.Next() {
            row, shape := shapefile.Shape()

            // print feature
            //fmt.Println(reflect.TypeOf(shape).Elem(), shape.BBox())
            outputFile.Write(shape)

            // print attributes
            for k/*, field*/ := range fields {
                val := shapefile.ReadAttribute(row, k)
                //fmt.Printf("\t%v: %v\row", f, val)
                err = outputFile.WriteAttribute(row, k, val)
                if err != nil {
                    fmt.Println("Failed writing attribute, skipping. ", k, val, err)
                    continue
                }
            }
        }

        err = shapefile.Close()
        if err != nil {
            fmt.Println("Failed closing shapefile", shapePath, err)
            continue
        }
    }

    outputFile.Close()
}
