# 🚀 Mini BI Pipeline

Go + Docker + Jenkins ile mini CI/CD pipeline.

---

## 📁 Dosya Yapısı

```
mini-bi-pipeline/
├── main.go        ← Go mikroservis
├── go.mod         ← Go modül tanımı
├── Dockerfile     ← Multi-stage Docker build
├── Jenkinsfile    ← CI/CD pipeline tanımı
└── README.md
```

---

## ▶️ 1. Lokal Çalıştırma (Docker'sız)

```bash
go run main.go
# → http://localhost:8081
```

---

## 🐳 2. Docker ile Çalıştırma

```bash
# Build
docker build -t mini-bi-pipeline .

# Çalıştır
docker run -d --name mini-bi-app -p 8081:8081 mini-bi-pipeline

# Kontrol
curl http://localhost:8081/health

# Loglar
docker logs mini-bi-app

# Durdur
docker stop mini-bi-app && docker rm mini-bi-app
```

---

## 🔧 3. Jenkins Kurulumu

### Jenkins'i Docker ile başlat:

```bash
docker run -d \
  --name jenkins \
  -p 8080:8080 \
  -p 50000:50000 \
  -v jenkins_home:/var/jenkins_home \
  -v /var/run/docker.sock:/var/run/docker.sock \
  jenkins/jenkins:lts
```

### İlk şifreyi al:
```bash
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

### Jenkins'te Pipeline oluşturma adımları:

1. `http://localhost:8080` → Jenkins aç
2. **New Item** → isim ver → **Pipeline** seç
3. **Pipeline Definition:** `Pipeline script from SCM` seç
4. **SCM:** Git → repo URL'ini gir
5. **Script Path:** `Jenkinsfile` (default)
6. **Save** → **Build Now**

> 💡 **Jenkins'in Docker'a erişmesi için:**
> ```bash
> docker exec -u root jenkins chmod 666 /var/run/docker.sock
> ```

---

## 🌐 Endpointler

| URL | Açıklama |
|-----|----------|
| `http://localhost:8081/` | Hoş geldiniz sayfası |
| `http://localhost:8081/health` | Health check (JSON) |
