# Kasir API

REST API sederhana untuk manajemen produk dan kategori (in-memory).

## Base URL

- Lokal: `http://localhost:8080`
- Vercel: `https://<your-app>.vercel.app`

## Menjalankan Lokal

```bash
go run .
```

## Health Check

### GET /health

Response:

```json
{
  "status": "ok"
}
```

## Products

### GET /api/product

Response:

```json
[
  {
    "id": 1,
    "name": "Beras",
    "price": 15000,
    "stock": 20
  }
]
```

### GET /api/product/{id}

Response:

```json
{
  "id": 1,
  "name": "Beras",
  "price": 15000,
  "stock": 20
}
```

Jika ID tidak valid:

```json
{
  "error": "invalid id"
}
```

Jika tidak ditemukan:

```json
{
  "error": "product not found"
}
```

### POST /api/product

Request body:

```json
{
  "name": "Beras",
  "price": 15000,
  "stock": 20
}
```

Response:

```json
{
  "id": 1,
  "name": "Beras",
  "price": 15000,
  "stock": 20
}
```

### PUT /api/product/{id}

Request body:

```json
{
  "name": "Beras Premium",
  "price": 20000,
  "stock": 12
}
```

Response:

```json
{
  "id": 1,
  "name": "Beras Premium",
  "price": 20000,
  "stock": 12
}
```

Jika ID tidak ditemukan:

```json
{
  "error": "product not found"
}
```

### DELETE /api/product/{id}

Response:

```json
{
  "message": "product deleted successfully"
}
```

Jika ID tidak ditemukan:

```json
{
  "error": "product not found"
}
```

## Categories

### GET /api/category

Response:

```json
[
  {
    "id": 1,
    "name": "Sembako",
    "description": "Produk kebutuhan sehari-hari"
  }
]
```

### GET /api/category/{id}

Response:

```json
{
  "id": 1,
  "name": "Sembako",
  "description": "Produk kebutuhan sehari-hari"
}
```

Jika ID tidak valid:

```json
{
  "error": "invalid id"
}
```

Jika tidak ditemukan:

```json
{
  "error": "category not found"
}
```

### POST /api/category

Request body:

```json
{
  "name": "Sembako",
  "description": "Produk kebutuhan sehari-hari"
}
```

Response:

```json
{
  "id": 1,
  "name": "Sembako",
  "description": "Produk kebutuhan sehari-hari"
}
```

### PUT /api/category/{id}

Request body:

```json
{
  "name": "Sembako",
  "description": "Produk kebutuhan pokok"
}
```

Response:

```json
{
  "id": 1,
  "name": "Sembako",
  "description": "Produk kebutuhan pokok"
}
```

Jika ID tidak ditemukan:

```json
{
  "error": "category not found"
}
```

### DELETE /api/category/{id}

Response:

```json
{
  "message": "category deleted successfully"
}
```

Jika ID tidak ditemukan:

```json
{
  "error": "category not found"
}
```

## Catatan

- Data disimpan di memori (bukan database), akan reset setiap restart.
- Semua response menggunakan JSON.
