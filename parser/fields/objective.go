package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseObjective ค้นหาและกำหนดค่า Objective (วัตถุประสงค์)
func ParseObjective(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "2.2 วัตถุประสงค์") || output.Objective != "" {
		return i
	}

	// ไล่หาบรรทัดถัดไปจนกว่าจะถึง 2.3 โครงสร้างหรือเนื้อหาของหลักสูตร หรือจบไฟล์
	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "2.3 โครงสร้างหรือเนื้อหาของหลักสูตร"); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		output.Objective = content
		return j
	}
	return i
}
