package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bluele/mecab-golang"
)

const BOSEOS = "BOS/EOS"
const constPartsNoun = "名詞"

type molphology struct {
	letter string                   // letter
	err    error                    // error
	rows   []map[string]interface{} // result as morphologic
	json   []byte
}

func (item *molphology) analyze() bool {
	//	rtn := item.init()
	m, err := mecab.New("-Owakati")
	if err != nil {
		item.err = err
		return false
	}
	defer m.Destroy()

	tg, err := m.NewTagger()
	if err != nil {
		item._setErr(err)
		return false
	}
	defer tg.Destroy()

	fmt.Printf("letter; %s\n", item.letter)
	lt, err := m.NewLattice(item.letter)
	if item.err != nil {
		item._setErr(err)
		return false
	}
	defer lt.Destroy()

	if !item._parseToNode(tg, lt) {
		return false
	}
	return item._convJSON()
}

func (item *molphology) _parseToNode(tg *mecab.Tagger, lt *mecab.Lattice) bool {

	node := tg.ParseToNode(lt)

	for {
		feature := _getFeatures(node)
		if _checkFeature(feature) {
			data := make(map[string]interface{})
			data["surface"] = node.Surface()
			data["feature"] = feature
			item._append(data)
		}
		if node.Next() != nil {
			break
		}
	}
	return true
}

func _getFeatures(node *mecab.Node) []string {
	return strings.Split(node.Feature(), ",")
}

func _checkFeature(ft []string) bool {
	if ft[0] != BOSEOS {
		return strings.Compare(ft[0], constPartsNoun) == 0
	}
	return false
}

func (item *molphology) _append(data map[string]interface{}) {
	if item.rows == nil {
		item.rows = []map[string]interface{}{}
	}
	item.rows = append(item.rows, data)
}

func (item *molphology) _setErr(err error) {
	item.err = err
}

func (item *molphology) _convJSON() bool {
	rtn := false
	cv, err := json.Marshal(item.rows)
	if err != nil {
		item._setErr(err)
	} else {
		item.json = cv
		rtn = true
	}
	return rtn
}

func (item *molphology) getJSON() []byte {
	fmt.Println(string(item.json))
	return item.json
}
