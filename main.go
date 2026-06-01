package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <title>Mini BI Pipeline</title>
    <style>
        body {
            font-family: 'Segoe UI', sans-serif;
            background: #0f172a;
            color: #e2e8f0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .card {
            background: #1e293b;
            border: 1px solid #334155;
            border-radius: 12px;
            padding: 40px 60px;
            text-align: center;
            box-shadow: 0 4px 32px rgba(0,0,0,0.4);
        }
        h1 { color: #38bdf8; font-size: 2.5rem; margin-bottom: 8px; }
        p  { color: #94a3b8; margin: 4px 0; }
        .badge {
            display: inline-block;
            background: #0ea5e9;
            color: white;
            border-radius: 20px;
            padding: 4px 16px;
            font-size: 0.85rem;
            margin-top: 16px;
        }
    </style>
</head>
<body>
    <div class="card">
        <h1>🚀 Hoş Geldiniz!</h1>
        <p>Mini BI Pipeline Mikroservisi</p>
        <p>🕐 %s</p>
        <p>🖥️ Host: %s</p>
        <div class="badge">Port 8081 · Go + Docker + Jenkins</div>
    </div>
</body>
</html>`, time.Now().Format("2006-01-02 15:04:05"), hostname)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, html)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok","service":"mini-bi-pipeline"}`)
}

func main() {
	port := "8081"

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("✅ Servis başlatıldı → http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Servis başlatılamadı: %v", err)
	}
}
