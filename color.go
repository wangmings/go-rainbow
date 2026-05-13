package rainbow

import (
	"fmt"
	"strconv"
	"strings"
)

func colorCode(bg bool, values ...any) (string, error) {
	if len(values) != 1 && len(values) != 3 {
		return "", fmt.Errorf("color expects 1 name/index/hex value or 3 RGB values")
	}
	if len(values) == 3 {
		r, rok := asInt(values[0])
		g, gok := asInt(values[1])
		b, bok := asInt(values[2])
		if !rok || !gok || !bok {
			return "", fmt.Errorf("RGB color expects integers")
		}
		return rgbCode(bg, r, g, b)
	}

	switch v := values[0].(type) {
	case int:
		if v < 0 || v > 9 {
			return "", fmt.Errorf("indexed ANSI color outside 0-9")
		}
		return strconv.Itoa(base(bg) + v), nil
	case string:
		name := strings.ToLower(strings.TrimSpace(v))
		if n, ok := ansiColors[name]; ok {
			return strconv.Itoa(base(bg) + n), nil
		}
		if rgb, ok := x11Colors[name]; ok {
			return rgbCode(bg, rgb[0], rgb[1], rgb[2])
		}
		if rgb, ok := parseHex(name); ok {
			return rgbCode(bg, rgb[0], rgb[1], rgb[2])
		}
	}

	return "", fmt.Errorf("unknown color %v", values[0])
}

func rgbCode(bg bool, r, g, b int) (string, error) {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return "", fmt.Errorf("RGB value outside 0-255")
	}
	prefix := 38
	if bg {
		prefix = 48
	}
	code := 16 + cube(r)*36 + cube(g)*6 + cube(b)
	return fmt.Sprintf("%d;5;%d", prefix, code), nil
}

func cube(v int) int { return 6 * v / 256 }

func base(bg bool) int {
	if bg {
		return 40
	}
	return 30
}

func asInt(v any) (int, bool) {
	n, ok := v.(int)
	return n, ok
}

func parseHex(s string) ([3]int, bool) {
	var zero [3]int
	s = strings.TrimPrefix(s, "#")
	if len(s) != 6 {
		return zero, false
	}
	n, err := strconv.ParseUint(s, 16, 24)
	if err != nil {
		return zero, false
	}
	return [3]int{int(n >> 16), int((n >> 8) & 255), int(n & 255)}, true
}
