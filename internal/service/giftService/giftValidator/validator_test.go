package giftValidator

import (
	"testing"

	"gift-buyer/internal/config"

	"github.com/gotd/td/tg"
	"github.com/stretchr/testify/assert"
)

func TestNewGiftValidator(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	assert.NotNil(t, validator)
	assert.Equal(t, criterias, validator.criteria)
	assert.Equal(t, int64(5000), validator.totalStarCap)
	assert.False(t, validator.testMode)
	assert.False(t, validator.LimitedStatus)
}

func TestGiftValidator_IsEligible_ValidGift(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	gift := &tg.StarGift{
		ID:    1,
		Stars: 500, // Within price range
	}

	giftRequire, result := validator.IsEligible(gift)
	assert.True(t, result)
	assert.NotNil(t, giftRequire)
	assert.Equal(t, int64(10), giftRequire.CountForBuy)
	assert.Equal(t, []int{1}, giftRequire.ReceiverType)
}

func TestGiftValidator_IsEligible_PriceTooLow(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	gift := &tg.StarGift{
		ID:    1,
		Stars: 50, // Below minimum price
	}

	_, result := validator.IsEligible(gift)
	assert.False(t, result)
}

func TestGiftValidator_IsEligible_PriceTooHigh(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	gift := &tg.StarGift{
		ID:    1,
		Stars: 1500, // Above maximum price
	}

	_, result := validator.IsEligible(gift)
	assert.False(t, result)
}

func TestGiftValidator_IsEligible_ExactMinPrice(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	gift := &tg.StarGift{
		ID:    1,
		Stars: 100, // Exactly minimum price
	}

	_, result := validator.IsEligible(gift)
	assert.True(t, result)
}

func TestGiftValidator_IsEligible_ExactMaxPrice(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, true, true)

	gift := &tg.StarGift{
		ID:      1,
		Stars:   1000, // Exactly maximum price
		Limited: true, // Must match LimitedStatus = true
	}

	_, result := validator.IsEligible(gift)
	assert.True(t, result)
}

func TestGiftValidator_IsEligible_ZeroCriteria(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     0,
			MaxPrice:     0,
			TotalSupply:  0,
			Count:        0,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 0, true, true)

	gift := &tg.StarGift{
		ID:      1,
		Stars:   0,    // Zero price should be valid for zero criteria
		Limited: true, // Must match LimitedStatus = true
	}

	_, result := validator.IsEligible(gift)
	assert.True(t, result)
}

func TestGiftValidator_IsEligible_NegativePrice(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     1000,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	gift := &tg.StarGift{
		ID:    1,
		Stars: -100, // Negative price
	}

	_, result := validator.IsEligible(gift)
	assert.False(t, result)
}

func TestGiftValidator_MultipleCriterias(t *testing.T) {
	criterias := []config.Criterias{
		{
			MinPrice:     100,
			MaxPrice:     500,
			TotalSupply:  50,
			Count:        10,
			ReceiverType: []int{1},
		},
		{
			MinPrice:     600,
			MaxPrice:     1000,
			TotalSupply:  100,
			Count:        20,
			ReceiverType: []int{2},
		},
	}
	validator := NewGiftValidator(criterias, 5000, false, false)

	tests := []struct {
		name     string
		stars    int64
		expected bool
	}{
		{"price in first range", 300, true},
		{"price in second range", 800, true},
		{"price between ranges", 550, false},
		{"price below all ranges", 50, false},
		{"price above all ranges", 1500, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gift := &tg.StarGift{
				ID:    1,
				Stars: tt.stars,
			}

			_, result := validator.IsEligible(gift)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGiftValidator_PriceValid_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		criterias []config.Criterias
		stars     int64
		expected  bool
	}{
		{
			name: "price within range",
			criterias: []config.Criterias{
				{MinPrice: 100, MaxPrice: 1000, TotalSupply: 100, ReceiverType: []int{1}},
			},
			stars:    500,
			expected: true,
		},
		{
			name: "price at minimum boundary",
			criterias: []config.Criterias{
				{MinPrice: 100, MaxPrice: 1000, TotalSupply: 100, ReceiverType: []int{1}},
			},
			stars:    100,
			expected: true,
		},
		{
			name: "price at maximum boundary",
			criterias: []config.Criterias{
				{MinPrice: 100, MaxPrice: 1000, TotalSupply: 100, ReceiverType: []int{1}},
			},
			stars:    1000,
			expected: true,
		},
		{
			name: "price below minimum",
			criterias: []config.Criterias{
				{MinPrice: 100, MaxPrice: 1000, TotalSupply: 100, ReceiverType: []int{1}},
			},
			stars:    99,
			expected: false,
		},
		{
			name: "price above maximum",
			criterias: []config.Criterias{
				{MinPrice: 100, MaxPrice: 1000, TotalSupply: 100, ReceiverType: []int{1}},
			},
			stars:    1001,
			expected: false,
		},
		{
			name: "zero range with zero price",
			criterias: []config.Criterias{
				{MinPrice: 0, MaxPrice: 0, TotalSupply: 0, ReceiverType: []int{1}},
			},
			stars:    0,
			expected: true,
		},
		{
			name: "zero range with non-zero price",
			criterias: []config.Criterias{
				{MinPrice: 0, MaxPrice: 0, TotalSupply: 0, ReceiverType: []int{1}},
			},
			stars:    1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewGiftValidator(tt.criterias, 5000, true, false)

			gift := &tg.StarGift{
				ID:    1,
				Stars: tt.stars,
			}

			_, result := validator.IsEligible(gift)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGiftValidator_StarCapValidation(t *testing.T) {
	tests := []struct {
		name        string
		giftStars   int64
		totalSupply int
		starCap     int64
		testMode    bool
		expected    bool
	}{
		{
			name:        "gift with zero availability total",
			giftStars:   100,
			totalSupply: 0,
			starCap:     1000,
			testMode:    false,
			expected:    true,
		},
		{
			name:        "gift with zero stars",
			giftStars:   0,
			totalSupply: 100,
			starCap:     1000,
			testMode:    false,
			expected:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			criterias := []config.Criterias{
				{
					MinPrice:     0,
					MaxPrice:     10000,
					TotalSupply:  1000,
					Count:        10,
					ReceiverType: []int{1},
				},
			}
			validator := NewGiftValidator(criterias, tt.starCap, tt.testMode, false)

			gift := &tg.StarGift{
				ID:    1,
				Stars: tt.giftStars,
			}
			if tt.totalSupply > 0 {
				totalSupply := tt.totalSupply
				gift.AvailabilityTotal = totalSupply
			}

			_, result := validator.IsEligible(gift)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGiftValidator_TestMode(t *testing.T) {
	tests := []struct {
		name        string
		testMode    bool
		giftLimited bool
		totalSupply int
		remains     int
		expected    bool
		description string
	}{
		{
			name:        "test mode - accepts unlimited gift",
			testMode:    true,
			giftLimited: false,
			totalSupply: 0,
			remains:     0,
			expected:    true,
			description: "Test mode should accept unlimited gifts",
		},
		{
			name:        "test mode - accepts limited gift",
			testMode:    true,
			giftLimited: true,
			totalSupply: 10,
			remains:     5,
			expected:    true,
			description: "Test mode should accept limited gifts regardless of supply",
		},
		{
			name:        "test mode - rejects sold out gift",
			testMode:    true,
			giftLimited: true,
			totalSupply: 10,
			remains:     0,
			expected:    false,
			description: "Even in test mode, sold out gifts should be rejected",
		},
		{
			name:        "production mode - rejects limited gift with low supply",
			testMode:    false,
			giftLimited: true,
			totalSupply: 10,
			remains:     5,
			expected:    false,
			description: "Production mode should reject limited gifts with supply below criteria",
		},
		{
			name:        "production mode - rejects limited gift in production mode due to complex validation",
			testMode:    false,
			giftLimited: true,
			totalSupply: 10,
			remains:     5,
			expected:    false,
			description: "Production mode may reject limited gifts due to complex validation rules",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			criterias := []config.Criterias{
				{
					MinPrice:     100,
					MaxPrice:     1000,
					TotalSupply:  150,
					Count:        10,
					ReceiverType: []int{1},
				},
			}
			validator := NewGiftValidator(criterias, 100000, tt.testMode, tt.giftLimited)

			gift := &tg.StarGift{
				ID:      1,
				Stars:   500,
				Limited: tt.giftLimited,
			}

			if tt.giftLimited {
				totalSupply := tt.totalSupply
				remains := tt.remains
				gift.AvailabilityTotal = totalSupply
				gift.AvailabilityRemains = remains
				if tt.remains == 0 {
					gift.SoldOut = true
				}
			}

			_, result := validator.IsEligible(gift)
			assert.Equal(t, tt.expected, result, tt.description)
		})
	}
}
