package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	InitDB()
	defer db.Close()

	r := Routers()
	log.Println("Starting the HTTP server on port 9080")
	http.Handle("/", &CORSRouterDecorator{r})
	http.ListenAndServe(":9080", nil)
}

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/inventory", GetPeminjamanBuku).Methods("GET")
	router.HandleFunc("/api/inventory", CreatePeminjamanBuku).Methods("POST")
	router.HandleFunc("/api/inventory/{id}", GetPeminjamanBukuByID).Methods("GET")
	router.HandleFunc("/api/inventory/{id}", UpdatePeminjamanBuku).Methods("PUT")
	router.HandleFunc("/api/inventory/{id}", DeletePeminjamanBuku).Methods("DELETE")
	return router
}

/***************************************************/

// Get all users
func GetPeminjamanBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var peminjams []PeminjamanBuku
	result, err := db.Query("SELECT id, judul_buku, jumlah, nama_peminjam, alamat_peminjam, noHp_peminjam, tanggal_pinjam, tanggal_pengembalian, lama_pinjam FROM peminjamanbuku_dzulfikrinf")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var peminjam PeminjamanBuku
		err := result.Scan(&peminjam.ID, &peminjam.JudulBuku, &peminjam.Jumlah,
			&peminjam.NamaPeminjam, &peminjam.AlamatPeminjam, &peminjam.NoHpPeminjam,
			&peminjam.TanggalPinjam, &peminjam.TanggalPengembalian, &peminjam.LamaPinjam)
		if err != nil {
			panic(err.Error())
		}
		peminjams = append(peminjams, peminjam)
	}
	json.NewEncoder(w).Encode(peminjams)
}

// Create user
func CreatePeminjamanBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := db.Prepare("INSERT INTO peminjamanbuku_dzulfikrinf(judul_buku, jumlah, nama_peminjam, alamat_peminjam, noHp_peminjam, tanggal_pinjam, tanggal_pengembalian, lama_pinjam) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	judulBuku := requestData["judul_buku"].(string)
	jumlah := requestData["jumlah"].(string)
	namaPeminjam := requestData["nama_peminjam"].(string)
	alamatPeminjam := requestData["alamat_peminjam"].(string)
	noHpPeminjam := requestData["noHp_peminjam"].(string)
	tanggalPinjam := requestData["tanggal_pinjam"].(string)
	tanggalPengembalian := requestData["tanggal_pengembalian"].(string)
	lamaPinjam := requestData["lama_pinjam"].(string)

	_, err = stmt.Exec(judulBuku, jumlah, namaPeminjam, alamatPeminjam, noHpPeminjam, tanggalPinjam, tanggalPengembalian, lamaPinjam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "New user was created")
}

// Get user by ID
func GetPeminjamanBukuByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var peminjam PeminjamanBuku

	result := db.QueryRow("SELECT id, judul_buku, jumlah, nama_peminjam, alamat_peminjam, noHp_peminjam, tanggal_pinjam, tanggal_pengembalian, lama_pinjam FROM peminjamanbuku_dzulfikrinf WHERE id = ?", params["id"])

	err := result.Scan(&peminjam.ID, &peminjam.JudulBuku, &peminjam.Jumlah, &peminjam.NamaPeminjam, &peminjam.AlamatPeminjam, &peminjam.NoHpPeminjam, &peminjam.TanggalPinjam, &peminjam.TanggalPengembalian, &peminjam.LamaPinjam)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintf(w, "Peminjaman buku dengan ID %s tidak ditemukan", params["id"])
		} else {
			panic(err.Error())
		}
		return
	}

	json.NewEncoder(w).Encode(peminjam)
}

// Update user
func UpdatePeminjamanBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE peminjamanbuku_dzulfikrinf SET judul_buku = ?," +
		"jumlah = ?, nama_peminjam = ?, alamat_peminjam = ?, noHp_peminjam = ?, tanggal_pinjam = ?, tanggal_pengembalian = ?, lama_pinjam = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	judulBuku := requestData["judul_buku"].(string)
	jumlah := requestData["jumlah"].(string)
	namaPeminjam := requestData["nama_peminjam"].(string)
	alamatPeminjam := requestData["alamat_peminjam"].(string)
	noHpPeminjam := requestData["noHp_peminjam"].(string)
	tanggalPinjam := requestData["tanggal_pinjam"].(string)
	tanggalPengembalian := requestData["tanggal_pengembalian"].(string)
	lamaPinjam := requestData["lama_pinjam"].(string)

	_, err = stmt.Exec(judulBuku, jumlah, namaPeminjam, alamatPeminjam, noHpPeminjam, tanggalPinjam, tanggalPengembalian, lamaPinjam, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Data peminjaman dengan ID = %s telah diupdate", params["id"])
}

// Delete User
func DeletePeminjamanBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM peminjamanbuku_dzulfikrinf WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Peminjaman buku dengan ID = %s berhasil dihapus", params["id"])
}

/***************************************************/

type PeminjamanBuku struct {
	ID                  string `json:"id"`
	JudulBuku           string `json:"judul_buku"`
	Jumlah              int    `json:"jumlah"`
	NamaPeminjam        string `json:"nama_peminjam"`
	AlamatPeminjam      string `json:"alamat_peminjam"`
	NoHpPeminjam        string `json:"noHp_peminjam"`
	TanggalPinjam       string `json:"tanggal_pinjam"`
	TanggalPengembalian string `json:"tanggal_pengembalian"`
	LamaPinjam          string `json:"lama_pinjam"`
}

// Db configuration
var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_2204083_dzulfikrinf_uas_pilkomB")
	if err != nil {
		log.Fatal(err)
	}
}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}
