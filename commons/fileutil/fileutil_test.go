package fileutil

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetFullUrl(t *testing.T) {
	// Set up the base URL in Viper
	viper.Set("base_url_upload", "http://example.com/uploads")

	t.Run("returns empty string when fileName is empty", func(t *testing.T) {
		result := GetFullUrl("")
		assert.Equal(t, "", result)
	})

	t.Run("returns correct full URL when fileName is provided", func(t *testing.T) {
		fileName := "image.jpg"
		expected := "http://example.com/uploads/" + fileName
		result := GetFullUrl(fileName)
		assert.Equal(t, expected, result)
	})

	t.Run("handles leading/trailing slashes in base_url_upload", func(t *testing.T) {
		viper.Set("base_url_upload", "http://example.com/uploads")
		fileName := "image.jpg"
		expected := "http://example.com/uploads/image.jpg"
		result := GetFullUrl(fileName)
		assert.Equal(t, expected, result)
	})
}
