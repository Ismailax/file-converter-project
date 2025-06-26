package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"word2json/parser/fields"
	"word2json/types"

	"code.sajari.com/docconv"
)

// Output, Contact struct ประกาศใน types.go

func ParseDocx(docPath string, outputPath string) {
	// 1. เปิดไฟล์ DOCX
	file, err := os.Open(docPath)
	if err != nil {
		log.Fatalf("❌ Failed to open file: %v", err)
	}
	defer file.Close()

	// 2. แปลงเป็นข้อความ plain text
	result, _, err := docconv.ConvertDocx(file)
	if err != nil {
		log.Fatalf("❌ Failed to convert docx: %v", err)
	}

	// 3. ตัด text เป็นแต่ละบรรทัด
	lines := strings.Split(result, "\n")

	// 4. เตรียม struct สำหรับ output
	var output types.Output

	// 5. Loop ทีละบรรทัด พร้อมระบุตำแหน่ง i เพื่อส่งต่อให้แต่ละ field
	for i := 0; i < len(lines); i++ {
		// แต่ละฟังก์ชันคืนค่า index ใหม่ (กรณี multi-line)
		i = fields.ParseTitle(lines, i, &output)
		i = fields.ParseOrganizedBy(lines, i, &output)
		i = fields.ParseEnrollLimit(lines, i, &output)
		i = fields.ParseTarget(lines, i, &output)
		i = fields.ParseRationale(lines, i, &output)
		i = fields.ParseObjective(lines, i, &output)
		i = fields.ParseContent(lines, i, &output)
		i = fields.ParseEvaluation(lines, i, &output)
		i = fields.ParseKeywordsOverview(lines, i, &output)
		i = fields.ParseEnrollPeriod(lines, i, &output)
		i = fields.ParsePayment(lines, i, &output)
		i = fields.ParseFee(lines, i, &output)
		i = fields.ParseContacts(lines, i, &output)
		fields.ParseCategories(docPath, &output)
	}

	// 6. Marshal เป็น JSON และเขียนลงไฟล์
	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("❌ Failed to encode JSON: %v", err)
	}
	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		log.Fatalf("❌ Failed to write file: %v", err)
	}
	fmt.Println("✅ JSON saved to", outputPath)
}
