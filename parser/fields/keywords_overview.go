package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseKeywordsOverview ดึง "คำสำคัญสำหรับการสืบค้น (keyword)" และ "คำอธิบายหลักสูตรอย่างย่อ"
func ParseKeywordsOverview(lines []string, i int, output *types.Output) int {
	clean := utils.CleanText(lines[i])
	if !strings.Contains(clean, "3. คำสำคัญสำหรับการสืบค้น (keyword) และคำอธิบายหลักสูตรอย่างย่อ") {
		return i
	}

	var keywords []string
	var overviewLines []string
	var inKeyword, inOverview bool

	for j := i + 1; j < len(lines); j++ {
		content := utils.CleanText(lines[j])
		if content == "" {
			continue
		}
		// เจอหัวข้อถัดไป ให้หยุด
		if strings.HasPrefix(content, "4.") || strings.Contains(content, "4. ช่วงวัน-เวลาของการรับสมัคร") {
			break
		}
		// เริ่มอ่าน keyword
		if strings.Contains(content, "คำสำคัญสำหรับการสืบค้น (keyword)") || strings.Contains(content, "keyword") {
			inKeyword = true
			inOverview = false
			continue
		}
		// เริ่มอ่าน overview
		if strings.Contains(content, "คำอธิบายหลักสูตรอย่างย่อ") {
			inOverview = true
			inKeyword = false
			continue
		}
		// ดึง keyword
		if inKeyword {
			kws := strings.Split(content, ",")
			for _, kw := range kws {
				cleaned := utils.CleanText(kw)
				if cleaned != "" {
					keywords = append(keywords, cleaned)
				}
			}
			inKeyword = false // เอาแค่บรรทัดเดียว
			continue
		}
		// ดึง overview (รองรับหลายบรรทัด)
		if inOverview {
			overviewLines = append(overviewLines, content)
			continue
		}
	}

	output.Keywords = keywords
	output.Overview = strings.Join(overviewLines, " ")
	return i // ฟังก์ชันนี้ไม่จำเป็นต้องข้าม index
}
