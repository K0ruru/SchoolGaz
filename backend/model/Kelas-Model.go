package model

type Kelas struct {
    
    Id_kelas  int    `gorm:"primary_key"`
    NamaKelas string `gorm:"varchar(50)"`
    WalasID   uint  
    Walas     Guru 
 }

