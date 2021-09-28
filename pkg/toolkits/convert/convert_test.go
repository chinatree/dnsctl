package convert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToJSON(t *testing.T) {
	type People struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	tests := []struct {
		name    string
		input   People
		wantErr bool
	}{
		{
			input: People{
				Name: "张三",
				Age:  18,
			},
		},
		{
			input: People{
				Name: "李四",
				Age:  20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJSON(tt.input)
			assert.Nil(t, err)
			fmt.Println(got)
		})
	}
}

func TestStructToJSONWithIndent(t *testing.T) {
	type People struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	tests := []struct {
		name    string
		input   People
		wantErr bool
	}{
		{
			input: People{
				Name: "张三",
				Age:  18,
			},
		},
		{
			input: People{
				Name: "李四",
				Age:  20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJSONWithIndent(tt.input)
			assert.Nil(t, err)
			fmt.Println(got)
		})
	}
}

func TestStringToJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			input: `{"name":"张三","age":18}`,
		},
		{
			input: "{\"name\":\"李四\",\"age\":20}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToJSON(tt.input)
			assert.Nil(t, err)
			fmt.Println(got)
		})
	}
}

func TestStringToJSONWithIndent(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			input: `{"name":"张三","age":18}`,
		},
		{
			input: "{\"name\":\"李四\",\"age\":20}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToJSONWithIndent(tt.input)
			assert.Nil(t, err)
			fmt.Println(got)
		})
	}
}
