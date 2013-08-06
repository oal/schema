package schema

import (
	"encoding/json"
	"testing"
)

func TestBasic(t *testing.T) {
	data := []byte(`{
		"Title": "A basic test",
		"IntValue": 50
	}`)

	var v map[string]interface{}
	err := json.Unmarshal(data, &v)

	schema := map[string]interface{}{
		"Title":    Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{25}},
		"IntValue": Validator{required: true, fun: Builtin.Integer},
	}

	err = Validate(v, schema)
	if err != nil {
		panic(err)
	}
}

func TestBasicRecursive(t *testing.T) {
	data := []byte(`{
		"Title": "About me",
		"Websites": {
			"Twitter": {
				"Norwegian": "@oav",
				"English": "@lindekleiv"
			},
			"Github": "oal",
			"BitBucket": "lindekleiv"
		}
	}`)

	var v map[string]interface{}
	err := json.Unmarshal(data, &v)

	schema := map[string]interface{}{
		"Title": Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{25}},
		"Websites": map[string]interface{}{
			"Twitter": map[string]interface{}{
				"Norwegian": Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{20}},
				"English":   Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{20}},
			},
			"Tumblr":    Validator{required: false, fun: Builtin.MaxLength, args: []interface{}{100}},
			"Github":    Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{39}},
			"BitBucket": Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{39}},
		},
	}

	err = Validate(v, schema)
	if err != nil {
		panic(err)
	}
}
