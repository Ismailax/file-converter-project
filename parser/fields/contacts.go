package fields

import (
	"strings"
	"word2json/types"
	"word2json/utils"
)

// ParseContacts ดึงข้อมูลผู้ประสานงานหลักสูตรทั้งหมด
func ParseContacts(lines []string, i int, output *types.Output) int {
	clean := strings.TrimSpace(lines[i])

	if !(strings.Contains(clean, "ผู้ประสานงานหลักสูตร") || strings.Contains(clean, "ติดต่อสอบถาม") || strings.Contains(clean, "ข้อมูลในการติดต่อสอบถาม")) {
		return i
	}

	var contactBlock []string
	for j := i + 1; j < len(lines); j++ {
		c := strings.TrimSpace(lines[j])
		// เจอหัวข้อถัดไป ให้หยุด
		if strings.HasPrefix(c, "11.") || strings.HasPrefix(c, "11") {
			output.Contacts = utils.ParseContactBlock(contactBlock)
			return j - 1
		}
		contactBlock = append(contactBlock, c)
	}
	if len(contactBlock) > 0 {
		output.Contacts = utils.ParseContactBlock(contactBlock)
	}
	return len(lines) - 1
}
