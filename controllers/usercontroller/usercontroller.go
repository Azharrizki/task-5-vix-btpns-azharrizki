package usercontroller

import (
	"net/http"
	"task-5-vix-btpns/database"
	"task-5-vix-btpns/helpers"
	"task-5-vix-btpns/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var userInput models.User

	// cek apakah terdapat request berupa json
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// melakukan pencarian data berdasarkan username
	var user models.User
	if err := database.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &helpers.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "task-5-vix-btpns",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma untuk login
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(helpers.JWT_KEY)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// set token ke cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "login berhasil"})
}

func Register(c *gin.Context) {
	var user models.User

	// bind data json yang dikirim
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// encrypt password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Register berhasil"})
}

func Update(c *gin.Context) {
	var user models.User
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// hashing password baru
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	if database.DB.Model(&user).Where("id = ?", userId).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data user diperbarui"})
}

func Delete(c *gin.Context) {
	var user models.User

	userId := c.Param("userId")

	if err := database.DB.Delete(&user, userId).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
