package main

import "testing"

func Test_isValidCountryCode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid country code",
			input:    "VN",
			expected: true,
		},
		{
			name:     "Invalid country code",
			input:    "CA",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isValidCountryCode(tc.input)
			if actual != tc.expected {
				t.Errorf("Test %s: expected %t, got %t", tc.name, tc.expected, actual)
			}
		})
	}
}

func Test_getFullNameFromCountryCode(t *testing.T) {
	testCases := []struct {
		name         string
		inputRawName []string
		inputCode    string
		expected     string
	}{
		{
			name:         "no middle name - US code",
			inputRawName: []string{"Binh", "Vu"},
			inputCode:    "US",
			expected:     "Binh Vu",
		},
		{
			name:         "no middle name - VN code",
			inputRawName: []string{"Binh", "Vu"},
			inputCode:    "VN",
			expected:     "Vu Binh",
		},
		{
			name:         "1 middle name - US code",
			inputRawName: []string{"Binh", "Vu", "Duy"},
			inputCode:    "US",
			expected:     "Binh Duy Vu",
		},
		{
			name:         "1 middle name - VN code",
			inputRawName: []string{"Binh", "Vu", "Duy"},
			inputCode:    "VN",
			expected:     "Vu Duy Binh",
		},
		{
			name:         "more than 1 middle name - US code",
			inputRawName: []string{"Binh", "Vu", "Tran Ngoc Duy"},
			inputCode:    "US",
			expected:     "Binh Tran Ngoc Duy Vu",
		},
		{
			name:         "more than 1 middle name - VN code",
			inputRawName: []string{"Binh", "Vu", "Tran Ngoc Duy"},
			inputCode:    "VN",
			expected:     "Vu Tran Ngoc Duy Binh",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFullNameFromCountryCode(tc.inputRawName, tc.inputCode)
			if actual != tc.expected {
				t.Errorf("Test %s: expected %s, got %s", tc.name, tc.expected, actual)
			}
		})
	}
}
