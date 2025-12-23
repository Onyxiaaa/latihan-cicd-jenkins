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
                    snykTokenId: 'SNYK-ANOM'
                )
            }
        } 
        
        stage('DAST Scan (OWASP ZAP)') {
            steps {
                script {
                    // FIX PERMISSIONS:
                    // 1. Buat file kosong dulu biar owned by Jenkins user
                    sh 'touch zap_report.html'
                    // 2. Beri izin baca/tulis ke semua user (biar container docker bisa nulis)
                    sh 'chmod 777 zap_report.html'
                    
                    // JALANKAN ZAP:
                    // Tambahkan parameter -u 0 (run as root) jika masih gagal, 
                    // tapi trik chmod di atas biasanya sudah cukup.
                    sh 'docker run --rm -v $(pwd):/zap/wrk/:rw -t ghcr.io/zaproxy/zaproxy:stable zap-baseline.py -t http://ip-aplikasi-anda:3000 -r zap_report.html || true'
                }
            }
        }   
        
        stage('Publish Report') {
            steps {
                // Pastikan Plugin "HTML Publisher" sudah diinstall sebelum jalanin ini
                publishHTML (target: [
                    allowMissing: false,
                    alwaysLinkToLastBuild: false,
                    keepAll: true,
                    reportDir: '.',
                    reportFiles: 'zap_report.html',
                    reportName: 'ZAP Security Report'
                ])
            }
        }
    }
}
