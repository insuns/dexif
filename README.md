# dexif
dexif write in [go](https://golang.org/).It use to remove extension info from image, like Exif, TIFF.

# How to build
System request go 1.11+

## macbook or linux
If you pc is macosx or linux, you can make file yourself, simple run `make build` from your terminal. After processed, there are three file in bin folder:
```txt
bin
├── dexif            for macosx
├── dexif.exe        for windows(x64)
└── dexif_linux      for linux(x64)
```

## windows or others
If you pc system is windows or others, you can compile by go:
```shell
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -v -a -ldflags -s -installsuffix cgo  main.go bin/dexif.exe
```

more go compile option see [go compile](https://golang.org/cmd/compile/)

# How to use
Download bin file from [dexif-amd64.tar.gz](https://github.com/insuns/dexif/releases/download/v0.0.1/dexif-amd64.tar.gz), or build yourself, and then:

run `./dexif -f  /path/to/your/image` to remove extension information from image.

run `./dexif -d /path/to/your/image/folder` to batch remove extension information from image.

run `./dexif -h` to see more help info.