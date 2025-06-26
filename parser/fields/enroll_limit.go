package fields

import (
	"regexp"
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseEnrollLimit ค้นหาและกำหนดค่า EnrollLimit (จำนวนรับสมัคร)
func ParseEnrollLimit(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "1.4 จำนวนรับสมัคร") || output.EnrollLimit != 0 {
		return i
	}

	// ไล่หาบรรทัดถัดไปจนกว่าจะถึง 1.5 หรือจบไฟล์
	for j := i + 1; j < len(lines) && !strings.Contains(lines[j], "1.5 กลุ่มเป้าหมาย"); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		if strings.Contains(content, "ไม่จำกัด") {
			output.EnrollLimit = 999999999
		} else {
			nums := regexp.MustCompile(`\d+`).FindAllString(content, -1)
			if len(nums) > 0 {
				output.EnrollLimit = utils.Atoi(nums[0])
			}
		}
		return j
	}
	return i
}
