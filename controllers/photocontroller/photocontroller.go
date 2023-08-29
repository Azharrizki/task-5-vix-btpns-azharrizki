package photocontroller

import (
	"net/http"
	"strconv"
	"task-5-vix-btpns/database"
	"task-5-vix-btpns/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var photos []models.Photo

	database.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func Create(c *gin.Context) {
	file, err := c.FormFile("photo_url")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	filename := uuid.New().String() + ".jpg"
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	userIdStr := c.PostForm("user_id")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User ID tidak valid"})
		return
	}

	photo := models.Photo{
		Title:    c.PostForm("title"),
		Caption:  c.PostForm("caption"),
		PhotoUrl: "uploads/" + filename,
		UserId:   userId,
	}

	database.DB.Create(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil upload gambar"})
}

func Update(c *gin.Context) {
	// Mengambil nilai dari parameter url
	photoId := c.Param("photoId")

	// Mengambil file gambar yang diunggah dari photo_url
	// Jika gagal mengambil maka akan menampilkan error
	file, err := c.FormFile("photo_url")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Membuat random name menggunakan uuid
	filename := uuid.New().String() + ".jpg"

	// Menyimpan file gambar yang diunggah ke folder uploads dengan nama yang telah dibuat
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Mengambil nilai user_id dari permintaan POST
	userIdStr := c.PostForm("user_id")
	// Melakukan konversi data user_id menjadi int agar sesuai dengan tipe data dalam model photo
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User ID tidak valid"})
		return
	}

	// Membuat objek foto untuk memperbarui data foto yang ada.
	photo := models.Photo{
		Title:    c.PostForm("title"),
		Caption:  c.PostForm("caption"),
		PhotoUrl: "uploads/" + filename,
		UserId:   userId,
	}

	// Memperbarui data foto dalam database dengan data yang baru dibuat.
	if database.DB.Model(&photo).Where("id = ?", photoId).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil mengupdate photo"})
}

func Delete(c *gin.Context) {
	var photo models.Photo

	photoId := c.Param("photoId")

	if err := database.DB.Delete(&photo, photoId).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Id photo tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
