package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type kurikulum struct {
	Semester    int
	Kode        string
	Mata_kuliah string
	Sks         int
}

const harga_sks int = 90000
const harga_spp int = 600000

var data = []kurikulum{
	kurikulum{1, "IF4002", "Dasar Pemrograman", 4},
	kurikulum{1, "ST1018", "Pendidikan Kewarganegaraan", 2},
	kurikulum{1, "FT2001", "Kalkulus", 3},
	kurikulum{1, "FT3001", "Pengantar Teknologi Informasi", 3},
	kurikulum{1, "IF4001", "Pengantar Sistem Digital", 3},
	kurikulum{1, "FT3009", "Basis Data", 4},
	kurikulum{1, "ST1020", "Bahasa Indonesia", 2},
	kurikulum{2, "FT3003", "Arsitektur & Organisasi Komputer", 3},
	kurikulum{2, "IF2004", "Kalkulus Lanjut", 3},
	kurikulum{2, "IF2005", "Logika Informatika", 2},
	kurikulum{2, "ST1005", "Ilmu Alamiah Dasar", 2},
	kurikulum{2, "IF2006", "Fisika Teknik", 3},
	kurikulum{3, "FT3104", "Sistem Basis Data", 4},
	kurikulum{3, "FT3012", "Algoritma dan Struktur Data", 3},
	kurikulum{3, "FT2002", "Statistika dan Probabilitas", 3},
	kurikulum{3, "FT2003", "Aljabar Vektor dan Matriks", 3},
	kurikulum{3, "FT3004", "Sistem Operasi", 3},
	kurikulum{3, "IF2007", "Matematika Diskrit", 3},
	kurikulum{3, "ST1008", "Kewirausahaan", 2},
	kurikulum{4, "FT3005", "Jaringan Komputer", 3},
	kurikulum{4, "FT3006", "Pemrograman Berorientasi Objek", 4},
	kurikulum{4, "ST1007", "Pengantar Manajemen dan Bisnis", 2},
	kurikulum{4, "ST1016", "Etika Profesi", 2},
	kurikulum{4, "IF4003", "Dasar Rekayasa Perangkat Lunak", 2},
	kurikulum{4, "IF4004", "Grafika Komputer", 3},
	kurikulum{4, "IF4005", "Strategi Algoritma", 2},
	kurikulum{4, "IF4006", "Teori Bahasa Formal & Automata", 3},
	kurikulum{5, "ST1010", "Hukum Bisnis & TI", 2},
	kurikulum{5, "ST1012", "Interpersonal Skill", 3},
	kurikulum{5, "FT3106", "Manajemen Proyek Perangkat Lunak", 4},
	kurikulum{5, "FT3011", "Sistem Informasi", 3},
	kurikulum{5, "FT3008", "Interaksi Manusia & Komputer", 3},
	kurikulum{5, "FT3108", "Kecerdasan Buatan", 3},
	kurikulum{5, "IF4007", "Sistem Paralel & Terdistribusi", 3},
	kurikulum{6, "FT3109", "Pengembangan Aplikasi Berbasis Web", 4},
	kurikulum{6, "IF4008", "Pengembangan Aplikasi Platform Khusus", 4},
	kurikulum{6, "IF4009", "Sistem Multimedia", 3},
	kurikulum{6, "ST1002", "Pendidikan Pancasila & Kewarganegaraan", 3},
	kurikulum{6, "ST1011", "Komputer Masyarakat", 3},
	kurikulum{6, "IF4014", "Proyek Perangkat Lunak", 4},
	kurikulum{7, "IF4010", "Pengembangan Aplikasi Terdistribusi", 4},
	kurikulum{7, "IF4011", "Jaringan Komputer Lanjut", 4},
	kurikulum{7, "IF4012", "Visualisasi Data dan Informasi", 4},
	kurikulum{7, "IF4013", "Pengembangan Aplikasi Berbasis 3D", 4},
	kurikulum{7, "IF1080", "Kerja Praktek", 2},
	kurikulum{8, "IF2071", "Skripsi / Tugas Akhir", 6},
	kurikulum{8, "IF4015", "Perangkat Lunak Berbasis Komponen", 4},
	kurikulum{8, "IF4016", "Perangkat Lunak Berbasis Service", 4},
	kurikulum{8, "IF4017", "Sistem Informasi Lanjut", 4},
	kurikulum{8, "IF4018", "Rekayasa Interaksi", 4},
}

func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

var tmpl = template.Must(template.New("foo").Funcs(template.FuncMap{"N": N}).ParseGlob("form/*"))

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// nId := r.URL.Query().Get("id")
		semester, err1 := strconv.Atoi(r.URL.Query().Get("semester"))
		if err1 != nil {
			http.Error(w, "Bukan Angka", http.StatusBadRequest)
			return
		}
		var result []byte
		var err error
		var kur []kurikulum
		for _, each := range data {
			if each.Semester == semester {
				kur = append(kur, each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
		result, err = json.Marshal(kur)
		w.Write(result)
	}
	http.Error(w, "", http.StatusBadRequest)
}

func Index(w http.ResponseWriter, r *http.Request) {
	type M map[string]interface{}
	var data = M{
		"page":  "home",
		"judul": "Home",
	}
	tmpl.ExecuteTemplate(w, "Index", data)
}

func Tentang(w http.ResponseWriter, r *http.Request) {
	type M map[string]interface{}
	var data = M{
		"page":  "tentang",
		"judul": "Tentang Aplikasi",
	}
	tmpl.ExecuteTemplate(w, "Tentang", data)
}

func Pembayaran_sks(w http.ResponseWriter, r *http.Request) {
	type M map[string]interface{}
	var data = M{
		"page":  "pembayaran_sks",
		"judul": "Pembayaran SKS",
		"jurusan": map[string]string{
			"ti": "Teknik Informatika",
			"mi": "Sistem Informasi",
			"si": "Manajemen Informatika",
		},
		"harga_sks": harga_sks,
		"sks": []map[string]string{map[string]string{"name": "chicken	blue", "gender": "male"}, map[string]string{"name": "chicken	red", "gender": "male"}, map[string]string{"name": "chicken	yellow", "gender": "female"}}}
	tmpl.ExecuteTemplate(w, "Pembayaran_sks", data)
}

func Pembayaran_spp(w http.ResponseWriter, r *http.Request) {
	type M map[string]interface{}
	var data = M{
		"page":      "pembayaran_spp",
		"judul":     "Pembayaran SPP",
		"harga_spp": harga_spp,
		"bulan": map[int]string{
			1:  "Januari",
			2:  "Februari",
			3:  "Maret",
			4:  "April",
			5:  "Mei",
			6:  "Juni",
			7:  "Juli",
			8:  "Agustus",
			9:  "September",
			10: "Oktober",
			11: "November",
			12: "Desember",
		},
	}
	tmpl.ExecuteTemplate(w, "Pembayaran_spp", data)
}

func post_sks(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nim := r.FormValue("nim")
		nama := r.FormValue("nama")
		r.ParseForm()
		kode := r.Form["kode"]
		type User struct {
			Status      bool
			Nim         string
			Nama        string
			Semester    int
			Jumlah_sks  int
			Total_bayar int
			Sekarang    string
			Arr         []kurikulum
		}

		var jumlah_sks int = 0

		semester, err1 := strconv.Atoi(r.FormValue("semester"))
		if err1 != nil {
			http.Error(w, "Bukan Angka", http.StatusBadRequest)
			return
		}

		var kur []kurikulum

		// log.Println(kode)
		for _, ko := range kode {
			// jumlah_sks += 1
			// log.Println(ko)
			for _, each := range data {
				if each.Kode == ko {
					jumlah_sks += each.Sks
					kur = append(kur, each)
				}
			}
		}

		var total_bayar = jumlah_sks * harga_sks

		w.Header().Set("Content-Type", "application/json")
		user := &User{Sekarang: sekarang(), Arr: kur, Nim: nim, Status: true, Nama: nama, Semester: semester, Jumlah_sks: jumlah_sks, Total_bayar: total_bayar}
		b, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}

func post_spp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nim := r.FormValue("nim")
		nama := r.FormValue("nama")
		bulan := bulan(r.FormValue("bulan"))
		type User struct {
			Status      bool
			Nim         string
			Nama        string
			Bulan       string
			Total_bayar int
			Sekarang    string
		}

		var total_bayar = harga_spp

		w.Header().Set("Content-Type", "application/json")
		user := &User{Sekarang: sekarang(), Nim: nim, Status: true, Nama: nama, Bulan: bulan, Total_bayar: total_bayar}
		b, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}

func bulan(no_bln string) string {
	var nama_bulan string
	switch no_bln {
	case "1":
		nama_bulan = "Januari"
	case "2":
		nama_bulan = "Februari"
	case "3":
		nama_bulan = "Maret"
	case "4":
		nama_bulan = "April"
	case "5":
		nama_bulan = "Mei"
	case "6":
		nama_bulan = "Juni"
	case "7":
		nama_bulan = "Juli"
	case "8":
		nama_bulan = "Agustus"
	case "9":
		nama_bulan = "September"
	case "10":
		nama_bulan = "Oktober"
	case "11":
		nama_bulan = "November"
	case "12":
		nama_bulan = "Desember"
	default:
		nama_bulan = "null"
	}

	return nama_bulan

}

func sekarang() string {
	t := time.Now().Local()
	jam := t.Format("15:04")
	tanggal := t.Format("02")
	bulan := bulan(t.Format("1"))
	tahun := t.Format("2006")
	// var result string = strings.Join(tanggal, nama_bulan, tahun, jam)
	result := strings.Join([]string{tanggal, bulan, tahun, jam}, " ")
	return result

}

func get_total(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		kode := r.Form["kode"]
		type Json struct {
			Total_bayar int
		}

		var jumlah_sks int = 0

		// log.Println(kode)
		for _, ko := range kode {
			// jumlah_sks += 1
			// log.Println(ko)
			for _, each := range data {
				if each.Kode == ko {
					jumlah_sks += each.Sks
				}
			}
		}

		var total_bayar = jumlah_sks * harga_sks

		w.Header().Set("Content-Type", "application/json")
		user := &Json{Total_bayar: total_bayar}
		b, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	log.Println("Server telah dijalankan di : http://localhost:3001")
	http.HandleFunc("/", Index)
	http.HandleFunc("/pembayaran_sks", Pembayaran_sks)
	http.HandleFunc("/pembayaran_spp", Pembayaran_spp)
	http.HandleFunc("/tentang", Tentang)
	http.HandleFunc("/users", users)
	http.HandleFunc("/post_sks", post_sks)
	http.HandleFunc("/post_spp", post_spp)
	http.HandleFunc("/get_total", get_total)
	http.ListenAndServe(":3001", nil)
}
