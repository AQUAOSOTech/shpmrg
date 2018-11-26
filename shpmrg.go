package main

import (
    "flag"
    "fmt"
    "github.com/jonas-p/go-shp"
    "os"
    "path/filepath"
    "sync"
    "sync/atomic"
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

    var allFields []shp.Field
    fieldNameToIndex := make(map[string]int)

    // pass 1, get all possible allFields
    for i, shapePath := range fileMatches {
        shapefile, err := shp.Open(shapePath)
        if err != nil {
            fmt.Println("Problem reading", shapePath, ", skipping. ", err)
            continue
        }

        fmt.Println("Processing shapefile allFields", shapePath, "(", i+1, "of", len(fileMatches), ")")

        localFields := shapefile.Fields()
        var fieldName string
        var fieldIndex int
        for _, localField := range localFields {
            fieldName = string(localField.Name[:11])
            if _, exists := fieldNameToIndex[fieldName]; !exists {
                allFields = append(allFields, localField)
                fieldIndex = len(allFields) - 1
                fieldNameToIndex[fieldName] = fieldIndex
            }
        }
    }

    err = outputFile.SetFields(allFields)
    if err != nil {
        fmt.Println("Failed setting output shapefile allFields, aborting!", err, allFields)
        os.Exit(1)
    }

    var rowCursor int64
    var filesProcessed int64

    // pass 2, copy shapefiles
    var wg sync.WaitGroup
    var writelock sync.Mutex
    for _, shPath := range fileMatches {
        sf, err := shp.Open(shPath)
        if err != nil {
            fmt.Println("Problem reading", shPath, ", skipping. ", err)
            continue
        }

        wg.Add(1)

        go func(shapePath string, shapefile *shp.Reader){
            localFields := shapefile.Fields()

            // loop through all features in the shapefile
            for shapefile.Next() {
                localRow, shape := shapefile.Shape()

                // print feature
                //fmt.Println(reflect.TypeOf(shape).Elem(), shape.BBox())
                writelock.Lock()
                outputFile.Write(shape)
                writelock.Unlock()

                // print attributes
                var remoteKey int
                var fieldName string
                for localKey, field := range localFields {
                    val := shapefile.ReadAttribute(localRow, localKey)
                    //fmt.Printf("\t%v: %v\row", f, val)
                    fieldName = string(field.Name[:11])
                    remoteKey = fieldNameToIndex[fieldName]
                    writelock.Lock()
                    err = outputFile.WriteAttribute(int(rowCursor), remoteKey, val)
                    writelock.Unlock()
                    if err != nil {
                        fmt.Println("Failed writing attribute, skipping. ", localKey, val, err)
                        continue
                    }
                }

                atomic.AddInt64(&rowCursor, 1)
                if rowCursor % 10000 == 0 {
                    fmt.Println("Total shapes processed:", rowCursor)
                }
            }

            final := atomic.AddInt64(&filesProcessed, 1)
            err = shapefile.Close()
            if err != nil {
                fmt.Println("Failed closing shapefile", shapePath, err)
            } else {
                fmt.Println("Finished", shapePath, "(", final, "of", len(fileMatches), ")")
            }
        }(shPath, sf)
    }

    wg.Wait()

    outputFile.Close()
}
