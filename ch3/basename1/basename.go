// basenameはディレクトリ要素、接尾辞を取り除きます
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package basename

func basename(s string) string {
	// 最後の’/’とその前の全てを破棄する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 最後の'.'より前の全てを保持する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
