1. <div align="justify"><strong>Clean Architecture</strong><p>Clean Architecture adalah sebuah pendekatan dalam pengembangan perangkat lunak yang bertujuan untuk menjaga struktur dan kualitas kode agar tetap bersih, terpisah, dan mudah dipahami. </p> </div> 

2. <div align="justify"><strong>Semua kode dalam projek akan menjadi independent</strong><p>
  
     * Independent Framework
     * Testable
     * Independent UI
     * Independent Database
     * Independent any external
       
   Dengan begini program menjadi flesibel, jika terjadi perubahan disalah satunya tidak akan mempengaruhi yang lainnya.</p></div>

2. <div align="justify"><strong>Context Golang</strong><p>Package untuk mengelola informasi yang berkaitan dengan eksekusi suatu tugas. Package ini berguna untuk mengendalikan pembatalan, timeout, dan nilai-nilai pada request API. API tanpa implementasi Context ketika melakukan cancel request data biasanya akan tetap terkirim ke database.</p> </div>

3. <div align="justify"><strong>Unit Testing Di Clean Arsitektur</strong><p>Penerapan unit testing biasanya mencakup fungsi-fungsi kecil dari sebuah projek seperti function atau model. Unit testing memerlukan Mocking data atau program tiruan yang menyerupai aslinya, misal database kita tidak perlu melakukan testing ke database karena akan membuang banyak waktu dan memang tujuan unit testing adalah menguji fungsi-fungsi model kodenya. Misal UseCase atau Controller.</p> </div>
