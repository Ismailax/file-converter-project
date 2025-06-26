package fields

import (
	"strings"
	"word2json/types"
)

// ParseOrganizedBy ค้นหาและกำหนดค่า OrganizedBy (หน่วยงานที่จัด)
func ParseOrganizedBy(lines []string, i int, output *types.Output) int {
	clean := strings.TrimSpace(lines[i])
	if !strings.Contains(clean, "1.2 ดำเนินการโดย") || output.OrganizedBy != "" {
		return i
	}

	// ไล่หาบรรทัดที่ไม่ว่างถัดไปจนกว่าจะเจอ "1.3 ผู้รับผิดชอบหลักสูตร" หรือจบไฟล์
	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "1.3 ผู้รับผิดชอบหลักสูตร"); j++ {
		content := strings.TrimSpace(lines[j])
		if content == "" {
			continue
		}
		output.OrganizedBy = content
		return j
	}
	return i
}
