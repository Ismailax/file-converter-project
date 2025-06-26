package fields

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"log"
	"strings"

	"word2json/types"
)

// ParseCategories: เติม categories ลงใน output โดยอ่าน document.xml โดยตรง
func ParseCategories(docxPath string, output *types.Output) {
	// เปิด .docx (zip)
	r, err := zip.OpenReader(docxPath)
	if err != nil {
		log.Fatalf("❌ Failed to open DOCX(zip): %v", err)
	}
	defer r.Close()

	var xmlContent []byte
	for _, f := range r.File {
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				log.Fatalf("❌ Failed to open document.xml: %v", err)
			}
			defer rc.Close()
			xmlContent, err = io.ReadAll(rc)
			if err != nil {
				log.Fatalf("❌ Failed to read document.xml: %v", err)
			}
		}
	}
	if xmlContent == nil {
		log.Println("⚠️ No document.xml found")
		return
	}

	var doc types.Document // ****
	if err := xml.Unmarshal(xmlContent, &doc); err != nil {
		log.Fatalf("❌ Failed to parse document.xml: %v", err)
	}

	output.Categories = extractCheckedCategories(doc.Body.Paragraphs)
}

// -- Extract only checked categories (☑/Wingdings F0FE) --
func extractCheckedCategories(paragraphs []types.Paragraph) []string {
	var categories []string
	foundHeader := false
	foundAnyCheckbox := false
	var plainTextCandidates []string

	for _, p := range paragraphs {
		// รวม text ทั้ง paragraph
		text := ""
		for _, r := range p.Runs {
			if r.Text != nil {
				text += strings.TrimSpace(r.Text.Text)
			}
		}
		// หาหัวข้อ
		if !foundHeader && strings.Contains(text, "หมวดหมู่การเรียนรู้") {
			foundHeader = true
			continue
		}
		if !foundHeader {
			continue
		}

		// ------ (1) ปกติ: มี checkbox ------
		expectLabel := false
		hasCheckbox := false
		for _, r := range p.Runs {
			if r.Sym != nil && r.Sym.Font == "Wingdings" && strings.ToUpper(r.Sym.Char) == "F0FE" {
				expectLabel = true
				hasCheckbox = true
				foundAnyCheckbox = true
				continue
			}
			if expectLabel && r.Text != nil {
				label := strings.TrimSpace(r.Text.Text)
				if label != "" {
					categories = append(categories, label)
					expectLabel = false
				}
			}
		}

		// ------ (2) กรณีพิเศษ: ไม่มี checkbox แต่ paragraph มีข้อความ ------
		// (เก็บเป็น candidate เผื่อเป็น category เดียว)
		if !hasCheckbox && len(text) > 0 {
			plainTextCandidates = append(plainTextCandidates, text)
		}
		// ถ้า paragraph ว่างหลังเจอ header ให้หยุดเก็บ plain text (optional, กรณีมีหลายบรรทัด)
		if foundHeader && len(text) == 0 && len(plainTextCandidates) > 0 {
			break
		}
	}

	// ถ้าไม่เจอ checkbox เลย ให้คืน plain text ที่เก็บมาเป็น category
	if !foundAnyCheckbox && len(plainTextCandidates) > 0 {
		return plainTextCandidates
	}
	return categories
}
