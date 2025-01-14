package strinterp

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// String expands variables within a string using the syntax ${var} or ${var:format}.
func String(input string, variables map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	i := 0
	for i < len(input) {
		if input[i] == '$' {
			if i+1 < len(input) {
				if input[i+1] == '{' {
					// Potential variable expansion
					j := i + 2
					for j < len(input) && input[j] != '}' {
						j++
					}
					if j == len(input) {
						return "", fmt.Errorf("unclosed variable expansion at position %d", i)
					}

					varName, format := parseVar(input[i+2 : j])
					if val, ok := variables[varName]; ok {
						expanded, err := formatValue(val, format)
						if err != nil {
							return "", err
						}
						buf.WriteString(expanded)
					} else {
						return "", fmt.Errorf("undefined variable '%s' at position %d", varName, i)
					}

					i = j + 1
				} else if input[i+1] == '$' {
					// Escaped dollar sign
					buf.WriteByte('$')
					i += 2
				} else {
					buf.WriteByte(input[i])
					i++
				}
			} else {
				buf.WriteByte(input[i])
				i++
			}
		} else {
			buf.WriteByte(input[i])
			i++
		}
	}
	return buf.String(), nil
}

// parseVar extracts the variable name and optional format from the expression.
func parseVar(expression string) (name string, format string) {
	parts := strings.SplitN(expression, ":", 2)
	name = parts[0]
	if len(parts) > 1 {
		format = parts[1]
	}
	return
}

// formatValue formats the variable value based on the optional format.
func formatValue(value interface{}, format string) (string, error) {
	if format == "" {
		return fmt.Sprintf("%v", value), nil
	}

	switch v := value.(type) {
	case int:
		return formatInt(int64(v), format)
	case int8:
		return formatInt(int64(v), format)
	case int16:
		return formatInt(int64(v), format)
	case int32:
		return formatInt(int64(v), format)
	case int64:
		return formatInt(v, format)
	case float32:
		return formatFloat(float64(v), format)
	case float64:
		return formatFloat(v, format)
	case string:
		return formatString(v, format)
	default:
		return "", fmt.Errorf("unsupported format '%s' for value type %T", format, value)
	}
}

func formatInt(value int64, format string) (string, error) {
	switch format {
	case "x":
		return strconv.FormatInt(value, 16), nil
	case "X":
		return strings.ToUpper(strconv.FormatInt(value, 16)), nil
	case "o":
		return strconv.FormatInt(value, 8), nil
	case "b":
		return strconv.FormatInt(value, 2), nil
	case "d":
		return strconv.FormatInt(value, 10), nil
	default:
		return "", fmt.Errorf("unsupported integer format '%s'", format)
	}
}

func formatFloat(value float64, format string) (string, error) {
	if strings.HasPrefix(format, ".") {
		end := 1
		for ; end < len(format); end++ {
			if format[end] < '0' || format[end] > '9' {
				break
			}
		}

		if precision, err := strconv.Atoi(format[1:end]); err == nil {
			// Default format character is 'f'
			formatChar := byte('f')
			if end < len(format) {
				// Ensure the format character is a byte
				formatChar = format[end]
			}
			return strconv.FormatFloat(value, formatChar, precision, 64), nil
		} else {
			return "", fmt.Errorf("invalid precision in float format '%s'", format)
		}
	}
	switch format {
	case "e":
		return strconv.FormatFloat(value, 'e', -1, 64), nil
	case "E":
		return strconv.FormatFloat(value, 'E', -1, 64), nil
	case "f":
		return strconv.FormatFloat(value, 'f', -1, 64), nil
	case "g":
		return strconv.FormatFloat(value, 'g', -1, 64), nil
	case "G":
		return strconv.FormatFloat(value, 'G', -1, 64), nil
	default:
		return "", fmt.Errorf("unsupported float format '%s'", format)
	}
}
func formatString(value string, format string) (string, error) {
	switch format {
	case "upper":
		return strings.ToUpper(value), nil
	case "lower":
		return strings.ToLower(value), nil
	case "title":
		return strings.ToTitle(value), nil
	default:
		return "", fmt.Errorf("unsupported string format '%s'", format)
	}
}
