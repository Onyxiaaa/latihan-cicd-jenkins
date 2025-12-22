pipeline {
    agent any
    // environment {
    //     // Mengambil token dari dompet Jenkins dan menyimpannya sebagai variable lingkungan
    //     // SNYK_TOKEN = credentials('snyk-api-token')
    //     // SNYK_TOKEN = credentials('SNYK-ANOM')
    // }
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

        stage('Test-SNYK') {
          steps {
            echo 'Testing...'
            snykSecurity(
              snykInstallation: 'snyk@latest',
              snykTokenId: 'SNYK-ANOM',
            )
          }
        }

        stage('Security Scan (Snyk)') {
            steps {
                script {
                    // 1. Scan Dependencies (Library pihak ketiga)
                    // --severity-threshold=high artinya cuma lapor kalau ada bahaya TINGGI
                    sh 'snyk test --severity-threshold=high || true' 
                    
                    // Catatan: "|| true" di akhir berguna agar Pipeline TIDAK BERHENTI/GAGAL 
                    // meskipun ketemu virus. Kalau mau pipeline gagal jika ada virus, hapus "|| true".

                    echo 'Scanning Docker Image...'
                    // 2. Scan Docker Image (Ganti nama image sesuai project Anda)
                    // Pastikan image sudah di-build sebelumnya
                    // sh 'snyk container test nama-image-anda:latest || true'
                }
            }
        }
    }
}
