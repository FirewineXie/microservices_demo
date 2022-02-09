package biz

import (
	"fmt"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"time"
)

type ShippingUseCase struct {
	logger *zap.Logger
}
type Quote struct {
	Dollars uint32
	Cents   uint32
}

func (c *ShippingUseCase) CreateQuoteFromCount(count int) Quote {
	return c.CreateQuoteFromFloat(c.quoteByCountFloat(count))
}

// quoteByCountFloat takes a number of items and generates a price quote represented as a float.
func (c *ShippingUseCase) quoteByCountFloat(count int) float64 {
	if count == 0 {
		return 0
	}
	count64 := float64(count)
	var p = 1 + (count64 * 0.2)
	return count64 + math.Pow(3, p)
}

// CreateQuoteFromFloat takes a price represented as a float and creates a Price struct.
func (c *ShippingUseCase) CreateQuoteFromFloat(value float64) Quote {
	units, fraction := math.Modf(value)
	return Quote{
		uint32(units),
		uint32(math.Trunc(fraction * 100)),
	}
}
// seeded determines if the random number generator is ready.
var seeded bool = false


func (c *ShippingUseCase) CreateTrackingId(salt string) string {
	if !seeded {
		rand.Seed(time.Now().UnixNano())
		seeded = true
	}

	return fmt.Sprintf("%c%c-%d%s-%d%s",
		getRandomLetterCode(),
		getRandomLetterCode(),
		len(salt),
		getRandomNumber(3),
		len(salt)/2,
		getRandomNumber(7),
	)
}
// getRandomLetterCode generates a code point value for a capital letter.
func getRandomLetterCode() uint32 {
	return 65 + uint32(rand.Intn(25))
}

// getRandomNumber generates a string representation of a number with the requested number of digits.
func getRandomNumber(digits int) string {
	str := ""
	for i := 0; i < digits; i++ {
		str = fmt.Sprintf("%s%d", str, rand.Intn(10))
	}

	return str
}

func NewShippingUseCase(logger *zap.Logger) *ShippingUseCase {
	return &ShippingUseCase{
		logger: logger,
	}
}
