# 🧰 Prompt-Powered Kickstart: Beginner Web Server Toolkit in Go (Golang)

---

## 📌 1. Title & Objective

### Title

Prompt-Powered Kickstart: Building a Beginner Web API Toolkit in Go (Golang)

### Objective

This project demonstrates how to build a simple web server using Go. It helps beginners understand backend development concepts such as routing, handling HTTP requests, and returning JSON responses.

### Why Go?

Go is a fast, simple, and powerful backend language used in cloud systems, APIs, and large-scale applications.

---

## 🧠 2. Technology Overview

### What is Go (Golang)?

Go is an open-source programming language created by Google designed for simplicity and high performance.

### Where it is used:

* Backend APIs
* Cloud computing systems
* Microservices

### Real-world example:

* Kubernetes is built using Go

---

## ⚙️ 3. System Requirements

* Windows / Linux / macOS
* Go installed ([https://go.dev/dl/](https://go.dev/dl/))
* VS Code or any code editor
* Terminal / Command Prompt

### Verify installation:

```bash
go version
```

---

## 📦 4. Project Setup

### Step 1: Create project folder

```bash
mkdir go_project
cd go_project
```

### Step 2: Initialize Go module

```bash
go mod init go_project
```

---

## 🚀 5. Minimal Working Example

### File: main.go

```go
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go Web Server!")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok","service":"go_project"}`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
```

### Run project:

```bash
go run main.go
```

### Output:

* Browser: Hello World message
* JSON response at /status

---

## 🧪 6. AI Prompt Journal

### Prompt 1

**Prompt:** How do I create a simple web server in Go?

**Outcome:** Learned how to use net/http and routing.

---

### Prompt 2

**Prompt:** How do I return JSON in Go HTTP server?

**Outcome:** Learned about encoding/json and headers.

---

### Prompt 3

**Prompt:** Common Go errors and fixes?

**Outcome:** Learned debugging techniques (port conflicts, PATH issues).

---

## ❗ 7. Common Issues & Fixes

### Issue 1: "go not recognized"

✔ Fix: Install Go and add PATH

### Issue 2: Port already in use

✔ Fix: Change port (e.g. 8081)

### Issue 3: Blank browser page

✔ Fix: Ensure server is running

---

## 📚 8. References

* [https://go.dev/doc/](https://go.dev/doc/)
* [https://pkg.go.dev/net/http](https://pkg.go.dev/net/http)
* [https://go.dev/tour/](https://go.dev/tour/)
* StackOverflow Go section

---

## 📁 9. Project Structure

```
go_project/
│
├── go.mod
├── main.go
└── Toolkit Document.md
```

---

## 🌐 10. Endpoints

| Route   | Description |
| ------- | ----------- |
| /       | Home page   |
| /status | JSON API    |

---

## 🧠 11. Learning Outcome

* Built a working web server in Go
* Learned routing and HTTP handling
* Understood JSON responses in APIs
* Used AI prompts to accelerate learning

---

## 👨‍💻 12. Conclusion

This project demonstrates how AI-assisted learning can help beginners quickly understand backend development using Go. The toolkit provides a foundation for building more advanced APIs and web services.

---

🚀 End of Toolkit Document
