/**
 * The programme use to remove image extension info, like Exif, TIFF, IPTC.
 * Support format: jpg, jpeg, png, gif, bmp, tiff
 *
 * File: main.go
 * Package: dexif
 * Author: Lee (lys@ezool.cn)
 * Date: 2020-12-04 10:08:17
 */
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// "github.com/chai2010/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	"golang.org/x/image/tiff"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		help()
		return
	}
	imgPath, localDir := "", ""
	flag.StringVar(&imgPath, "f", "", "The image path to remove extension info")
	flag.StringVar(&localDir, "d", "", "Remove all image extension info from this local dir")
	flag.Parse()
	var err error

	if imgPath != "" {
		if !exists(imgPath) {
			fmt.Println("[E] Image file not exist:", imgPath)
			return
		}
		err = Dexif(imgPath)
		if err == nil {
			fmt.Println("[I]", imgPath, " processed.")
		}
	} else if localDir != "" {
		if !exists(localDir) {
			fmt.Println("[E] Folder not exist:", localDir)
			return
		}
		dexifFromDir(localDir)
	} else {
		help()
		return
	}

	if err != nil {
		fmt.Println("[E]", err.Error())
	} else {
		fmt.Println("[I] finished.")
	}
}

func help() {
	fmt.Println(`
The programme will remove all extension info(like exif, tiff) from image.
IMPORTANT: backup your file first!!!

Usage: ./dexif [option]

option:
	-f: image file path 
	-d: the programme will scan this local dir and subdir

example:
	./dexif -f /path/to/image.jpg
	./dexif -d /path/to/folder/that/contain/image
	`)
}

//dexifFromDir remove extension info of image file from folder
func dexifFromDir(localDir string) {
	files, err := ioutil.ReadDir(localDir)
	if err == nil {
		for _, f := range files {
			if f.IsDir() {
				dexifFromDir(path.Join(localDir, f.Name()))
			} else {
				ext := getExt(f.Name())
				if isImage(ext) {
					err := Dexif(path.Join(localDir, f.Name()), ext)
					if err == nil {
						fmt.Println("[I]", path.Join(localDir, f.Name()), " processed.")
					} else {
						fmt.Println("[E]", err.Error())
					}
				}
			}
		}
	}
}

//Dexif 移除图片文件中的exif信息
func Dexif(filepath string, ex ...string) error {
	originFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer originFile.Close()

	var ext string
	if len(ex) > 0 {
		ext = ex[0]
	} else {
		ext = getExt(filepath)
	}

	decode, err := decode(originFile, ext)
	if err != nil {
		return err
	}

	// 缩略图
	destpath := filepath + ".tmp"
	thumb, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer thumb.Close()

	dst := image.NewRGBA(decode.Bounds())
	draw.ApproxBiLinear.Scale(dst, decode.Bounds(), decode, decode.Bounds(), draw.Over, nil)
	if ext == ".jpg" || ext == ".jpeg" {
		err = save(ext, thumb, dst, &jpeg.Options{Quality: 100})
	} else {
		err = save(ext, thumb, dst)
	}
	if err == nil {
		os.Rename(destpath, filepath)
	}
	return err
}

//save image
func save(ext string, f *os.File, m image.Image, opt ...*jpeg.Options) error {
	switch ext {
	case ".jpg", ".jpeg":
		var option *jpeg.Options
		if len(opt) > 0 {
			option = opt[0]
		} else {
			option = &jpeg.Options{Quality: 100}
		}
		return jpeg.Encode(f, m, option)

	case ".png":
		return png.Encode(f, m)

	case ".gif":
		o := &gif.Options{}
		return gif.Encode(f, m, o)

	case ".bmp":
		return bmp.Encode(f, m)

	case ".tiff":
		o := &tiff.Options{}
		return tiff.Encode(f, m, o)

		// case ".webp":
		// 	o := webp.Options{
		// 		Lossless: true,
		// 		Quality:  100.0,
		// 		Exact:    true,
		// 	}
		// 	webp.Encode(f, m, &o)
	}
	return errors.New("File not support")
}

//decode get the image decode object
func decode(f *os.File, ext string) (image.Image, error) {
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Decode(f)

	case ".png":
		return png.Decode(f)

	case ".gif":
		return gif.Decode(f)

	case ".bmp":
		return bmp.Decode(f)

	case ".tiff":
		return tiff.Decode(f)

		// case ".webp":
		// 	return webp.Decode(f)
	}
	return nil, errors.New("File not support")
}

//exists check the path if exist
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//getExt get file ext
func getExt(p string) string {
	return strings.ToLower(path.Ext(p))
}

//isImage check whether the file is allow image
func isImage(ext string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".git", ".bmp", ".tiff":
		return true
	}
	return false
}
