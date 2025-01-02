package viewer

import (
	"github.com/schollz/progressbar/v3"
)

func Loading(text string) {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription(text),
		progressbar.OptionSetWidth(20),
		progressbar.OptionFullWidth(), // important for clearing the line completely
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerHead:    "█",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)	
	bar.Add(1)
}
