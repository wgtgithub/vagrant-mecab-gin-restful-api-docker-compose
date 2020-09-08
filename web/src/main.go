package main

import (
	"fmt"
	"strings"

	"github.com/bluele/mecab-golang"
	"github.com/gin-gonic/gin"

//	"morphlogic/mymod"
)
const BOSEOS = "BOS/EOS"
const const-def-target = "NiziUが好きな５０のWithU"

type molphology_data struct {
	letter string // letter
	list   array  // result as morphologic
	meCab    mecab  // mecab
	err    error  // error
	tagger     Tagger // tagger
	lattice     Lattice // lattice
}

func (item *molphology_data) config() {
	m, item.err := mecab.New("-Owakati")
//	item.meCab, item.err := mecab.New("-Owakati")
	if item.err != nil {
		panic(data.err)
	}
	defer m.Destroy()
	item.meCab = m
//	defer item.meCab.Destroy()
	item.parseToNode()
}

func (item *molphology_data) newobj() {
	rtn := false
	item.tagger, item.err := item.obj.NewTagger()
	if item.err == nil {
		item.tagger.Destroy()
		rtn = true
	}
	return rtn
}

func (item *molphology_data) set(input string) bool {
	rtn := false
	item.lattice, item.err := item.obj.NewLattice(input)
	if item.err == nil {
		item.lattice.Destroy()
		rtn = true
	}
	return rtn
}

func (item *molphology_data) parseToNode() {
	node := item.tagger.ParseToNode(item.lattice)

	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != BOSEOS {
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
		}
		if node.Next() != nil {
			break
		}
	}
}


func (item *molphology_data) parseToNode() error {
	item.err = nil
	tg, err := item.obj.NewTagger()
	if  err != nil { 
		item.err = err
	}
	defer tg.Destroy()

	lt, err := item.obj.NewLattice(item.letter)
	if err != nill {
		item.err = err
	}
	defer lt.Destroy()
	node := tg.ParseToNode(lt)

	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != BOSEOS {
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
		}
		if node.Next() != nil {
			break
		}
	}
}

func _get(c *gin.Context) {
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	letter := c.Query("letter")
	err = _parseToNode(m, letter)

	c.JSON(200, gin.H{
		"letter": letter,
	})
}

func _parseToNode(m *mecab.MeCab, letter string) error {
	tg, err := m.NewTagger()
	if err != nil {
		return err
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(letter)
	if err != nil {
		return err
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)

	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != BOSEOS {
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
		}
		if node.Next() != nil {
			break
		}
	}
	return nil
}

func main() {
	fmt.Print("Hello mecab")

	r := gin.Default()
	r.GET("/ping", _get)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
