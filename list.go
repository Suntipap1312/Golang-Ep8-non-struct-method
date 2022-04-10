package main

import (
	"fmt"
)

/// สร้าง list type ทำหน้าที่ในการเก็บ ข้อมูล []*game
// interface สร้างมาเพื่อใช้ต้องการใช้ฟังชั่นร่วมกัน (Ex/Print()) ในต่าง struct type
// interfacer ใช้ตรงๆไม่ได้ ต้องสร้าง var list []printer

type printer interface {
	print()
}

type list []printer

// *** method เหมือนใช้ได้ method ไม่เหมือนกันใน struct ใช้ไม่ได้
func (l list) print() { // func เสมือน method
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery")
		return
	}

	for _, it := range l {
		it.print()
	}
}

// ใช้ method ร่วมกัน โดยที่ไม่ต้องสร้างที่ struct อื่น
// เป็นการ extract struct
// only the games are discounted

// *** method ไม่เหมือนกันใช้ได้
func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64) // ใช้ method เฉพาะ
	}

	for _, it := range l {

		// g, ok := it.(discounter) //it.(interface{ discount(float64)})
		// if !ok {
		// 	continue
		// }
		// g.discount(ratio)
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}

	}
}
