package json2keys

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type json2Keys struct{}

func NewJson2Keys() *json2Keys {
	return &json2Keys{}
}

func (j *json2Keys) Parse(jsonStr string) (keys map[string]interface{}, err error) {
	var mapa interface{}
	if err = json.Unmarshal([]byte(jsonStr), &mapa); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return j.describeMap(mapa, "")
}

func (j *json2Keys) describeMap(mapa interface{}, index string) (secrets map[string]interface{}, err error) {
	secrets = make(map[string]interface{})
	if reflect.TypeOf(mapa) == reflect.TypeOf(map[string]interface{}{}) {
		return j.getSecretsFromMap(mapa.(map[string]interface{}), index)
	} else if reflect.TypeOf(mapa) == reflect.TypeOf(make([]interface{}, 0)) {
		for i, value := range mapa.([]interface{}) {
			newIndex := fmt.Sprintf("[%v]", i)
			if index != "" {
				newIndex = fmt.Sprintf("%s[%v]", index, i)
			}
			secr, err := j.describeMap(value, newIndex)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			for k, v := range secr {
				secrets[k] = v
			}
		}
	} else {
		return map[string]interface{}{
			index: mapa,
		}, nil
	}
	return secrets, nil
}

func (j *json2Keys) getSecretsFromMap(mapa map[string]interface{}, previousKey string) (secrets map[string]interface{}, err error) {
	secrets = make(map[string]interface{})
	for key, val := range mapa {
		mapa2 := make(map[string]interface{})
		if reflect.TypeOf(val) == reflect.TypeOf(mapa2) {
			bytes, _ := json.Marshal(val)
			err := json.Unmarshal(bytes, &mapa2)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			newMap, err := j.getSecretsFromMap(mapa2, key)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			for k, v := range newMap {
				if previousKey != "" {
					secrets[previousKey+"."+k] = v
				} else {
					secrets[k] = v
				}
			}
		} else if reflect.TypeOf(val) == reflect.TypeOf(make([]interface{}, 0)) {
			newMap, err := j.describeMap(val, key)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			for k, v := range newMap {
				if previousKey != "" {
					secrets[previousKey+"."+k] = v
				} else {
					secrets[k] = v
				}
			}
		} else {
			if previousKey != "" {
				secrets[previousKey+"."+key] = val
			} else {
				secrets[key] = val
			}
		}
	}
	return secrets, nil
}
