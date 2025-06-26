package types

type Output struct {
	// ข้อมูลหลักสูตร
	TitleTH         string   `json:"title_th"`         // ชื่อหลักสูตร (ภาษาไทย)
	TitleEN         string   `json:"title_en"`         // ชื่อหลักสูตร (ภาษาอังกฤษ)
	OrganizedBy     string   `json:"organized_by"`     // หน่วยงานที่จัด
	EnrollLimit     int      `json:"enroll_limit"`     // จำนวนรับสมัคร
	Target          string   `json:"target"`           // กลุ่มเป้าหมาย
	Rationale       string   `json:"rationale"`        // หลักการและเหตุผล
	Objective       string   `json:"objective"`        // วัตถุประสงค์
	Content         string   `json:"content"`          // โครงสร้างหรือเนื้อหาของหลักสูตร
	Evaluation      string   `json:"evaluation"`       // Course Evaluation
	Keywords        []string `json:"keywords"`         // คำสำคัญ
	Overview        string   `json:"overview"`         // คำอธิบายหลักสูตรอย่างย่อ
	StartEnroll     string   `json:"start_enroll"`     // วันเปิดรับสมัคร
	EndEnroll       string   `json:"end_enroll"`       // วันปิดรับสมัคร
	PaymentDeadline *string  `json:"payment_deadline"` // วันสิ้นสุดชำระเงิน (nullable)
	Fee             int      `json:"fee"`              // ค่าธรรมเนียม

	// ผู้ประสานงานหลักสูตร
	Contacts []Contact `json:"contacts"`

	// หมวดหมู่การเรียนรู้
	Categories []string `json:"categories"`
}

type Contact struct {
	Prefix  string `json:"prefix"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
