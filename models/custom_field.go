package msaevents

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type CustomFieldMapLabelValue struct {
	Label string
	Value string
}

type CustomField struct {
	Field CustomFieldConfig `json:"field"`
	Data  *CustomFieldData  `json:"data"`
}

type CustomFieldData struct {
	FieldId    int64       `json:"field_id"`
	FieldIdStr string      `json:"field_id_str"`
	Value      interface{} `json:"value"`
}

func (c *CustomFieldData) StringValue() string {
	if c.Value == nil {
		return ""
	}
	v := reflect.ValueOf(c.Value)
	if v.IsNil() {
		return ""
	}
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Slice, reflect.Array:
		output := []string{}
		for i := 0; i < v.Len(); i++ {
			output = append(output, fmt.Sprint(v.Index(i)))
		}
		return strings.Join(output, ",")
	case reflect.Int, reflect.Int64:
		return strconv.Itoa(int(v.Int()))
	case reflect.Float32, reflect.Float64:
		return strconv.Itoa(int(v.Float()))
	default:
		//fmt.Printf("fail to found type of %+v\n", v.String())
		return ""
	}
}

type CustomFieldConfig struct {
	Id          int64
	IdStr       string
	Labels      map[string]string
	Enabled     bool
	FieldType   string
	Position    *int
	Description string
}

func (c *CustomFieldConfig) Label(lang string) string {
	if label, ok := c.Labels[lang]; ok {
		return label
	}
	if label, ok := c.Labels[fallbackLanguage]; ok {
		return label
	}
	return ""
}
