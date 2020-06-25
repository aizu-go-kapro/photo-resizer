package converters

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/aizu-go-kapro/photo-resizer/pkg/photos/domain/photos"
)

// edited: https://github.com/golang/tour/blob/master/pic/pic.go#L35
func writeImage(m image.Image) (*os.File, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*-temp-photo.png")
	if err != nil {
		return nil, err
	}

	err = png.Encode(file, m)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getAnimatedPhoto() (photos.Photo, error) {
	var animatedImages []image.Image
	fileInfos, err := ioutil.ReadDir(filepath.Join("test-resources", "motions"))
	if err != nil {
		return photos.Photo{}, err
	}

	for _, fileInfo := range fileInfos {
		file, err := os.Open(filepath.Join("test-resources", "motions", fileInfo.Name()))
		if err != nil {
			return photos.Photo{}, err
		}

		pngImage, err := png.Decode(file)
		if err != nil {
			return photos.Photo{}, err
		}
		animatedImages = append(animatedImages, pngImage)
	}
	return photos.NewPhoto(animatedImages), nil
}

func getOptimizedGifResource() (*os.File, error) {
	file, err := os.Open(filepath.Join("test-resources", "test-img.gif"))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getJpegResource() (*os.File, error) {
	file, err := os.Open(filepath.Join("test-resources", "smptebars.jpg"))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getPngResource() (*os.File, error) {
	file, err := os.Open(filepath.Join("test-resources", "rgbtest.png"))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func TestUploadImageToImgur(t *testing.T) {
	imageData, err := getJpegResource()
	if err != nil {
		t.Fatal(err)
	}

	imageUrl, err := uploadImageToImgur(imageData, "946667977ad0113")
	t.Log("Image Uploaded To", imageUrl)
}

// upload to imgur ref: https://stackoverflow.com/questions/53426576/can-i-upload-image-to-imgur-via-golang
func uploadImageToImgur(image io.Reader, token string) (string, error) {
	var buf = new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, _ := writer.CreateFormFile("image", "dont care about name")
	_, err := io.Copy(part, image)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", "https://api.imgur.com/3/image", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Client-ID "+token)

	client := &http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	decodedResponse := ImgurUploadResponse{}
	err = json.NewDecoder(res.Body).Decode(&decodedResponse)
	if err != nil {
		return "", err
	}

	return decodedResponse.Data.Link, nil
}

type ImgurUploadResponse struct {
	Data    ImgurDataResponse `json:"data"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
}

type ImgurDataResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateTime    int    `json:"datetime"`
	Type        string `json:"type"`
	Animated    bool   `json:"animated"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Size        int    `json:"height"`
	Views       int    `json:"views"`
	BandWidth   int    `json:"bandwidth"`
	Vote        string `json:"vote"`
	Favorite    bool   `json:"favorite"`
	NSFW        string `json:"nsfw"`
	Section     string `json:"section"`
	AccountUrl  string `json:"account_url"`
	AccountId   int    `json:"account_id"`
	IsAd        bool   `json:"is_ad"`
	InMostViral bool   `json:"in_most_viral"`
	AdType      int    `json:"ad_type"`
	AdUrl       string `json:"ad_url"`
	InGallery   bool   `json:"in_gallery"`
	DeleteHash  string `json:"x70po4w7BVvSUzZ"`
	Name        string `json:"name"`
	Link        string `json:"link"`
}
