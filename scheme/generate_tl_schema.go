package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://github.com/telegramdesktop/tdesktop/blob/dev/Telegram/Resources/scheme.tl
// https://github.com/danog/MadelineProto/tree/master/src/danog/MadelineProto

// https://core.telegram.org/mtproto/TL
// https://core.telegram.org/mtproto/TL-combinator
// identifier#name attr:type attr:type = resultType;

// Pq Ids

type Field struct {
	name     string
	typeName string
	flagBit  int
}

func (f Field) isFlag() bool {
	return f.flagBit >= 0
}

type Combinator struct {
	id       string
	name     string
	fields   []Field
	typeName string
}

func (c Combinator) hasFlags() bool {
	for _, f := range c.fields {
		if f.typeName == "#" {
			return true
		}
	}
	return false
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

func normalizeName(nameStr string, base int) string {
	nameInt, err := strconv.ParseInt(nameStr, base, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("0x%08x", uint32(nameInt))
}

func normalizeAttr(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	if strings.HasSuffix(s, "Id") {
		s = s[:len(s)-2] + "ID"
	}
	return s
}

func maybeFlagged(_type string, isFlag bool, flagBit int, args ...string) string {
	argsStr := strings.Join(args, ",")
	if isFlag {
		return fmt.Sprintf("m.Flagged%s(flags, %d, %s),\n", _type, flagBit, argsStr)
	} else {
		return fmt.Sprintf("m.%s(%s),\n", _type, argsStr)
	}
}

func makeField(name, typeName string) Field {
	flagBit := -1
	if strings.HasPrefix(typeName, "flags.") { //flags.2?string
		var err error
		qPos := strings.Index(typeName, "?")
		dPos := strings.Index(typeName, ".")
		flagBit, err = strconv.Atoi(typeName[dPos+1 : qPos])
		typeName = typeName[qPos+1:]
		if err != nil {
			log.Fatalf("parsing %s: %s", typeName, err)
		}
	}
	return Field{normalize(name), normalize(typeName), flagBit}
}

func parseJsonSchema(fpath string) []*Combinator {
	// reading json file
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	// parsings json
	var parsed interface{}
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(&parsed); err != nil {
		log.Fatal(err)
	}

	// processing constructors
	combinators := []*Combinator{}

	parsefunc := func(data []interface{}, kind string) {
		for _, data := range data {
			data := data.(map[string]interface{})

			// name
			name := normalizeName(data["id"].(string), 10)

			// identifier
			id := normalize(data[kind].(string))
			if id == "vector" {
				continue
			}

			// fields
			fields := make([]Field, 0, 16)
			srcFields := data["params"].([]interface{})
			for _, srcField := range srcFields {
				srcField := srcField.(map[string]interface{})
				name := srcField["name"].(string)
				typeName := srcField["type"].(string)
				fields = append(fields, makeField(name, typeName))
			}

			// type
			typeName := normalize(data["type"].(string))

			combinators = append(combinators, &Combinator{id, name, fields, typeName})
		}
	}
	parsefunc(parsed.(map[string]interface{})["constructors"].([]interface{}), "predicate")
	parsefunc(parsed.(map[string]interface{})["methods"].([]interface{}), "method")
	return combinators
}

func parseTLSchema(fpath string) []*Combinator {
	// reading tl file
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	// processing constructors
	combinators := []*Combinator{}

	lineRegexp := regexp.MustCompile(`^(.*?)#([a-f0-9]*)? (.*)= (.*);$`)
	fieldRegexp := regexp.MustCompile(`^(.*?):(.*)$`)
	for lineNum, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if line == "---functions---" || line == "---types---" {
			continue
		}

		match := lineRegexp.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Printf("line %d: wrong combinator: %s", lineNum, line)
			continue
		}
		id := normalize(match[1])
		name := normalizeName(match[2], 16)
		fieldsStr := match[3]
		typeName := normalize(match[4])

		if id == "vector" {
			continue
		}
		// if strings.HasPrefix(fieldsStr, "{") {
		// 	log.Printf("line %d: skipping: %s", lineNum, line)
		// 	continue
		// }

		fields := make([]Field, 0, 16)
		for _, fieldStr := range strings.Split(fieldsStr, " ") {
			fieldStr = strings.TrimSpace(fieldStr)
			if fieldStr == "" || strings.HasPrefix(fieldStr, "{") { //TODO
				continue
			}
			match := fieldRegexp.FindStringSubmatch(fieldStr)
			if len(match) == 0 {
				log.Fatalf("line %d: wrong field: %s", lineNum, fieldStr)
			}
			name, typeName := match[1], match[2]
			fields = append(fields, makeField(name, typeName))
		}

		combinators = append(combinators, &Combinator{id, name, fields, typeName})
	}
	return combinators
}

func main() {
	if len(os.Args) != 4 {
		println("Usage: " + os.Args[0] + " layer tl_schema.<json|tl> tl_schema.go")
		os.Exit(2)
	}
	layer, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fpath := os.Args[2]

	// parsing
	var combinators []*Combinator
	if strings.HasSuffix(fpath, ".json") {
		combinators = parseJsonSchema(fpath)
	} else if strings.HasSuffix(fpath, ".tl") {
		combinators = parseTLSchema(fpath)
	} else {
		log.Fatalf("unknown file format, expected .json or .tl: %s", fpath)
	}

	// opening out file
	outFile, err := os.Create(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	write := func(format string, a ...interface{}) {
		if _, err := fmt.Fprintf(outFile, format, a...); err != nil {
			log.Fatal(err)
		}
	}

	// constants
	write("package mtproto\nimport \"fmt\"\nconst (\n")
	write("TL_Layer = %d\n", layer)
	for _, c := range combinators {
		write("CRC_%s = %s\n", c.id, c.name)
	}
	write(")\n\n")

	// type structs
	for _, c := range combinators {
		write("type TL_%s struct {\n", c.id)
		for _, t := range c.fields {
			write("%s\t", normalizeAttr(t.name))
			switch t.typeName {
			case "true": //flags only
				write("bool")
			case "int", "#":
				write("int32")
			case "long":
				write("int64")
			case "int128":
				write("[]byte")
			case "int256":
				write("[]byte")
			case "string":
				write("string")
			case "double":
				write("float64")
			case "bytes":
				write("[]byte")
			case "Vector<int>":
				write("[]int32")
			case "Vector<long>":
				write("[]int64")
			case "Vector<string>":
				write("[]string")
			case "Vector<double>":
				write("[]float64")
			case "!X":
				write("TL")
			default:
				var inner string
				n, _ := fmt.Sscanf(t.typeName, "Vector<%s", &inner)
				if n == 1 {
					write("[]TL // %s", inner[:len(inner)-1])
				} else {
					write("TL // %s", t.typeName)
				}
			}
			if t.isFlag() {
				write(" //flag")
			}
			write("\n")
		}
		write("}\n\n")
	}

	// encode funcs
	for _, c := range combinators {
		write("func (e TL_%s) encode() []byte {\n", c.id)
		write("x := NewEncodeBuf(512)\n")
		write("x.UInt(CRC_%s)\n", c.id)
		for _, t := range c.fields {
			attrName := normalizeAttr(t.name)
			if t.isFlag() && t.typeName != "true" {
				write("if e.Flags & %d != 0 {\n", 1<<uint(t.flagBit))
			}
			switch t.typeName {
			case "true": //flags only
				write("//flag %s\n", attrName)
			case "int", "#":
				write("x.Int(e.%s)\n", attrName)
			case "long":
				write("x.Long(e.%s)\n", attrName)
			case "int128":
				write("x.Bytes(e.%s)\n", attrName)
			case "int256":
				write("x.Bytes(e.%s)\n", attrName)
			case "string":
				write("x.String(e.%s)\n", attrName)
			case "double":
				write("x.Double(e.%s)\n", attrName)
			case "bytes":
				write("x.StringBytes(e.%s)\n", attrName)
			case "Vector<int>":
				write("x.VectorInt(e.%s)\n", attrName)
			case "Vector<long>":
				write("x.VectorLong(e.%s)\n", attrName)
			case "Vector<string>":
				write("x.VectorString(e.%s)\n", attrName)
			case "Vector<double>":
				write("x.VectorDouble(e.%s)\n", attrName)
			case "!X":
				write("x.Bytes(e.%s.encode())\n", attrName)
			default:
				var inner string
				n, _ := fmt.Sscanf(t.typeName, "Vector<%s", &inner)
				if n == 1 {
					write("x.Vector(e.%s)\n", attrName)
				} else {
					write("x.Bytes(e.%s.encode())\n", attrName)
				}
			}
			if t.isFlag() && t.typeName != "true" {
				write("}\n")
			}
		}
		write("return x.buf\n")
		write("}\n\n")
	}

	// decode funcs
	write(`
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	for _, c := range combinators {
		write("case CRC_%s:\n", c.id)
		if c.hasFlags() {
			write("flags := m.Int()\n")
		}
		write("r = TL_%s{\n", c.id)
		for _, t := range c.fields {
			isFlag := t.isFlag()
			switch t.typeName {
			case "true": //flags only
				write("false, //flag\n")
			case "#":
				write("flags,\n")
			case "int":
				write(maybeFlagged("Int", isFlag, t.flagBit))
			case "long":
				write(maybeFlagged("Long", isFlag, t.flagBit))
			case "int128":
				write(maybeFlagged("Bytes", isFlag, t.flagBit, "16"))
			case "int256":
				write(maybeFlagged("Bytes", isFlag, t.flagBit, "32"))
			case "string":
				write(maybeFlagged("String", isFlag, t.flagBit))
			case "double":
				write(maybeFlagged("Double", isFlag, t.flagBit))
			case "bytes":
				write(maybeFlagged("StringBytes", isFlag, t.flagBit))
			case "Vector<int>":
				write(maybeFlagged("VectorInt", isFlag, t.flagBit))
			case "Vector<long>":
				write(maybeFlagged("VectorLong", isFlag, t.flagBit))
			case "Vector<string>":
				write(maybeFlagged("VectorString", isFlag, t.flagBit))
			case "Vector<double>":
				write(maybeFlagged("VectorDouble", isFlag, t.flagBit))
			case "!X":
				write(maybeFlagged("Object", isFlag, t.flagBit))
			default:
				var inner string
				n, _ := fmt.Sscanf(t.typeName, "Vector<%s", &inner)
				if n == 1 {
					write(maybeFlagged("Vector", isFlag, t.flagBit))
				} else {
					write(maybeFlagged("Object", isFlag, t.flagBit))
				}
			}
		}
		write("}\n\n")
	}

	write(`
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
