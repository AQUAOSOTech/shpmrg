# shpmrg

`shpmrg` is a utility for merging shapefiles.

The utility is simplistic and assumes all the shapefiles have the same fields and types of objects. It will not work if a shapefile mixes points and polygons, for example. Or if one of the shapefiles has fields that the other shapefiles lack.

## Usage

Releases are available for mac, windows and linux at:

- https://github.com/AQUAOSOTech/shpmrg/releases

Example of merging shapefiles in the folder "myshapefiles" into the file "output_file.shp"

```bash
shpmrg -i myshapefiles/*.shp -o output_file.shp
```

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
