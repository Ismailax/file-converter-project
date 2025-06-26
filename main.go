package main

import (
	"fmt"
	"log"
	"os"

	"word2json/parser"

	"code.sajari.com/docconv"
)

// ฟังก์ชันแปลง docx เป็นไฟล์ .txt (plain text)
func ConvertDocxToTxt(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	result, _, err := docconv.ConvertDocx(file)
	if err != nil {
		return fmt.Errorf("failed to convert docx: %v", err)
	}
	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("failed to write txt: %v", err)
	}
	return nil
}

func main() {
	inputFile := "docs/04 การดูแลสุขภาพสัตว์เลี้ยง-28112024.docx"
	outputFile := "output/text_output.json"
	textOutput := "output/text_output.txt"

	// 1. แปลง docx เป็น text (.txt)

	if err := ConvertDocxToTxt(inputFile, textOutput); err != nil {
		log.Fatalf("❌ %v", err)
	}
	fmt.Println("✅ แปลง docx เป็น text แล้ว:", textOutput)

	// 2. แปลง docx เป็น json ตาม logic เดิม
	parser.ParseDocx(inputFile, outputFile)
}
