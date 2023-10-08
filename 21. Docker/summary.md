### 3 poin yang dipelajari dari materi
---

1. <div align="justify"><strong>Docker</strong><p>Docker adalah platform yang memungkinkan untuk mengemas, menjalankan, dan mendistribusikan aplikasi dan layanan dalam container ringan dapat diisolasi. container Docker bisa mengemas kode, dependensi, dan konfigurasi ke dalam unit yang dapat dijalankan di berbagai lingkungan. sehingga mudah diatur dan juga dipindahkan<p>Cara install Docker via ubuntu Apt :

    * Set up Docker's Apt repository
      ```
      # Add Docker's official GPG key:
      sudo apt-get update
      sudo apt-get install ca-certificates curl gnupg
      sudo install -m 0755 -d /etc/apt/keyrings
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
      sudo chmod a+r /etc/apt/keyrings/docker.gpg
      
      # Add the repository to Apt sources:
      echo \
        "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
        "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
        sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
      sudo apt-get update
      ```

   * Install
      ```
      sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
      ```
</p>
</p></div>

2. <div align="justify"><strong>Docker Compose</strong><p>Docker Compose adalah alat yang membuat program docker mudah untuk menjalankan beberapa program dalam container Docker sekaligus, dengan konfigurasi yang diatur dalam satu berkas. yaitu pada berkas docker-compose.yml<p>Cara install Docker Compose :

    * Linux Ubuntu
      ```
      sudo apt-get install docker-compose-plugin
      ```
    * Melakukan running docker compose
      ```
      docker-compose up
      ```
    * Menghentikan running docker compose
      ```
      docker-compose down
      ```
</p></p>
</div>

3. <div align="justify"><strong>Perintah-perintah dasar pada docker</strong><p>Perintah Image : 
  
      * docker pull: Mengunduh image Docker dari Docker Hub.
      * docker images/docker image ls: Menampilkan image Docker yang ada di sistem.
      * docker rmi: Menghapus image Docker.
      * docker build: Membuat image Docker dari Dockerfile.
   </br>
   
   Perintah Container :
      * docker run: Membuat dan menjalankan kontainer dari image Docker.
      * docker ps: Menampilkan container yang sedang berjalan.
      *  docker ps -a: Menampilkan semua container, termasuk yang sudah berhenti.
      * docker start: Menjalankan container yang berhenti.
      * docker stop: Menghentikan container yang berjalan.
      * docker restart: Mengulang container.
      * docker rm: Menghapus container.
      * docker exec: Menjalankan perintah di dalam kontainer yang sedang berjalan.
   </br>
   
   Perintah Lainnya :
      * docker --version: Menampilkan versi Docker yang terpasang.
      * docker-compose version: Menampilkan versi Docker compose yang terpasang.
      * docker login: Masuk ke akun Docker Hub.
      * docker logout: Keluar dari akun Docker Hub.

</p></div>