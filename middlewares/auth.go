package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk melindungi endpoint yang memerlukan autentikasi.
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Di sini Anda bisa memeriksa apakah pengguna telah login atau memiliki token yang sah
        // Misalnya, Anda dapat memeriksa cookie atau token JWT

        // Jika pengguna tidak terautentikasi, beri respons kesalahan
        if !IsAuthenticated(c) {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Tidak terautentikasi"})
            c.Abort() // Hentikan eksekusi lebih lanjut dari handler
            return
        }

        // Jika pengguna telah terautentikasi, lanjutkan ke handler endpoint
        c.Next()
    }
}

// IsAuthenticated adalah fungsi yang memeriksa apakah pengguna telah login atau memiliki token yang sah.
func IsAuthenticated(c *gin.Context) bool {
    // Di sini Anda dapat memeriksa cookie atau token JWT dan menentukan apakah pengguna telah terautentikasi.
    // Anda bisa menggunakan library JWT atau mekanisme autentikasi sesuai kebutuhan Anda.

    // Contoh sederhana: Cek apakah ada cookie dengan nama "token"
    _, err := c.Request.Cookie("token")
    return err == nil
}
