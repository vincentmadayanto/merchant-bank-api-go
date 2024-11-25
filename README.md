# Merchant Bank API

## Fitur
- **Login**: Memverifikasi kredensial pelanggan dan menghasilkan token JWT.
- **Pembayaran**: Memproses pembayaran untuk pelanggan yang sudah login.
- **Logout**: Mengakhiri sesi pelanggan.
- **Pencatatan Aktivitas**: Semua aktivitas dicatat ke dalam file JSON untuk histori.

## Alur Kerja Aplikasi

### 1. Login
- **Endpoint**: POST /login
- **Input**: ID pelanggan dan password
- **Proses**:
  - Autentikasi dilakukan dengan mencocokkan ID dan password terhadap data pelanggan di customers.json
  - Jika berhasil, token JWT dibuat dan dikembalikan ke klien
- **Output**: Token JWT

### 2. Pembayaran
- **Endpoint**: POST /payment
- **Input**: 
  - Header CustomerID 
  - Body JSON dengan jumlah pembayaran
- **Proses**:
  - Validasi token JWT melalui middleware
  - Memproses pembayaran dan mencatat transaksi di transactions.json
- **Output**: Status sukses atau pesan error

### 3. Logout
- **Endpoint**: POST /logout
- **Proses**: Mengakhiri sesi pengguna (untuk API ini hanya simbolis)
- **Output**: Pesan logout sukses

## Langkah Instalasi dan Penggunaan

### 1. Kloning Repository
```bash
git clone <repository_url>
cd merchant-bank-api
```

### 2. Install Dependency
Pastikan untuk menginstall dependency berikut:
- Gorilla Mux: Routing untuk API
- JWT-Go: Library untuk manajemen token JWT
- Godotenv: Membaca file .env

### 3. Jalankan Aplikasi
Pastikan Anda telah menginstal Go dan semua dependensi proyek:
```bash
go run cmd/main.go
```

### 4. Uji API
Gunakan Postman atau cURL untuk mengakses endpoint:

#### Login API:
```http
POST http://localhost:8080/login

Body:
{
    "id": "customer1",
    "password": "password123"
}
```

#### Payment API:
```http
POST http://localhost:8080/payment

Header:
CustomerID: <customer_id>

Body:
{
    "amount": 100.0
}
```

#### Logout API:
```http
POST http://localhost:8080/logout
```

## Penggunaan API

### Endpoint Login
- URL: `/login`
- Metode: `POST`
- Body Request:
```json
{
    "id": "customer1",
    "password": "password"
}
```
- Response Sukses:
```json
{
    "token": "[JWT_TOKEN]"
}
```

### Endpoint Pembayaran
- URL: `/payment`
- Metode: `POST`
- Headers:
  - CustomerID: [ID_PELANGGAN]
- Body Request:
```json
{
    "amount": 100.0
}
```
- Response Sukses:
```json
{
    "message": "Payment Successful"
}
```

### Endpoint Logout
- URL: `/logout`
- Metode: `POST`
- Response Sukses:
```json
{
    "message": "Logout Successful"
}
```
