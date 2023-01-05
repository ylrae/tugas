package main

import "fmt"

func main() {
	var tgl1, tgl2, thn1, bln2, thn2 int // pengajuan
	var hr1, hr2, bln1 string            // pengambilan

	fmt.Scan(&hr1, &tgl1, &bln1, &thn1)
	// panggil subprogram untuk penentuan tanggal pengambilan
	pengambilan(tgl1, angkaBulan(bln1), thn1, hr1, &tgl2, &bln2, &thn2, &hr2)
	// tampilkan tanggal pengambilan visa
	fmt.Println(hr2, tgl2, bulanAngka(bln2), thn2)
}

func kabisat(tahun int) bool {
	return tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
}

func angkaBulan(bulan string) int {
	if bulan == "januari" {
		return 1
	} else if bulan == "februari" {
		return 2
	} else if bulan == "februari" {
		return 2
	} else if bulan == "maret" {
		return 3
	} else if bulan == "april" {
		return 4
	} else if bulan == "mei" {
		return 5
	} else if bulan == "juni" {
		return 6
	} else if bulan == "juli" {
		return 7
	} else if bulan == "agustus" {
		return 8
	} else if bulan == "september" {
		return 9
	} else if bulan == "oktober" {
		return 10
	} else if bulan == "november" {
		return 11
	} else if bulan == "desember" {
		return 12
	} else {
		return 0
	}

}

func bulanAngka(angka int) string {
	/* Mengembalikan nama bulan berdasarkan urutan angka bulan pada kalender masehi (1 hingga 12)."invalid" untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua. Contoh: 1 menjadi "januari" */
	if angka == 1 {
		return "januari"
	} else if angka == 2 {
		return "februari"
	} else if angka == 3 {
		return "maret"
	} else if angka == 4 {
		return "april"
	} else if angka == 5 {
		return "mei"
	} else if angka == 6 {
		return "juni"
	} else if angka == 7 {
		return "juli"
	} else if angka == 8 {
		return "agustus"
	} else if angka == 9 {
		return "september"
	} else if angka == 10 {
		return "oktober"
	} else if angka == 11 {
		return "november"
	} else if angka == 12 {
		return "desember"
	} else {
		return "invalid"
	}
}

func jumlahHari(bulan, tahun int) int {
	/* Mengembalikan jumlah hari berdasarkan bulan dan tahun yang terdefinisi,hati-hati pada bulan februari pada saat kabisat. -1 apabila bulan tidak valid */
	if bulan == 2 {
		if kabisat(tahun) {
			return 29
		} else {
			return 28
		}
	} else if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		return 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		return 30
	} else {
		return -1
	}
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	/* I.S. terdefinisi hari1 yang menyatakan hari pengajuan string, asumsi huruf kecil semuaF.S. hari2 berisi hari pengambilan dan durasi berisi lama pembuatan visa, karena sabtu dan minggu tidak dihitung */
	*durasi = 2
	if hari1 == "senin" {
		*hari2 = "rabu"
	} else if hari1 == "selasa" {
		*hari2 = "kamis"
	} else if hari1 == "rabu" {
		*hari2 = "jumat"
	} else if hari1 == "kamis" {
		*hari2 = "senin"
		*durasi = 4
	} else if hari1 == "jumat" {
		*hari2 = "selasa"
		*durasi = 4
	}

}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	/* I.S. terdefinisi waktu pengajuan visa, yaitu tanggal1, bulan1, tahun1 dan hari1F.S. tanggal2, bulan2, tahun2 dan hari2 berisi waktu pengambilan visa */
	var durasi, lamaBulan int

	mencariDurasi(hari1, hari2, &durasi)
	lamaBulan = jumlahHari(bulan1, tahun1)

	*tanggal2 = tanggal1 + durasi
	*bulan2 = bulan1
	*tahun2 = tahun1

	if *tanggal2 > lamaBulan {
		*tanggal2 = *tanggal2 - lamaBulan
		*bulan2 = bulan1 + 1
		if *bulan2 == 12 {
			*bulan2 = 1
			*tahun2 = tahun1 + 1
		}
	}
}
