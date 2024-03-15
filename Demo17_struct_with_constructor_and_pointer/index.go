package main

import "demo16/employee"

func main() {
	// initialize จะทำงานครั้งเดียวเพราะใช้ Pointer เช็ค Null pointer และทำ Singleton ในระบบ
	e := employee.Init(
		"Lek",
		"Somjit",
		30,
		20,
	)
	e.LeavesRemaining()

	// initialize จะไม่ทำงาน
	e = employee.Init(
		"Xsaxin",
		"Somjit",
		30,
		20,
	)

	e.LeavesRemaining()
}
