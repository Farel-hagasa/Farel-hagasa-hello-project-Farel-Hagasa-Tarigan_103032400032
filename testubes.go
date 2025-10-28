package main

import "fmt"

type FoodItem struct {
	ID         int
	Name       string
	Quantity   int
	Unit       string
	ExpiryDate string
}

type UsageLog struct {
	ItemID   int
	ItemName string
	UsedQty  int
	Unit     string
	Purpose  string
}

const NMAX int = 100

type foodStock [NMAX]FoodItem
type usageLogs [NMAX]UsageLog

func inputFoodItems(stock *foodStock, count *int) {
	var jumlah int
	fmt.Print("Masukkan jumlah bahan makanan: ")
	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Printf("\n-- Bahan #%d --\n", i+1)
		fmt.Print("Nama bahan: ")
		fmt.Scanln(&stock[*count].Name)
		fmt.Print("Jumlah: ")
		fmt.Scanln(&stock[*count].Quantity)
		fmt.Print("Satuan (kg, liter, pcs, dll): ")
		fmt.Scanln(&stock[*count].Unit)

		fmt.Print("Tanggal Kadaluarsa (DD-MM-YYYY): ")
		fmt.Scanln(&stock[*count].ExpiryDate)

		stock[*count].ID = *count + 1
		*count++
	}
}

func bubbleSortFoodItems(stock *foodStock, count int) {
	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			if stock[j].ID > stock[j+1].ID {
				stock[j], stock[j+1] = stock[j+1], stock[j]
			}
		}
	}
}

func sequentialSearchFood(stock foodStock, count, id int) int {
	for i := 0; i < count; i++ {
		if stock[i].ID == id {
			return i
		}
	}
	return -1
}

func binarySearchFood(stock foodStock, count, id int) int {
	low := 0
	high := count - 1
	for low <= high {
		mid := (low + high) / 2
		if stock[mid].ID == id {
			return mid
		} else if stock[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSortByQuantity(stock *foodStock, count int) {
	for i := 0; i < count-1; i++ {
		maxIdx := i
		for j := i + 1; j < count; j++ {
			if stock[j].Quantity > stock[maxIdx].Quantity {
				maxIdx = j
			}
		}
		stock[i], stock[maxIdx] = stock[maxIdx], stock[i]
	}
}

func isEarlierDate(date1, date2 string) bool {
	d1 := parseDate(date1)
	d2 := parseDate(date2)

	if d1[2] != d2[2] {
		return d1[2] < d2[2]
	}
	if d1[1] != d2[1] {
		return d1[1] < d2[1]
	}
	return d1[0] < d2[0]
}

func parseDate(dateStr string) [3]int {
	var d, m, y int
	fmt.Sscanf(dateStr, "%d-%d-%d", &d, &m, &y)
	return [3]int{d, m, y}
}

func insertionSortByExpiryDate(stock *foodStock, count int) {
	for i := 1; i < count; i++ {
		key := stock[i]
		j := i - 1
		for j >= 0 && isEarlierDate(key.ExpiryDate, stock[j].ExpiryDate) {
			stock[j+1] = stock[j]
			j = j - 1
		}
		stock[j+1] = key
	}
}

func displayFoodItems(stock foodStock, count int) {
	if count == 0 {
		fmt.Println("Belum ada data bahan.")
		return
	}
	for i := 0; i < count; i++ {
		fmt.Printf("[%d] %s - %d %s (Kadaluarsa: %s)\n", stock[i].ID, stock[i].Name, stock[i].Quantity, stock[i].Unit, stock[i].ExpiryDate)
	}
}

func searchFoodItemMenu(stock *foodStock, count int) {
	if count == 0 {
		fmt.Println("Belum ada data bahan untuk dicari.")
		return
	}

	var searchChoice int
	fmt.Println("\nPilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (membutuhkan data terurut berdasarkan ID)")
	fmt.Print("Pilih (1-2): ")
	fmt.Scanln(&searchChoice)

	var idToSearch int
	fmt.Print("Masukkan ID bahan yang ingin dicari: ")
	fmt.Scanln(&idToSearch)

	var index int
	switch searchChoice {
	case 1:
		index = sequentialSearchFood(*stock, count, idToSearch)
	case 2:
		fmt.Println("Mengurutkan data untuk Binary Search (berdasarkan ID)...")
		bubbleSortFoodItems(stock, count)
		index = binarySearchFood(*stock, count, idToSearch)
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	if index != -1 {
		fmt.Println("\n--- Bahan Ditemukan ---")
		fmt.Printf("ID: %d\n", stock[index].ID)
		fmt.Printf("Nama: %s\n", stock[index].Name)
		fmt.Printf("Jumlah: %d %s\n", stock[index].Quantity, stock[index].Unit)
		fmt.Printf("Kadaluarsa: %s\n", stock[index].ExpiryDate)
	} else {
		fmt.Println("Bahan dengan ID tersebut tidak ditemukan.")
	}
}

func sortFoodItemMenu(stock *foodStock, count int) {
	if count == 0 {
		fmt.Println("Belum ada data bahan untuk diurutkan.")
		return
	}

	var sortChoice int
	fmt.Println("\nPilih metode pengurutan:")
	fmt.Println("1. Selection Sort (berdasarkan stok terbanyak)")
	fmt.Println("2. Insertion Sort (berdasarkan kadaluarsa terpendek)")
	fmt.Print("Pilih (1-2): ")
	fmt.Scanln(&sortChoice)

	switch sortChoice {
	case 1:
		selectionSortByQuantity(stock, count)
		fmt.Println("Data bahan makanan berhasil diurutkan berdasarkan stok terbanyak.")
	case 2:
		insertionSortByExpiryDate(stock, count)
		fmt.Println("Data bahan makanan berhasil diurutkan berdasarkan kadaluarsa terpendek.")
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}
	displayFoodItems(*stock, count)
}

func listFoodItemsMenu(stock *foodStock, count int) {
	if count == 0 {
		fmt.Println("Belum ada data bahan.")
		return
	}

	for {
		fmt.Println("\n--- Menu Lihat Daftar Bahan ---")
		fmt.Println("1. Tampilkan Semua Bahan")
		fmt.Println("2. Cari Bahan")
		fmt.Println("3. Urutkan Bahan")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih (0-3): ")

		var subChoice int
		fmt.Scanln(&subChoice)

		switch subChoice {
		case 1:
			displayFoodItems(*stock, count)
		case 2:
			searchFoodItemMenu(stock, count)
		case 3:
			sortFoodItemMenu(stock, count)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func updateFoodItem(stock *foodStock, count int) {
	displayFoodItems(*stock, count)
	var id int
	fmt.Print("Masukkan ID bahan yang ingin diubah: ")
	fmt.Scanln(&id)

	bubbleSortFoodItems(stock, count)
	index := binarySearchFood(*stock, count, id)

	if index != -1 {
		fmt.Print("Nama baru: ")
		fmt.Scanln(&stock[index].Name)
		fmt.Print("Jumlah baru: ")
		fmt.Scanln(&stock[index].Quantity)
		fmt.Print("Satuan baru: ")
		fmt.Scanln(&stock[index].Unit)
		fmt.Print("Tanggal Kadaluarsa baru (DD-MM-YYYY): ")
		fmt.Scanln(&stock[index].ExpiryDate)
		fmt.Println("Data bahan berhasil diubah.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func deleteFoodItem(stock *foodStock, count *int) {
	displayFoodItems(*stock, *count)
	var id int
	fmt.Print("Masukkan ID bahan yang ingin dihapus: ")
	fmt.Scanln(&id)

	bubbleSortFoodItems(stock, *count)
	index := binarySearchFood(*stock, *count, id)

	if index != -1 {
		for i := index; i < *count-1; i++ {
			stock[i] = stock[i+1]
			stock[i].ID = i + 1 
		}
		*count--
		fmt.Println("Bahan berhasil dihapus.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func useFoodItem(stock *foodStock, stockCount int, logs *usageLogs, logCount *int) {
	displayFoodItems(*stock, stockCount)
	var id int
	fmt.Print("Masukkan ID bahan yang digunakan: ")
	fmt.Scanln(&id)

	bubbleSortFoodItems(stock, stockCount)
	index := binarySearchFood(*stock, stockCount, id)

	if index == -1 {
		fmt.Println("Bahan tidak ditemukan.")
		return
	}

	var usedQty int
	var purpose string
	fmt.Print("Jumlah yang digunakan: ")
	fmt.Scanln(&usedQty)

	if usedQty > stock[index].Quantity {
		fmt.Println("Jumlah tidak mencukupi di stok.")
		return
	}

	fmt.Print("Untuk keperluan apa digunakan?: ")
	fmt.Scanln(&purpose)

	stock[index].Quantity -= usedQty
	logs[*logCount] = UsageLog{id, stock[index].Name, usedQty, stock[index].Unit, purpose}
	*logCount++
	fmt.Println("Penggunaan bahan tercatat.")
}

func showUsageLogs(logs usageLogs, count int) {
	fmt.Println("\n--- Riwayat Penggunaan Bahan ---")
	if count == 0 {
		fmt.Println("Belum ada penggunaan tercatat.")
		return
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s - %d %s untuk %s\n", i+1, logs[i].ItemName, logs[i].UsedQty, logs[i].Unit, logs[i].Purpose)
	}
}

func menu() {
	fmt.Println("\n=== Menu Manajemen Stok Bahan Makanan ===")
	fmt.Println("1. Tambah Bahan Makanan")
	fmt.Println("2. Lihat Daftar Bahan")
	fmt.Println("3. Gunakan Bahan")
	fmt.Println("4. Lihat Riwayat Penggunaan")
	fmt.Println("5. Ubah Data Bahan")
	fmt.Println("6. Hapus Bahan")
	fmt.Println("7. Keluar")
	fmt.Print("Pilih (1-7): ")
}

func main() {
	var stock foodStock
	var logs usageLogs
	var stockCount, logCount int

	for {
		menu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			inputFoodItems(&stock, &stockCount)
		case 2:
			listFoodItemsMenu(&stock, stockCount)
		case 3:
			useFoodItem(&stock, stockCount, &logs, &logCount)
		case 4:
			showUsageLogs(logs, logCount)
		case 5:
			updateFoodItem(&stock, stockCount)
		case 6:
			deleteFoodItem(&stock, &stockCount)
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi manajemen stok.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}