package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sampah struct {
	Jenis         string
	Jumlah        int
	DaurUlang     bool
	MetodeDaurUlang string
}

var dataSampah []Sampah

func tambahSampah(jenis string, jumlah int, daurUlang bool, metode string) {
	dataSampah = append(dataSampah, Sampah{jenis, jumlah, daurUlang, metode})
}

func ubahSampah(index int, jenis string, jumlah int, daurUlang bool, metode string) {
	dataSampah[index] = Sampah{jenis, jumlah, daurUlang, metode}
}

func hapusSampah(index int) {
	dataSampah = append(dataSampah[:index], dataSampah[index+1:]...)
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

func displayDataHeader() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ %-5s ║ %-15s ║ %-10s ║ %-11s ║ %-17s ║\n", "Index", "Jenis", "Jumlah", "Daur Ulang", "Metode Daur Ulang")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")
}

func displayDataRow(index int, s Sampah) {
	metode := s.MetodeDaurUlang
	if metode == "" {
		metode = "-"
	}
	fmt.Printf("║ %-5d ║ %-15s ║ %-8dkg ║ %-11v ║ %-17s ║\n", index, toLower(s.Jenis), s.Jumlah, s.DaurUlang, toLower(metode))
}

func displayDataFooter() {
	fmt.Println("╚════════════════════════════════════════════════════════════════════════╝")
}

func tampilkanSemuaData() {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}

	fmt.Println("\n📋 DAFTAR SAMPAH")
	displayDataHeader()
	for i, s := range dataSampah {
		displayDataRow(i, s)
	}
	displayDataFooter()
}

func sequentialSearch(jenis string) {
	ditemukan := false
	jenis = toLower(jenis)
	var hasilPencarian []Sampah

	for _, s := range dataSampah {
		if toLower(s.Jenis) == jenis {
			ditemukan = true
			hasilPencarian = append(hasilPencarian, s)
		}
	}

	if ditemukan {
		fmt.Println("\n✅ Data ditemukan:")
		displayDataHeader()
		for i, s := range hasilPencarian {
			displayDataRow(i, s)
		}
		displayDataFooter()
	} else {
		fmt.Println("❌ Data tidak ditemukan.")
	}
}

func insertionSortByJenis() {
	for i := 1; i < len(dataSampah); i++ {
		key := dataSampah[i]
		j := i - 1
		for j >= 0 && toLower(dataSampah[j].Jenis) > toLower(key.Jenis) {
			dataSampah[j+1] = dataSampah[j]
			j--
		}
		dataSampah[j+1] = key
	}
}

func binarySearch(jenis string) {
	insertionSortByJenis()

	jenis = toLower(jenis)
	indeksAwal, indeksAkhir := 0, len(dataSampah)-1
	ditemukan := -1

	for indeksAwal <= indeksAkhir {
		indeksTengah := (indeksAwal + indeksAkhir) / 2
		if toLower(dataSampah[indeksTengah].Jenis) == jenis {
			ditemukan = indeksTengah
			indeksAkhir = indeksTengah - 1
		} else if toLower(dataSampah[indeksTengah].Jenis) < jenis {
			indeksAwal = indeksTengah + 1
		} else {
			indeksAkhir = indeksTengah - 1
		}
	}

	if ditemukan == -1 {
		fmt.Println("❌ Data tidak ditemukan.")
		return
	}

	fmt.Println("\n✅ Data ditemukan:")
	displayDataHeader()
	for i := ditemukan; i < len(dataSampah); i++ {
		if toLower(dataSampah[i].Jenis) == jenis {
			displayDataRow(i, dataSampah[i])
		} else {
			break
		}
	}
	displayDataFooter()
}

func tampilkanStatistik() { 
    total := 0
    totalDaurUlang := 0
    for _, s := range dataSampah {
        total += s.Jumlah
        if s.DaurUlang {
            totalDaurUlang += s.Jumlah
        }
    }

    fmt.Println("\n📈 STATISTIK SAMPAH")
    fmt.Println("╔═════════════════════════════════════════════╗")
    fmt.Printf("║ %-30s ║ %-8dkg ║\n", "Total sampah yang terkumpul", total)
    fmt.Printf("║ %-30s ║ %-8dkg ║\n", "Total sampah yang didaur ulang", totalDaurUlang)
    persentaseDaurUlang := 0.0
    persentaseTidakDaurUlang := 0.0
    if total > 0 {
        persentaseDaurUlang = float64(totalDaurUlang) / float64(total) * 100
        persentaseTidakDaurUlang = float64(total-totalDaurUlang) / float64(total) * 100
    }

    fmt.Printf("║ %-30s ║ %-9.2f%% ║\n", "Persentase didaur ulang", persentaseDaurUlang)
    fmt.Printf("║ %-30s ║ %-9.2f%% ║\n", "Persentase tidak didaur ulang", persentaseTidakDaurUlang)
    fmt.Println("╚═════════════════════════════════════════════╝")
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

func handleTambahSampah(scanner *bufio.Scanner) {
	fmt.Print("Masukkan jenis sampah (Organik, Anorganik, atau B3): ")
	scanner.Scan()
	jenisSampah := scanner.Text()
	if (toLower(jenisSampah) != "organik" && toLower(jenisSampah) != "anorganik" && toLower(jenisSampah) != "b3") || toLower(jenisSampah) == "" {
		fmt.Println("❌ Jenis sampah tidak valid.")
		return
	}

	fmt.Print("Masukkan jumlah sampah (kg): ")
	scanner.Scan()
	jumlahStr := scanner.Text()
	jumlahSampah, err := strconv.Atoi(jumlahStr)
	if err != nil || jumlahSampah <= 0 {
		fmt.Println("❌ Jumlah sampah tidak valid")
		return
	}

	fmt.Print("Apakah sampah akan didaur ulang? (y/n): ")
	scanner.Scan()
	daurUlangInput := strings.ToLower(scanner.Text())
	daurUlang := true
	metodeDaurUlang := ""
	if daurUlangInput == "y" {
		fmt.Print("Masukkan metode daur ulang sampah (A, B, atau C): ")
		scanner.Scan()
		metodeDaurUlang = scanner.Text()
		if (toLower(metodeDaurUlang) != "a" && toLower(metodeDaurUlang) != "b" && toLower(metodeDaurUlang) != "c") || toLower(metodeDaurUlang) == "" {
			fmt.Println("❌ Metode daur ulang tidak valid")
			return
		}

	} else if daurUlangInput == "n" {
        daurUlang = false
    } else {
        fmt.Println("❌ Input tidak valid")
        return
    }

    tambahSampah(jenisSampah, jumlahSampah, daurUlang, metodeDaurUlang)
    fmt.Println("\n✅ Data sampah berhasil ditambahkan.")
}

func handleUbahSampah(scanner *bufio.Scanner) {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}
	tampilkanSemuaData()

	fmt.Print("\nMasukkan Index data yang ingin diubah: ")
	scanner.Scan()
	idxStr := scanner.Text()
	idx, err := strconv.Atoi(idxStr)
	if err != nil || idx < 0 || idx >= len(dataSampah) {
		fmt.Println("❌ Index tidak valid")
		return
	}

	fmt.Print("Masukkan jenis sampah yang baru (Organik, Anorganik, atau B3): ")
	scanner.Scan()
	jenisSampahBaru := scanner.Text()
	if (toLower(jenisSampahBaru) != "organik" && toLower(jenisSampahBaru) != "anorganik" && toLower(jenisSampahBaru) != "b3") || 
		toLower(jenisSampahBaru) == "" {
		fmt.Println("❌ Jenis sampah yang baru tidak valid")
		return
	}

	fmt.Print("Masukkan jumlah sampah yang baru (kg): ")
	scanner.Scan()
	jumlahStr := scanner.Text()
	jumlahSampahBaru, err := strconv.Atoi(jumlahStr)
	if err != nil || jumlahSampahBaru <= 0 {
		fmt.Println("❌ Jumlah sampah yang baru tidak valid, harus lebih dari 0")
		return
	}

	fmt.Print("Apakah sampah akan didaur ulang? (y/n): ")
	scanner.Scan()
	daurUlangBaruInput := strings.ToLower(scanner.Text())
	daurUlangBaru := true
	metodeBaru := ""
	if daurUlangBaruInput == "y" {
		fmt.Print("Masukkan metode daur ulang yang baru (A, B, atau C): ")
		scanner.Scan()
		metodeBaru = scanner.Text()
		if (toLower(metodeBaru) != "a" && toLower(metodeBaru) != "b" && toLower(metodeBaru) != "c") || toLower(metodeBaru) == "" {
			fmt.Println("❌ Metode daur ulang yang baru tidak valid")
			return
		}
	} else if daurUlangBaruInput == "n" {
        daurUlangBaru = false
    } else {
        fmt.Println("❌ Input tidak valid") 
        return
    }

    ubahSampah(idx, jenisSampahBaru, jumlahSampahBaru, daurUlangBaru, metodeBaru)
    fmt.Println("\n✅ Data sampah berhasil diubah.")
}

func handleHapusSampah(scanner *bufio.Scanner) {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}
	tampilkanSemuaData()

	fmt.Print("\nMasukkan index data yang ingin dihapus: ")
	scanner.Scan()
	idxStr := scanner.Text()
	idx, err := strconv.Atoi(idxStr)
	if err != nil || idx < 0 || idx >= len(dataSampah) {
		fmt.Println("❌ Index tidak valid")
		return
	}

	hapusSampah(idx)
	fmt.Println("\n✅ Data sampah berhasil dihapus.")
}

func handleTampilkanStatistik() {
    if len(dataSampah) == 0 {
        fmt.Println("❌ Tidak ada statistik data sampah untuk ditampilkan")
        return
    }

    tampilkanStatistik()
}

func handleSequentialSearch(scanner *bufio.Scanner) {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}

	fmt.Print("Cari data dari jenis sampah (Organik, Anorganik, atau B3) (sequential): ")
	scanner.Scan()
	jenisSampah := scanner.Text()
	if (toLower(jenisSampah) != "organik" && toLower(jenisSampah) != "anorganik" && toLower(jenisSampah) != "b3") ||
		toLower(jenisSampah) == "" {
		fmt.Println("❌ Jenis sampah tidak valid")
		return
	}

	sequentialSearch(jenisSampah)
}

func handleBinarySearch(scanner *bufio.Scanner) {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}

	fmt.Print("Cari data dari jenis sampah (Organik, Anorganik, atau B3) (binary): ")
	scanner.Scan()
	jenisSampah := scanner.Text()
	if (toLower(jenisSampah) != "organik" && toLower(jenisSampah) != "anorganik" && toLower(jenisSampah) != "b3") || 
		toLower(jenisSampah) == "" {
		fmt.Println("❌ Jenis sampah tidak valid")
		return
	}

	binarySearch(jenisSampah)
}

func handleUrutkanByJumlah() {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}
	selectionSortByJumlah()
    tampilkanSemuaData()
	fmt.Println("\n✅ Data sampah diurutkan berdasarkan jumlah sampah (selection sort).")
}

func handleUrutkanByJenis() {
	if len(dataSampah) == 0 {
		fmt.Println("❌ Tidak ada data sampah")
		return
	}
	insertionSortByJenis()
    tampilkanSemuaData()
	fmt.Println("\n✅ Data sampah diurutkan berdasarkan jenis (insertion sort).")
}

func handleTampilkanSemuaData() {
	tampilkanSemuaData()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Data Dummy
	dataSampah = append(dataSampah,
		Sampah{"Anorganik", 10, true, "a"},
		Sampah{"Organik", 5, true, "b"},
		Sampah{"B3", 8, false, ""},
		Sampah{"Organik", 15, true, "c"},
		Sampah{"Anorganik", 6, false, ""},
	)

	for {
		fmt.Println("\n╔═════════════════════════════════════╗")
		fmt.Println("║     APLIKASI PENGELOLAAN SAMPAH     ║")
		fmt.Println("╠═════════════════════════════════════╣")
		fmt.Println("║ 1. Tambah Data Sampah               ║")
		fmt.Println("║ 2. Ubah Data Sampah                 ║")
		fmt.Println("║ 3. Hapus Data Sampah                ║")
		fmt.Println("║ 4. Tampilkan Statistik              ║")
		fmt.Println("║ 5. Cari Sampah (Sequential)         ║")
		fmt.Println("║ 6. Cari Sampah (Binary)             ║")
		fmt.Println("║ 7. Urutkan Berdasarkan Jumlah       ║")
		fmt.Println("║ 8. Urutkan Berdasarkan Jenis        ║")
		fmt.Println("║ 9. Tampilkan Semua Data             ║")
		fmt.Println("║ 0. Keluar                           ║")
		fmt.Println("╚═════════════════════════════════════╝")
		fmt.Print("Pilih menu (0-9): ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			handleTambahSampah(scanner)
		case "2":
			handleUbahSampah(scanner)
		case "3":
			handleHapusSampah(scanner)
		case "4":
			handleTampilkanStatistik()
		case "5":
			handleSequentialSearch(scanner)
		case "6":
			handleBinarySearch(scanner)
		case "7":
			handleUrutkanByJumlah()
		case "8":
			handleUrutkanByJenis()
		case "9":
			handleTampilkanSemuaData()
		case "0":
			fmt.Println("Keluar aplikasi.")
			return
		case "info":
			fmt.Println("\nAplikasi ini adalah aplikasi pengelolaan sampah yang dapat digunakan untuk menambah, mengubah, menghapus, dan mencari data sampah. \nSelain itu, aplikasi ini juga dapat menampilkan statistik pengelolaan sampah dan mengurutkan data berdasarkan jumlah atau jenis sampah.")
		default:
			fmt.Println("\n❌ Pilihan tidak valid.")
		}
	}
}