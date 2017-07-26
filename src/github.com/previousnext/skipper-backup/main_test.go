package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
)

// Ensure timestamp tokenisation works.
func TestTokenReplacementTimestamp(t *testing.T) {
	timestampPlaceholder := "%TIMESTAMP%"
	knownTimestamp := "2020-10-11_12-13-14"
	knownTime, err := time.Parse("2006-01-02_15-04-05", knownTimestamp)
	if err != nil {
		panic(err)
	}

	// Test timestamp replacement returns expected format.
	assert.Equal(t, knownTimestamp, tokensReplace(timestampPlaceholder, knownTime, 11, time.Thursday))
}

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
