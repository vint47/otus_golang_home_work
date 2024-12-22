package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		offset   int64
		limit    int64
		expected string
	}{
		{offset: 0, limit: 0, expected: "testdata/out_offset0_limit0.txt"},
		{offset: 0, limit: 10, expected: "testdata/out_offset0_limit10.txt"},
		{offset: 0, limit: 1000, expected: "testdata/out_offset0_limit1000.txt"},
		{offset: 0, limit: 10000, expected: "testdata/out_offset0_limit10000.txt"},
		{offset: 100, limit: 1000, expected: "testdata/out_offset100_limit1000.txt"},
		{offset: 6000, limit: 1000, expected: "testdata/out_offset6000_limit1000.txt"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.expected, func(t *testing.T) {
			err := Copy("testdata/input.txt", "out.txt", tc.offset, tc.limit)
			require.NoError(t, err)

			actual, _ := os.ReadFile("out.txt")
			expected, _ := os.ReadFile(tc.expected)

			require.Equal(t, expected, actual)
		})
	}
}

func TestCopyError(t *testing.T) {
	t.Run("UnsupportedFile", func(t *testing.T) {
		err := Copy("random", "random", 0, 0)
		require.ErrorAs(t, err, &ErrUnsupportedFile)
	})

	t.Run("OffsetExceedsFileSize", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 20000, 0)
		require.ErrorAs(t, err, &ErrOffsetExceedsFileSize)
	})

	t.Run("OffsetOrLimitLessThanZero", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", -1, 0)
		require.ErrorAs(t, err, &ErrOffsetOrLimitLessThanZero)
	})

	t.Run("OffsetOrLimitLessThanZero", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 0, -1)
		require.ErrorAs(t, err, &ErrOffsetOrLimitLessThanZero)
	})
}
