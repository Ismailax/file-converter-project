package utils

import (
	"regexp"
	"strings"
	"word2json/types"
)

// ParseContactBlock : แปลง block ข้อความเป็น []Contact
func ParseContactBlock(lines []string) []types.Contact {
	var contacts []types.Contact
	var current types.Contact
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		switch {
		case strings.HasPrefix(line, "ชื่อ-สกุล"):
			if current.Name != "" || current.Surname != "" {
				contacts = append(contacts, current)
				current = types.Contact{}
			}
			if i+1 < len(lines) {
				nameLine := strings.TrimSpace(lines[i+1])
				current.Prefix, current.Name, current.Surname = SplitContactFullName(nameLine)
				i++
			}
		case strings.HasPrefix(line, "เบอร์โทร"):
			if i+1 < len(lines) {
				phone := strings.TrimSpace(lines[i+1])
				current.Phone = strings.ReplaceAll(phone, "-", "")
				i++
			}
		case strings.HasPrefix(line, "อีเมล"):
			if i+1 < len(lines) {
				emails := ExtractContactEmails(lines[i+1])
				if len(emails) > 0 {
					current.Email = SelectContactEmail(emails)
				}
				i++
			}
		}
	}
	if current.Name != "" || current.Surname != "" {
		contacts = append(contacts, current)
	}
	return contacts
}

// SelectContactEmail : เลือกอีเมลที่เหมาะสมที่สุด (เช่น @cmu.ac.th)
func SelectContactEmail(emails []string) string {
	for _, e := range emails {
		if strings.HasSuffix(strings.ToLower(e), "@cmu.ac.th") {
			return e
		}
	}
	if len(emails) > 0 {
		return emails[0]
	}
	return ""
}

// SplitContactFullName : แยกคำนำหน้า ชื่อ นามสกุล
func SplitContactFullName(full string) (prefix, name, surname string) {
	full = strings.TrimSpace(full)
	parts := strings.Fields(full)
	prefixes := []string{
		"รศ.พญ.", "รศ.นพ.", "ผศ.ดร.", "ผศ.พญ.", "ผศ.นพ.",
		"ศ.ดร.", "ศ.", "ดร.", "รศ.", "ผศ.",
		"พญ.", "นพ.",
		"คุณ", "นาย", "นางสาว", "นาง", "น.ส.",
	}
	for _, p := range prefixes {
		if strings.HasPrefix(full, p) {
			remain := strings.TrimPrefix(full, p)
			remain = strings.TrimSpace(remain)
			subparts := strings.Fields(remain)
			if len(subparts) >= 2 {
				return p, subparts[0], strings.Join(subparts[1:], " ")
			} else if len(subparts) == 1 {
				return p, subparts[0], ""
			}
		}
	}
	if len(parts) >= 3 && stringInSlice(parts[0], prefixes) {
		return parts[0], parts[1], strings.Join(parts[2:], " ")
	} else if len(parts) == 2 && stringInSlice(parts[0], prefixes) {
		return parts[0], parts[1], ""
	} else if len(parts) == 2 {
		return "", parts[0], parts[1]
	} else if len(parts) == 1 {
		return "", parts[0], ""
	}
	return "", "", ""
}

func stringInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// ExtractContactEmails : ดึง email ทั้งหมดจาก string
func ExtractContactEmails(s string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	seps := []string{" และ ", ",", " "}
	for _, sep := range seps {
		parts := strings.Split(s, sep)
		var emails []string
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if emailRegex.MatchString(part) {
				emails = append(emails, part)
			}
		}
		if len(emails) > 0 {
			return emails
		}
	}
	return emailRegex.FindAllString(s, -1)
}
