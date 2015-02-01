package main

import (
	"bytes"
    "encoding/base64"
    "fmt"
    "image"
    "image/png"
)

func Show(f func(int, int) [][]uint8) {
        const (
                dx = 256
                dy = 256
        )
        data := f(dx, dy)
        m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
        for y := 0; y < dy; y++ {
                for x := 0; x < dx; x++ {
                        v := data[y][x]
                        i := y*m.Stride + x*4
                        m.Pix[i] = v
                        m.Pix[i+1] = v
                        m.Pix[i+2] = 255
                        m.Pix[i+3] = 255
                }
        }
        ShowImage(m)
}

func ShowImage(m image.Image) {
        var buf bytes.Buffer
        err := png.Encode(&buf, m)
        if err != nil {
                panic(err)
        }
        enc := base64.StdEncoding.EncodeToString(buf.Bytes())
        fmt.Println("IMAGE:" + enc)
}

/* Everything above this line was part of the source code provided in the
 * required library. NOT written by me as part of the problem.
*/ 

/* Problem statement
 * Implement Pic. It should return a slice of length dy, each element of which
 * is a slice of dx 8-bit unsigned integers. When you run the program, it will
 * display your picture, interpreting the integers as grayscale (well,
 * bluescale) values.
 * The choice of image is up to you. Interesting functions include (x+y)/2,
 * x*y, and x^y (to compute the latter function, use math.Pow).
 * You need to use a loop to allocate each []uint8 inside the [][]uint8.
 * Use uint8(intValue) to convert between types.
*/
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	/* This needs to be interpreted like ([]([]uint8), dy) .i.e. slice of length
	 * "dy" where each element is a slice of uint8 (the size of this inner slice
	 * is not defined in this step. I tried p:= make([dy][dx]uint8) which is not
	 * legal. 
	*/
		for i := range p { // value is skipped, equivalent to i, _
			p[i] = make([]uint8, dx) // This is the inner slice of uint8
        }
        for y, row := range p { // row is the entire inner slice, this loop
                                // repeats dy times
			for x := range row { // x is index of inner slice, equivalent to
								 // x, _. This loop repeats dx times.
				row[x] = uint8(x * y) // don't need to use append here as we
				                      // used make to create the needed sizes
            }
        }
	return p // this data structure has dy number of slices, each of size dx.
}

func main () {
	Show(Pic)
}
