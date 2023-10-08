### 1. Membuat settingan koneksi dengan domain
#### setting nginx di `/etc/nginx/sites-available/<namadomain>.conf`
  * Setting koneksi
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/5e8297a7-587a-4fbb-a6ad-cd7b754db76f" width="800">
  * Lalu setting DNS domain ke ip publik instansce
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/e46caf23-61bf-44c4-abbb-b287966e1b2c" width="800">
  * Hasilnya dengan `api.kodeteks.com:8080/users`
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/26a4a147-6d51-4bba-8fa1-290ca640824a" width="800">

### 2. Alternatif lain dengan Ngrok
#### Download dan install [Ngrok](https://dashboard.ngrok.com/get-started/setup)
  * Setting ngrok.yaml, dan pastikan sudah melakukan running docker 
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/776c22c6-b22d-45fc-8731-ceeb4ccc40b3" width="800">
    
  * Lalu melakukan running ngrok menggunakan tmux dengan perintah `tmux new-session -d -s ngrok './ngrok start api1'`
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/95d94b37-c091-45c7-bbf2-b79ce20366c8" width="800">
    
  * Melihat Link Ngrok dari dasboard
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/f5e856fb-3254-4876-a471-52f0053e479a" width="800">

  * Hasilnya dengan `https://1941-8-222-232-160.ngrok-free.app/users`
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/455ea6ee-368b-4a0e-8a06-ff00580752b8" width="800">

