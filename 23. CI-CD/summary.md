### 3 poin yang dipelajari dari materi
---

1. <div align="justify"><strong>Strategi Deployment</strong><p>

    * **Big Bang Deployment**
      
      Big Bang Deployment adalah strategi di mana seluruh versi perangkat lunak baru diterapkan secara sekaligus pada seluruh infrastruktur produksi mengganti seluruh sistem sekaligus dengan yang baru
    * **Rolling Deployment**
      
      Rolling Deployment adalah metode dengan versi perangkat lunak baru diterapkan secara bertahap (satu persatu) pada sebagian kecil dari infrastruktur produksi.
    * **Canary Deployment**
      
      Canary Deployment adalah cara untuk menerapkan versi perangkat lunak baru dengan menguji secara bertahap pada bagian kecil pengguna atau server.
    * **Blue-Green Deployment**
      
      Blue-Green adalah pendekatan untuk memperbarui aplikasi tanpa downtime. dengan membuat dua lingkungan: Blue (yang sedang berjalan) dan Green (yang baru). Aplikasi diuji di Green sebelum lalu lintas dialihkan dari Blue ke Green
</p></div>

2. <div align="justify"><strong>CI/CD (Continuous Integration/Continuous Deployment)</strong><p>Praktik pengembangan perangkat lunak untuk memngerjakan tugas secara otomatis. Continuous Integration (CI) adalah proses otomatisasi penggabungan dan pengujian kode yang baru, sedangkan Continuous Deployment (CD) adalah proses otomatisasi penyebaran aplikasi ke produksi setelah melakukan uji coba. CI/CD membantu mempercepat siklus pengembangan, meningkatkan kualitas perangkat lunak, dan mengurangi risiko kesalahan saat development. <p>Misal ingin melakukan deploy code ke github sekaligus membuat docker image lalu melakukan push ke registry dan mendeploy ke server secara otomatis menggunakan github action yang memungkin kan melakukan semunya dalam satu konfigurasi file. Pada github kita harus menulis kode konfigurasi pada folder ( .github/workflows ) agar github action bisa mengeksekusinya</p></p>
</div>

3. <div align="justify"><strong>Kubernetes</strong><p>Kubernetes adalah platform open source untuk mengelola kumpulan kontainer dalam suatu cluster server. Kontainer sendiri adalah environment dengan sumber daya, CPU, dan sistem file untuk satu aplikasi. didalam kubernetes ada beberapa komponen-komponen yaitu diantaranya : </p><p>
  
    * **Namespace**
    
      Namespace mengelompokkan dan mengisolasi sumber daya dalam Kubernetes.
    * **Pod**
    
      Pod adalah unit terkecil yang berisi satu atau beberapa kontainer dan menjalankan aplikasi.
    * **Label**
    
      Label adalah penanda kunci-nilai untuk mengorganisasi dan seleksi sumber daya.
    * **Deployment**
    
      Deployment mengelola aplikasi, termasuk pembaruan bertahap (rolling updates) dan rollback.
    * **Rolling Updates**
    
      Rolling updates adalah pembaruan aplikasi secara bertahap.
    * **Secrets**
    
      Secrets adalah objek untuk menyimpan informasi sensitif.
    * **Service**
    
      Service adalah abstraksi untuk mengakses aplikasi dan menyediakan load balancing.
    * **Ingress**
    
      Ingress mengatur lalu lintas masuk eksternal ke dalam cluster.
</p></div>