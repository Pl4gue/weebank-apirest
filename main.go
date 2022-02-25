package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB = map[string]int{}

func main() {

	r := gin.Default()

	r.GET("/:wallet", func(c *gin.Context) {
		wallet := c.Param("wallet")
		c.JSON(http.StatusOK, gin.H{"value": DB[wallet]})
	})

	r.POST("/", func(c *gin.Context) {
		user := genarateKeys()
		DB[user.Pub] = 10
		c.JSON(http.StatusOK, user)
	})

	r.POST("/send", func(c *gin.Context) {
		transaction := Transaction{}
		c.Bind(&transaction)

		if transaction.From.Pub != reverseString(transaction.From.Pvt) || DB[transaction.From.Pub] < transaction.Value {
			c.Status(http.StatusBadRequest)
			return
		}

		DB[transaction.From.Pub] -= transaction.Value
		DB[transaction.To] += transaction.Value

		c.Status(http.StatusOK)
	})

	r.Run()
}
