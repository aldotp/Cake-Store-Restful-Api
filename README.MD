# Cakestore API

**Base URL :**
> http:127.0.0.1:8080


# Instalation
Requirements
* Golang 
* Docker

# Use
- docker compose build (build image)
- docker compose up (untuk menjalankan)

# Note
* Migration auto migrate saat docker image dibuat
* Api bisa ditest menggunakan file postman (folder postman)

# API Documentation
GET /cakes
Menampilkan semua data

GET /cakes/{id}
Menampilkan detail cakes berdasarkan id

POST /cakes 
Menambah data cakes

PATCH /cakes/{id}
Edit data cakes

DELETE /cakes{id}
Hapus data cakes berdasarkan id

