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
``` Challange count the character
Buatlah looping dengan variable yang berisi string suatu kalimat dan pecahlah kalimat tersebut menjadi 1 per 1 kata

Setelah sudah dipecah, lakukan perhitungan munculnya kata dari variable tersebut dengan cara mapping golang

Personal Notes: Komputer menyimpan suatu data pada memori, setiap memori memiliki alamat tersendiri. Pointer pada Go memberikan akses bagi pengguna untuk memanipulasi memori. Ketika ada Variabel A dengan Value 1 disimpan pada memori 0xc00001, kita bisa membuat sebuah Variabel PointerA dengan Value berisi alamat memori Variabel A. Sehingga, ketika kita memanggil PointerA, akan mengembalikan sebuah alamat: 0xc00001. Ampersan (*) digunakan untuk melihat nilai yang tersimpan pada suatu alamat memori. Sehingga ketika kita memanggil *PointerA akan mengembalikan nilai dari Variabel A. Karena kita sedang menunjuk alamat memori dari 0xc00001 dari value yang tersimpan pada Variabel PointerA.
```
