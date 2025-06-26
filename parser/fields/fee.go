package fields

import (
	"regexp"
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseFee ดึงค่าธรรมเนียมการอบรม (output.Fee)
func ParseFee(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "8. ค่าธรรมเนียมในการอบรม") {
		return i
	}

	for j := i + 1; j < len(lines) && !strings.HasPrefix(lines[j], "9. แหล่งที่มาของงบประมาณการเปิดหลักสูตร"); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		// ไม่มีการเก็บค่าธรรมเนียม
		if strings.Contains(content, "ไม่มีการเก็บค่าธรรมเนียม") {
			output.Fee = 0
			return j
		}
		// ถ้ามี "บาท" ให้หาเลขที่มาก่อน
		if strings.Contains(content, "บาท") {
			re := regexp.MustCompile(`([\d,]+)\s*บาท`)
			if m := re.FindStringSubmatch(content); len(m) >= 2 {
				fee := strings.ReplaceAll(m[1], ",", "")
				output.Fee = utils.Atoi(fee)
				return j
			}
		}
	}
	return i
}
