package utils

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

type GifWriter struct {
	File       *os.File
	GIF        *gif.GIF
	FileName   string
	Recording  bool
	FrameCount int
	StopCount  int
}

func NewGifWriter(fileName string, stopCount int) *GifWriter {
	return &GifWriter{
		FileName:  fileName,
		StopCount: stopCount,
		Recording: true,
	}
}

func (writer *GifWriter) RecordGif(img image.Image) error {
	if !writer.Recording {
		return nil
	}
	if writer.FrameCount > writer.StopCount {
		err := gif.EncodeAll(writer.File, writer.GIF)
		defer writer.File.Close()
		if err != nil {
			return err
		}
		writer.Recording = false
		fmt.Println("Done gif!")
		return nil
	}
	if writer.File == nil {
		writer.GIF = &gif.GIF{}
		file, err := os.Create(writer.FileName)
		if err != nil {
			writer.Recording = false
			return err
		}
		writer.File = file
	}
	pImage := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(pImage, pImage.Rect, img, img.Bounds().Min, draw.Over)
	writer.GIF.Image = append(writer.GIF.Image, pImage)
	writer.GIF.Delay = append(writer.GIF.Delay, 0)
	writer.FrameCount += 1
	return nil
}
