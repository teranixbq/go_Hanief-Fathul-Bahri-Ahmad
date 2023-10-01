### poin yang dipelajari dari materi
---

1. <div align="justify"><strong>Unit Testing</strong><p>Testing adalah proses analisis untuk mendeteksi perbedaan antara kondisi saat ini dan kondisi yang diinginkan oleh pengguna. Tujuannya adalah untuk memastikan bahwa sistem berjalan dengan baik dan bebas dari bug.</p></div>


2. <div align="justify"><strong>Tujuan Unit Testing</strong><p>
  
   * Mengubah kode tanpa mengubah fungsionalitasnya
   * Memikirkan berbagai kasus yang mungkin terjadi
   * Membuat dokumentasi yang baik
   * Memastikan bahwa kode dapat dipahami dengan mudah
     
</p>
</div>

3. <div align="justify"><strong>Level Testing</strong><p>
  
   * **UI Testing**
     
     Pengujian yang berfokus pada pengujian antarmuka pengguna (UI). Termasuk menguji elemen-elemen grafis, tampilan, dan interaksi pengguna dengan aplikasi.
     
   * **Integration Testing**

     Pengujian yang fokus pada pengujian interaksi antara komponen-komponen yang telah diuji secara terpisah
     
   * **Unit Testing**
   
     Pengujian yang fokus pada pengujian unit individu. Biasanya pada fungsi, metode, atau kelas kelas kecil yang bisa diuji secara terpisah.

</p>
</div>

4. <div align="justify"><strong>Beberapa kode yang tidak perlu dilakukan testing</strong><p>
  
   * 3rd party modules
   * Database
   * 3rd party api or other external
   * System
   * Object class
</p>
</div>

5. <div align="justify"><strong>Cara melakukan testing</strong><p>

    * Membuat file testing dengan nama file `NamaController_test.go`
    * Buat penamaan function `func TestNamaController(t *testing.T){}`
    * Melakukan testing :

      Ada banyak cara untuk melakukan testing, Jika ingin melakukan testing pada seluruh test file yang berada di satu folder :
      ```
      go test -v
      ```
      Jika ingin melakukan test pada function tertentu gunakan `-run=Namafunction` :
      ```
      go test -v -run=TestUserController
      ```
      Jika ingin menampilkan coverage dari hasil test tambahkan `-cover` :
      ```
      //coverage dari semua test
      go test -v -cover
      ```
      Jika ingin menampilkan output dari coveragenya bisa gunakan :
      ```
      go test -coverprofile=cover.out && go tool cover  -html=cover.out
      ```
      Jika ingin menjadikannya html bisa gunakan :
      ```
      go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html && open cover.html
      ```

</p></div>