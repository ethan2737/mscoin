package logic

import (
	"strconv"
	"strings"
)

func buildSwapRateCacheKey(symbol string) string {
	replacer := strings.NewReplacer("/", "::", "-", "::")
	return "SWAP::" + replacer.Replace(symbol) + "::SWAP::RATE"
}

func parseSwapPriceString(value string) float64 {
	price, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0
	}
	return price
}
