package fields

import (
	"strings"
	"word2json/types"
)

func ParseTitle(lines []string, i int, output *types.Output) int {
	clean := strings.TrimSpace(lines[i])
	if !strings.Contains(clean, "1.1 ชื่อหลักสูตร") || output.TitleTH != "" {
		return i
	}

	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "1.2 ดำเนินการโดย"); j++ {
		th := strings.TrimSpace(lines[j])
		if th == "" {
			continue
		}
		output.TitleTH = th

		if j+1 < len(lines) {
			enCandidate := strings.TrimSpace(lines[j+1])
			if strings.HasPrefix(enCandidate, "(") && strings.HasSuffix(enCandidate, ")") {
				output.TitleEN = strings.Trim(enCandidate, "()")
			}
		}
		return j + 1
	}
	return i
}
