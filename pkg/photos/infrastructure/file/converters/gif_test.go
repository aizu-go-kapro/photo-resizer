package converters

import (
	"fmt"
	"github.com/pkg/browser"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestGifConverter_Open(t *testing.T) {
	m := Image{}
	filePath := writeImage(m)

	fmt.Println("The file path : ", filePath)
	_ = browser.OpenURL(filePath)

	//gifConverter := NewGifConverter()
	//gifFile := getOptimizedGifResource()

	//photos, err := gifConverter.Open(gifFile)
	//if err != nil {
	//	t.Error(err)
	//}

	//for _, image := range photos.Images() {
	//	url := base64Image(image)
	//	fmt.Println(url)
	//	browser.OpenURL("http://localhost:63342/photo-resizer/test/resources/b64viewer.html?png=" + url)
	//}
}

// edited: https://github.com/golang/tour/blob/master/pic/pic.go#L35
func writeImage(m image.Image) string {
	file, err := ioutil.TempFile(os.TempDir(), "tempPhoto.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, m)
	if err != nil {
		panic(err)
	}
	return file.Name()
}

func getOptimizedGifResource() *os.File {
	file, _ := os.Open("test-resources/nyancat.gif")
	return file
}

/*
About test cases:
All frame; test-resources/rgb.gif: http://tech.nitoyon.com/ja/blog/2016/01/07/go-animated-gif-gen/
Optimized; test-resources/nyancat.gif: https://blog.zhaytam.com/2018/08/21/creating-gifs-using-python-pillow/
*/

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(y), uint8(y)}
}
