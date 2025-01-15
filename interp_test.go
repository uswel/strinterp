package strinterp

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		variables map[string]interface{}
		want      string
		wantErr   bool
	}{
		{
			name:      "no expansion",
			input:     "Hello, world!",
			variables: nil,
			want:      "Hello, world!",
			wantErr:   false,
		},
		{
			name:      "simple expansion",
			input:     "Hello, ${name}!",
			variables: map[string]interface{}{"name": "Alice"},
			want:      "Hello, Alice!",
			wantErr:   false,
		},
		{
			name:      "multiple expansions",
			input:     "Name: ${name}, Age: ${age}",
			variables: map[string]interface{}{"name": "Bob", "age": 30},
			want:      "Name: Bob, Age: 30",
			wantErr:   false,
		},
		{
			name:      "escaped dollar sign",
			input:     "Price: $$${price}",
			variables: map[string]interface{}{"price": 10},
			want:      "Price: $10",
			wantErr:   false,
		},
		{
			name:      "unclosed expansion",
			input:     "Hello, ${name",
			variables: map[string]interface{}{"name": "Alice"},
			want:      "",
			wantErr:   true,
		},
		{
			name:      "undefined variable",
			input:     "Hello, ${undefined}!",
			variables: map[string]interface{}{},
			want:      "",
			wantErr:   true,
		},
		{
			name:      "integer formatting",
			input:     "Decimal: ${value:d}, Hex: ${value:x}, Octal: ${value:o}, Binary: ${value:b}",
			variables: map[string]interface{}{"value": 42},
			want:      "Decimal: 42, Hex: 2a, Octal: 52, Binary: 101010",
			wantErr:   false,
		},
		{
			name:      "float formatting",
			input:     "Float: ${value:.2f}, Scientific: ${value:e}",
			variables: map[string]interface{}{"value": 3.14159},
			want:      "Float: 3.14, Scientific: 3.14159e+00",
			wantErr:   false,
		},
		{
			name:      "string formatting",
			input:     "Upper: ${value:upper}, Lower: ${value:lower}, Title: ${value:title}",
			variables: map[string]interface{}{"value": "hello world"},
			want:      "Upper: HELLO WORLD, Lower: hello world, Title: HELLO WORLD",
			wantErr:   false,
		},
		{
			name:      "invalid format",
			input:     "Invalid: ${value:xyz}",
			variables: map[string]interface{}{"value": 10},
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := String(tt.input, tt.variables)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
