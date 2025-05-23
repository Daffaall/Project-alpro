package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Sampah struct {
    Jenis            string
    Jumlah           int
    DaurUlang        bool
    MetodeDaurUlang  string
}

var dataSampah []Sampah

func tambahSampah(jenis string, jumlah int, daurUlang bool, metode string) {
    dataSampah = append(dataSampah, Sampah{jenis, jumlah, daurUlang, metode})
    fmt.Println("\n✅ Data sampah berhasil ditambahkan.")
}

func ubahSampah(index int, jenis string, jumlah int, daurUlang bool, metode string) {
    if index >= 0 && index < len(dataSampah) {
        dataSampah[index] = Sampah{jenis, jumlah, daurUlang, metode}
        fmt.Println("\n✅ Data berhasil diubah.")
    } else {
        fmt.Println("❌ Index tidak valid.")
    }
}

func hapusSampah(index int) {
    if index >= 0 && index < len(dataSampah) {
        dataSampah = append(dataSampah[:index], dataSampah[index+1:]...)
        fmt.Println("\n✅ Data berhasil dihapus.")
    } else {
        fmt.Println("❌ Index tidak valid.")
    }
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


func tampilkanStatistik() {
    total := 0
    totalDaurUlang := 0
    for _, s := range dataSampah {
        total += s.Jumlah
        if s.DaurUlang {
            totalDaurUlang += s.Jumlah
        }
    }
    fmt.Println("\n📈 Total sampah:", total ,"Kg")
    fmt.Println("♻️  Total yang didaur ulang:", totalDaurUlang, "Kg")
    fmt.Println("♻️  Persentase yang didaur ulang:", float64(totalDaurUlang)/float64(total)*100, "%") //New
    fmt.Println("♻️  Persentase yang tidak didaur ulang:", float64(total-totalDaurUlang)/float64(total)*100, "%")
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
            fmt.Print("Jumlah (kg): ")
            scanner.Scan()
            jumlah, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Didaur ulang? (y/n): ")
            scanner.Scan()
            du := strings.ToLower(scanner.Text()) == "y"
            metode := ""
            if du {
                fmt.Print("Metode daur ulang: ")
                scanner.Scan()
                metode = scanner.Text()
            }
            tambahSampah(jenis, jumlah, du, metode)

        case "2":
            fmt.Print("Index data yang ingin diubah: ")
            scanner.Scan()
            idx, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Jenis baru: ")
            scanner.Scan()
            jenis := scanner.Text()
            fmt.Print("Jumlah baru (kg): ")
            scanner.Scan()
            jumlah, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Didaur ulang? (y/n): ")
            scanner.Scan()
            du := strings.ToLower(scanner.Text()) == "y"
            metode := ""
            if du {
                fmt.Print("Metode daur ulang: ")
                scanner.Scan()
                metode = scanner.Text()
            }
            ubahSampah(idx, jenis, jumlah, du, metode)

        case "3":
            fmt.Print("Index data yang ingin dihapus: ")
            scanner.Scan()
            idx, _ := strconv.Atoi(scanner.Text())
            hapusSampah(idx)

        case "4":
            tampilkanStatistik()

        case "5":
            fmt.Print("Cari jenis (sequential): ")
            scanner.Scan()
            jenis := scanner.Text()
            index := sequentialSearch(jenis)
            if index != -1 {
                fmt.Println("Ditemukan di index:", index, dataSampah[index])
            } else {
                fmt.Println("Tidak ditemukan.")
            }

        case "6":
            fmt.Print("Cari jenis (binary): ")
            scanner.Scan()
            jenis := scanner.Text()
            index := binarySearch(jenis)
            if index != -1 {
                fmt.Println("Ditemukan di index:", index, dataSampah[index])
            } else {
                fmt.Println("Tidak ditemukan.")
            }

        case "7":
            selectionSortByJumlah()
            fmt.Println("\n✅ Diurutkan berdasarkan jumlah (selection sort).")

        case "8":
            insertionSortByJenis()
            fmt.Println("\n✅ Diurutkan berdasarkan jenis (insertion sort).")
            for i, s := range dataSampah {
                fmt.Printf("[%d] %s - %dkg - Daur ulang: %v - Metode: %s\n", i, s.Jenis, s.Jumlah, s.DaurUlang, s.MetodeDaurUlang)
            }

        case "9":
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
