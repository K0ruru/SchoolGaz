package handlers

import(
  "github.com/gin-gonic/gin"
)


type user struct{
  nis int `json:"nis"`
  nama string `json:"nama"`
  passphrase string `json:"passphrase"`
  email string `json:"email"`
  gender string `json:"gender"`
  agama string `json:"agama"`
  status string `json:"status"`  
}



