package models

import (
	"Self-Money_BE/database"
	"gorm.io/gorm"
)

type Pengeluaran struct {
	gorm.Model
	ID               int    `gorm:"primaryKey"`
	JenisPengeluaran string `json:"jenis_pengeluaran" binding:"required"`
	Harga            string `json:"harga" binding:"required"`
	Tanggal          string `json:"tanggal" binding:"required"`
	Keterangan       string `json:"keterangan" binding:"required"`
}

func (pengeluaran *Pengeluaran) CreatePengeluaranRecord() error {
	result := database.GlobalDB.Create(&pengeluaran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pengeluaran *Pengeluaran) UpdatePengeluaranRecord() error {
	result := database.GlobalDB.Save(&pengeluaran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pengeluaran *Pengeluaran) DeletePengeluaranRecord() error {
	result := database.GlobalDB.Delete(&pengeluaran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllPengeluaran() ([]Pengeluaran, error) {
	var pengeluaran []Pengeluaran
	result := database.GlobalDB.Find(&pengeluaran)
	if result.Error != nil {
		return nil, result.Error
	}
	return pengeluaran, nil
}

func GetPengeluaranByJenisPengeluaran(jenisPengeluaran string) (Pengeluaran, error) {
	var pengeluaran Pengeluaran
	result := database.GlobalDB.Where("jenis_pengeluaran = ?", jenisPengeluaran).First(&pengeluaran)
	if result.Error != nil {
		return pengeluaran, result.Error
	}
	return pengeluaran, nil
}

func GetPengeluaranByTanggal(tanggal string) (Pengeluaran, error) {
	var pengeluaran Pengeluaran
	result := database.GlobalDB.Where("tanggal = ?", tanggal).First(&pengeluaran)
	if result.Error != nil {
		return pengeluaran, result.Error
	}
	return pengeluaran, nil
}

func GetPengeluarnById(id int) (Pengeluaran, error) {
	var pengeluaran Pengeluaran
	result := database.GlobalDB.Where("id = ?", id).First(&pengeluaran)
	if result.Error != nil {
		return pengeluaran, result.Error
	}
	return pengeluaran, nil
}
