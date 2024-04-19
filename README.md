# go-practice

## Go-Introduction
``` ### Data types, variable, condition, loopings ```
``` Challange : SquareCube
Buatlah looping berkondisi dengan variable n sebagai penentu maksimal number loopingnya (Iterasi dari 1 sampai n). Input n dinamis, tidak statis/hardcode.

Jika angka tersebut adalah kuadrat sempurna, cetak "Square".

Jika angka tersebut adalah kubus sempurna, cetak "Cube".

Jika angka tersebut adalah kuadrat dan kubus sempurna (contoh: 64), cetak "SquareCube".

Jika tidak, cetak angka tersebut.

Notes:
1. Kuadrat sempurna: Hasil dari pengkuadratan bilangan bulat. Dengan kata lain, sebuah bilangan adalah kuadrat sempurna jika ia bisa ditulis sebagai kuadrat dari bilangan bulat lainnya. Contoh, 16 adalah kuadrat sempurna karena 4² = 16. Demikian pula, 25 adalah kuadrat sempurna karena 5² = 25.
2. Kubus sempurna adalah hasil dari perkalian bilangan bulat sebanyak tiga kali (pengkubusan). Sebuah bilangan adalah kubus sempurna jika ia bisa ditulis sebagai kubus dari bilangan bulat lainnya. Contoh, 27 adalah kubus sempurna karena 3³ = 27.
```
## Fundamental-1
``` ### Data types collections ```
```Personal Notes: Komputer menyimpan suatu data pada memori, setiap memori memiliki alamat tersendiri. Pointer pada Go memberikan akses bagi pengguna untuk memanipulasi memori. Ketika ada Variabel A dengan Value 1 disimpan pada memori 0xc00001, kita bisa membuat sebuah Variabel PointerA dengan Value berisi alamat memori Variabel A. Sehingga, ketika kita memanggil PointerA, akan mengembalikan sebuah alamat: 0xc00001. Ampersan (*) digunakan untuk melihat nilai yang tersimpan pada suatu alamat memori. Sehingga ketika kita memanggil *PointerA akan mengembalikan nilai dari Variabel A. Karena kita sedang menunjuk alamat memori dari 0xc00001 dari value yang tersimpan pada Variabel PointerA. ```
``` Challange count the character
Buatlah looping dengan variable yang berisi string suatu kalimat dan pecahlah kalimat tersebut menjadi 1 per 1 kata

Setelah sudah dipecah, lakukan perhitungan munculnya kata dari variable tersebut dengan cara mapping golang
```

## Fundamental-2
``` ### Struct, Func, Method, Reflect, Export-Unexported ```
```
Buatlah sebuah service berupa CLI untuk menampilkan data teman-teman kalian di kelas.

Contohnya, ketika kalian menjalankan perintah “go run main.go Fitri” maka data yang akan muncul adalah data teman kalian dengan nama "Fitri". Data yang harus ditampilkan yaitu: 
Nama
Alamat
Pekerjaan
Alasan memilih kelas Golang

Gunakanlah struct dan function untuk menampilkan data tersebut. Kalian bisa menggunakan os.Args untuk mendapatkan argument pada terminal.

Flow teknis coding :
Buat logic untuk menampilkan/generate data peserta dari cli by name, contoh cli : go run main.go Fitri
Buatlah looping yang memproses slice of string dari beberapa nama peserta, dimana jika nama yg di ketik pada cli sama dengan yang ada di data peserta maka akan dapat index dari data pesertanya dan di append pada data baru untuk ditampilkan pada CLI atau terminal.
Jika data nama atau absen yang diinput pada CLI tidak tersedia pada program, misal go run main.go Fitri atau go run main.go 1 maka bisa dihandle dengan pesan "Data dengan nama/absen tsb tidak tersedia".
```

## WEB-SERVER
``` ### net/http, gin ```
```
Buatlah login dan halaman biodata dengan data statis yang sudah diatur pada backend golang kalian. Dan buatlah kondisi pada backend nya dimana data statis akan tergenerate pada halaman index sesuai dengan email yang diinput pada saat form login.
```

## FINAL-PROJECT
``` LINK DEPLOYMENT RAILWAY ADA PADA COMMENT CLASSROOM ```
```
Postman Environtmant: BASIC API - FAHRI
base_url
access_token
```