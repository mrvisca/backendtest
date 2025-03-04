package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Struct untuk merepresentasikan bahasa pemrograman
type ProgrammingLanguage struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"` // Daftar pencipta bahasa pemrograman
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"` // Objek yang merepresentasikan hubungan dengan bahasa lain
}

// Struct Relation untuk menangani objek "relation" yang bersarang
type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

// Slice global untuk menyimpan data bahasa pemrograman
var languages []ProgrammingLanguage

// Fungsi untuk mendapatkan daftar bahasa pemrograman
func GetLanguange(c *gin.Context) {
	// Jika slice languages kosong, kembalikan data default
	if len(languages) == 0 {
		defaultLanguage := ProgrammingLanguage{
			Language:       "C",
			Appeared:       1972,
			Created:        []string{"Dennis Ritchie"},
			Functional:     true,
			ObjectOriented: false,
			Relation: Relation{
				InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
				Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
			},
		}
		// Menambahkan data default ke dalam slice global
		languages = append(languages, defaultLanguage)
		// Mengembalikan data default dalam format JSON
		c.JSON(http.StatusOK, defaultLanguage)
	} else {
		// // Jika ada data dalam slice, kembalikan data tersebut
		// c.JSON(http.StatusOK, languages)

		// Mendapatkan query pencarian dari parameter URL
		query := c.Query("query")
		if query == "" {
			// Jika tidak ada query, kembalikan seluruh daftar bahasa pemrograman
			c.JSON(http.StatusOK, languages)
		} else {
			// Konversi query menjadi huruf kecil
			queryLower := strings.ToLower(query)
			// Slice untuk menyimpan hasil pencarian
			var results []ProgrammingLanguage
			// Loop melalui setiap bahasa pemrograman
			for _, lang := range languages {
				// Konversi nama bahasa pemrograman menjadi huruf kecil
				langLower := strings.ToLower(lang.Language)
				// Periksa apakah nama bahasa pemrograman mengandung query
				if strings.Contains(langLower, queryLower) {
					results = append(results, lang)
				}
			}
			// Jika ada hasil pencarian, kembalikan hasil tersebut
			if len(results) > 0 {
				c.JSON(http.StatusOK, results)
			} else {
				// Jika tidak ada hasil yang cocok, kembalikan pesan tidak ditemukan
				c.JSON(http.StatusNotFound, gin.H{"message": "Tidak ada data yang ditemukan"})
			}
		}
	}
}

// Fungsi untuk mendapatkan bahasa pemrograman berdasarkan ID
func GetByIdLanguage(c *gin.Context) {
	// Mengambil ID dari parameter URL
	id := c.Param("id")
	// Mengonversi ID ke indeks dalam slice
	index := convertIDToIndex(id)
	if index < 0 || index >= len(languages) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found"})
		return
	}
	c.JSON(http.StatusOK, languages[index])
}

// Fungsi untuk menambahkan bahasa pemrograman baru
func PostLanguage(c *gin.Context) {
	var newLanguage ProgrammingLanguage
	if err := c.ShouldBindJSON(&newLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	languages = append(languages, newLanguage)
	c.JSON(http.StatusCreated, newLanguage)
}

// Fungsi untuk memperbarui data bahasa pemrograman berdasarkan ID
func UpdateLanguage(c *gin.Context) {
	// Mengambil ID dari parameter URL
	id := c.Param("id")
	// Mengonversi ID ke indeks dalam slice
	index := convertIDToIndex(id)
	if index < 0 || index >= len(languages) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found"})
		return
	}

	var updatedLanguage ProgrammingLanguage
	if err := c.ShouldBindJSON(&updatedLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memperbarui data bahasa dalam slice
	languages[index] = updatedLanguage
	c.JSON(http.StatusOK, updatedLanguage)
}

// Fungsi untuk menghapus bahasa pemrograman berdasarkan ID
func DeleteLanguage(c *gin.Context) {
	// Mengambil ID dari parameter URL
	id := c.Param("id")
	// Mengonversi ID ke indeks dalam slice
	index := convertIDToIndex(id)
	if index < 0 || index >= len(languages) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not found"})
		return
	}

	languages = append(languages[:index], languages[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Bahasa pemrograman berhasil dihapus"})
}

// Fungsi untuk mengonversi ID ke indeks slice
func convertIDToIndex(id string) int {
	var index int
	// Mengonversi string ID ke integer
	_, err := fmt.Sscanf(id, "%d", &index)
	if err != nil {
		// Mengembalikan -1 jika terjadi error
		return -1
	}
	// Mengembalikan indeks yang sesuai
	return index
}
