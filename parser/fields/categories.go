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

	for _, p := range paragraphs {
		// รวม text ทั้ง paragraph
		text := ""
		for _, r := range p.Runs {
			if r.Text != nil {
				text += strings.TrimSpace(r.Text.Text)
			}
		}
		// หาหัวข้อ "หมวดหมู่การเรียนรู้"
		if !foundHeader && strings.Contains(text, "หมวดหมู่การเรียนรู้") {
			foundHeader = true
			continue // เริ่มเก็บตั้งแต่ paragraph ถัดไป
		}

		// ถ้ายังไม่เจอหัวข้อ ก็ข้ามไป
		if !foundHeader {
			continue
		}

		// หา checkbox ที่ติ๊ก
		for i, r := range p.Runs {
			if r.Sym != nil && r.Sym.Font == "Wingdings" && strings.ToUpper(r.Sym.Char) == "F0FE" {
				if i+1 < len(p.Runs) && p.Runs[i+1].Text != nil {
					category := strings.TrimSpace(p.Runs[i+1].Text.Text)
					if category != "" {
						categories = append(categories, category)
					}
				}
			}
		}
	}
	return categories
}
