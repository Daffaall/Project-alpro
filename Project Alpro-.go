package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sampah struct {
	Jenis           string
	Jumlah          int
	DaurUlang       bool
	MetodeDaurUlang string
}

var dataSampah []Sampah

func bukanHurufSaja(s string) bool {
	for _, c := range s {
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return true 
		}
	}
	return len(s) == 0 
}

func tambahSampah(jenis string, jumlah int, daurUlang bool, metode string) error {
	if jenis == "" {
		return fmt.Errorf("jenis sampah tidak valid: tidak boleh kosong atau hanya angka")
	}
	if jumlah <= 0 {
		return fmt.Errorf("jumlah sampah harus lebih dari 0")
	}
	if daurUlang && metode == "" {
		return fmt.Errorf("metode daur ulang harus diisi jika sampah didaur ulang")
	}

	dataSampah = append(dataSampah, Sampah{jenis, jumlah, daurUlang, metode})
	return nil
}

func ubahSampah(index int, jenis string, jumlah int, daurUlang bool, metode string) error {
	if index < 0 || index >= len(dataSampah) {
		return fmt.Errorf("index tidak valid")
	}
	if jenis == "" {
		return fmt.Errorf("jenis sampah tidak boleh kosong")
	}
	if jumlah <= 0 {
		return fmt.Errorf("jumlah sampah harus lebih dari 0")
	}
	if daurUlang && metode == "" {
		return fmt.Errorf("metode daur ulang harus diisi jika sampah didaur ulang")
	}

	dataSampah[index] = Sampah{jenis, jumlah, daurUlang, metode}
	return nil
}

func hapusSampah(index int) error {
	if index < 0 || index >= len(dataSampah) {
		return fmt.Errorf("index tidak valid")
	}
	dataSampah = append(dataSampah[:index], dataSampah[index+1:]...)
	return nil
}

func toLower(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		result += string(c)
	}
	return result
}

func sequentialSearch(jenis string) int {
	jenis = toLower(jenis)
	for i, s := range dataSampah {
		if toLower(s.Jenis) == jenis {
			return i
		}
	}
	return -1
}

func urutkanDataSampah() {
	for i := 1; i < len(dataSampah); i++ {
		key := dataSampah[i]
		j := i - 1
		for j >= 0 && dataSampah[j].Jenis > key.Jenis {
			dataSampah[j+1] = dataSampah[j]
			j--
		}
		dataSampah[j+1] = key
	}
}

func binarySearch(jenis string) int {
	urutkanDataSampah()

	left, right := 0, len(dataSampah)-1
	for left <= right {
		mid := (left + right) / 2
		if dataSampah[mid].Jenis == jenis {
			return mid
		} else if dataSampah[mid].Jenis < jenis {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func tampilkanStatistik() error {
	if len(dataSampah) == 0 {
		return fmt.Errorf("tidak ada data sampah")
	}

	total := 0
	totalDaurUlang := 0
	for _, s := range dataSampah {
		total += s.Jumlah
		if s.DaurUlang {
			totalDaurUlang += s.Jumlah
		}
	}
	fmt.Println("\n📈 Total sampah:", total, "Kg")
	fmt.Println("♻  Total yang didaur ulang:", totalDaurUlang, "Kg")
	fmt.Println("♻  Persentase yang didaur ulang:", float64(totalDaurUlang)/float64(total)*100, "%")
	fmt.Println("♻  Persentase yang tidak didaur ulang:", float64(total-totalDaurUlang)/float64(total)*100, "%")
	return nil
}

func selectionSortByJumlah() {
	for i := 0; i < len(dataSampah); i++ {
		min := i
		for j := i + 1; j < len(dataSampah); j++ {
			if dataSampah[j].Jumlah < dataSampah[min].Jumlah {
				min = j
			}
		}
		dataSampah[i], dataSampah[min] = dataSampah[min], dataSampah[i]
	}
}

func insertionSortByJenis() {
	for i := 1; i < len(dataSampah); i++ {
		key := dataSampah[i]
		j := i - 1
		for j >= 0 && dataSampah[j].Jenis > key.Jenis {
			dataSampah[j+1] = dataSampah[j]
			j--
		}
		dataSampah[j+1] = key
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== APLIKASI PENGELOLAAN SAMPAH =====")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Ubah Data Sampah")
		fmt.Println("3. Hapus Data Sampah")
		fmt.Println("4. Tampilkan Statistik")
		fmt.Println("5. Cari Sampah (Sequential)")
		fmt.Println("6. Cari Sampah (Binary)")
		fmt.Println("7. Urutkan Berdasarkan Jumlah (Selection Sort)")
		fmt.Println("8. Urutkan Berdasarkan Jenis (Insertion Sort)")
		fmt.Println("9. Tampilkan Semua Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Print("Jenis sampah: ")
			scanner.Scan()
			jenis := scanner.Text()
            if bukanHurufSaja(jenis) || jenis == "" {
	            fmt.Println("❌ Jenis tidak valid.")
	            continue
            }
			fmt.Print("Jumlah (kg): ")
			scanner.Scan()
			jumlahStr := scanner.Text()
			jumlah, err := strconv.Atoi(jumlahStr)
			if err != nil || jumlah <= 0 {
				fmt.Println("❌ Jumlah harus angka positif")
				continue
			}
			fmt.Print("Didaur ulang? (y/n): ")
			scanner.Scan()
			du := strings.ToLower(scanner.Text()) == "y"
			metode := ""
			if du {
				fmt.Print("Metode daur ulang: ")
				scanner.Scan()
				metode = scanner.Text()
			}
			if err := tambahSampah(jenis, jumlah, du, metode); err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("\n✅ Data sampah berhasil ditambahkan.")
			}

		case "2":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			fmt.Print("Index data yang ingin diubah: ")
			scanner.Scan()
			idxStr := scanner.Text()
			idx, err := strconv.Atoi(idxStr)
			if err != nil || idx < 0 || idx >= len(dataSampah) {
				fmt.Println("❌ Index tidak valid")
				continue
			}
			fmt.Print("Jenis baru: ")
			scanner.Scan()
			jenis := scanner.Text()
			fmt.Print("Jumlah baru (kg): ")
			scanner.Scan()
			jumlahStr := scanner.Text()
			jumlah, err := strconv.Atoi(jumlahStr)
			if err != nil || jumlah <= 0 {
				fmt.Println("❌ Jumlah harus angka positif")
				continue
			}
			fmt.Print("Didaur ulang? (y/n): ")
			scanner.Scan()
			du := strings.ToLower(scanner.Text()) == "y"
			metode := ""
			if du {
				fmt.Print("Metode daur ulang: ")
				scanner.Scan()
				metode = scanner.Text()
			}
			if err := ubahSampah(idx, jenis, jumlah, du, metode); err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("\n✅ Data berhasil diubah.")
			}

		case "3":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			fmt.Print("Index data yang ingin dihapus: ")
			scanner.Scan()
			idxStr := scanner.Text()
			idx, err := strconv.Atoi(idxStr)
			if err != nil || idx < 0 || idx >= len(dataSampah) {
				fmt.Println("❌ Index tidak valid")
				continue
			}
			if err := hapusSampah(idx); err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("\n✅ Data berhasil dihapus.")
			}

		case "4":
			if err := tampilkanStatistik(); err != nil {
				fmt.Println("❌", err)
			}

		case "5":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			fmt.Print("Cari jenis (sequential): ")
			scanner.Scan()
			jenis := scanner.Text()
			if jenis == "" {
				fmt.Println("❌ Jenis sampah tidak boleh kosong")
				continue
			}
			index := sequentialSearch(jenis)
			if index != -1 {
				fmt.Println("Ditemukan di index:", index, dataSampah[index])
			} else {
				fmt.Println("Tidak ditemukan.")
			}

		case "6":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			fmt.Print("Cari jenis (binary): ")
			scanner.Scan()
			jenis := scanner.Text()
			if jenis == "" {
				fmt.Println("❌ Jenis sampah tidak boleh kosong")
				continue
			}
			index := binarySearch(jenis)
			if index != -1 {
				fmt.Println("Ditemukan di index:", index, dataSampah[index])
			} else {
				fmt.Println("Tidak ditemukan.")
			}

		case "7":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			selectionSortByJumlah()
			fmt.Println("\n✅ Diurutkan berdasarkan jumlah (selection sort).")

		case "8":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			insertionSortByJenis()
			fmt.Println("\n✅ Diurutkan berdasarkan jenis (insertion sort).")
			for i, s := range dataSampah {
				fmt.Printf("[%d] %s - %dkg - Daur ulang: %v - Metode: %s\n", i, s.Jenis, s.Jumlah, s.DaurUlang, s.MetodeDaurUlang)
			}

		case "9":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			for i, s := range dataSampah {
				fmt.Printf("\n[%d] %s - %dkg - Daur ulang: %v - Metode: %s\n", i, s.Jenis, s.Jumlah, s.DaurUlang, s.MetodeDaurUlang)
			}

		case "0":
			fmt.Println("Keluar aplikasi.")
			return

		case "alamak":
			fmt.Println("\nAplikasi ini adalah aplikasi pengelolaan sampah yang dapat digunakan untuk menambah, mengubah, menghapus, dan mencari data sampah. \nSelain itu, aplikasi ini juga dapat menampilkan statistik pengelolaan sampah dan mengurutkan data berdasarkan jumlah atau jenis sampah.")

		default:
			fmt.Println("\n❌ Pilihan tidak valid.")
		}
	}
}