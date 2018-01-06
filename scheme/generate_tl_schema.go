package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type nametype struct {
	name  string
	_type string
}

type constuctor struct {
	id        string
	predicate string
	params    []nametype
	_type     string
}

func normalize(s string) string {
	x := []byte(s)
	for i, r := range x {
		if r == '.' {
			x[i] = '_'
		}
	}
	y := string(x)
	if y == "type" {
		return "_type"
	}
	return y
}

func maybeFlagged(_type string, isFlag bool, flagBit int) string {
	if isFlag {
		return fmt.Sprintf("m.Flagged%s(flags, %d),\n", _type, flagBit)
	} else {
		return fmt.Sprintf("m.%s(),\n", _type)
	}
}

func main() {
	var err error
	var parsed interface{}

	// read json file from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// parse json
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&parsed)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process constructors
	_order := make([]string, 0, 1000)
	_cons := make(map[string]constuctor, 1000)
	_types := make(map[string][]string, 1000)

	parsefunc := func(data []interface{}, kind string) {
		for _, data := range data {
			data := data.(map[string]interface{})

			// id
			idx, err := strconv.Atoi(data["id"].(string))
			if err != nil {
				fmt.Println(err)
				return
			}
			_id := fmt.Sprintf("0x%08x", uint32(idx))

			// predicate
			_predicate := normalize(data[kind].(string))

			if _predicate == "vector" {
				continue
			}

			// params
			_params := make([]nametype, 0, 16)
			params := data["params"].([]interface{})
			for _, params := range params {
				params := params.(map[string]interface{})
				_params = append(_params, nametype{normalize(params["name"].(string)), normalize(params["type"].(string))})
			}

			// type
			_type := normalize(data["type"].(string))

			_order = append(_order, _predicate)
			_cons[_predicate] = constuctor{_id, _predicate, _params, _type}
			if kind == "predicate" {
				_types[_type] = append(_types[_type], _predicate)
			}
		}
	}
	parsefunc(parsed.(map[string]interface{})["constructors"].([]interface{}), "predicate")
	parsefunc(parsed.(map[string]interface{})["methods"].([]interface{}), "method")

	// constants
	fmt.Print("package mtproto\nimport \"fmt\"\nconst (\n")
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Print(")\n\n")

	// type structs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("type TL_%s struct {\n", c.predicate)
		for _, t := range c.params {
			isFlag := false
			typeName := t._type
			if strings.HasPrefix(t._type, "flags") {
				isFlag = true
				typeName = t._type[strings.Index(t._type, "?")+1:]
			}
			fmt.Printf("%s\t", strings.Title(t.name))
			switch typeName {
			case "true": //flags only
				fmt.Print("bool")
			case "int", "#":
				fmt.Print("int32")
			case "long":
				fmt.Print("int64")
			case "string":
				fmt.Print("string")
			case "double":
				fmt.Print("float64")
			case "bytes":
				fmt.Print("[]byte")
			case "Vector<int>":
				fmt.Print("[]int32")
			case "Vector<long>":
				fmt.Print("[]int64")
			case "Vector<string>":
				fmt.Print("[]string")
			case "Vector<double>":
				fmt.Print("[]float64")
			case "!X":
				fmt.Print("TL")
			default:
				var inner string
				n, _ := fmt.Sscanf(typeName, "Vector<%s", &inner)
				if n == 1 {
					fmt.Printf("[]TL // %s", inner[:len(inner)-1])
				} else {
					fmt.Printf("TL // %s", typeName)
				}
			}
			if isFlag {
				fmt.Print(" //flag")
			}
			fmt.Print("\n")
		}
		fmt.Print("}\n\n")
	}

	// encode funcs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("func (e TL_%s) encode() []byte {\n", c.predicate)
		fmt.Print("x := NewEncodeBuf(512)\n")
		fmt.Printf("x.UInt(crc_%s)\n", c.predicate)
		for _, t := range c.params {
			typeName := t._type
			if strings.HasPrefix(t._type, "flags") {
				typeName = t._type[strings.Index(t._type, "?")+1:]
			}
			switch typeName {
			case "true": //flags only
				fmt.Printf("//flag %s\n", t.name)
			case "int", "#":
				fmt.Printf("x.Int(e.%s)\n", strings.Title(t.name))
			case "long":
				fmt.Printf("x.Long(e.%s)\n", strings.Title(t.name))
			case "string":
				fmt.Printf("x.String(e.%s)\n", strings.Title(t.name))
			case "double":
				fmt.Printf("x.Double(e.%s)\n", strings.Title(t.name))
			case "bytes":
				fmt.Printf("x.StringBytes(e.%s)\n", strings.Title(t.name))
			case "Vector<int>":
				fmt.Printf("x.VectorInt(e.%s)\n", strings.Title(t.name))
			case "Vector<long>":
				fmt.Printf("x.VectorLong(e.%s)\n", strings.Title(t.name))
			case "Vector<string>":
				fmt.Printf("x.VectorString(e.%s)\n", strings.Title(t.name))
			case "Vector<double>":
				fmt.Printf("x.VectorDouble(e.%s)\n", strings.Title(t.name))
			case "!X":
				fmt.Printf("x.Bytes(e.%s.encode())\n", strings.Title(t.name))
			default:
				var inner string
				n, _ := fmt.Sscanf(typeName, "Vector<%s", &inner)
				if n == 1 {
					fmt.Printf("x.Vector(e.%s)\n", strings.Title(t.name))
				} else {
					fmt.Printf("x.Bytes(e.%s.encode())\n", strings.Title(t.name))
				}
			}
		}
		fmt.Print("return x.buf\n")
		fmt.Print("}\n\n")

	}

	// decode funcs
	fmt.Println(`
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("case crc_%s:\n", c.predicate)
		for _, t := range c.params {
			if t._type == "#" {
				fmt.Print("flags := m.Int()\n")
				//fmt.Print("_ = flags\n")
				break
			}
		}
		fmt.Printf("r = TL_%s{\n", c.predicate)
		for _, t := range c.params {
			isFlag := false
			flagBit := 0
			typeName := t._type
			if strings.HasPrefix(t._type, "flags") {
				isFlag = true
				flagBit, _ = strconv.Atoi(string(t._type[strings.Index(t._type, "_")+1 : strings.Index(t._type, "?")]))
				typeName = t._type[strings.Index(t._type, "?")+1:]
			}
			switch typeName {
			case "true": //flags only
				fmt.Print("false, //flag\n")
			case "#":
				fmt.Print("flags,\n")
			case "int":
				fmt.Print(maybeFlagged("Int", isFlag, flagBit))
			case "long":
				fmt.Print(maybeFlagged("Long", isFlag, flagBit))
			case "string":
				fmt.Print(maybeFlagged("String", isFlag, flagBit))
			case "double":
				fmt.Print(maybeFlagged("Double", isFlag, flagBit))
			case "bytes":
				fmt.Print(maybeFlagged("StringBytes", isFlag, flagBit))
			case "Vector<int>":
				fmt.Print(maybeFlagged("VectorInt", isFlag, flagBit))
			case "Vector<long>":
				fmt.Print(maybeFlagged("VectorLong", isFlag, flagBit))
			case "Vector<string>":
				fmt.Print(maybeFlagged("VectorString", isFlag, flagBit))
			case "Vector<double>":
				fmt.Print(maybeFlagged("VectorDouble", isFlag, flagBit))
			case "!X":
				fmt.Print(maybeFlagged("Object", isFlag, flagBit))
			default:
				var inner string
				n, _ := fmt.Sscanf(typeName, "Vector<%s", &inner)
				if n == 1 {
					fmt.Print(maybeFlagged("Vector", isFlag, flagBit))
				} else {
					fmt.Print(maybeFlagged("Object", isFlag, flagBit))
				}
			}
		}
		fmt.Print("}\n\n")
	}

	fmt.Println(`
	default:
		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}`)

}
