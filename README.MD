# Cakestore API

**Base URL :**
> http:127.0.0.1:8080

# Instalation
Requirements
* Golang 
* Docker
* Xampp (Optional)

Use
- docker compose build 
- docker compose up 

# Notes
* Testing dan API sudah berjalan ketika docker compose up dijalankan
* Migration sql sudah auto saat docker image dibuat
* Api bisa ditest menggunakan Postman dengan export file (folder postman)
* Jika pakai manual atau local ubah settingan .env

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

