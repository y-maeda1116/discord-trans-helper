package translator

import (
	"testing"
)

func TestDetectLanguage(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "Japanese text",
			text:     "こんにちは、世界",
			expected: "JA",
		},
		{
			name:     "English text",
			text:     "Hello, World",
			expected: "EN",
		},
		{
			name:     "Japanese with kanji",
			text:     "今日は良い天気です",
			expected: "JA",
		},
		{
			name:     "English sentence",
			text:     "This is a test message",
			expected: "EN",
		},
		{
			name:     "Empty text",
			text:     "",
			expected: "EN",
		},
		{
			name:     "Mixed text with Japanese",
			text:     "Hello こんにちは",
			expected: "JA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectLanguage(tt.text); got != tt.expected {
				t.Errorf("DetectLanguage() = %v, want %v", got, tt.expected)
			}
		})
	}
}
