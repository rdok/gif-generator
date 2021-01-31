package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/go-playground/colors.v1"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"strconv"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	buffer := lissajous(request)
	body := buffer.String()
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(body))

	return events.APIGatewayProxyResponse{
		Body:            base64Encoded,
		StatusCode:      200,
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "image/gif",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}

func lissajous(request events.APIGatewayProxyRequest) bytes.Buffer {
	query := request.QueryStringParameters

	backgroundColorInput := "#000000"
	if value, ok := query["background-color"]; ok {
		backgroundColorInput = "#" + value
	}

	lineColorInput := "#00FF00"
	if value, ok := query["line-color"]; ok {
		lineColorInput = "#" + value
	}
	imgCanvasSize := 100
	if value, ok := query["image-size"]; ok {
		imgCanvasSize, _ = strconv.Atoi(value)
	}

	backgroundColorParsed, _ := colors.Parse(backgroundColorInput)
	backgroundColorRGBA := backgroundColorParsed.ToRGBA()

	var backgroundColor = color.RGBA{
		R: backgroundColorRGBA.R,
		G: backgroundColorRGBA.G,
		B: backgroundColorRGBA.B,
		A: uint8(backgroundColorRGBA.A),
	}

	parseLineColor, _ := colors.Parse(lineColorInput)
	lineColorRGBA := parseLineColor.ToRGBA()
	var lineColor = color.RGBA{
		R: lineColorRGBA.R,
		G: lineColorRGBA.G,
		B: lineColorRGBA.B,
		A: uint8(lineColorRGBA.A),
	}

	var palette = []color.Color{backgroundColor, lineColor}

	const (
		blackIndex = 1 // next color in palette
	)

	const (
		oscillatorRevolutions = 5
		angularResolution     = 0.001
		totalAnimationFrames  = 64
		frameDelaysIn10sMS    = 8
	)

	relativeFreqYOscillator := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: totalAnimationFrames}
	phaseDifferences := 0.0

	for frameIndex := 0; frameIndex < totalAnimationFrames; frameIndex++ {
		rect := image.Rect(0, 0, 2*imgCanvasSize+1, 2*imgCanvasSize+1)
		img := image.NewPaletted(rect, palette)
		for phaseIndex := 0.0; phaseIndex < oscillatorRevolutions*2*math.Pi; phaseIndex += angularResolution {
			x := math.Sin(phaseIndex)
			y := math.Sin(phaseIndex * relativeFreqYOscillator * phaseDifferences)
			img.SetColorIndex(
				imgCanvasSize+int(x*float64(imgCanvasSize)+0.5),
				imgCanvasSize+int(y*float64(imgCanvasSize)+0.5),
				blackIndex,
			)
		}
		phaseDifferences += 0.1
		animation.Delay = append(animation.Delay, frameDelaysIn10sMS)
		animation.Image = append(animation.Image, img)
	}
	var buffer = bytes.Buffer{}
	_ = gif.EncodeAll(&buffer, &animation) // Ignore encoding errors
	return buffer
}
