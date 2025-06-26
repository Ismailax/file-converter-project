package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseEnrollPeriod ดึงวันเวลาเปิด-ปิดรับสมัคร
func ParseEnrollPeriod(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "4. ช่วงวัน-เวลาของการรับสมัคร") {
		return i
	}

	var startDate, startTime, endDate, endTime string
	var parsingStart, parsingEnd bool

	for j := i + 1; j < len(lines); j++ {
		line := utils.CleanText(lines[j])
		if line == "" {
			continue
		}
		// ถ้าเจอหัวข้อถัดไป ให้หยุด
		if strings.Contains(line, "5. ช่วงวัน-เวลาของการชำระค่าธรรมเนียม") {
			i = j - 1 // ชี้ index ให้หยุดหลังบรรทัดนี้
			break
		}
		if strings.Contains(line, "เปิดรับสมัคร") {
			parsingStart = true
			parsingEnd = false
			continue
		}
		if strings.Contains(line, "ปิดรับสมัคร") {
			parsingStart = false
			parsingEnd = true
			continue
		}
		// Start enroll
		if parsingStart && startDate == "" && strings.Contains(line, "วันที่") {
			startDate = strings.TrimSpace(strings.TrimPrefix(line, "วันที่"))
			continue
		}
		if parsingStart && startTime == "" && strings.HasPrefix(line, "เวลา") {
			startTime = strings.TrimSpace(strings.TrimPrefix(line, "เวลา"))
			continue
		}
		// End enroll
		if parsingEnd && endDate == "" && strings.Contains(line, "วันที่") {
			endDate = strings.TrimSpace(strings.TrimPrefix(line, "วันที่"))
			continue
		}
		if parsingEnd && endTime == "" && strings.HasPrefix(line, "เวลา") {
			endTime = strings.TrimSpace(strings.TrimPrefix(line, "เวลา"))
			continue
		}
	}
	// รวมวันและเวลาเป็น string เดียว ส่งไป parser
	output.StartEnroll = utils.ParseThaiDateTime("วันที่ " + startDate + " เวลา " + startTime)
	output.EndEnroll = utils.ParseThaiDateTime("วันที่ " + endDate + " เวลา " + endTime)

	return i
}
