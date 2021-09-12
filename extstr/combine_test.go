package extstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecombine(t *testing.T) {
	type args struct {
		str string
		sep byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符",
			args{
				str: "",
				sep: '_',
			},
			"",
		},
		{
			"大小写",
			args{
				str: "HelloWorld",
				sep: '_',
			},
			"hello_world",
		},
		{
			"大小写并带分隔符",
			args{
				str: "Hello_World",
				sep: '_',
			},
			"hello_world",
		},
		{
			"小写带分隔符",
			args{
				str: "HelloWor_ldID",
				sep: '_',
			},
			"hello_wor_ld_id",
		},
		{
			"小写带分隔符",
			args{
				str: "HelloWor_ldA",
				sep: '_',
			},
			"hello_wor_ld_a",
		},
		{
			"特殊分隔IDCom",
			args{
				str: "IDCom",
				sep: '_',
			},
			"id_com",
		},
		{
			"特殊分隔IDcom",
			args{
				str: "IDcom",
				sep: '_',
			},
			"idcom",
		},
		{
			"特殊分隔nameIDCom",
			args{
				str: "nameIDCom",
				sep: '_',
			},
			"name_id_com",
		},
		{
			"特殊分隔nameIDcom",
			args{
				str: "nameIDcom",
				sep: '_',
			},
			"name_idcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Recombine(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("Recombine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnRecombine(t *testing.T) {
	type args struct {
		str string
		sep byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符",
			args{
				str: "",
				sep: '_',
			},
			"",
		},
		{
			"以_分隔符",
			args{
				str: "hello_world",
				sep: '_',
			},
			"HelloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnRecombine(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("UnRecombine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符IDCom",
			args{str: "IDCom"},
			"id_com",
		},
		{
			"特殊字符IDcom",
			args{str: "IDcom"},
			"idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "nameIDCom"},
			"name_id_com",
		},
		{
			"特殊字符nameIDcom",
			args{str: "nameIDcom"},
			"name_idcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.str); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKebab(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符IDCom",
			args{str: "IDCom"},
			"id-com",
		},
		{
			"特殊字符IDcom",
			args{str: "IDcom"},
			"idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "nameIDCom"},
			"name-id-com",
		},
		{
			"特殊字符nameIDcom",
			args{str: "nameIDcom"},
			"name-idcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kebab(tt.args.str); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符IDCom",
			args{str: "id_com"},
			"IDCom",
		},
		{
			"特殊字符IDcom",
			args{str: "idcom"},
			"Idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "name_id_com"},
			"NameIDCom",
		},
		{
			"特殊字符nameIDcom",
			args{str: "name_idcom"},
			"NameIdcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.str); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNames(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		fieldName    string
		expectedName string
	}{
		{
			name:         "simple lowercase",
			fieldName:    "name",
			expectedName: "Name",
		},
		{
			name:         "snake case",
			fieldName:    "field_name",
			expectedName: "FieldName",
		},
		{
			name:         "long snake",
			fieldName:    "field__name",
			expectedName: "FieldName",
		},
		{
			name:         "snake at ends",
			fieldName:    "_field_name_",
			expectedName: "FieldName",
		},
		{
			name:         "long snake at ends",
			fieldName:    "__field_name",
			expectedName: "FieldName",
		},
		{
			name:         "snake case with initialism",
			fieldName:    "html_field_name",
			expectedName: "HTMLFieldName",
		},
		{
			name:         "camel case",
			fieldName:    "camelCaseName",
			expectedName: "CamelCaseName",
		},
		{
			name:         "mixed case",
			fieldName:    "mixed_caseName_test",
			expectedName: "MixedCaseNameTest",
		},
		{
			name:         "mixed case with initialism",
			fieldName:    "mixed_caseName_htmlTest",
			expectedName: "MixedCaseNameHTMLTest",
		},
		{
			name:         "special chars",
			fieldName:    "$field_$name日本語",
			expectedName: "FieldName",
		},
		{
			name:         "garbage",
			fieldName:    "$@!%^&*()",
			expectedName: "",
		},
		{
			name:         "starting with digits",
			fieldName:    "123key",
			expectedName: "Key",
		},
		{
			name:         "name with digits",
			fieldName:    "key_666",
			expectedName: "Key666",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedName, CamelCase(tc.fieldName))
		})
	}
}

func TestSmallCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符idCom",
			args{str: "id_com"},
			"idCom",
		},
		{
			"特殊字符idcom",
			args{str: "idcom"},
			"idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "name_id_com"},
			"nameIDCom",
		},
		{
			"特殊字符nameIDcom",
			args{str: "name_idcom"},
			"nameIdcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallCamelCase(tt.args.str); got != tt.want {
				t.Errorf("SmallCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowTitle(t *testing.T) {
	var LowTitleTests = []struct {
		in, out string
	}{
		{"", ""},
		{"A", "a"},
		{" aaa aaa aaa ", " aaa aaa aaa "},
		{" Aaa Aaa Aaa ", " aaa aaa aaa "},
		{"123a456", "123a456"},
		{"Double-Blind", "double-blind"},
		{"Ÿøû", "ÿøû"},
		{"With_underscore", "with_underscore"},
		{"Unicode \xe2\x80\xa8 Line Separator", "unicode \xe2\x80\xa8 line separator"},
	}
	for _, tt := range LowTitleTests {
		if s := LowTitle(tt.in); s != tt.out {
			t.Errorf("LowTitle(%q) = %q, want %q", tt.in, s, tt.out)
		}
	}
}
