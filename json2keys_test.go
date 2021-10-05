package json2keys

import (
	"reflect"
	"testing"
)

func TestNewJson2Keys(t *testing.T) {
	tests := []struct {
		name string
		want *json2Keys
	}{
		{
			name: "Success",
			want: &json2Keys{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJson2Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJson2Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_json2Keys_Parse(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name     string
		j        *json2Keys
		args     args
		wantKeys map[string]interface{}
		wantErr  bool
	}{
		{
			name: "Error on unmarshal",
			j:    &json2Keys{},
			args: args{
				jsonStr: `{"id": 0,"nome" "leandro"}`,
			},
			wantErr: true,
		},
		{
			name: "Success whitout recursion",
			j:    &json2Keys{},
			args: args{
				jsonStr: `{"id": 0,"nome": "leandro"}`,
			},
			wantKeys: map[string]interface{}{
				"id":   float64(0),
				"nome": "leandro",
			},
			wantErr: false,
		},
		{
			name: "Success whit recursion",
			j:    &json2Keys{},
			args: args{
				jsonStr: `[{"id": 0,"nome": "leandro", "subJson": {"subid": 1, "subnome": "leandro", "subsubjson": {"subsubid":2, "subsubnome":"subsubnome"}}}]`,
			},
			wantKeys: map[string]interface{}{
				"[0].id":                            float64(0),
				"[0].nome":                          "leandro",
				"[0].subJson.subid":                 float64(1),
				"[0].subJson.subnome":               "leandro",
				"[0].subJson.subsubjson.subsubid":   float64(2),
				"[0].subJson.subsubjson.subsubnome": "subsubnome",
			},
			wantErr: false,
		},
		{
			name: "Success whit recursion and more than one nested block",
			j:    &json2Keys{},
			args: args{
				jsonStr: `{
					"object": [
						{
							"tokenA": "asdfasdfasdfasdfA"
						},
						{
							"tokenB": "asdfasdfasdfasdfB"
						},
						[
							{
								"tokenC": "asdfasdfasdfasdfC"
							},
							{
								"tokenD": "asdfasdfasdfasdfD"
							}
						]
					]
				}`,
			},
			wantKeys: map[string]interface{}{
				"object[0].tokenA":    "asdfasdfasdfasdfA",
				"object[1].tokenB":    "asdfasdfasdfasdfB",
				"object[2][0].tokenC": "asdfasdfasdfasdfC",
				"object[2][1].tokenD": "asdfasdfasdfasdfD",
			},
			wantErr: false,
		},
		{
			name: "Success whit recursion and more than one nested block, initializing with an array",
			j:    &json2Keys{},
			args: args{
				jsonStr: `[{
					"object": [
						{
							"tokenA": "asdfasdfasdfasdfA"
						},
						{
							"tokenB": "asdfasdfasdfasdfB"
						},
						[
							{
								"tokenC": "asdfasdfasdfasdfC"
							},
							{
								"tokenD": "asdfasdfasdfasdfD"
							}
						]
					]
				}]`,
			},
			wantKeys: map[string]interface{}{
				"[0].object[0].tokenA":    "asdfasdfasdfasdfA",
				"[0].object[1].tokenB":    "asdfasdfasdfasdfB",
				"[0].object[2][0].tokenC": "asdfasdfasdfasdfC",
				"[0].object[2][1].tokenD": "asdfasdfasdfasdfD",
			},
			wantErr: false,
		},
		{
			name: "Success whit an array of string",
			j:    &json2Keys{},
			args: args{
				jsonStr: `{
					"tags": [
					  "laborum",
					  "enim",
					  "consequat"
					]}`,
			},
			wantKeys: map[string]interface{}{
				"tags[0]": "laborum",
				"tags[1]": "enim",
				"tags[2]": "consequat",
			},
			wantErr: false,
		},
		{
			name: "Success whit an array of string and more objects",
			j:    &json2Keys{},
			args: args{
				jsonStr: `{
					"tags": [
						"laborum",
						"enim",
						"consequat"
					],
					"array1": [
						{
							"object11": "name1"
						},
						{
							"object12": "name2"
						}
					],
					"array2": [
						{
							"object21": "name1"
						},
						{
							"object22": "name2"
						}
					],
					"outsideobj": "value"
				}`,
			},
			wantKeys: map[string]interface{}{
				"tags[0]":            "laborum",
				"tags[1]":            "enim",
				"tags[2]":            "consequat",
				"array1[0].object11": "name1",
				"array1[1].object12": "name2",
				"array2[0].object21": "name1",
				"array2[1].object22": "name2",
				"outsideobj":         "value",
			},
			wantErr: false,
		},
		{
			name: "Success whit nested objects",
			j:    &json2Keys{},
			args: args{
				jsonStr: `{"data":{"parameter": "second"}}`,
			},
			wantKeys: map[string]interface{}{
				"data.parameter": "second",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &json2Keys{}
			gotKeys, err := j.Parse(tt.args.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("json2Keys.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotKeys) == 0 && !tt.wantErr {
				t.Errorf("json2Keys.Parse() got nothing")
			}

			for k, v := range gotKeys {
				if tt.wantKeys[k] != v {
					t.Errorf("json2Keys.Parse() got %v, want %v", v, tt.wantKeys[k])
				}
			}
		})
	}
}
