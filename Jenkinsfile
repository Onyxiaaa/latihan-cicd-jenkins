pipeline {
    agent any
    stages {
        stage('Build Image') {
            steps {
                // Build Docker Image dari folder saat ini (.)
                sh 'docker build -t go-image .'
            }
        }
        
        stage('Deploy Container') {
            steps {
                // Hapus container lama
                sh 'docker stop go-container || true'
                sh 'docker rm go-container || true'
                
                // Jalankan Container Baru
                sh 'docker run -d -p 3000:8080 --name go-container go-image'
            }
        }
    }
}
