# NVI Identity Verification API

This project provides an API for identity verification of Turkish citizens and foreign individuals. It is written in **Go** and can be easily run using Docker.

---

## Setup and Run

### 1. Clone the Project
```bash
git clone https://github.com/veliardabalci/NVI-Identity-Verification-Service.git
cd NVI-Identity-Verification-Service
```

### 2. Run with Docker
```bash
docker build -t nvi-api .
docker run -d -p 8080:8080 --name nvi-api nvi-api
```

The API will now be accessible at `http://localhost:8080`.

---

## API Usage

### Turkish Citizens Verification
- **Endpoint:** `POST /check/turkish`
- **Body:**
```json
{
  "id": "12345678901",
  "name": "Ahmet",
  "surname": "Yılmaz",
  "birth_year": "1985"
}
```

- **cURL:**
```bash
curl -X POST http://localhost:8080/check/turkish \
-H "Content-Type: application/json" \
-d '{"id": "12345678901", "name": "Ahmet", "surname": "Yılmaz", "birth_year": "1985"}'
```
- **Response:**
```json
{
  "is_valid": true
}
```

### Foreign Identity Verification
- **Endpoint:** `POST /check/foreign`
- **Body:**
```json
{
  "id": "987654321",
  "name": "John",
  "surname": "Doe",
  "birth_day": 15,
  "birth_month": 6,
  "birth_year": 1990
}
```

- **cURL:**
```bash
curl -X POST http://localhost:8080/check/foreign \
-H "Content-Type: application/json" \
-d '{"id": "987654321", "name": "John", "surname": "Doe", "birth_day": 15, "birth_month": 6, "birth_year": 1990}'
```
- **Response:**
```json
{
  "is_valid": true
}
```