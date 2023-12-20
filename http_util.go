package cronet

func isTokenChar(c byte) bool {
	return !(c >= 0x7F || c <= 0x20 || c == '(' || c == ')' || c == '<' ||
		c == '>' || c == '@' || c == ',' || c == ';' || c == ':' ||
		c == '\\' || c == '"' || c == '/' || c == '[' || c == ']' ||
		c == '?' || c == '=' || c == '{' || c == '}')
}

func isToken(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, c := range str {
		if !isTokenChar(byte(c)) {
			return false
		}
	}
	return true
}

func IsValidHeaderName(value string) bool {
	return isToken(value)
}

func IsValidHeaderValue(value string) bool {
	if len(value) == 0 {
		return false
	}
	for _, c := range value {
		if c == '\x00' || c == '\r' || c == '\n' {
			return false
		}
	}
	return true
}
