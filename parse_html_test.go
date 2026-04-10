package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{
			name:     "h1 simple test",
			actual:   getHeadingFromHTML("<html><body><h1>Test Title h1</h1></body></html>"),
			expected: "Test Title h1",
		},
		{
			name:     "h2 simple test",
			actual:   getHeadingFromHTML("<html><body><h2>Test Title h2</h2></body></html>"),
			expected: "Test Title h2",
		},
		{
			name:     "no heading",
			actual:   getHeadingFromHTML("<html><body><p>Test Title</p></body></html>"),
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, tc.actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{
			name: "P top level",
			actual: getFirstParagraphFromHTML(`<html><body>
			<p>Outside paragraph.</p>
			</body></html>`),
			expected: "Outside paragraph.",
		},
		{
			name: "P first",
			actual: getFirstParagraphFromHTML(`<html><body>
			<p>First paragraph.</p>
			<p>Secound paragraph.</p>
			</body></html>`),
			expected: "First paragraph.",
		},
		{
			name: "P in main",
			actual: getFirstParagraphFromHTML(`<html><body>
			<p>Outside paragraph.</p>
			<main>
				<p>Main paragraph.</p>
			</main>
			</body></html>`),
			expected: "Main paragraph.",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, tc.actual)
			}
		})
	}
}
