# shpmrg

`shpmrg` is a multithreaded utility for working with shapefiles.

The utility is simplistic and assumes all the shapefiles have the same fields and types of objects. It will not work if a shapefile mixes points and polygons, for example. Different attribute fields between shapefiles should be handled, but may not work perfectly.

## Usage

Releases are available for mac, windows and linux at:

- https://github.com/AQUAOSOTech/shpmrg/releases

### Example of merging shapefiles

Get shapefiles in the folder "myshapefiles" and extract into the file "output_file.shp"

```bash
shpmrg -i myshapefiles/*.shp -o output_file.shp merge
```

```text
Total shapes processed: 10000
Finished /Users/jpx/myshapefiles/a.shp ( 1 of 3 )
Finished /Users/jpx/myshapefiles/b.shp ( 2 of 3 )
Total shapes processed: 20000
Total shapes processed: 30000
Total shapes processed: 40000
Total shapes processed: 50000
Finished /Users/jpx/myshapefiles/c.shp ( 3 of 3 )
Processed 51002
Done
```

Note: this program consumes approximately 2 CPU cores and is primarily limited by disk speed.

### Example of extracting attributes into CSV

Get shapefiles in the folder "myshapefiles" and extract into the file "attrs.csv"

```bash
shpmrg -i myshapefiles/*.shp -o attrs.csv extract-attrs
```

```text
Total shapes processed: 10000
Finished /Users/jpx/myshapefiles/a.shp ( 1 of 3 )
Total shapes processed: 20000
Finished /Users/jpx/myshapefiles/b.shp ( 2 of 3 )
Total shapes processed: 30000
Total shapes processed: 40000
Total shapes processed: 50000
Finished /Users/jpx/myshapefiles/c.shp ( 3 of 3 )
Processed 51002
Done
```

Note: this program consumes nearly all CPU, but is partly limited by disk speed.

## Shape Type flag `-t`

By default, the utility expects the shapes to be polygons.
The `-t < integer >` flag allows this to be changed.

The list of types is at https://godoc.org/github.com/jonas-p/go-shp#ShapeType

Here is a recent list, but this may change:

| Shape       | int|
|-------------|---|
| NULL        | 0 |
| POINT       | 1 |
| POLYLINE    | 3 |
| POLYGON     | 5 |
| MULTIPOINT  | 8 |
| POINTZ      | 11 |
| POLYLINEZ   | 13 |
| POLYGONZ    | 15 |
| MULTIPOINTZ | 18 |
| POINTM      | 21 |
| POLYLINEM   | 23 |
| POLYGONM    | 25 |
| MULTIPOINTM | 28 |
| MULTIPATCH  | 31 |


## License

MIT

Copyright (c) 2018 - present AQUAOSO Technologies, PBC

See the LICENSE file in this repository.

Credits to:

- https://github.com/jonas-p/go-shp
