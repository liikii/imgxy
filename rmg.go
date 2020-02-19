package main

import "os"
import "flag"
import "log"
import "fmt"
import "image"
import _ "image/jpeg"
import "image/png"
import "github.com/nfnt/resize"
import "path/filepath"


func img_resize(fp string, w uint, h uint) {
    file, err := os.Open(fp)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)        
    }
    img2 := resize.Resize(w, h, img, resize.Bilinear)
    fp = filepath.Join(filepath.Dir(fp), "rmg_" + filepath.Base(fp))
    out, err := os.Create(fp)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()
    // write new image to file
    png.Encode(out, img2)
    fmt.Println(fp + " ok")
}


func main() {
    var w uint
    var h uint
    var fl string


    flag.UintVar(&w, "w", 0, "max width")
    flag.UintVar(&h, "h", 0, "max height")
    flag.StringVar(&fl, "f", "", "images file path")
    flag.Parse()

    img_resize(fl, w, h)
}
