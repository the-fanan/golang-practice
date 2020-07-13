package main

import (
	"fmt"
	"strings"
	"strconv"
	"mime/multipart"
)

type NestedFormData struct {
	Value *ValueNode
	File *FileNode
}

type ValueNode struct {
	Value []string
	Children map[string]*ValueNode
}

type FileNode struct {
	Value []multipart.File
	Children map[string]*FileNode
}

func (fd *NestedFormData) nestValues(n *ValueNode, k *[]string, v []string) {
	var key string
	key, *k = (*k)[0], (*k)[1:]
	if len(*k) == 0 {
		if _, ok := n.Children[key]; ok {
			n.Children[key].Value = append(n.Children[key].Value, v...)
		} else {
			cn := &ValueNode{
				Value: v,
				Children: make(map[string]*ValueNode),
			}
			n.Children[key] = cn
		}
	} else {
		if _, ok := n.Children[key]; ok {
			fd.nestValues(n.Children[key], k,v)
		} else {
			cn := &ValueNode{
				Children: make(map[string]*ValueNode),
			}
			n.Children[key] = cn
			fd.nestValues(cn, k,v)
		}
	}
}

func (fd *NestedFormData) ParseValues(m map[string][]string){
	n := &ValueNode{
		Children: make(map[string]*ValueNode),
	}
	for key, val := range m {
		keys := strings.Split(key,".")
		fd.nestValues(n, &keys, val)
	}
	fd.Value = n
}

//multipart.File

func main() {
	//tr := "fanan.dala.moses"
	fd := map[string][]string{
		"fanan.dala.moses": []string{"boy"},
		"fanan.bala.lop": []string{"bro"},
		"fanan.dala": []string{"girl"},
		"fanan": []string{"trans"},
		"tickets.0.name": []string{"fanan"},
		"tickets.1.name": []string{"flenjor"},
	}
	f := &NestedFormData{}
	f.ParseValues(fd)
	for key, child := range f.Value.Children {
		fmt.Println("Key:", key, "Value:", child.Value)
		if child.Children != nil {
			for k, c := range child.Children {
				fmt.Println(key + "'s child key:", k, "Value:", c.Value)
			}
		}
	}
	fmt.Println(f.Value.Children["fanan"].Value[0])
	//fmt.Println(mn.Children["fanan"].Children["dala"].Value)
	//fmt.Println(mn.Children["fanan"].Children["bala"].Value)
	//fmt.Println(mn.Children["fanan"].Children["dala"].Children["moses"].Value)
	//fmt.Println(mn.Children["fanan"].Children["bala"].Children["lop"].Value)
	//fmt.Println(mn.Children["tickets"].Children["0"].Children["name"].Value)
	fmt.Println(f.Value.Children["tickets"].Children[strconv.Itoa(1)].Children["name"].Value)
}