package vtex

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatPrice(price float32) string {
	return strings.ReplaceAll(fmt.Sprintf("%.2f", price), ".", ",")
}

func ConvertToWebpURL(imageURL string) string {
	lastIndex := strings.LastIndex(imageURL, ".")
	if lastIndex == -1 {
		return imageURL
	}

	if imageURL[lastIndex:] != ".svg" {
		baseURL := imageURL[:lastIndex]
		return baseURL + ".webp"
	}

	return imageURL
}

func ChangeImageSize(imageURL string, newWidth, newHeight int) string {
	parts := strings.Split(imageURL, "/")

	for i, part := range parts {
		if part == "ids" {
			// fmt.Print(part, parts[i+1])
			sizeParts := strings.Split(parts[i+1], "-")
			if len(sizeParts) != 3 {
				// Se não houver exatamente três partes (incluindo os traços), ignorar e continuar
				parts[i+1] = fmt.Sprintf("%s-%d-%d", parts[i+1], newWidth, newHeight)
				continue
			}

			// Substituir a largura e a altura pelos novos valores
			sizeParts[1] = strconv.Itoa(newWidth)
			sizeParts[2] = strconv.Itoa(newHeight)
			parts[i+1] = strings.Join(sizeParts, "-")
		}
	}

	// Reunir as partes de volta em uma única URL
	return strings.Join(parts, "/")
}

func OptimizeImageURL(imageURL string, width, height int) string {
	webpImage := ConvertToWebpURL(imageURL)

	return ChangeImageSize(webpImage, width, height)
}
