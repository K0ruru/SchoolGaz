package handlers

import (
	"net/http"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)


func CreateMapel(c *gin.Context){
  if dbConn == nil{
    c.JSON(http.StatusInternalServerError,gin.H{"error":"erorr database connection or from the server error"})
    return
  }

  var newMapel model.Mapel
  if err := c.ShouldBindJSON(&newMapel); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
    return
  }
  if err := dbConn.Create(&newMapel).Error; err != nil{
    c.JSON(http.StatusNotAcceptable, gin.H{"error": err})
    return
  }
  c.JSON(http.StatusCreated,newMapel)

}

func UpdateMapel(c *gin.Context) {
    if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

    ID := c.Param("id_mapel")

    _, err := strconv.Atoi(ID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kelasID"})
        return
    }
    var UpdateMapel model.Mapel

    if err := dbConn.Where("id_mapel = $1", ID).First(&UpdateMapel).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    }

    if err := c.ShouldBindJSON(&UpdateMapel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    if err := dbConn.Save(&UpdateMapel).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
    }

    c.JSON(http.StatusOK, UpdateMapel)
}
func DeleteMapel(c *gin.Context){
  var DeleteMapel model.Mapel

  if dbConn == nil{
    c.JSON(http.StatusInternalServerError, gin.H{"error" : "database connection error or from the server error"})
    return
  }

  ID := c.Param("id_mapel")
  if err := dbConn.Where("id_mapel = ?", ID).Delete(&DeleteMapel).Error; err != nil{
    c.JSON(http.StatusBadRequest,gin.H{"error": "invalid data or can not delete"})
  }
  c.JSON(http.StatusOK, gin.H{"succes" : "successfully deleted"})
}
func GetALLMapel(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
	var GetMapel []model.Mapel
	if err := dbConn.Find(&GetMapel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, GetMapel)
}
func GetMapel(c * gin.Context){
  if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
var GetMapel model.Mapel
id := c.Param("id_mapel")

  if err := dbConn.Where("id_mapel = ?", id).First(&GetMapel).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error" : "cannot find mapel"})
    return
  }
  c.JSON(http.StatusOK, GetMapel)
}




/*

func CreateMapel(c *gin.Context){
var newMapel Mapel
  if err := c.ShouldBindJSON(&newMapel);err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
    fmt.Println(err)
  } 

  _, err := dbConn.Exec("INSERT INTO mapel (nama_mapel, guru) VALUES ($1, $2)",
    newMapel.Nama_mapel, newMapel.Guru)
  if err != nil{
    c.JSON(http.StatusInternalServerError, gin.H{"erorr": err})
    fmt.Println(err)
  }
  c.JSON(http.StatusCreated, gin.H{"succes":"success create mapel"})
}
// Update user handler
func UpdateMapel(c *gin.Context) {
    // Mendapatkan ID pengguna dari parameter
    id_mapel := c.Param("id_mapel")

    // Memeriksa apakah nilai "nis" ada
    if id_mapel == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "mapel ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    ID, err := strconv.Atoi(id_mapel)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kelas ID format"})
        return
    }

    // Mendapatkan data pengguna yang ingin diperbarui dari body permintaan
    var UpdateMapel Mapel
    if err := c.ShouldBindJSON(&UpdateMapel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        fmt.Println(err)
        return
    }

    // Lakukan operasi pembaruan pengguna sesuai dengan kebutuhan aplikasi Anda
    _, err = dbConn.Exec("UPDATE mapel SET nama_mapel=$1, guru=$2  WHERE id_mapel=$3",
    UpdateMapel.Nama_mapel,UpdateMapel.Guru, ID) 
  
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update mapel"})
        fmt.Println(err)
        return
    }

    // Respon berhasil
    c.JSON(http.StatusOK, gin.H{"message": "mapel updated successfully"})
}
// Get all users handler
func GetMapel(c *gin.Context) {
    rows, err := dbConn.Query("SELECT id_mapel,nama_mapel,guru FROM mapel")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch mapel"})
        return
    }
    defer rows.Close()

    var mapels []Mapel
    for rows.Next() {
        var mapel Mapel
        if err := rows.Scan(&mapel.ID, &mapel.Nama_mapel,&mapel.Guru); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan users"})
            fmt.Println(err)
            return
        }
        mapels = append(mapels, mapel)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan fetch users"})
        return
    }

    c.JSON(http.StatusOK, mapels)
}
// DeleteUser handles the deletion of a user
func DeleteMapel(c *gin.Context) {
    ID := c.Param("id_mapel")
    
    // Memeriksa apakah nilai "nis" ada
    if ID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "mapel ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    id_mapel, err := strconv.Atoi(ID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
        return
    }
   _, err = dbConn.Exec("DELETE FROM mapel WHERE id_kelas=$1", id_mapel)
      if err != nil {
         fmt.Println("Failed to delete mapel: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete mapel"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "mapel deleted successfully"})
}
*/
