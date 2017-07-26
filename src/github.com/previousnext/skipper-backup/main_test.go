package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
)

// Ensure frequency tokenisation works.
func TestTokenReplacementFrequency(t *testing.T) {
	frequencyPlaceholder := "%FREQUENCY%"
	currentTime := time.Now().Local()

	// Test first day of month returns "monthly".
	assert.Equal(t, "monthly", tokensReplace(frequencyPlaceholder, currentTime, 1, time.Monday))

	// Test first day of month returns "monthly" even when taken on a Sunday.
	assert.Equal(t, "monthly", tokensReplace(frequencyPlaceholder, currentTime, 1, time.Sunday))

	for i := 2; i <= 31; i++ {
		// Test all days of month (except 1st) return "daily" if not a Sunday.
		assert.Equal(t, "daily", tokensReplace(frequencyPlaceholder, currentTime, i, time.Friday))

		// Test all Sundays return "weekly" on every dates except 1st of the month.
		assert.Equal(t, "weekly", tokensReplace(frequencyPlaceholder, currentTime, i, time.Sunday))
	}
}
