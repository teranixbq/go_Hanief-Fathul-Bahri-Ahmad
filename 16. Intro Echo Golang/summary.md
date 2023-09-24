### 3 poin yang dipelajari dari materi
---

1. <div align="justify"><strong>Third-Party Library Di Golang</strong><p>Banyak framework golang yang berfungsi untuk mempermudah pembuatan sebuah API seperti Echo, GORM.io, Go Kit, Cobra, dan Logrus. Echo dan Gorm adalah framework yang sering digunakan dalam pembuatan API, dikarenakan memupunyai latensi yang cepat sertam mendukung middleware dan dengan menggunakan GORM.io tidak perlu lagi repot membuat query SQL secara manual bahkan bisa melakukan auto migration structure database yang dibuat didalam code.</p> </div>

2. <div align="justify"><strong>Import Module</strong><p>Import module adalah hal penting jika menggunakan framework atau ingin melakukan eksport function tertentu ke file lain.</p><p>contohnya import framework module GORM:

    ```
    import (
       "gorm.io/gorm"
       "gorm.io/driver/mysql"
    )
    ```
    di letakkan diawal code</p>

    ```
    import (
       "praktikum/config"
       "praktikum/controller
    )
    
    ```
    ini ketika kita ingin menggunakan function dari folder lain dalam projek yang dibuat.</p>
</div>

3. <div align="justify"><strong>API dan Response HTTP</strong><p>API adalah suatu mekanisme yang membuat dua program bisa berinteraksi satu sama lain contoh nya client dengan server melalui API. Ketika pembuatan API harus memahami juga macam-macam kode response HTTP :
  
    * **200 OK** : permintaan berhasil dan server mengirimkan data yang diminta.
    * **400 Bad Request** : terjadi dikarenakan klien mengirim data yang tidak lengkap atau tidak sesuai dengan format yang diharapkan oleh server, parameter yang dibutuhkan dalam permintaan tidak disertakan atau salah dan permintaan tidak sesuai dengan aturan yang sudah dibuat
    * **404 Not Found** : sumber daya yang diminta oleh klien tidak dapat ditemukan di server.
    * **500 Internal Server Error** : terjadi kesalahan di pihak server saat mencoba memproses permintaan klien.
    * **401 Unauthorized** : menunjukkan bahwa klien tidak memiliki izin atau autentikasi untuk mengakses suatu sumber.
    * **403 Forbidden** : server menolak permintaan klien untuk mengakses sumber daya tertentu, bahkan jika klien memiliki autentikasi.</p></div>