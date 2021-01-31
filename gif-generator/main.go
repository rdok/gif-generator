package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
)

func handler(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	buffer := lissajous()
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

func lissajous() bytes.Buffer {
	const (
		oscillatorRevolutions = 5
		angularResolution     = 0.001
		imgCanvasSize         = 100
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
				imgCanvasSize+int(x*imgCanvasSize+0.5),
				imgCanvasSize+int(y*imgCanvasSize+0.5),
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

var palette = []color.Color{color.White, color.Black}

const (
	blackIndex = 1 // next color in palette
)
