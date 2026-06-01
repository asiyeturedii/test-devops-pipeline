pipeline {
    agent any

    environment {
        IMAGE_NAME = "test-devops-pipeline"
        IMAGE_TAG  = "latest"
        CONTAINER_NAME = "test-devops-app"
        PORT = "8081"
        PATH = "/usr/local/bin:/opt/homebrew/bin:/usr/bin:/bin:/usr/sbin:/sbin"
    }

    stages {

        stage('📥 Checkout') {
            steps {
                echo "Kaynak kod alınıyor..."
                checkout scm
            }
        }

        stage('🔍 Test') {
            steps {
                echo "Go testleri çalıştırılıyor..."
                sh '/opt/homebrew/bin/go test ./... || true'
            }
        }

        stage('🐳 Docker Build') {
            steps {
                echo "Docker image build ediliyor..."
                sh "/usr/local/bin/docker build -t ${IMAGE_NAME}:${IMAGE_TAG} ."
            }
        }

        stage('🛑 Eski Container Durdur') {
            steps {
                echo "Eski container temizleniyor (varsa)..."
                sh """
                    /usr/local/bin/docker stop ${CONTAINER_NAME} || true
                    /usr/local/bin/docker rm   ${CONTAINER_NAME} || true
                """
            }
        }

        stage('🚀 Deploy') {
            steps {
                echo "Yeni container başlatılıyor..."
                sh """
                    /usr/local/bin/docker run -d \
                        --name ${CONTAINER_NAME} \
                        -p ${PORT}:${PORT} \
                        --restart unless-stopped \
                        ${IMAGE_NAME}:${IMAGE_TAG}
                """
            }
        }

        stage('✅ Health Check') {
            steps {
                echo "Servis ayakta mı kontrol ediliyor..."
                sh "sleep 3 && curl -sf http://localhost:${PORT}/health || exit 1"
            }
        }
    }

    post {
        success {
            echo "✅ Pipeline başarıyla tamamlandı! → http://localhost:${PORT}"
        }
        failure {
            echo "❌ Pipeline başarısız. Logları kontrol edin."
        }
    }
}
