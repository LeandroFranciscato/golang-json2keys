# golang-json2keys
Returns a map of key/value from given json

# Usage Example

```golang 
 
 jsonString := 
 `{
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
}`
 
 json2Keys := json2Keys.NewJson2Keys()
 keys, err := json2Keys.Parse(jsonStr)
 if err != nil {
    panic(err)
 }
 
 fmt.Println(keys)

```

This snipet should return a map like this:
```golang

map[string]interface{}{
  "tags[0]":            "laborum",
  "tags[1]":            "enim",
  "tags[2]":            "consequat",
  "array1[0].object11": "name1",
  "array1[1].object12": "name2",
  "array2[0].object21": "name1",
  "array2[1].object22": "name2",
  "outsideobj":         "value",
}
```

