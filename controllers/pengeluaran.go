package controllers

import (
	"Self-Money_BE/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PengeluaranPayload struct {
	Id               int    `json:"id"`
	JenisPengeluaran string `json:"jenis_pengeluaran" binding:"required"`
	Harga            string `json:"harga" binding:"required"`
	Tanggal          string `json:"tanggal" binding:"required"`
	Keterangan       string `json:"keterangan" binding:"required"`
}

type PengeluaranResponse struct {
	JenisPengeluaran string `json:"jenis_pengeluaran"`
	Harga            string `json:"harga"`
	Tanggal          string `json:"tanggal"`
	Keterangan       string `json:"keterangan"`
}

func CreatePengeluaran(c *gin.Context) {
	fmt.Print("masuk" + c.GetHeader("Authorization"))
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"Error": "Token Not Found",
		})
		c.Abort()
		return
	}

	var payload PengeluaranPayload
	var pengeluaran models.Pengeluaran
	c.BindJSON(&payload)
	pengeluaran.JenisPengeluaran = payload.JenisPengeluaran
	pengeluaran.Harga = payload.Harga
	pengeluaran.Tanggal = payload.Tanggal
	pengeluaran.Keterangan = payload.Keterangan
	err := pengeluaran.CreatePengeluaranRecord()
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Create Pengeluaran",
		})
		c.Abort()
		return
	}
	c.JSON(200, pengeluaran)
}

func UpdatePengeluaran(c *gin.Context) {
	var payload PengeluaranPayload
	var pengeluaran models.Pengeluaran
	c.BindJSON(&payload)
	pengeluaran.ID = payload.Id
	pengeluaran.JenisPengeluaran = payload.JenisPengeluaran
	pengeluaran.Harga = payload.Harga
	pengeluaran.Tanggal = payload.Tanggal
	pengeluaran.Keterangan = payload.Keterangan
	err := pengeluaran.UpdatePengeluaranRecord()
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Update Pengeluaran",
		})
		c.Abort()
		return
	}
	c.JSON(200, pengeluaran)
}

func DeletePengeluaran(c *gin.Context) {
	var payload PengeluaranPayload
	var pengeluaran models.Pengeluaran
	c.BindJSON(&payload)
	pengeluaran.ID = payload.Id
	err := pengeluaran.DeletePengeluaranRecord()
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Delete Pengeluaran",
		})
		c.Abort()
		return
	}
	c.JSON(200, pengeluaran)
}

func GetAllPengeluaran(c *gin.Context) {
	pengeluaran, err := models.GetAllPengeluaran()
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get All Pengeluaran",
		})
		c.Abort()
		return
	}
	var response []PengeluaranResponse
	for _, pengeluaran := range pengeluaran {
		response = append(response, PengeluaranResponse{
			JenisPengeluaran: pengeluaran.JenisPengeluaran,
			Harga:            pengeluaran.Harga,
			Tanggal:          pengeluaran.Tanggal,
			Keterangan:       pengeluaran.Keterangan,
		})
	}
	c.JSON(200, response)
}

func GetPengeluaranByJenisPengeluaran(c *gin.Context) {
	jenisPengeluaran := c.Param("jenis_pengeluaran")
	pengeluaran, err := models.GetPengeluaranByJenisPengeluaran(jenisPengeluaran)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get Pengeluaran",
		})
		c.Abort()
		return
	}
	response := PengeluaranResponse{
		JenisPengeluaran: pengeluaran.JenisPengeluaran,
		Harga:            pengeluaran.Harga,
		Tanggal:          pengeluaran.Tanggal,
		Keterangan:       pengeluaran.Keterangan,
	}
	c.JSON(200, response)
}

func GetPengeluaranByTanggal(c *gin.Context) {
	tanggal := c.Param("tanggal")
	pengeluaran, err := models.GetPengeluaranByTanggal(tanggal)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get Pengeluaran",
		})
		c.Abort()
		return
	}
	response := PengeluaranResponse{
		JenisPengeluaran: pengeluaran.JenisPengeluaran,
		Harga:            pengeluaran.Harga,
		Tanggal:          pengeluaran.Tanggal,
		Keterangan:       pengeluaran.Keterangan,
	}
	c.JSON(200, response)
}

func GetPengeluaranById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid ID",
		})
		return
	}

	pengeluaran, err := models.GetPengeluarnById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get Pengeluaran",
		})
		c.Abort()
		return
	}

	response := PengeluaranResponse{
		JenisPengeluaran: pengeluaran.JenisPengeluaran,
		Harga:            pengeluaran.Harga,
		Tanggal:          pengeluaran.Tanggal,
		Keterangan:       pengeluaran.Keterangan,
	}

	c.JSON(200, response)
}

func GetPengeluaranFilterByJenisPengeluaran(c *gin.Context) {
	jenisPengeluaran := c.Param("jenis_pengeluaran")
	pengeluaran, err := models.GetPengeluaranByJenisPengeluaran(jenisPengeluaran)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get Pengeluaran",
		})
		c.Abort()
		return
	}
	response := PengeluaranResponse{
		JenisPengeluaran: pengeluaran.JenisPengeluaran,
		Harga:            pengeluaran.Harga,
		Tanggal:          pengeluaran.Tanggal,
		Keterangan:       pengeluaran.Keterangan,
	}
	c.JSON(200, response)
}
