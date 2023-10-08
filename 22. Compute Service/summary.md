### poin yang dipelajari dari materi
---

1. <div align="justify"><strong>Deploy Program To Cloud</strong><p>Melakukan deploy dicloud membutuhkan instansce atau server yang disewa di berbagai platform pilihan seerti AWS, Google Cloud Platform, Alibaba Cloud, digitalocean, penyedia hosting vps dan lainnya. cara melakukan deploy hampir sama karena konfigurasi dari setiap os pasti sama.<p>Hal pertama ketika sudah membuat instansce adalah mengatur security group atau firewall untuk mengatur lalulintas yang diinginkan. lalu mengatur ssh untuk remote, di Google cloud kita harus mengenerate kode itu sendiri di linux lokal kita menggunakan perintah `ssh-keygen` atau `ssh-keygen -t ecdsa` atau apapun itu jenis keygennya yang mendukung Google Cloud. Sedangkan di Alibaba cloud kita hanya perlu melakukan generate otomatis keygennya didalam console dasboard key pair lalu mendownloadnya ke linux lokal.</p><p>Untuk cara melakukan remote instansce semuanya sama dari berbagai platform menggunakan perintah 
  ```
  ssh -i <path kode keygen berada> <usernameinstansce>@<ippublic instansce>
  ```
</p></div>

2. <div align="justify"><strong>Database</strong><p>Banyak cara untuk melakukan penyimpanan database untuk program yang akan kita running, bisa menggunakan Cloud SQl, RDS aws, Polardb, atau penyedia hosting mysql. Jika ingin bisa menggunakan docker mysql atau menggunakan mysql server di instansce itu sendiri. Untuk melakukan koneksi dengan program kita sama seperti biasa hanya perlu mengubah settingannya dari penyedia database seperti ip hostnya dan mengatur network mana yang akan disambungkan dengan database</p>
</div>

3. <div align="justify"><strong>Pointing Domain</strong><p>Proses mengarahkan atau menghubungkan sebuah nama domain (URL) dengan alamat IP dari server. kita jadi bisa mengakses situs web dengan mengetikkan nama domain dalam browser tanpa harus mengingat nama IP nya.</p><p>Cara yang populer dengan menggunakan apache2 atau nginx, yang saya gunakan adalah nginx. Cara installnya : 
  
    * Install Nginx
      ```
      sudo apt update
      sudo apt install nginx
      ```
    * Membuat settingan server,di folder /etc/nginx/sites-available/<namadomain>.conf
      ```
      server {
          listen 80;
          server_name 8.222.232.160; # Ganti dengan alamat IP publik instanasce
      
          location / {
              proxy_pass http://api.kodeteks.com; # Ganti dengan domain
              proxy_set_header Host $host;
              proxy_set_header X-Real-IP $remote_addr;
          }
      }
      ```
    * Restart Nginx
      ```
      sudo ln -s /etc/nginx/sites-available/api.kodeteks.com.conf /etc/nginx/sites-enabled/
      sudo systemctl restart nginx
      ```
    * Melakukan koneksi dns dari domain ke ip server

      <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/b79acea4-734c-4a8a-b4d1-300ada20373d" width="700">

    * Hasilnya http://api.kodeteks.com:8080/users
</p>

<p>Alternatif Pointing Domain dengan Ngrok :
  
  * Install Ngrok di https://dashboard.ngrok.com/get-started/setup
  * lalu ekstrak file ke direktori yang diinginkan
  * masuk ke difolder yang sudah diekstrak tadi dan lakukan perintah, jangan lupa menambahkan token yang ada di dasboard Ngrok
    ```
    ./ngrok config add-authtoken <token ngrok kalin>
    ```
  * Jika sudah maka hasilnya sepert ini, akan otmatis disimpan di path .config/ngrok
    ```
    root@teranix:~/bind# ./ngrok config add-authtoken 2WMBaRivCfuo8q6A2viVsCIuv74_6iYJBpyF91dEwXqbEENcA
    Authtoken saved to configuration file: /root/.config/ngrok/ngrok.yml
    root@teranix:~/bind# 
    ```
  * Lalu setting ngrok.yaml untuk melakukan port mana yang akan di publish
    ```
    version: "2"
    authtoken: 2WMBaRivCfuo8q6A2viVsCIuv74_6iYJBpyF91dEwXqbEENcA
    tunnels:
       api1:
         proto: http
         addr: 8080
    ```
    bagian version dan authtoken sudah ada default dari hasil generat, jadi hanya perlu menambahkan tunnels.
  * Start Ngrok
    start ngrok di folder hasil ekstrak sebelumnya dengan melakukan perintah ./ngrok start (nama tunnels kalian)
    ```
    ./ngrok start api1
    ```
    akan otomatis menampilkan link public yang bisa diakses

  * Untuk bisa berjalan di cloud secara terus menerus bisa menggunakan modul tambahan menggunakan tmux, lakukan install tmux terlebih dahulu
    ```
    sudo apt tmux
    ```
  * kembali ke folder ekstrak ngrok dan lakukan pemanggilan start ngrok dengan perintah tmux
    ```
    tmux new-session -d -s ngrok './ngrok start api1'
    ```
    untuk menghentikan sesi gunakan perintah :
    ```
    tmux detach -s ngrok
    ```

</p></div>