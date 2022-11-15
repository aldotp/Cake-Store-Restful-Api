# Cakestore API

**Base URL :**
> http:127.0.0.1:8080

# Instalation
Requirements
* Golang 
* Docker
* Xampp (Optional)

Use API, MySQL, Phpmyadmin (docker)
- docker compose up 

Use Unit Testing (docker)
- docker exec -it -e DB_HOST=golang-db -e DB_PASS=root -e DB_NAME=cakestore  golang-api go test -v ./...
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

