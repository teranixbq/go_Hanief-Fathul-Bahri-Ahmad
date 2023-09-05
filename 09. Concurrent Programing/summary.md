### 3 poin yang dipelajari dari materi
---

1. <div align="justify"><strong>Goroutine</strong><p>Go menggunakan goroutine sebagai unit konkurensi. Goroutine sangat ringan, sehingga Anda dapat membuat ribuan goroutine dalam program Anda tanpa mengalami overhead yang signifikan. Goroutine dapat dibuat dengan menggunakan kata kunci "go" sebelum pemanggilan fungsi. Memungkinkan eksekusi fungsi tersebut secara konkuren tanpa harus menunggu selesai.</p> </div>

2. <div align="justify"><strong>Channel</strong><p>Channel adalah mekanisme komunikasi untuk berbagi data antara goroutine. Channel dapat menghindari masalah race condition dan deadlock. Membuat channel dengan menggunakan tanda panah (<-chan) untuk membaca dari channel atau (chan<-) untuk menulis ke channel. Jenis Channel ada dua yaitu;
* Unbuffered Channel: Tidak memiliki buffer, data dikirim dan diterima secara langsung. Pengirim akan terblokir jika tidak ada penerima yang siap
* Buffered Channel: Memiliki buffer dengan kapasitas tertentu, sehingga pengirim dapat mengirim beberapa data sebelum penerima harus membacanya.</p>
</div>

3. <div align="justify"><strong>Sync Package</strong><p>Paket "sync" di Go menyediakan alat-alat untuk mengkoordinasikan eksekusi goroutine, seperti WaitGroup, Mutex, dan RWMutex.WaitGroup digunakan untuk menunggu sekelompok goroutine selesai sebelum melanjutkan eksekusi program. Mutex dan RWMutex digunakan untuk mengamankan akses bersama ke data yang digunakan oleh goroutine.</p>
</div>