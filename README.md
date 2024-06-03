# govipsblank
Small utility for [https://github.com/davidbyttow/govips](govips) which can be used to create a blank image with a background color. (useful for layering/compositing)  

## Usage
```go
import "govipsblank"

func main() {
  image_ref, err := govipsblank.MakeBlankImage(1200, 800, vips.ColorRBA{
    R: 255,
    G: 0,
    B: 0,
    A: 255
  })
}
```
