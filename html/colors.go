package html

import (
	"fmt"
	"html/template"
	"strings"
)

func dayBackground(colors []string) template.CSS {

	if len(colors) == 1 {
		return template.CSS("background: " + colors[0] + ";")
	}

	gradients := make([]string, 0, len(colors)-1)

	for i, color := range colors[1:] {
		offset := i * 8

		gradients = append(gradients, fmt.Sprintf(
			`repeating-linear-gradient(
                135deg,
                transparent %dpx,
                transparent %dpx,
                %s %dpx,
                %s %dpx
            )`,
			offset,
			offset+4,
			color,
			offset+4,
			color,
			offset+6,
		))
	}

	return template.CSS(
		fmt.Sprintf(
			"background-color: %s; background-image: %s;",
			colors[0],
			strings.Join(gradients, ", "),
		),
	)
}