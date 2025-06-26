package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseEvaluation ดึง "การประเมินผลตลอดหลักสูตร" (Evaluation) จากไฟล์
func ParseEvaluation(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "2.4 การประเมินผลตลอดหลักสูตร") || output.Evaluation != "" {
		return i
	}

	// ไล่หาบรรทัดแรกที่ไม่ว่างหลังหัวข้อ และหยุดเมื่อเจอหัวข้อถัดไป หรือจบไฟล์
	for j := i + 1; j < len(lines); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		// หยุดถ้าเจอหัวข้อถัดไป (หัวข้อ 3)
		if strings.HasPrefix(content, "3. คำสำคัญสำหรับการสืบค้น (keyword) และคำอธิบายหลักสูตรอย่างย่อ  ") {
			break
		}
		output.Evaluation = content
		return j
	}
	return i
}
