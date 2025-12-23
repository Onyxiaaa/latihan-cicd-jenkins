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
                    // Jalankan ZAP dalam container, scan target, lalu mati
                    // -t: Target URL
                    // -r: Nama file laporan
                    // JANGAN LUPA: Ganti 'ip-aplikasi-anda' dengan IP Public atau Internal VM Anda
                    sh 'docker run --rm -v $(pwd):/zap/wrk/:rw -t ghcr.io/zaproxy/zaproxy:stable zap-baseline.py -t http://ip-aplikasi-anda:3000 -r zap_report.html || true'
                }
            }
        }   
        
        stage('Publish Report') {
            steps {
                // Menampilkan laporan di Dashboard Jenkins
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
