pipeline {
    agent any
    
    stages {
        // --- TAHAP 1: KUALITAS KODE (Unit Test) ---
        stage('Unit Test') {
            steps {
                script {
                    echo 'Running Unit Tests...'
                    // Menggunakan image golang:alpine sementara untuk test logika kode
                    // Pipeline akan GAGAL jika unit test error
                    sh 'docker run --rm -v $(pwd):/app -w /app golang:alpine go test -v ./...'
                }
            }
        }

        // --- TAHAP 2: BUILD & DEPLOY ---
        stage('Build Image') {
            steps {
                echo 'Building Docker Image...'
                sh 'docker build -t go-image .'
            }
        }
        
        stage('Deploy Container') {
            steps {
                echo 'Deploying Application...'
                // Bersihkan container lama agar tidak bentrok nama
                sh 'docker stop go-container || true'
                sh 'docker rm go-container || true'
                
                // Jalankan container baru di port 3000
                sh 'docker run -d -p 3000:8080 --name go-container go-image'
            }
        }

        // --- TAHAP 3: KEAMANAN (SAST - Snyk) ---
        stage('Security SAST (Snyk)') {
            steps {
                echo 'Scanning Code & Dependencies...'
                // Pastikan ID Credential sesuai dengan yang ada di Jenkins Anda
                snykSecurity(
                    snykInstallation: 'snyk@latest',
                    snykTokenId: 'SNYK-ANOM',
                    severity: 'high', // Hanya lapor jika ada High/Critical severity
                    failOnIssues: false // Set true jika ingin pipeline stop saat ada virus
                )
            }
        } 
        
        // --- TAHAP 4: KEAMANAN (DAST - OWASP ZAP) ---
        stage('Security DAST (OWASP ZAP)') {
            steps {
                script {
                    echo 'Scanning Running Application...'
                    // Fix Permission: Buat file laporan kosong & beri izin tulis
                    sh 'touch zap_report.html'
                    sh 'chmod 777 zap_report.html'
                    
                    // Jalankan ZAP Baseline Scan
                    // GANTI 'ip-aplikasi-anda' DENGAN IP PUBLIC/PRIVATE SERVER ANDA
                    sh 'docker run --rm -v $(pwd):/zap/wrk/:rw -t ghcr.io/zaproxy/zaproxy:stable zap-baseline.py -t http://34.101.251.163:3000 -r zap_report.html || true'
                }
            }
        }   
        
        stage('Publish DAST Report') {
            steps {
                // Menampilkan laporan ZAP di Dashboard Jenkins
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

        // --- TAHAP 5: PERFORMA (k6 Load Test) ---
        stage('Performance Test (k6)') {
            steps {
                script {
                    echo 'Running Load Test...'
                    // Membuat script test k6 sederhana (10 user selama 10 detik)
                    sh '''
                    echo "import http from 'k6/http';
                    import { sleep } from 'k6';
                    export let options = {
                        vus: 10,
                        duration: '10s',
                    };
                    export default function () {
                        // GANTI IP DI BAWAH INI JUGA
                        http.get('http://34.101.251.163:3000');
                        sleep(1);
                    }" > loadtest.js
                    '''
                    
                    // Menjalankan k6 via Docker
                    sh 'docker run --rm -v $(pwd):/scripts -i grafana/k6 run /scripts/loadtest.js || true'
                }
            }
        }
    }
}
