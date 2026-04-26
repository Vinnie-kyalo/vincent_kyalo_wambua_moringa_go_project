# 🧰 Prompt-Powered Kickstart: Beginner Web Server Toolkit in Go (Golang)

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

---

## 📌 Overview
This project is a beginner-friendly toolkit for building a simple web server using **Go (Golang)**. It introduces core backend concepts such as routing, handling HTTP requests, and returning JSON responses.

---

## 🎯 Objectives
- Build a simple web server in Go  
- Understand HTTP request handling  
- Learn how to return JSON responses  
- Practice backend development fundamentals  

---

## 🧠 Why Go?
Go is a fast, simple, and efficient programming language widely used in:
- Backend APIs  
- Cloud computing  
- Microservices  

**Real-world example:** Kubernetes is built using Go.

---

## ⚙️ System Requirements
- Windows / Linux / macOS  
- Go installed → https://go.dev/dl/  
- VS Code or any code editor  
- Terminal / Command Prompt  

### ✅ Verify Installation
```bash
go version
```

---

## 📁 Project Structure
```
go-web-server/
│
├── main.go
├── go.mod
├── handlers/
│   └── handlers.go
├── routes/
│   └── routes.go
└── utils/
    └── response.go
```

---

## 🚀 Getting Started

### 📥 Clone the Repository
```bash
git clone https://github.com/Vinnie-kyalo/vincent_kyalo_wambua_moringa_go_project.git
cd vincent_kyalo_wambua_moringa_go_project
```

---

### ▶️ Run the Server
```bash
go run main.go
```

---

### 🌐 Access the Server
Open in your browser:
```
http://localhost:8080
```

---

## 🌐 API Endpoints

### 🏠 Home
- **GET /**  
```json
{
  "message": "Welcome to Go Web Server!"
}
```

---

### 👋 Hello
- **GET /api/hello**  
```json
{
  "message": "Hello, World!"
}
```

---

## 🧪 Testing
Test the API using:

### cURL
```bash
curl http://localhost:8080
```

### Tools
- Postman  
- Browser  

---

## 📸 Example Output
```json
{
  "status": "success",
  "data": "Server is running"
}
```

---

## 📜 License
This project is licensed under the MIT License.

---

## 👨‍💻 Author
**Vincent Kyalo**  
🔗 https://github.com/Vinnie-kyalo  
