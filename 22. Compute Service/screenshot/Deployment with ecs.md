### 1. Membuat ECS di Alibaba, dan Implementasi Security Group.
<img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/f4f132a2-4c25-48b1-a8b2-abbeb0feb05d" width="800">

<img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/2b695830-ce9b-4717-95e8-7ed4af9321a9" width="800">

### 2. Melakukan SSH Remote ke ECS (Dengan Key) serta dijelaskan key dan password.

  * Membuat Key Pair terlebih dahulu, jika sudah akan download otomatis
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/23478e1b-0e82-4560-86e9-0420763a97e5" width="800">
  * Simpan kunci ssh di folder yang aman
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/241adc22-424f-46d9-8756-eaa4f0af69d4" width="800">
  * Connect remote ssh dengan cara `ssh -i <path kunci key berada> <username>@<ippublic instansce>`
    <img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/a0ff3ace-8e35-4993-a33f-28940e269f23" width="800">
    
### 3. Deploy your Program to ECS
#### Melakukan pull docker dari repositori dan melakukan running dengan mysql lokal instansce
<img src="https://github.com/teranixbq/go_Hanief-Fathul-Bahri-Ahmad/assets/66883583/0ad13cda-2168-45ed-991f-8087326b1a2f" width="800">