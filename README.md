## CRUD dengan Golang

1. Clone repository: `git clone https://github.com/williamy101/go_crud.git` dan masuk ke folder proyek.
2. Jalankan file `Script.sql` di MySQL untuk membuat database, tabel, dan data awal. (Ganti bagian your_password di file db.go dengan password MySQL masing-masing)
3. Jalankan server dengan perintah `go run main.go`. Server akan aktif di `http://localhost:8080`.
4. Gunakan Postman untuk uji API:
   - Produk/Product:
     - GET `/product`: Melihat semua produk.
     - POST `/product`: Menambahkan produk.
     - PUT `/product/:product_id`: Memperbarui produk berdasarkan ProductID.
     - DELETE `/product/:product_id`: Menghapus produk berdasarkan ProductID.
   - Inventaris/Inventory:
     - GET `/inventory`: Melihat tingkat stok produk.
     - PUT `/inventory/:product_id`: Memperbarui tingkat stok produk.
   - Pesanan/Order:
     - GET `/orders/:order_id`: Melihat detail pesanan berdasarkan OrderID.
     - POST `/orders`: Membuat pesanan baru.
   - Gambar Produk:
     - POST `/product/:product_id/image`: Mengunggah gambar produk.
     - GET `/product/:product_id/image`: Mengunduh gambar produk.
5. Gambar yang diunggah akan disimpan di folder `uploads/`.
