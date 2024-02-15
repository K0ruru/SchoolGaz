// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes

package handlers
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes
// Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes Belum di tes

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"server/model"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

// get all file
func IndexFileJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
	var file []model.UploadJawaban
	if err := dbConn.Find(&file).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, file)
}

// get file by id
func GetFileJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}

	id_UT := c.Param("id_UT")
	var Getfile model.UploadJawaban

	if err := dbConn.Where("id_UT = ?", id_UT).First(&Getfile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user no found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, Getfile)

}

// Upload File ke Cloudinary dan database
func UploadFileJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	// Parse form
	if err := c.Request.ParseMultipartForm(1 << 50); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		fmt.Println(err)
		return
	}

	// Get user name from form
	name := c.Request.FormValue("name")

	// Get file from form
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Upload file to Cloudinary
	ctx := context.Background()
	uploadResult, err := UploadToCloudinaryJawaban(ctx, file, fileHeader.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload file"})
		fmt.Println(err)
		return
	}

	ut := model.UploadJawaban{Name: name, File: uploadResult.SecureURL}
	err = dbConn.Create(&ut).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create data upload"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, ut)

}

func UploadToCloudinaryJawaban(ctx context.Context, file multipart.File, filename string) (*uploader.UploadResult, error) {
	// Define upload parameters
	params := uploader.UploadParams{
		PublicID: filename,
		Folder:   "jawaban",
	}

	// Upload file to Cloudinary
	uploadResult, err := cld.Upload.Upload(ctx, file, params)
	if err != nil {
		return nil, err
	}

	return uploadResult, nil
}
