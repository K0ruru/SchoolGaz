package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"server/model"

// 	"github.com/gin-gonic/gin"
// )

// func IndexJawaban(c *gin.Context) {
// 	if dbConn == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
// 	}

// 	var jawaban []model.Jawaban
// 	dbConn.Find(&jawaban)

// 	if len(tugas) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "tidak ada jawaban"})
// 		fmt.Println(err)
// 		return
// 	}
// }
