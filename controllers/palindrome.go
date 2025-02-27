package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// isPalindrome memeriksa apakah sebuah string adalah palindrom
func isPalindrome(s string) bool {
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
}

func CheckPalindrome(c *gin.Context) {
	kueri := c.Query("text")

	cek := isPalindrome(kueri)

	if cek {
		c.JSON(200, gin.H{
			"status":      "Sukses",
			"text":        kueri,
			"palindrome?": "Palindrome",
		})
	} else {
		c.JSON(400, gin.H{
			"status":      "Elor",
			"text":        kueri,
			"palindrome?": "Not Palindrome",
		})
	}
}
