package utils

import "log"

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	log.Println(string(color), message, string(ColorReset))
}

func ColorizeBlack(msg string) {
	colorize(ColorBlack, msg)
}

func ColorizeRed(msg string) {
	colorize(ColorRed, msg)
}

func ColorizeGreen(msg string) {
	colorize(ColorGreen, msg)
}

func ColorizeYellow(msg string) {
	colorize(ColorYellow, msg)
}

func ColorizeBlue(msg string) {
	colorize(ColorBlue, msg)
}

func ColorizeReset(msg string) {
	colorize(ColorReset, msg)
}
