package main

// Non-struct method สร้าง type ขึ้นมาอีกหน้าเช่น type money float64, type list []game
// interface สร้างมาเพื่อใช้ต้องการใช้ฟังชั่นร่วมกัน (Ex/Print()) ในต่าง struct type
// interface ไปใส่ไว้ใน list
// interface คือ list og structs
// interface is group of structs that shared common methods.

func main() {
	var (
		mobydick = book{title: "Moby Dick", price: 10}
		dota     = game{title: "Dota2", price: 20}
		tetris   = game{title: "Tetris", price: 10}
		rubik    = puzzle{title: "Rubik's cube", price: 5}
		yoda     = toy{title: "Yoda", price: 150}
	)

	///////  Non-struct Methos ///////n,m/
	//**** สร้าง slice หรือ list ขึ้นมาเพื่อเก็บข้อกลุ่มของ struct
	// var p printer
	var store list
	store = append(store, &dota, &tetris, mobydick, rubik, &yoda)
	store.discount(.5)
	store.print()

	///// สร้าง method อีกหน้า สำหรับใช้งานฟังชั่นในส่วนของ slice

}

//**** ถ้าใช้ method เดี่ยวๆ ให้ใช้ struct แต่ถ้าใช้ method แบบรวม struct ให้ไปเปลี่ยนที่ list []
