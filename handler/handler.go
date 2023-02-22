package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"module/model"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

const (
	// Qris Endpoint
	qrisEndpoint string = "https://api.sandbox.midtrans.com/v2/charge"
)

func QrisHandler(c *gin.Context) {

	paymentType := c.Query("paymentType")
	midtransData := model.NewMidtransData(paymentType)
	result := midtransData.InitDummyData()
	data, err := json.Marshal(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	payload := strings.NewReader(string(data))
	fmt.Println("json = ", string(data))
	/*
		Request untuk qris diharuskan POST
	*/
	req, err := http.NewRequest("POST", qrisEndpoint, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	/*
		Nilai header Authorization adalah adalah string yang diencode ke base-64.
		Kalian bisa mendapatkan nilai encoded string nya dengan cara
		memasukkan username dan password kalian ke link berikut ini,
		setelah itu copas Authorization Value yang kalian dapatkan

		https://docs.midtrans.com/reference/qris
		username = SERVER_KEY
		password = kosong
	*/
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(os.Getenv("AUTHORIZATION_VALUE")))))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	var responseBody any
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "proses menghubungkan data dengan midtrans berhasil",
		"data":    responseBody,
	})
}
