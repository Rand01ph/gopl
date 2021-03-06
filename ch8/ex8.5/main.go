// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
    "flag"
    "image"
    "image/color"
    "math/cmplx"
    "sync"
)

var workers int

func init() {
    flag.IntVar(&workers, "worker", 4, "goroutines number")
}

func main() {
    flag.Parse()
    Render(workers)
}

func Render(workers int) {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )

    ch := make(chan int, height)
    for py := 0; py < height; py++ {
        ch <- py
    }
    close(ch)

    wg := sync.WaitGroup{}

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for c := range ch {
                y := float64(c)/height*(ymax-ymin) + ymin
                for px := 0; px < width; px++ {
                    x := float64(px)/width*(xmax-xmin) + xmin
                    z := complex(x, y)
                    // Image point (px, py) represents complex value z.
                    img.Set(px, c, mandelbrot(z))
                }
            }
        }()
    }
    wg.Wait()
    //png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}
