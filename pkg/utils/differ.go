package utils

import (
	"bytes"

	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func DiffPrettyText(a, b string) string {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(a, b, false)

	var buff bytes.Buffer
	green := color.New(color.FgGreen).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()
	for _, diff := range diffs {
		text := diff.Text

		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			_, _ = buff.WriteString(green(text))
		case diffmatchpatch.DiffDelete:
			_, _ = buff.WriteString(red(text))
		case diffmatchpatch.DiffEqual:
			_, _ = buff.WriteString(text)
		}
	}

	return buff.String()
}
