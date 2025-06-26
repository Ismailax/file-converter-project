package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseContent ดึง "โครงสร้างหรือเนื้อหาของหลักสูตร" (Content) จากไฟล์
func ParseContent(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "2.3 โครงสร้างหรือเนื้อหาของหลักสูตร") || output.Content != "" {
		return i
	}

	// หากมีข้อมูลในบรรทัดถัดไป ให้ใช้บรรทัดแรกที่ไม่ว่างเป็นค่า Content
	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "2.4 การประเมินผลตลอดหลักสูตร"); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		output.Content = content
		return j
	}
	return i
}
