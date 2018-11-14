```go
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
```