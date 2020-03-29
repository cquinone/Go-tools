package main


import (
    "image"
    "image/gif"
    "image/png"
    "fmt"
    "os"
    "path/filepath"
    "bytes"
)

func main() {
    directory := "." + string(filepath.Separator)
    d, err := os.Open(directory)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()

    files, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("Reading PNGS from: "+ directory)

    file_names := make([]string, 0)
    for _, file := range files {
        if file.Mode().IsRegular() {
            if filepath.Ext(file.Name()) == ".png" {
              file_names = append(file_names, file.Name())
              fmt.Println(file.Name())
            }
        }
    }

    // load PNGS, convert to Image.image, convert Image.image to .gif
    // make palette, and construct out.gif
    delay := 100 
    final_gif := &gif.GIF{}
    for _, name := range file_names {
        f, _ := os.Open(name)
        temp_png_img,_ := png.Decode(f)
        buf := bytes.Buffer{}
        if err := gif.Encode(&buf, temp_png_img, nil); err != nil {
            fmt.Printf("Skipping file %s due to error in gif encoding:%s", name, err)
            continue
        }
        gif_img, _ := gif.Decode(&buf)
        f.Close()
        fmt.Println("ON IMAGE: ", name)
        final_gif.Image = append(final_gif.Image, gif_img.(*image.Paletted))
        final_gif.Delay = append(final_gif.Delay, delay)
    }
 
    // save to out.gif
    f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, final_gif)
}