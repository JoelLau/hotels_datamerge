package hotels_test

import (
	"testing"

	hotels "hotels_data_merge"

	"github.com/stretchr/testify/require"
)

func TestLowerCaseWithSpaces(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "single word", in: "Pool", want: "pool"},
		{name: "two words", in: "BusinessCenter", want: "business center"},
		{name: "three words", in: "DryCleaningService", want: "dry cleaning service"},
		{name: "already lower case", in: "aircon", want: "aircon"},
		{name: "empty string", in: "", want: ""},
		{name: "consecutive capitals", in: "WiFi", want: "wi fi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, hotels.ToLowerCaseWithSpaces(tt.in))
		})
	}
}

func TestNilIfEmpty(t *testing.T) {
	t.Parallel()

	t.Run("empty string returns nil", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, hotels.ToNilIfEmpty(""))
	})

	t.Run("non-empty string returns pointer", func(t *testing.T) {
		t.Parallel()

		got := hotels.ToNilIfEmpty("hello")
		require.NotNil(t, got)
		require.Equal(t, "hello", *got)
	})
}

func TestLongestString(t *testing.T) {
	t.Parallel()

	t.Run("returns longest string", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, "ccc", hotels.LongestString([]string{"a", "bb", "ccc"}))
	})

	t.Run("empty slice returns empty string", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, "", hotels.LongestString(nil))
	})
}
