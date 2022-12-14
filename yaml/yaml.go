/**
  @author: 1043193460@qq.com
  @date: 2022/12/14 12:02
  @note: 这个包首先使用goyaml将YAML转换为JSON，然后使用json。Marshal和json。要转换为结构或从结构转换为结构的反编组
		另请参见http://ghodss.com/2014/the-right-way-to-handle-yaml-in-golang
**/
package yaml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"reflect"
	"strconv"
)

type YamlUtils struct {
}

//将对象编组为JSON，然后将JSON转换为YAML并返回YAML。
func (y *YamlUtils) Marshal(o interface{}) ([]byte, error) {
	j, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("error marshaling into JSON: %v", err)
	}

	b, err := y.JSONToYAML(j)
	if err != nil {
		return nil, fmt.Errorf("error converting JSON to YAML: %v", err)
	}

	return b, nil
}

//JSONOpt是用于从JSON格式解码的解码选项。
type JSONOpt func(*json.Decoder) *json.Decoder

//解组将YAML转换为JSON，然后使用JSON将其解组为对象，可选地配置JSON解组的行为。
func (y *YamlUtils) Unmarshal(b []byte, o interface{}, opts ...JSONOpt) error {
	return unmarshal(yaml.Unmarshal, b, o, opts)
}

//UnmarshallStrict与Unmarshall类似重复将导致错误。要严格控制未知字段，请添加DisableUnknownFields选项。
func (y *YamlUtils) UnmarshalStrict(b []byte, o interface{}, opts ...JSONOpt) error {
	return unmarshal(yaml.UnmarshalStrict, b, o, opts)
}

func unmarshal(f func(in []byte, out interface{}) (err error), y []byte, o interface{}, opts []JSONOpt) error {
	vo := reflect.ValueOf(o)
	j, err := yamlToJSON(y, &vo, f)
	if err != nil {
		return fmt.Errorf("error converting YAML to JSON: %v", err)
	}

	err = jsonUnmarshal(bytes.NewReader(j), o, opts...)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return nil
}

//jsonUnmarshal将来自给定读取器的JSON字节流解组到对象，可选地在解码之前应用解码器选项。
//我们不是使用json。直接取消编组，因为我们希望机会以非默认方式传递选项。
func jsonUnmarshal(r io.Reader, o interface{}, opts ...JSONOpt) error {
	d := json.NewDecoder(r)
	for _, opt := range opts {
		d = opt(d)
	}
	if err := d.Decode(&o); err != nil {
		return fmt.Errorf("while decoding JSON: %v", err)
	}
	return nil
}

//将JSON转换为YAML。
func (y *YamlUtils) JSONToYAML(j []byte) ([]byte, error) {
	var jsonObj interface{}
	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(jsonObj)
}

//YAMLToJSON将YAML转换为JSON。由于JSON是YAML的子集，通过此方法传递JSON应该是一个错误。
func (y *YamlUtils) YAMLToJSON(b []byte) ([]byte, error) {
	return yamlToJSON(b, nil, yaml.Unmarshal)
}

//YAMLToJSONStrict类似于YAMLToJSON，但允许严格的YAML解码，返回任何重复字段名的错误。
func (y *YamlUtils) YAMLToJSONStrict(b []byte) ([]byte, error) {
	return yamlToJSON(b, nil, yaml.UnmarshalStrict)
}

func yamlToJSON(y []byte, jsonTarget *reflect.Value, yamlUnmarshal func([]byte, interface{}) error) ([]byte, error) {
	var yamlObj interface{}
	err := yamlUnmarshal(y, &yamlObj)
	if err != nil {
		return nil, err
	}
	jsonObj, err := convertToJSONableObject(yamlObj, jsonTarget)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonObj)
}

func convertToJSONableObject(yamlObj interface{}, jsonTarget *reflect.Value) (interface{}, error) {
	var err error
	if jsonTarget != nil {
		ju, tu, pv := indirect(*jsonTarget, false)
		if ju != nil || tu != nil {
			jsonTarget = nil
		} else {
			jsonTarget = &pv
		}
	}
	switch typedYAMLObj := yamlObj.(type) {
	case map[interface{}]interface{}:
		strMap := make(map[string]interface{})
		for k, v := range typedYAMLObj {
			// Resolve the key to a string first.
			var keyString string
			switch typedKey := k.(type) {
			case string:
				keyString = typedKey
			case int:
				keyString = strconv.Itoa(typedKey)
			case int64:
				keyString = strconv.FormatInt(typedKey, 10)
			case float64:
				s := strconv.FormatFloat(typedKey, 'g', -1, 32)
				switch s {
				case "+Inf":
					s = ".inf"
				case "-Inf":
					s = "-.inf"
				case "NaN":
					s = ".nan"
				}
				keyString = s
			case bool:
				if typedKey {
					keyString = "true"
				} else {
					keyString = "false"
				}
			default:
				return nil, fmt.Errorf("Unsupported map key of type: %s, key: %+#v, value: %+#v",
					reflect.TypeOf(k), k, v)
			}

			if jsonTarget != nil {
				t := *jsonTarget
				if t.Kind() == reflect.Struct {
					keyBytes := []byte(keyString)
					var f *field
					fields := cachedTypeFields(t.Type())
					for i := range fields {
						ff := &fields[i]
						if bytes.Equal(ff.nameBytes, keyBytes) {
							f = ff
							break
						}
						if f == nil && ff.equalFold(ff.nameBytes, keyBytes) {
							f = ff
						}
					}
					if f != nil {
						jtf := t.Field(f.index[0])
						strMap[keyString], err = convertToJSONableObject(v, &jtf)
						if err != nil {
							return nil, err
						}
						continue
					}
				} else if t.Kind() == reflect.Map {
					jtv := reflect.Zero(t.Type().Elem())
					strMap[keyString], err = convertToJSONableObject(v, &jtv)
					if err != nil {
						return nil, err
					}
					continue
				}
			}
			strMap[keyString], err = convertToJSONableObject(v, nil)
			if err != nil {
				return nil, err
			}
		}
		return strMap, nil
	case []interface{}:
		var jsonSliceElemValue *reflect.Value
		if jsonTarget != nil {
			t := *jsonTarget
			if t.Kind() == reflect.Slice {
				ev := reflect.Indirect(reflect.New(t.Type().Elem()))
				jsonSliceElemValue = &ev
			}
		}

		arr := make([]interface{}, len(typedYAMLObj))
		for i, v := range typedYAMLObj {
			arr[i], err = convertToJSONableObject(v, jsonSliceElemValue)
			if err != nil {
				return nil, err
			}
		}
		return arr, nil
	default:
		if jsonTarget != nil && (*jsonTarget).Kind() == reflect.String {
			var s string
			switch typedVal := typedYAMLObj.(type) {
			case int:
				s = strconv.FormatInt(int64(typedVal), 10)
			case int64:
				s = strconv.FormatInt(typedVal, 10)
			case float64:
				s = strconv.FormatFloat(typedVal, 'g', -1, 32)
			case uint64:
				s = strconv.FormatUint(typedVal, 10)
			case bool:
				if typedVal {
					s = "true"
				} else {
					s = "false"
				}
			}
			if len(s) > 0 {
				yamlObj = interface{}(s)
			}
		}
		return yamlObj, nil
	}

	return nil, nil
}
