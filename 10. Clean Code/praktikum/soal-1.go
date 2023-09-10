/*
1. Berapa banyak kekurangan dalam penulisan kode tersebut? 5
2. Bagian mana saja terjadi kekurangan tersebut? 
	- Tipe data pada struct User
	- Variabel paa struct UserService
	- Receiver method
	- Penamaan Func dan variabel

3. Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.
*/


package main

/* 
Tipe data pada struct sebelumnya tidak deskriptif dikarenakan User adalah struct yang menampilkan sebuah data personal,
Maka dari itu variabel Username dan Password harusnya menggunakan tipe data String.
*/
type User struct {
    ID       int
    Username string	
    Password string
}

/*
Variabel slice pada struct UserService tidak deskriptif, sebelumnya menggunakan huruf "t" sebagai variabelnya 
yang mungkin akan sulit dipahami karena "t" tidak memiliki makna apapun dalam penulisannya. Maka dari itu variabel
"t" diubah menjadi "Users" agar mudah dipahami yang artinya variabel Users tipe slice yang datanya dari struct User
*/
type UserService struct {
    Users []User
}

/*
Variabel struct yang digunakan dalam receiver fungsi sebelumnya setiap kali method dipanggil salinan struct akan dibuat
dan ini akan mempengaruhi kinerja program. Maka dari itu menggunakan pointer pada receiver adalah pilihan paling tepat
dengan menghindari salinan data setiap method dipanggil dan memungkinkan perubahan data secara langsung serta
efisiensi memori juga mempengaruhi
*/
func (u *UserService) GetAllUsers() []User {
    return u.Users
}


/*
Mengubah penulisan menjadi lebih deskriptif, sebelumnya menggunakan huruf "r" untuk mengindikasikan value
dari sebuah slice. Namun penamaan itu akan membuat bingung dan sulit dimengerti maka dari itu diganti dengan
"user" yang artinya value dari sebuah user. Dan mengganti `return user{}` menjadi `return nil` karena pada struct
User variabel Username dan Password sudah menjadi string serta menghindari kebingungan dikarenakan `nil` sudah pasti
data adalah kosong artinya
*/
func (u *UserService) GetUserByID(id int) *User {
    for _, user := range u.Users {
        if id == user.ID {
            return &user
        }
    }
    return nil
}

/* 
Juga menerapkan kejelasan func dengan menggunakan uppercase, jika ingin bisa di ekspor berarti huruf depannya uppercase
jika tidak huruf depannya di lowercase namun sambungan kata dibelakangnya tetap uppercase, contoh :

getUserByID() agar mudah dibaca

dari pada -> getuserbyid()

*/







