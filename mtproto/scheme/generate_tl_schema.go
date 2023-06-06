package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://github.com/telegramdesktop/tdesktop/tree/dev/Telegram/Resources/tl (merge two files into one separated by "---types---" line)
// https://github.com/danog/MadelineProto/tree/master/src/danog/MadelineProto
// https://github.com/tdlib/td/tree/master/td/generate/scheme

// https://core.telegram.org/mtproto/TL
// https://core.telegram.org/mtproto/TL-combinator
// identifier#name attr:type attr:type = resultType;

// TODO:
// Pq Ids
// Silent  bool //flag (add flag number)
// Channel TL   // InputChannel (add possible types)
// Use decodeResponse of inner obj when decoding types like:
//  invokeWithLayer#da9b0d0d {X:Type} layer:int query:!X = X;

type Flag struct {
	fieldName string
	bit       int
}

type Field struct {
	name     string
	typeName string
	flag     *Flag
}

type Combinator struct {
	id         string
	name       uint32
	fields     []Field
	typeName   string
	isFunction bool
}

func (c Combinator) flagFieldNames() []string {
	var names []string
	for _, f := range c.fields {
		if f.typeName == "#" {
			names = append(names, f.name)
		}
	}
	return names
}

func (c Combinator) structName() string {
	return "TL_" + c.id
}

func findConstructorIDs(combinators []*Combinator, fieldType string) []string {
	var constructorIDs []string
	for _, comb := range combinators {
		if !comb.isFunction && comb.typeName == fieldType {
			constructorIDs = append(constructorIDs, comb.structName())
		}
	}
	return constructorIDs
}

func parseVectorType(typeName string) (string, int, bool) {
	nesting := 0
	for strings.HasPrefix(typeName, "Vector<") || strings.HasPrefix(typeName, "vector<") {
		typeName = typeName[len("Vector<") : len(typeName)-len(">")]
		nesting += 1
	}
	return typeName, nesting, nesting > 0
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

func normalizeName(nameStr string) uint32 {
	if nameStr == "" {
		return 0
	}
	if nameStr[0] == '#' {
		nameStr = nameStr[1:]
	}
	nameInt, err := strconv.ParseInt(nameStr, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return uint32(nameInt)
}

func normalizeFieldName(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	if strings.HasSuffix(s, "Id") {
		s = s[:len(s)-2] + "ID"
	}
	return s
}

func maybeFlagged(_type string, flag *Flag, args ...string) string {
	argsStr := strings.Join(args, ",")
	if flag != nil {
		return fmt.Sprintf("m.Flagged%s(%s, %d, %s),\n", _type, flag.fieldName, flag.bit, argsStr)
	} else {
		return fmt.Sprintf("m.%s(%s),\n", _type, argsStr)
	}
}

func makeField(name, typeName string, knownFields []Field) Field {
	var flag *Flag
	for _, knownField := range knownFields {
		if knownField.typeName == "#" {
			// expecting smth like `flags.2?string`, assuming regular field names have no dots
			if strings.HasPrefix(typeName, knownField.name+".") {
				var err error
				qPos := strings.Index(typeName, "?")
				dPos := strings.Index(typeName, ".")
				flagFieldName := typeName[:dPos]
				flagBit, err := strconv.Atoi(typeName[dPos+1 : qPos])
				typeName = typeName[qPos+1:]
				flag = &Flag{fieldName: flagFieldName, bit: flagBit}
				if err != nil {
					log.Fatalf("parsing %s: %s", typeName, err)
				}
				break
			}
		}
	}
	return Field{normalize(name), normalize(typeName), flag}
}

var fieldsFixForCrcRegexp = regexp.MustCompile(`([Vv])ector<(.*?)>`)

func makeCombinatorDescription(id, fieldsStr, typeName string) string {
	if fieldsStr != "" {
		var fiteredFields []string
		for _, f := range strings.Split(fieldsStr, " ") {
			// for some reason if flagged field has type "true", it is NOT used in crc32 completely
			if strings.HasSuffix(f, "?true") {
				continue
			}
			fiteredFields = append(fiteredFields, f)
		}
		fieldsStr = strings.Join(fiteredFields, " ")
		// for some reason if field is "name:bytes" crc32 will be calculated from "name:string"
		fieldsStr = strings.Replace(fieldsStr, ":bytes", ":string", -1)
		fieldsStr = strings.Replace(fieldsStr, "?bytes", "?string", -1) //same for flags
		// for some reason... again
		fieldsStr = strings.Replace(fieldsStr, "{X:Type}", "X:Type", -1)
	}

	descr := id
	if fieldsStr != "" {
		descr += " " + fieldsStr
	}
	descr += " = " + typeName

	// for some reason if type is "Vector<subtype>" crc32 will be calculated from "Vector subtype"
	// and SOME TIMES it is named "vector<subtype>" (with lower "v")
	descr = fieldsFixForCrcRegexp.ReplaceAllString(descr, "${1}ector $2")
	return descr
}

func parseTLSchema(fpath string) []*Combinator {
	// reading tl file
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	// processing constructors
	combinators := []*Combinator{}
	isFunction := false

	lineRegexp := regexp.MustCompile(`^(.*?)(#[a-f0-9]*)? (.*)= (.*);$`)
	fieldRegexp := regexp.MustCompile(`^(.*?):(.*)$`)
	for lineNum, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if line == "---functions---" {
			isFunction = true
			continue
		}
		if line == "---types---" {
			isFunction = false
			continue
		}

		match := lineRegexp.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Printf("line %d: wrong combinator: %s", lineNum+1, line)
			continue
		}

		id := strings.TrimSpace(match[1])
		if id == "vector" {
			continue
		}
		name := normalizeName(match[2])
		fieldsStr := strings.TrimSpace(match[3])
		typeName := strings.TrimSpace(match[4])

		// making combinator description string (without id) and checking it's crc32
		descr := makeCombinatorDescription(id, fieldsStr, typeName)
		crc32sum := normalizeName(fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(descr))))
		if name == 0 {
			log.Printf("WARN: line %d: missing crc32 sum: %s", lineNum+1, line)
			name = crc32sum
		} else if name != crc32sum {
			log.Printf("WARN: line %d: wrong crc32 sum, expected %08x: %s", lineNum+1, crc32sum, line)
		}

		id = normalize(id)
		typeName = normalize(typeName)

		fields := make([]Field, 0, 16)
		for _, fieldStr := range strings.Split(fieldsStr, " ") {
			fieldStr = strings.TrimSpace(fieldStr)
			if fieldStr == "" {
				continue
			}
			if strings.HasPrefix(fieldStr, "{") { //if it is "{X:Type}", just skipping, "!X" type will be written as "TL" later
				continue
			}
			match := fieldRegexp.FindStringSubmatch(fieldStr)
			if len(match) == 0 {
				log.Fatalf("line %d: wrong field: %s", lineNum+1, fieldStr)
			}
			name, typeName := match[1], match[2]
			fields = append(fields, makeField(name, typeName, fields))
		}

		combinators = append(combinators, &Combinator{id, name, fields, typeName, isFunction})
	}
	return combinators
}

func main() {
	if len(os.Args) != 4 {
		println("Usage: " + os.Args[0] + " layer tl_schema.tl tl_schema.go")
		os.Exit(2)
	}
	layer, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fpath := os.Args[2]

	// parsing
	combinators := parseTLSchema(fpath)

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
	write(`package mtproto
import (
	"github.com/ansel1/merry/v2"
)
`)
	write("const (\n")
	write("TL_Layer = %d\n", layer)
	for _, c := range combinators {
		write("CRC_%s = 0x%08x\n", c.id, c.name)
	}
	write(")\n\n")

	// type structs
	for _, c := range combinators {
		write("type %s struct {\n", c.structName())
		for _, t := range c.fields {
			fieldComment := ""
			write("%s\t", normalizeFieldName(t.name))
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
			case "Vector<bytes>":
				write("[][]byte")
			case "!X":
				write("TL")
			default:
				innerTypeName, vecNesting, _ := parseVectorType(t.typeName)

				constructorIDs := findConstructorIDs(combinators, innerTypeName)
				if len(constructorIDs) == 0 {
					log.Printf("WARN: no constructors for type %s in %s { %s:%s }", innerTypeName, c.id, t.name, t.typeName)
				}

				write("%sTL", strings.Repeat("[]", vecNesting))
				write(" // %s", innerTypeName)
				// fieldComment += innerTypeName + ": " + strings.Join(constructorIDs, " | ") TODO
			}
			if t.flag != nil {
				write(" //flag")
				// fieldComment = fmt.Sprintf("(%s.%d) %s", t.flag.fieldName, t.flag.bit, fieldComment) TODO
			}

			if fieldComment != "" {
				// write("// %s", strings.TrimSpace(fieldComment)) TODO
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
			fieldName := normalizeFieldName(t.name)
			if t.flag != nil && t.typeName != "true" {
				write("if e.%s & %d != 0 {\n", normalizeFieldName(t.flag.fieldName), 1<<uint(t.flag.bit))
			}
			switch t.typeName {
			case "true": //flags only
				write(" //flag %s\n", fieldName)
				// write("// %s.%d %s\n", t.flag.fieldName, t.flag.bit, fieldName) TODO
			case "int", "#":
				write("x.Int(e.%s)\n", fieldName)
			case "long":
				write("x.Long(e.%s)\n", fieldName)
			case "int128":
				write("x.Bytes(e.%s)\n", fieldName)
			case "int256":
				write("x.Bytes(e.%s)\n", fieldName)
			case "string":
				write("x.String(e.%s)\n", fieldName)
			case "double":
				write("x.Double(e.%s)\n", fieldName)
			case "bytes":
				write("x.StringBytes(e.%s)\n", fieldName)
			case "Vector<int>":
				write("x.VectorInt(e.%s)\n", fieldName)
			case "Vector<long>":
				write("x.VectorLong(e.%s)\n", fieldName)
			case "Vector<string>":
				write("x.VectorString(e.%s)\n", fieldName)
			case "Vector<double>":
				write("x.VectorDouble(e.%s)\n", fieldName)
			case "Vector<bytes>":
				write("x.VectorBytes(e.%s)\n", fieldName)
			case "!X":
				write("x.Bytes(e.%s.encode())\n", fieldName)
			default:
				_, vecNesting, _ := parseVectorType(t.typeName)
				if vecNesting == 1 {
					write("x.Vector(e.%s)\n", fieldName)
				} else if vecNesting == 2 {
					write("x.Vector2d(e.%s)\n", fieldName)
				} else {
					write("x.Bytes(e.%s.encode())\n", fieldName)
				}
			}
			if t.flag != nil && t.typeName != "true" {
				write("}\n")
			}
		}
		write("return x.buf\n")
		write("}\n\n")
	}

	// request decode funcs (for funtions)
	for _, c := range combinators {
		if c.isFunction {
			write("func (e TL_%s) decodeResponse(dbuf *DecodeBuf) TL {\n", c.id)
			if c.typeName == "Vector<int>" {
				write("return VectorInt(dbuf.VectorInt())\n")
			} else if c.typeName == "Vector<long>" {
				write("return VectorLong(dbuf.VectorLong())\n")
			} else if strings.HasPrefix(c.typeName, "Vector<") {
				write("return VectorObject(dbuf.Vector())\n")
			} else {
				write("return dbuf.Object()\n")
			}
			write("}\n\n")
		}
	}

	// decode funcs
	write(`
func readFlags(m *DecodeBuf, flagsPtr *int32) int32 {
	flags := m.Int()
	*flagsPtr = flags
	return flags
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	objStartOffset := m.off - 4 //4 bytes of constructor name have been already read
	switch constructor {`)

	for _, c := range combinators {
		write("case CRC_%s:\n", c.id)
		flagNames := c.flagFieldNames()
		if len(flagNames) > 0 {
			write("var %s int32\n", strings.Join(flagNames, ","))
		}
		write("r = TL_%s{\n", c.id)
		for _, t := range c.fields {
			switch t.typeName {
			case "true": //flags only
				write("%s & %d != 0, //flag #%d\n", t.flag.fieldName, 1<<uint(t.flag.bit), t.flag.bit)
				// write("%s & %d != 0, //%s.%d\n", t.flag.fieldName, 1<<uint(t.flag.bit), t.flag.fieldName, t.flag.bit) TODO
			case "#":
				write("readFlags(m, &%s),\n", t.name)
			case "int":
				write(maybeFlagged("Int", t.flag))
			case "long":
				write(maybeFlagged("Long", t.flag))
			case "int128":
				write(maybeFlagged("Bytes", t.flag, "16"))
			case "int256":
				write(maybeFlagged("Bytes", t.flag, "32"))
			case "string":
				write(maybeFlagged("String", t.flag))
			case "double":
				write(maybeFlagged("Double", t.flag))
			case "bytes":
				write(maybeFlagged("StringBytes", t.flag))
			case "Vector<int>":
				write(maybeFlagged("VectorInt", t.flag))
			case "Vector<long>":
				write(maybeFlagged("VectorLong", t.flag))
			case "Vector<string>":
				write(maybeFlagged("VectorString", t.flag))
			case "Vector<double>":
				write(maybeFlagged("VectorDouble", t.flag))
			case "Vector<bytes>":
				write(maybeFlagged("VectorBytes", t.flag))
			case "!X":
				write(maybeFlagged("Object", t.flag))
			default:
				_, vecNesting, _ := parseVectorType(t.typeName)
				if vecNesting == 1 {
					write(maybeFlagged("Vector", t.flag))
				} else if vecNesting == 2 {
					write(maybeFlagged("Vector2d", t.flag))
				} else {
					write(maybeFlagged("Object", t.flag))
				}
			}
		}
		write("}\n\n")
	}

	write(`
	default:
		m.err = merry.Errorf("Unknown constructor: %%08x", constructor)
		return nil

	}

	if m.err != nil {
		m.pushToErrBufStack(objStartOffset, constructor)
		return nil
	}

	return
}`)
}
