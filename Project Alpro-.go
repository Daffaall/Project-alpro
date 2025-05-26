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

func displayDataHeader() {
	fmt.Println("==========================================================================")
	fmt.Printf("| %-5s | %-15s | %-10s | %-11s | %-17s |\n", "Index", "Jenis", "Jumlah", "Daur Ulang", "Metode Daur Ulang")
	fmt.Println("==========================================================================")
}

func displayDataRow(index int, s Sampah) {
	metode := s.MetodeDaurUlang
	if metode == "" {
		metode = "-"
	}
	fmt.Printf("| %-5d | %-15s | %-8dkg | %-11v | %-17s |\n", index, toLower(s.Jenis), s.Jumlah, s.DaurUlang, toLower(metode))
}

func displayDataFooter() {
	fmt.Println("==========================================================================")
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

func binarySearch(jenis string) {
	urutkanDataSampah()

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

	fmt.Println("\n📈 STATISTIK SAMPAH")
	fmt.Println("===============================================")
	fmt.Printf("| %-30s | %-8dkg |\n", "Total sampah yang terkumpul", total)
	fmt.Printf("| %-30s | %-8dkg |\n", "Total sampah yang didaur ulang", totalDaurUlang)
	fmt.Printf("| %-30s | %-9.2f%% |\n", "Persentase didaur ulang",
		float64(totalDaurUlang)/float64(total)*100)
	fmt.Printf("| %-30s | %-9.2f%% |\n", "Persentase tidak didaur ulang",
		float64(total-totalDaurUlang)/float64(total)*100)
	fmt.Println("===============================================")
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
		for j >= 0 && toLower(dataSampah[j].Jenis) > toLower(key.Jenis) {
			dataSampah[j+1] = dataSampah[j]
			j--
		}
		dataSampah[j+1] = key
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dataSampah = append(dataSampah,
		Sampah{"Anorganik", 10, true, "a"},
		Sampah{"Organik", 5, true, "b"},
		Sampah{"B3", 8, false, ""},
		Sampah{"Organik", 15, true, "c"},
		Sampah{"Anorganik", 6, false, ""},
	)

	for {
    	fmt.Println("\n╔═════════════════════════════════════╗")
    	fmt.Println("║    APLIKASI PENGELOLAAN SAMPAH      ║")
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
   		fmt.Print("  Pilih menu (0-9): ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Print("Masukkan jenis sampah (Organik, Anorganik, atau B3): ")
			scanner.Scan()
			jenisSampah := scanner.Text()
			if (toLower(jenisSampah) != "organik" && toLower(jenisSampah) != "anorganik" && toLower(jenisSampah) != "b3") || toLower(jenisSampah) == "" {
				fmt.Println("❌ Jenis sampah tidak valid.")
				continue
			}

			fmt.Print("Masukkan jumlah sampah (kg): ")
			scanner.Scan()
			jumlahStr := scanner.Text()
			jumlahSampah, err := strconv.Atoi(jumlahStr)
			if err != nil || jumlahSampah <= 0 {
				fmt.Println("❌ Jumlah sampah tidak valid")
				continue
			}

			fmt.Print("Apakah sampah akan didaur ulang? (y/n): ")
			scanner.Scan()
			daurUlang := strings.ToLower(scanner.Text())
			metodeDaurUlang := ""
			if daurUlang == "y" {
				fmt.Print("Masukkan metode daur ulang sampah (A, B, atau C): ")
				scanner.Scan()
				metodeDaurUlang = scanner.Text()
				if (toLower(metodeDaurUlang) != "a" && toLower(metodeDaurUlang) != "b" && toLower(metodeDaurUlang) != "c") || toLower(metodeDaurUlang) == "" {
					fmt.Println("❌ Metode daur ulang tidak valid")
					continue
				}
			} else if daurUlang != "y" && daurUlang != "n" {
				fmt.Println("❌ Input tidak valid")
				continue
			}

			if err := tambahSampah(jenisSampah, jumlahSampah, daurUlang == "y", metodeDaurUlang); err != nil {
				fmt.Println("\n❌ Proses tambah data sampah gagal", err)
			} else {
				fmt.Println("\n✅ Data sampah berhasil ditambahkan.")
			}

		case "2":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			tampilkanSemuaData()

			fmt.Print("\nMasukkan Index data yang ingin diubah: ")
			scanner.Scan()
			idxStr := scanner.Text()
			idx, err := strconv.Atoi(idxStr)
			if err != nil || idx < 0 || idx >= len(dataSampah) {
				fmt.Println("❌ Index tidak valid")
				continue
			}

			fmt.Print("Masukkan jenis sampah yang baru (Organik, Anorganik, atau B3): ")
			scanner.Scan()
			jenisSampahBaru := scanner.Text()
			if (toLower(jenisSampahBaru) != "organik" && toLower(jenisSampahBaru) != "anorganik" && toLower(jenisSampahBaru) != "b3") || toLower(jenisSampahBaru) == "" {
				fmt.Println("❌ Jenis sampah yang baru tidak valid")
				continue
			}

			fmt.Print("Masukkan jumlah sampah yang baru (kg): ")
			scanner.Scan()
			jumlahStr := scanner.Text()
			jumlahSampahBaru, err := strconv.Atoi(jumlahStr)
			if err != nil || jumlahSampahBaru <= 0 {
				fmt.Println("❌ Jumlah sampah yang baru tidak valid, harus lebih dari 0")
				continue
			}

			fmt.Print("Apakah sampah akan didaur ulang? (y/n): ")
			scanner.Scan()
			daurUlangBaru := strings.ToLower(scanner.Text())
			metodeBaru := ""
			if daurUlangBaru == "y" {
				fmt.Print("Masukkan metode daur ulang yang baru (A, B, atau C): ")
				scanner.Scan()
				metodeBaru = scanner.Text()
				if (toLower(metodeBaru) != "a" && toLower(metodeBaru) != "b" && toLower(metodeBaru) != "c") || toLower(metodeBaru) == "" {
					fmt.Println("❌ Metode daur ulang yang baru tidak valid")
					continue
				}
			} else if daurUlangBaru != "y" && daurUlangBaru != "n" {
				fmt.Println("❌ Input tidak valid")
				continue
			}

			if err := ubahSampah(idx, jenisSampahBaru, jumlahSampahBaru, daurUlangBaru == "y", metodeBaru); err != nil {
				fmt.Println("❌ Data sampah gagal diubah", err)
			} else {
				fmt.Println("\n✅ Data sampah berhasil diubah.")
			}

		case "3":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}
			tampilkanSemuaData()

			fmt.Print("\nMasukkan index data yang ingin dihapus: ")
			scanner.Scan()
			idxStr := scanner.Text()
			idx, err := strconv.Atoi(idxStr)
			if err != nil || idx < 0 || idx >= len(dataSampah) {
				fmt.Println("❌ Index tidak valid")
				continue
			}

			if err := hapusSampah(idx); err != nil {
				fmt.Println("❌ Data sampah gagal dihapus", err)
			} else {
				fmt.Println("\n✅ Data sampah berhasil dihapus.")
			}

		case "4":
			if err := tampilkanStatistik(); err != nil {
				fmt.Println("❌ Gagal menampilkan statistik", err)
			}

		case "5":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}

			fmt.Print("Cari data dari jenis sampah (Organik, Anorganik, atau B3) (sequential): ")
			scanner.Scan()
			jenisSampah := scanner.Text()
			if (toLower(jenisSampah) != "organik" && toLower(jenisSampah) != "anorganik" && toLower(jenisSampah) != "b3") || toLower(jenisSampah) == "" {
				fmt.Println("❌ Jenis sampah tidak valid")
				continue
			}

			sequentialSearch(jenisSampah)

		case "6":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}

			fmt.Print("Cari data dari jenis sampah (Organik, Anorganik, atau B3) (binary): ")
			scanner.Scan()
			jenisSampah := scanner.Text()
			if (toLower(jenisSampah) != "organik" && toLower(jenisSampah) != "anorganik" && toLower(jenisSampah) != "b3") || toLower(jenisSampah) == "" {
				fmt.Println("❌ Jenis sampah tidak valid")
				continue
			}

			binarySearch(jenisSampah)

		case "7":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}

			selectionSortByJumlah()
			fmt.Println("\n✅ Data sampah diurutkan berdasarkan jumlah sampah (selection sort).")

		case "8":
			if len(dataSampah) == 0 {
				fmt.Println("❌ Tidak ada data sampah")
				continue
			}

			insertionSortByJenis()
			fmt.Println("\n✅ Data sampah diurutkan berdasarkan jenis (insertion sort).")

		case "9":
			tampilkanSemuaData()

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