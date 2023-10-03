package main

import "fmt"

// defer function adalah function yg bisa kita jadwalkan eksekusi setelah function dieksekusi
// walaupun error akan tetap dijalankan
func loging() {
	fmt.Println("selesai memanggil function")
}

// defer function
func appRunning() {
	defer loging() // ini adalah defer
	fmt.Println("aplikasi telah berjalan")
}

// panic funtion bisa digunakan untuk menghentikan program
// panic fucntion biasa dipanggil ketika error
// saat panic function dipanggil dan program berhenti, defer akan tetap akan dieksekusi
func endApplication() {
	message := recover()
	if message != nil {
		fmt.Println("ERRO coeg aplikasi nya", message)
	}
	fmt.Println("aplikasi selesai")

}

// panic function
func appRun(error bool) {
	defer endApplication()
	if error {
		panic("aplikasi error")
	}
	fmt.Println("aplikasi berjalan")

}

// recover adalah function yang brfungsi menangkap data panic
// dengan recover data panci akan terhenti, aplikasi tetap berjalan

func main() {
	// defer
	//appRunning()

	// panic
	appRun(false)
	fmt.Println("eko")
}
