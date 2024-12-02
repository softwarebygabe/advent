package colors

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type Color string

const (
	Red    Color = "red"
	Green  Color = "green"
	Yellow Color = "yellow"
	Blue   Color = "blue"
	Purple Color = "purple"
	Cyan   Color = "cyan"
	White  Color = "white"
)

func Printf(c Color, format string, args ...interface{}) {
	var color string
	switch c {
	case Red:
		color = colorRed
	case Green:
		color = colorGreen
	case Yellow:
		color = colorYellow
	case Blue:
		color = colorBlue
	case Purple:
		color = colorPurple
	case Cyan:
		color = colorCyan
	case White:
		color = colorWhite
	default:
		color = colorWhite
	}
	newArgs := []interface{}{color}
	newArgs = append(newArgs, args...)
	newArgs = append(newArgs, colorReset)
	fmt.Printf("%s"+format+"%s", newArgs...)
}

func Sprintf(c Color, format string, args ...interface{}) string {
	var color string
	switch c {
	case Red:
		color = colorRed
	case Green:
		color = colorGreen
	case Yellow:
		color = colorYellow
	case Blue:
		color = colorBlue
	case Purple:
		color = colorPurple
	case Cyan:
		color = colorCyan
	case White:
		color = colorWhite
	default:
		color = colorWhite
	}
	newArgs := []interface{}{color}
	newArgs = append(newArgs, args...)
	newArgs = append(newArgs, colorReset)
	return fmt.Sprintf("%s"+format+"%s", newArgs...)
}
