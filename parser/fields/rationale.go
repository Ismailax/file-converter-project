package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseRationale ค้นหาและกำหนดค่า Rationale (หลักการและเหตุผล)
func ParseRationale(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "2.1 หลักการและเหตุผล") || output.Rationale != "" {
		return i
	}

	// ไล่หาบรรทัดถัดไปจนกว่าจะถึง 2.2 วัตถุประสงค์ หรือจบไฟล์
	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "2.2 วัตถุประสงค์"); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		output.Rationale = content
		return j
	}
	return i
}
