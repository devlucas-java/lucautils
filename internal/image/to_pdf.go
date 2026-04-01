package image

import (
	"os"

	"github.com/jung-kurt/gofpdf"
)

func ImagesToPDF(images []string, output string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, img := range images {
		if _, err := os.Stat(img); os.IsNotExist(err) {
			continue
		}

		pdf.AddPage()
		width, height := pdf.GetPageSize()
		pdf.ImageOptions(img, 0, 0, width, height, false, gofpdf.ImageOptions{ImageType: "", ReadDpi: true}, 0, "")
	}

	return pdf.OutputFileAndClose(output)
}
