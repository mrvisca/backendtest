package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// isPalindrome memeriksa apakah sebuah string adalah palindrom
func isPalindrome(s string) bool {
	// Kondisi string query kosong dan string query biasa
	if s != "" {
		// Mengubah string menjadi lowercase untuk memastikan pengecekan case-insensitive
		s = strings.ToLower(s)
		// Menghapus spasi dari string
		s = strings.ReplaceAll(s, " ", "")
		// Mendapatkan panjang string
		n := len(s)
		// Memeriksa setiap karakter dari awal dan akhir string
		for i := 0; i < n/2; i++ {
			if s[i] != s[n-i-1] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func CheckPalindrome(c *gin.Context) {
	// Variabel query text
	kueri := c.Query("text")

	// Kondisi jika variabel query kosong maupun tidak
	if kueri != "" {
		cek := isPalindrome(kueri)

		if cek {
			c.JSON(http.StatusOK, gin.H{
				"status":      "Sukses",
				"text":        kueri,
				"palindrome?": "Palindrome",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":      "Error",
				"text":        kueri,
				"palindrome?": "Not Palindrome",
			})
		}
	} else {
		// Response gagal kueri kosong
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Palindrome gagal, pastikan query text tidak boleh kosong",
		})
	}
}
