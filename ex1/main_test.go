package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	testCases := []struct {
		name   string
		args   []string
		output string
	}{
		{"not enough args", []string{"Binh", "TH"}, "Need at least 3 arguments passed, has only: 2"},
		{"invalid country code", []string{"Binh", "Vu", "CA"}, "Invalid country code: CA\n"},
		{"valid VN name", []string{"John", "Smith", "VN"}, "Output: Smith John"},
		{"valid US name", []string{"Emily", "Rose", "Watson", "US"}, "Output: Emily Watson Rose"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// redirect stdout to capture the output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			// run the main func with test args
			os.Args = append([]string{"main.go"}, tc.args...)
			main()
			// restore stdout and read the captured output
			os.Stdout = oldStdout
			output := make([]byte, 1024)
			n, _ := r.Read(output)
			// close connection of the new pipe
			r.Close()
			w.Close()

			if string(output[:n]) != tc.output {
				t.Errorf("Test %q: got %q, want %q", tc.name, output[:n], tc.output)
			}
		})
	}
}

func Test_isValidCountryCode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid country code",
			input:    "VN",
			expected: true,
		},
		{
			name:     "invalid country code",
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
