package main


import "log"
import "fmt"
import "os"
import "image"
import _ "image/jpeg"
import _ "image/png"


func print_img_size(fp string) {
    file, err := os.Open(fp)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)        
    }
    w := img.Bounds().Max.X
    h := img.Bounds().Max.Y
    fmt.Printf("w: %d h: %d\n", w, h)
}


func main() {
    args_exc_prog := os.Args[1:]
    len_args := len(args_exc_prog)
    if len_args > 0 {
        print_img_size(args_exc_prog[0])
    }else{
        fmt.Println("usage: ")
        fmt.Println("iwh imgage_file")
    }
}
