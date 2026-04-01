package main

import (
	"fmt"

	"github.com/devlucas-java/lucautils/internal/image"
)

func main() {
	images := []string{
		"./img1.png",
		"./img2.png",
		"./img3.png",
		"./img4.png",
	}

	output := "./output.pdf"

	err := image.ImagesToPDF(images, output)
	if err != nil {
		fmt.Println("Erro ao gerar PDF:", err)
		return
	}

	fmt.Println("PDF gerado com sucesso em:", output)
}
