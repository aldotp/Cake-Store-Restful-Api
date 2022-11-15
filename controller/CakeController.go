package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "time"

	DB "github.com/aldotp/cakestore/config"
	Res "github.com/aldotp/cakestore/helper"
	"github.com/aldotp/cakestore/models"
)

func GetCakes(w http.ResponseWriter, r *http.Request) {
	var cake models.Cakes
	var hasil []models.Cakes

	db := DB.ConnectToDatabase()

	sqlQuery := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes order by rating desc "
	rows, err := db.Query(sqlQuery)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, "Terjadi Kesalahan")
		log.Println("Terjadi error saat query data =", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			Res.ResponseError(w, http.StatusInternalServerError, err.Error())
			log.Println("Terjadi error saat scan data =", err.Error())
		} else {
			hasil = append(hasil, cake)
		}
	}
	Res.ResponseJSON(w, http.StatusOK, hasil)
	log.Println("Berhasil Menampilkan Data ")
}

func GetCakeDetail(w http.ResponseWriter, r *http.Request) {
	var cakes models.Cakes
	params := mux.Vars(r)

	db := DB.ConnectToDatabase()

	sqlQuery := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"
	rows, err := db.Query(sqlQuery, params["id"])
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat query data =", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cakes.Id, &cakes.Title, &cakes.Description, &cakes.Rating, &cakes.Image, &cakes.CreatedAt, &cakes.UpdatedAt)
		if err != nil {
			Res.ResponseError(w, http.StatusInternalServerError, err.Error())
			log.Println("Terjadi error saat scan data =", err.Error())
			return
		}
	}
	Res.ResponseJSON(w, http.StatusOK, cakes)
	log.Println("Berhasil menampilkan data dengan id", cakes.Id)

}

func AddNewCake(w http.ResponseWriter, r *http.Request) {

	db := DB.ConnectToDatabase()
	defer db.Close()

	var cakes models.Cakes

	if err := json.NewDecoder(r.Body).Decode(&cakes); err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat mendecode query =", err.Error())
		return
	}
	sqlQuery := "INSERT INTO cakes(title, description, rating, image, created_at, updated_at) VALUES (?,?,?,?, now(), now())"
	_, err := db.Exec(sqlQuery, &cakes.Title, &cakes.Description, &cakes.Rating, &cakes.Image)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, "Terdapat kesalahan")
		log.Println("Terjadi error saat proses query =", err.Error())
		return
	}
	Res.ResponseJSON(w, http.StatusOK, "Insert data Success")
	log.Println("Data berhasil di Tambahkan")

}

func UpdateCake(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	db := DB.ConnectToDatabase()
	// r.ParseForm()

	var cakes models.Cakes

	if err := json.NewDecoder(r.Body).Decode(&cakes); err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, "Tidak dapat mendecode body")
		log.Println("Terjadi error saat Mendecode =", err.Error())
		return
	}

	sqlQuery := "UPDATE cakes SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?"
	statement, err := db.Prepare(sqlQuery)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat Memprepare query =", err.Error())
		return
	}

	_, err = statement.Exec(&cakes.Title, &cakes.Description, &cakes.Rating, &cakes.Image, params["id"])
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat Mengesekusi query =", err.Error())
		return
	}
	Res.ResponseJSON(w, http.StatusOK, "Update Success")
	log.Println("Data berhasil diupdate")
}

func DeleteCake(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := DB.ConnectToDatabase()

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat memproses data =", err.Error())
		return
	}

	sqlQuery := "DELETE FROM cakes WHERE id = ?"
	statement, err := db.Prepare(sqlQuery)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat Memprepare query =", err.Error())
		return
	}

	_, err = statement.Exec(id)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err.Error())
		log.Println("Terjadi error saat mengeksekusi query =", err.Error())
		return
	}
	Res.ResponseJSON(w, http.StatusOK, "Delete Success")
	log.Println("Data berhasil dihapus")

}
