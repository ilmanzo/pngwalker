# pngwalker
simple program to extract informations from PNG files

# Usage:

```
$ go build
$ ./pngwalker example.png
header: [137 80 78 71 13 10 26 10]
Offset: 8 chunk len=13, type: IHDR
Offset: 25 chunk len=93, type: PLTE
Offset: 122 chunk len=2173, type: IDAT
Offset: 2299 chunk len=0, type: IEND
```
