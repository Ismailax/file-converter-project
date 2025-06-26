package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseTarget ค้นหาและกำหนดค่า Target (กลุ่มเป้าหมาย)
func ParseTarget(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "1.5 กลุ่มเป้าหมาย") || output.Target != "" {
		return i
	}

	// ไล่หาบรรทัดถัดไปจนกว่าจะถึง 2. ข้อมูลเฉพาะของหลักสูตร หรือจบไฟล์
	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "2. ข้อมูลเฉพาะของหลักสูตร"); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		output.Target = content
		return j
	}
	return i
}
