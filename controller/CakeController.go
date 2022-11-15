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

	db := DB.GetDbConnection()

	GetCakes := `
		SELECT id, title, description, rating, image, created_at, updated_at FROM cakes order by rating desc 
	`
	rows, err := db.QueryContext(r.Context(), GetCakes)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			Res.ResponseError(w, http.StatusInternalServerError, err)
		} else {
			hasil = append(hasil, cake)
		}

	}

	Res.ResponseJSON(w, http.StatusOK, hasil)
	log.Println("Berhasil menampilkan semua data ")
}

func GetCakeDetail(w http.ResponseWriter, r *http.Request) {
	var cakes models.Cakes
	params := mux.Vars(r)

	db := DB.GetDbConnection()

	getCakeQuery := `
		SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?
	`

	rows, err := db.QueryContext(r.Context(), getCakeQuery, params["id"])
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cakes.Id, &cakes.Title, &cakes.Description, &cakes.Rating, &cakes.Image, &cakes.CreatedAt, &cakes.UpdatedAt)
		if err != nil {
			Res.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
	}

	Res.ResponseJSON(w, http.StatusOK, cakes)
	log.Println("Berhasil menampilkan detail data")
}

func AddNewCake(w http.ResponseWriter, r *http.Request) {
	db := DB.GetDbConnection()
	var cakes models.Cakes

	if err := json.NewDecoder(r.Body).Decode(&cakes); err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	AddNewCake := `
		INSERT INTO cakes(id, title, description, rating, image, created_at, updated_at) VALUES (?,?,?,?,?, now(), now())
	`
	res, err := db.ExecContext(r.Context(), AddNewCake, &cakes.Id, &cakes.Title, &cakes.Description, &cakes.Rating, &cakes.Image)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	Res.ResponseJSON(w, http.StatusCreated, map[string]interface{}{"Insert Success, id:": lastID})
	log.Println("Data berhasil ditambahkan")
}

func UpdateCake(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var cakes models.Cakes
	db := DB.GetDbConnection()

	if err := json.NewDecoder(r.Body).Decode(&cakes); err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	updateCake := `
		UPDATE cakes SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?
	`
	_, err := db.ExecContext(r.Context(), updateCake, &cakes.Title, &cakes.Description, &cakes.Rating, &cakes.Image, params["id"])
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	Res.ResponseJSON(w, http.StatusCreated, "Update Success")
	log.Println("Data berhasil Terupdate")
}

func DeleteCake(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := DB.GetDbConnection()

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	deleteCake := `
		DELETE FROM cakes WHERE id = ?
	`

	_, err = db.ExecContext(r.Context(), deleteCake, id)
	if err != nil {
		Res.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	Res.ResponseJSON(w, http.StatusOK, "Delete Success")
	log.Println("Data berhasil dihapus")
}
