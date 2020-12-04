# dexif
Remove extension info from image, like Exif, TIFF.

# Install
System request go 1.11+

## macbook or linux
If you pc is macbook or linux, you can make file yourself, simple run `make build` from your terminal. After processed, there are three file in bin folder:
```txt
bin
├── dexif            for macosx
├── dexif.exe        for windows(x64)
└── dexif_linux      for linux(x64)
```

## windows or others
If you pc system is windows or others, you can build by go:
```shell
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -v -a -ldflags -s -installsuffix cgo  main.go bin/dexif.exe
```

# How to use
run `./dexif -f  /path/to/your/image` to remove extension information from image.

run `./dexif -d /path/to/your/image/folder` to batch remove extension information from image.
