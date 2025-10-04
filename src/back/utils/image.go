package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"

	"github.com/anthonynsimon/bild/transform"
)

type pngPool struct {
	pool *sync.Pool
}

func (p *pngPool) Get() *png.EncoderBuffer {
	return p.pool.Get().(*png.EncoderBuffer)
}

func (p *pngPool) Put(b *png.EncoderBuffer) {
	p.pool.Put(b)
}

var encoderPool = &pngPool{
	pool: &sync.Pool{New: func() any { return new(png.EncoderBuffer) }},
}

var pngEncoder = png.Encoder{
	CompressionLevel: png.NoCompression,
	BufferPool:       encoderPool,
}

func ExportImageToPng(img image.Image, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	if err := pngEncoder.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode image to %s: %w", filePath, err)
	}

	return nil
}

var zeroPoint = image.Point{0, 0}

func Tile2x2(topLeft, topRight, bottomLeft, bottomRight *image.RGBA) (*image.RGBA, error) {
	topLeftBounds := topLeft.Bounds()
	if topLeftBounds != topRight.Bounds() ||
		topLeftBounds != bottomLeft.Bounds() ||
		topLeftBounds != bottomRight.Bounds() {
		return nil, fmt.Errorf("all images must have the same dimensions")
	}

	width, height := topLeftBounds.Dx(), topLeftBounds.Dy()
	combinedImg := image.NewRGBA(image.Rect(0, 0, width*2, height*2))

	var wg sync.WaitGroup
	wg.Add(4)
	processingFn := func(rect image.Rectangle, img *image.RGBA) {
		defer wg.Done()
		draw.Draw(combinedImg, rect, img, zeroPoint, draw.Src)
	}

	go processingFn(image.Rect(0, 0, width, height), topLeft)
	go processingFn(image.Rect(width, 0, width*2, height), topRight)
	go processingFn(image.Rect(0, height, width, height*2), bottomLeft)
	go processingFn(image.Rect(width, height, width*2, height*2), bottomRight)

	wg.Wait()

	return combinedImg, nil
}

func Stack(images ...*image.RGBA) (*image.RGBA, error) {
	if len(images) == 0 {
		return nil, fmt.Errorf("no images provided to stack")
	}

	width := images[0].Bounds().Dx()
	totalHeight := 0

	for i, img := range images {
		if img.Bounds().Dx() != width {
			return nil, fmt.Errorf("image %d has a different width: expected %d, got %d", i, width, img.Bounds().Dx())
		}
		totalHeight += img.Bounds().Dy()
	}

	stackedImg := image.NewRGBA(image.Rect(0, 0, width, totalHeight))
	wg := sync.WaitGroup{}
	currentY := 0

	for _, img := range images {
		wg.Add(1)
		go func(imgToDraw *image.RGBA, yOffset int) {
			defer wg.Done()
			height := imgToDraw.Bounds().Dy()
			draw.Draw(stackedImg, image.Rect(0, yOffset, width, yOffset+height), imgToDraw, zeroPoint, draw.Src)
		}(img, currentY)
		currentY += img.Bounds().Dy()
	}
	wg.Wait()
	return stackedImg, nil
}

func ConvertValuesToRGBAImage(
	originalWidth, originalHeight int,
	imageData []float32,
	minValue, maxValue float32,
) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, originalWidth, originalHeight))
	scale := float32(0)
	if maxValue-minValue > 0 {
		scale = 255.0 / (maxValue - minValue)
	}

	const workersCount = 4
	wg := sync.WaitGroup{}
	rowsPerWorker := originalHeight / workersCount

	processRows := func(startY, endY int) {
		defer wg.Done()
		for y := startY; y < endY; y++ {
			for x := range originalWidth {
				sourceIdx := y*originalWidth + x
				pixelValue := imageData[sourceIdx]

				pixelValue = Clamp(pixelValue, minValue, maxValue)
				grayValue := uint8((pixelValue - minValue) * scale)
				rgbaColor := GetFalseColor(grayValue)

				destIdx := y*img.Stride + x*4
				img.Pix[destIdx+0] = rgbaColor.R
				img.Pix[destIdx+1] = rgbaColor.G
				img.Pix[destIdx+2] = rgbaColor.B
				img.Pix[destIdx+3] = rgbaColor.A
			}
		}
	}

	wg.Add(workersCount)
	for i := range workersCount {
		startY := i * rowsPerWorker
		endY := startY + rowsPerWorker
		if i == workersCount-1 {
			endY = originalHeight
		}

		go processRows(startY, endY)
	}

	wg.Wait()
	return img, nil
}

func TransformRotate180(src *image.RGBA) *image.RGBA {
	return transform.Rotate(src, 180, nil)
}

func GetFalseColor(value uint8) color.RGBA {
	g := min(uint8(-0.00043*float32(value)*float32(value)+1.087*float32(value)+3.633), 255)
	b := min(uint8(-0.00080*float32(value)*float32(value)+1.119*float32(value)+17.232), 255)
	return color.RGBA{value, g, b, 255}
}
