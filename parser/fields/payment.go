package fields

import (
	"regexp"
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParsePayment ดึงวันสิ้นสุดชำระเงินค่าธรรมเนียม (output.PaymentDeadline)
func ParsePayment(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "5. ช่วงวัน-เวลาของการชำระค่าธรรมเนียม") {
		return i
	}

	for j := i + 1; j < len(lines) && !strings.HasPrefix(lines[j], "6."); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		// กรณีไม่มีค่าธรรมเนียม
		if strings.Contains(content, "ไม่มีการเก็บค่าธรรมเนียม") {
			output.PaymentDeadline = nil
			return j
		}
		// กรณีมีระบุช่วงวันที่
		if strings.Contains(content, "ตั้งแต่วันที่") {
			// ตัวอย่าง: "ตั้งแต่วันที่ทำการสมัคร ถึง วันที่ 27 มิถุนายน 2568 เวลา 16.30 น."
			pattern := `ถึง วันที่ ([\d ]+[\p{Thai}]+ \d{4}) เวลา ([\d\.]+) น\.`
			re := regexp.MustCompile(pattern)
			if matches := re.FindStringSubmatch(content); len(matches) == 3 {
				thaiDate := strings.TrimSpace(matches[1])
				thaiTime := strings.TrimSpace(matches[2])
				datetime := "วันที่ " + thaiDate + " เวลา " + thaiTime
				val := utils.ParseThaiDateTime(datetime)
				output.PaymentDeadline = &val
			}
			return j
		}
	}
	return i
}
