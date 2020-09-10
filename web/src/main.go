package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const constDefTarget = "NiziUが好きな５０のWithU"

func _get(c *gin.Context) {
	mol := molphology{
		letter: _getLetter(c),
	}
	rtn := mol.analyze()
	if !rtn {
		panic(mol.err)
	}
	c.JSON(200, gin.H{
		"letter": mol.letter,
		"json":   string(mol.getJSON()),
	})
}

func _getLetter(c *gin.Context) string {
	letter := c.Query("letter")
	if len(letter) == 0 {
		letter = constDefTarget
	}
	return letter
}

func main() {
	fmt.Print("Hello mecab")

	r := gin.Default()
	r.GET("/targetwords", _get)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
