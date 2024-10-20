<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL IV</strong></h2>
<h2 align="center"><strong> PROSEDUR </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  LAILATUR RAHMAH / 2311102177<br>
  S1 IF 11 06
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Abednego Dwi Septiadi
  
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## DAFTAR ISI

- [DASAR TEORI](#dasar-teori)
- [GUIDED](#guided)
- [UNGUIDED](#unguided)

## DASAR TEORI

Prosedur (Fungsi) dalam Go

Prosedur adalah elemen dasar dalam pemrograman terstruktur yang berfungsi untuk membagi program menjadi bagian-bagian yang lebih kecil dan terorganisir. Dengan membagi kode menjadi prosedur, program dapat lebih mudah dikelola karena setiap prosedur atau fungsi memiliki tanggung jawab yang jelas dan spesifik. Hal ini tidak hanya membuat kode lebih rapi, tetapi juga meningkatkan keterbacaan dan pemahaman bagi pengembang lain yang mungkin bekerja pada kode yang sama.

Dalam bahasa pemrograman Go, istilah "prosedur" diubah menjadi "fungsi" (function). Fungsi di Go adalah blok kode yang dirancang untuk melakukan tugas tertentu. Dengan mendefinisikan logika program dalam fungsi, Anda dapat:
  1. **Modularisasi Kode:** Fungsi memungkinkan pemisahan kode menjadi modul-modul kecil yang dapat dipanggil kapan saja. Setiap fungsi dapat menangani satu bagian dari tugas, sehingga memperkecil kompleksitas yang dihadapi pengembang.
  2. **Reusabilitas:** Setelah Anda menulis fungsi, Anda dapat menggunakannya di berbagai bagian program tanpa perlu menulis ulang kode tersebut. Ini mengurangi redundansi dan memudahkan pemeliharaan.
  3. **Pengujian:** Fungsi dapat diuji secara terpisah. Dengan menguji setiap fungsi secara individual, Anda dapat dengan cepat menemukan dan memperbaiki kesalahan. Ini juga mendukung pengujian unit, yang merupakan praktik baik dalam pengembangan perangkat lunak.
  4. **Peningkatan Kolaborasi:** Dalam proyek besar, beberapa pengembang dapat bekerja pada fungsi yang berbeda secara bersamaan. Ini memudahkan kolaborasi dan mempercepat proses pengembangan.
  5. **Abstraksi:** Fungsi memungkinkan pengembang untuk fokus pada apa yang dilakukan kode, bukan bagaimana cara melakukannya. Ini membantu dalam memahami alur logika program tanpa terjebak dalam detail implementasi.
 
Fungsi ini memiliki komponen berikut:

- Nama Fungsi : Digunakan untuk mengidentifikasi fungsi dan memanggilnya.
- Parameter : Input yang diterima oleh fungsi untuk diproses. Fungsi dapat menerima beberapa parameter.
- Tipe Pengembalian : Nilai yang dikembalikan oleh fungsi. Fungsi dapat mengembalikan satu atau lebih nilai.
- Badan Fungsi : Kode yang dijalankan ketika fungsi dipanggil.

Contoh :
```go
func hitungLuasPersegiPanjang(panjang, lebar int) int {
    return panjang * lebar
}

func main() {
    luas := hitungLuasPersegiPanjang(5, 3)
    fmt.Println("Luas:", luas)
}
```

Dalam contoh ini, fungsi `hitungLuasPersegiPanjang` menerima dua parameter bertipe `int`,
dan mengembalikan nilai hasil perkalian panjang dan lebar sebagai integer.

**Parameter dan Argumen**

- Parameter adalah variabel yang dideklarasikan dalam definisi fungsi.
  Ini adalah nilai-nilai yang akan diterima oleh fungsi ketika dipanggil.
  Go memungkinkan penulisan parameter dengan tipe data eksplisit.
- Argumen adalah nilai yang diberikan ke parameter saat memanggil fungsi.
  Argumen ini bisa berupa variabel, konstanta, atau hasil dari ekspresi lain.

Parameter dengan Tipe yang Sama: Go memungkinkan penulisan parameter yang memiliki tipe data yang sama secara singkat.

Contoh :

```go
func tambah(a, b int) int {
    return a + b
}
```

**Pengembalian Nilai (Return Values)**

Salah satu fitur yang menonjol di Go adalah kemampuannya untuk mengembalikan lebih dari satu nilai dari fungsi. 
Ini sangat berguna ketika sebuah fungsi perlu mengembalikan hasil utama dan informasi tambahan (misalnya, status kesalahan).

Contoh :
```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("tidak bisa membagi dengan nol")
    }
    return a / b, nil
}

func main() {
    hasil, err := bagi(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Hasil:", hasil)
    }
}
```

Dalam contoh ini, fungsi `bagi` mengembalikan dua nilai : hasil pembagian dan potensi kesalahan.

**First-Class Function (Fungsi sebagai Objek Kelas Utama)**

Dalam Go, fungsi adalah `first-class citizen`, yang berarti fungsi dapat diperlakukan seperti objek lainnya, seperti variabel. 

Fungsi dapat:

- Disimpan dalam variabel.
- Diteruskan sebagai parameter ke fungsi lain.
- Dikembalikan sebagai hasil dari fungsi.

Contoh :
```go
func operasi(oper func(int, int) int, a int, b int) int {
    return oper(a, b)
}

func main() {
    penjumlahan := func(a, b int) int { return a + b }
    hasil := operasi(penjumlahan, 5, 3)
    fmt.Println(hasil) // Output: 8
}
```
Di sini, fungsi *penjumlahan* didefinisikan sebagai fungsi anonim dan disimpan dalam variabel *penjumlahan*. 
Variabel ini kemudian diteruskan sebagai argumen ke fungsi *operasi*.

**Closure**

Closure adalah fungsi yang "membungkus" lingkungannya, yaitu variabel-variabel dari cakupan luar yang dapat diakses di dalam fungsi tersebut, bahkan setelah lingkungannya selesai dieksekusi.

Contoh :
```go
func buatCounter() func() int {
    hitung := 0
    return func() int {
        hitung++
        return hitung
    }
}

func main() {
    counter := buatCounter()
    fmt.Println(counter()) // Output: 1
    fmt.Println(counter()) // Output: 2
}
```

Di sini, fungsi anonim di dalam `buatCounter` dapat mengakses variabel `hitung`, meskipun `hitung` dideklarasikan di luar fungsi tersebut.

**Variadic Function**

Go mendukung fungsi *variadic*, yang artinya fungsi dapat menerima sejumlah argumen yang bervariasi. 
Hal ini memungkinkan lebih banyak fleksibilitas saat memanggil fungsi dengan jumlah argumen yang berbeda-beda.

Contoh :
```go
func cetakAngka(angka ...int) {
    for _, a := range angka {
        fmt.Println(a)
    }
}

func main() {
    cetakAngka(1, 2, 3, 4, 5)
}
```

Fungsi `cetakAngka` di atas dapat menerima jumlah argumen yang tidak terbatas karena menggunakan parameter variadic (`...int`).

**Goroutine dan Concurrency**

Salah satu kekuatan utama Go adalah dukungan bawaan untuk *concurrency*. 
Dengan menggunakan *goroutine*, kita bisa menjalankan fungsi secara paralel tanpa overhead yang signifikan. 
Goroutine adalah fungsi ringan yang berjalan secara bersamaan dengan fungsi lain.

Contoh :
```go
func cetakAngka() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    go cetakAngka()  // Memanggil fungsi sebagai goroutine
    fmt.Scanln()      // Menunggu input untuk mencegah program selesai sebelum goroutine selesai
}
```

Dalam contoh ini, fungsi `cetakAngka` berjalan sebagai goroutine, yang dijalankan secara bersamaan dengan eksekusi utama.

**Error Handling : Panic dan Recover**

Dalam bahasa Go, kesalahan penanganan dapat dilakukan dengan menggunakan mekanisme panic dan recover. Ketika terjadi panic, program akan berhenti, tetapi kita bisa menangkap panic menggunakan recover untuk mencegah program crash.

Contoh :
```go
func bagi(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Kesalahan:", r)
        }
    }()
    if b == 0 {
        panic("pembagian dengan nol tidak diizinkan")
    }
    return a / b
}

func main() {
    fmt.Println(bagi(4, 2))  // Output: 2
    fmt.Println(bagi(4, 0))  // Output: Kesalahan: pembagian dengan nol tidak diizinkan
}
```

**Interface**

Go menggunakan konsep *interface* untuk mendukung pemrograman yang lebih fleksibel. 
*Interface* mendefinisikan sekumpulan metode yang harus diimplementasikan oleh tipe tertentu. 
Jika tipe mengimplementasikan metode-metode tersebut, maka tipe tersebut "memenuhi" interface.

Contoh :
```go
type penggerak interface {
    bergerak() string
}

type mobil struct{}

func (m mobil) bergerak() string {
    return "Mobil bergerak"
}

func jalankan(p penggerak) {
    fmt.Println(p.bergerak())
}

func main() {
    m := mobil{}
    jalankan(m)
}
```

Dalam contoh ini, tipe *mobil* mengimplementasikan metode *bergerak*, 
sehingga dapat digunakan di fungsi *jalankan* yang menerima tipe *penggerak*.

**Anonymous Function**

Fungsi anonim adalah fungsi tanpa nama yang biasanya didefinisikan dan dipanggil secara langsung. Fungsi ini berguna untuk operasi sederhana yang tidak perlu nama khusus.

Contoh :
```go
func main() {
    hasil := func(a, b int) int {
        return a + b
    }(3, 4)
    fmt.Println(hasil)  // Output: 7
}
```

## GUIDED

### 1. Buatlah sebuah program beserta prosedur yang digunakan untuk menampilkan suatu pesan error, warning atau informasi berdasarkan masukan dari user.

**Masukan:**
- flag: Bilangan bulat dengan nilai 0, 1, atau 2.
- M: Sebuah string berisi pesan.
  
**Keluaran:**
- String pesan M beserta jenis pesannya (error, warning, atau informasi) berdasarkan nilai flag:
  - Jika flag = 0, tampilkan sebagai pesan error.
  - Jika flag = 1, tampilkan sebagai pesan warning.
  - Jika flag = 2, tampilkan sebagai pesan informasi.
    
#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF-11-06

package main

import "fmt"

func main() {
	var bilangan int
	var pesan string
	fmt.Scan(&bilangan, &pesan)
	cetakPesan(pesan, bilangan)
}

func cetakPesan(M string, flag int) {
	var jenis string = ""
	if flag == 0 {
		jenis = "error"
	} else if flag == 1 {
		jenis = "warning"
	} else if flag == 2 {
		jenis = "informasi"
	}
	fmt.Println(M, jenis)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/3ecef633-a58f-4c90-a455-f21508145d3b)

#### Alur Program

**Fungsi Utama (main)**

* **Deklarasi Variabel**
  + bilangan: Variabel bertipe integer yang akan menyimpan nilai bilangan bulat.
  + pesan: Variabel bertipe string yang akan menyimpan sebuah pesan.
* **Input:**
  + fmt.Scan(&bilangan, &pesan): Membaca dua input dari pengguna, yaitu sebuah bilangan bulat dan sebuah string, lalu menyimpannya ke dalam variabel bilangan dan pesan masing-masing.
* **Panggilan Fungsi:**
  + cetakPesan(pesan, bilangan): Memanggil fungsi cetakPesan dengan dua argumen: pesan (string) dan bilangan (integer). Fungsi ini akan memproses pesan dan bilangan sesuai dengan logika yang telah didefinisikan.

**Fungsi cetakPesan**

* **Parameter:**
  + M: Variabel bertipe string yang mewakili pesan yang akan dicetak.
  + flag: Variabel bertipe integer yang berfungsi sebagai flag untuk menentukan jenis pesan (error, warning, atau informasi).

* **Logika:**
  + Kondisi: Menggunakan struktur if-else if-else untuk memeriksa nilai flag dan menentukan jenis pesan.
  + Jika flag bernilai 0, maka jenis pesan adalah "error".
  + Jika flag bernilai 1, maka jenis pesan adalah "warning".
  + Jika flag bernilai 2, maka jenis pesan adalah "informasi".
    
* **Output:**
  + fmt.Println(M, jenis): Mencetak pesan yang diberikan bersama dengan jenis pesan yang sesuai.

#### Deskripsi Program
Program ini membaca dua input dari pengguna: sebuah bilangan bulat yang berfungsi sebagai flag dan sebuah string sebagai pesan. Berdasarkan nilai flag (0, 1, atau 2), program menentukan jenis pesan yang sesuai—"error", "warning", atau "informasi"—melalui fungsi `cetakPesan`. Setelah menerima input, fungsi ini mencetak pesan yang dimasukkan pengguna bersama dengan jenis pesan yang sesuai. Dengan demikian, program ini mengilustrasikan cara menggunakan fungsi untuk mengelompokkan dan memformat output berdasarkan kondisi tertentu.

### 2. Buatlah sebuah program beserta prosedur yang digunakan untuk mengirim email pemberitahuan kepada pengguna baru yang telah mendaftar.

**Masukan:**
- emails: Daftar string yang berisi alamat email pengguna baru.

**Keluaran:**
- Mengirimkan pesan konfirmasi ke setiap alamat email dalam daftar. Format pesan yang ditampilkan adalah:

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF-11-06

package main

import "fmt"

// Procedure untuk mengirim email pemberitahuan
func sendEmailNotification(email string) {
	fmt.Printf("Mengirim email ke %s: Pendaftaran berhasil.\n", email)
}

func main() {
	// Daftar email pengguna baru
	emails := []string{"morcellaa@gmail.com", "sakezaile@gmail.com", "butterscoth@gmail.com"}

	// Mengirim email pemberitahuan ke setiap pengguna
	fmt.Println("Mengirim email ke pengguna yang baru terdaftar:")
	for _, email := range emails {
		sendEmailNotification(email)
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/3ccc2e7f-2b5f-41b9-867d-81612a394524)

#### Alur Program

* **Deklarasi Paket:**
  + package main: Menyatakan bahwa ini adalah program utama (executable).

* **Import Paket:**
  + import "fmt": Mengimpor paket fmt untuk melakukan operasi input dan output (misalnya, mencetak ke layar).

* **Definisi Fungsi sendEmailNotification:**
  + Fungsi ini menerima satu argumen berupa string yang merepresentasikan alamat email.
  + Fungsi ini mencetak pesan ke konsol yang mengindikasikan bahwa email telah dikirim ke alamat email tersebut. Ini adalah simulasi pengiriman email.

* **Fungsi Utama main:**
  + Deklarasi Slice: Membuat sebuah slice (array dinamis) bernama emails untuk menyimpan daftar alamat email.
  + Inisialisasi Slice: Memasukkan beberapa alamat email ke dalam slice emails.
  + Iterasi: Menggunakan for loop untuk mengulang setiap elemen dalam slice emails.
  + Panggilan Fungsi: Di dalam loop, fungsi sendEmailNotification dipanggil dengan elemen slice saat ini sebagai argumen. Artinya, untuk setiap alamat email dalam slice, fungsi sendEmailNotification akan dipanggil dan mencetak pesan ke konsol.

#### Deskripsi Program

Program ini bertujuan untuk mengirimkan pemberitahuan email kepada pengguna baru yang telah terdaftar. Dalam fungsi main, program mendeklarasikan sebuah slice yang berisi daftar alamat email pengguna baru. Selanjutnya, program mencetak pesan yang menunjukkan bahwa email sedang dikirim, lalu menggunakan loop untuk mengiterasi setiap alamat email dalam daftar. Untuk setiap email, program memanggil fungsi sendEmailNotification, yang mencetak pesan konfirmasi bahwa email telah berhasil dikirim. Dengan demikian, program ini mendemonstrasikan cara sederhana untuk mengelola dan memproses daftar email dengan menggunakan prosedur.

### 3. Buatlah sebuah program yang menghitung hasil dari fungsi matematis menggunakan dua pendekatan: satu dengan menggunakan pengembalian nilai dan satu lagi dengan menggunakan referensi.

**Masukan:**
- Dua bilangan bulat yang diambil dari pengguna.

**Keluaran:**
- Hasil dari fungsi f2 yang disimpan dalam variabel c.
- Hasil dari fungsi f1 ketika dipanggil dengan parameter yang terbalik.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF-11-06

package main

import "fmt"

func f1(x, y int) float64 { // pass by value
	var hasil float64
	hasil = float64(2*x) - 0.5*float64(y) + 3.0
	return hasil
}

func f2(x, y int, hasil *float64) { // pass by reference
	*hasil = float64(2*x) - 0.5*float64(y) + 3.0
	// hasil
}

func main() {
	var a, b int
	var c float64

	// take input for a and b
	fmt.Print("Enter two integers: ")
	fmt.Scan(&a, &b)

	// call f2 to calculate and store the result in
	f2(a, b, &c)

	// print the result from f2
	fmt.Println("Result from f2 (stored in c):", c)

	// call f1 and print the result
	resultF1 := f1(b, a)
	fmt.Println("Result from f1:", resultF1)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/c59c054d-00cb-43c5-964f-10f9851136b5)

#### Alur Program

* **Deklarasi Paket:**
  + package main: Menyatakan bahwa ini adalah program utama (executable).
    
* **Definisi Fungsi f1:**
  + Fungsi ini menerima dua parameter integer x dan y sebagai nilai (pass by value).
  + Di dalam fungsi, dihitung nilai hasil berdasarkan rumus 2*x - 0.5*y + 3.0.
  + Nilai hasil dikembalikan sebagai tipe float64.

* **Definisi Fungsi f2:**
  + Fungsi ini menerima dua parameter integer x dan y sebagai nilai (pass by value) dan satu parameter pointer hasil sebagai referensi (pass by reference).
  + Di dalam fungsi, dihitung nilai hasil berdasarkan rumus yang sama seperti di f1.
  + Nilai hasil yang dihitung disimpan ke alamat memori yang ditunjuk oleh pointer hasil.

* **Fungsi Utama main:**
  + Deklarasi variabel:
      - a dan b untuk menyimpan dua bilangan integer yang akan diinput pengguna.
      - c untuk menyimpan hasil perhitungan.
  + Input:
      - Meminta pengguna untuk memasukkan dua bilangan integer dan menyimpannya ke dalam a dan b.
  + Panggilan Fungsi f2:
      - Memanggil fungsi f2 dengan argumen a, b, dan alamat memori dari c. Karena c dilewatkan sebagai referensi, perubahan nilai di dalam f2 akan mempengaruhi nilai c di luar fungsi.
  + Output:
      - Mencetak nilai c yang telah dihitung oleh f2.
  + Panggilan Fungsi f1:
      - Memanggil fungsi f1 dengan argumen b dan a.
  + Output:
      - Mencetak hasil yang dikembalikan oleh f1.

#### Deskripsi Program

Program ini menghitung hasil dari fungsi matematis menggunakan dua metode: pass by value dan pass by reference. Fungsi f1 menerima dua bilangan bulat dan mengembalikan hasil perhitungan berdasarkan rumus \( 2 \cdot x - 0.5 \cdot y + 3.0 \), sementara fungsi f2 menghitung hasil yang sama tetapi menyimpannya melalui pointer, memungkinkan nilai disimpan langsung dalam variabel yang ditunjuk. Dalam fungsi main, program mengambil input dua bilangan dari pengguna, memanggil f2 untuk menghitung dan menyimpan hasilnya, kemudian mencetak hasil tersebut serta hasil dari f1 dengan parameter yang terbalik, menunjukkan perbedaan cara penyimpanan dan pengaksesan hasil.

## Unguided

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas, salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu, bersediakah kalian membantu Jonas?<br/>
### **Masukan** terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.<br/>
### **Keluaran** terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.<br/>
### **Catatan:** permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!<br/>
### P(n,r) = n! / (n-r)!, sedangkan C(n,r) = n! / (r! * (n-r)!)

**Contoh**

| No | Masukan | Keluaran | Penjelasan                               |
|---|---|---|---|
| 1  | 5 10 3 10 | 60 10 3628800 1 | P(5,3) = 5!/2! = 120/2 = 60 <br>C(5,3) = 5!/(3!x2!) = 120/12 = 10 <br>P(10,10) = 10!/0! = 3628800/1 = 3628800 <br>C(10,10) = 10!/(10!x10!) = 10!/10! = 1 |
| 2  | 8 0 2 0 | 56 28 1 1 |  |

Selesaikan program tersebut dengan memanfaatkan subprogram yang diberikan berikut ini!

```
procedure factorial(in n: integer, in/out hasil: integer)
{ I.S. terdefinisi bilangan bulat positif n
  F.S. hasil berisi nilai faktorial dari n }

procedure permutation(in n, r: integer, in/out hasil: integer)
{ I.S. terdefinisi bilangan bulat positif n dan r, dan n >= r
  F.S. hasil berisi nilai dari n permutasi r }

procedure combination(in n, r: integer, in/out hasil: integer)
{ I.S. terdefinisi bilangan bulat positif n dan r, dan n >= r
  F.S. hasil berisi nilai dari n kombinasi r }
```

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF1106

package main

import "fmt"

// Prosedur untuk input bilangan
func inputBilangan(prompt string, bilangan *int) {
	fmt.Print(prompt)
	fmt.Scan(bilangan)
}

// Prosedur untuk mengecek syarat a>=c dan b>=d
func cekSyarat(a, b, c, d int) bool {
	return a >= c && b >= d
}

// Prosedur untuk menampilkan hasil perhitungan
func tampilkanHasil(label string, permutasi, kombinasi int) {
	fmt.Printf("\nPermutasi %s: %d\n", label, permutasi)
	fmt.Printf("Kombinasi %s: %d\n", label, kombinasi)
}

// Prosedur untuk menghitung faktorial
func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi
func permutasi(n, r int, hasil *int) {
	var faktorialN, faktorialNR int
	faktorial(n, &faktorialN)
	faktorial(n-r, &faktorialNR)
	*hasil = faktorialN / faktorialNR
}

// Prosedur untuk menghitung kombinasi
func kombinasi(n, r int, hasil *int) {
	var faktorialN, faktR, faktorialNR int
	faktorial(n, &faktorialN)
	faktorial(r, &faktR)
	faktorial(n-r, &faktorialNR)
	*hasil = faktorialN / (faktR * faktorialNR)
}

// Prosedur untuk melakukan perhitungan permutasi dan kombinasi
func hitungPermutasiKombinasi(n, r int) (int, int) {
	var hasilPermutasi, hasilKombinasi int
	permutasi(n, r, &hasilPermutasi)
	kombinasi(n, r, &hasilKombinasi)
	return hasilPermutasi, hasilKombinasi
}

func main() {
	var a, b, c, d int

	// Input bilangan menggunakan prosedur
	inputBilangan("Bilangan a: ", &a)
	inputBilangan("Bilangan b: ", &b)
	inputBilangan("Bilangan c: ", &c)
	inputBilangan("Bilangan d: ", &d)

	// Cek syarat menggunakan prosedur
	if cekSyarat(a, b, c, d) {
		// Hitung permutasi dan kombinasi untuk (a,c)
		permutasiAC, kombinasiAC := hitungPermutasiKombinasi(a, c)
		tampilkanHasil("(a, c)", permutasiAC, kombinasiAC)

		// Hitung permutasi dan kombinasi untuk (b,d)
		permutasiBD, kombinasiBD := hitungPermutasiKombinasi(b, d)
		tampilkanHasil("(b, d)", permutasiBD, kombinasiBD)
	} else {
		fmt.Println("Input tidak sesuai dengan syarat yang ada")
	}
}
```
#### Screenshoot Output

![image](https://github.com/user-attachments/assets/b2e0b437-4ad0-4f8c-ba4b-c01340849483)

#### Alur Program

* **Input Bilangan:**
  + Program meminta pengguna untuk memasukkan empat bilangan bulat a, b, c, dan d.

* **Cek Syarat:**
  + Program memeriksa apakah a >= c dan b >= d. Jika syarat ini tidak terpenuhi, program akan menampilkan pesan kesalahan.

* **Hitung Permutasi dan Kombinasi:**
  + Jika syarat terpenuhi, program akan menghitung permutasi dan kombinasi untuk pasangan bilangan (a, c) dan (b, d).
  + Hasil perhitungan kemudian ditampilkan.

#### Deskripsi Program

Program ini dirancang untuk menghitung permutasi dan kombinasi dari dua pasangan bilangan yang dimasukkan pengguna, dengan syarat bahwa bilangan pertama dari setiap pasangan harus lebih besar atau sama dengan bilangan kedua. Melalui fungsi inputBilangan, pengguna diminta memasukkan empat bilangan: a, b, c, dan d. Program kemudian memeriksa syarat menggunakan fungsi cekSyarat, dan jika syarat terpenuhi, melakukan perhitungan permutasi dan kombinasi melalui fungsi hitungPermutasiKombinasi, yang memanfaatkan prosedur faktorial untuk menghitung nilai faktorial yang diperlukan. Hasil dari perhitungan kemudian ditampilkan menggunakan fungsi tampilkanHasil. Dengan struktur ini, program menunjukkan penggunaan prosedur untuk modularisasi dan pemisahan logika dalam perhitungan matematis.

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.
### Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.

```prosedure hitungSkor(in/out soal, skor : integer)```

### Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).
### Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

| No | Masukan                                      | Keluaran |
|---|-----------------------------------------------|----------|
| 1 | Astuti 28 58 381 381 61 71 75 18 <br>Bertha 25 47 381 26 58 68 65 21 | Bertha 7 294     |
|   | Selesai                                          |          |

**Keterangan**
Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF1106

package main

import (
    "fmt"
)

// Struct untuk menyimpan data peserta
type Peserta struct {
    nama      string
    waktu     [8]int
    soalBenar int
    totalWaktu int
}

// Prosedur untuk input jumlah peserta
func inputJumlahPeserta() int {
    var n int
    fmt.Print("Jumlah Peserta: ")
    fmt.Scan(&n)
    return n
}

// Prosedur untuk input data peserta
func inputDataPeserta(nomorPeserta int) Peserta {
    var peserta Peserta
    
    // Input nama peserta
    fmt.Printf("\nNama peserta %d: ", nomorPeserta+1)
    fmt.Scan(&peserta.nama)
    
    // Input waktu pengerjaan
    fmt.Print("Waktu Pengerjaan Soal (8 soal): ")
    for j := 0; j < 8; j++ {
        fmt.Scan(&peserta.waktu[j])
    }
    
    return peserta
}

// Prosedur untuk menghitung total soal yang dikerjakan dan total waktu
func hitungSkor(waktu [8]int, soal *int, totalWaktu *int) {
    *soal = 0
    *totalWaktu = 0
    for i := 0; i < 8; i++ {
        if waktu[i] <= 300 { // jika waktu pengerjaan kurang dari 300 menit, soal selesai
            *soal++
            *totalWaktu += waktu[i] // hanya tambahkan waktu soal yang selesai
        }
    }
}

// Prosedur untuk menentukan pemenang
func tentukanPemenang(peserta Peserta, pemenangSekarang *Peserta) {
    if peserta.soalBenar > pemenangSekarang.soalBenar || 
       (peserta.soalBenar == pemenangSekarang.soalBenar && peserta.totalWaktu < pemenangSekarang.totalWaktu) {
        *pemenangSekarang = peserta
    }
}

// Prosedur untuk menampilkan hasil akhir
func tampilkanHasil(pemenang Peserta) {
    fmt.Printf("\nNama pemenang: %s\n", pemenang.nama)
    fmt.Printf("Jumlah soal yang selesai: %d\n", pemenang.soalBenar)
    fmt.Printf("Total waktu yang dihabiskan: %d menit\n", pemenang.totalWaktu)
}

func main() {
    // Input jumlah peserta
    n := inputJumlahPeserta()
    
    // Inisialisasi data pemenang sementara
    pemenang := Peserta{
        soalBenar: -1,
        totalWaktu: 1000,
    }
    
    // Proses setiap peserta
    for i := 0; i < n; i++ {
        // Input data peserta
        peserta := inputDataPeserta(i)
        
        // Hitung skor peserta
        hitungSkor(peserta.waktu, &peserta.soalBenar, &peserta.totalWaktu)
        
        // Tentukan pemenang sementara
        tentukanPemenang(peserta, &pemenang)
    }
    
    // Tampilkan hasil akhir
    tampilkanHasil(pemenang)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/7a69e86e-8941-4cd3-a38e-94f61f6aee07)

#### Alur Program

* **Input Jumlah Peserta:**
  + Program meminta pengguna untuk memasukkan jumlah peserta yang akan diproses.

* **Inisialisasi Pemenang:**
  + Program menginisialisasi struktur pemenang dengan nilai awal yang sangat rendah untuk jumlah soal benar dan waktu yang sangat tinggi.

* **Proses Setiap Peserta:**
  + Untuk setiap peserta, program meminta input data (nama dan waktu pengerjaan soal).
  + Menghitung jumlah soal benar dan total waktu yang digunakan.
  + Membandingkan skor peserta dengan skor pemenang sementara dan memperbarui pemenang jika diperlukan.

* **Tampilkan Hasil:**
  + Setelah semua peserta diproses, program menampilkan nama pemenang, jumlah soal benar, dan total waktu yang digunakan.

#### Deskripsi Program

Program ini mengelola data peserta dalam kompetisi dengan menghitung skor dan menentukan pemenang berdasarkan jumlah soal yang selesai dikerjakan dan total waktu yang dihabiskan. Menggunakan struct Peserta, program meminta input jumlah peserta, kemudian mengumpulkan data seperti nama dan waktu pengerjaan untuk delapan soal. Melalui prosedur hitungSkor, program menghitung jumlah soal yang berhasil diselesaikan dan total waktu yang dihabiskan, sementara tentukanPemenang membandingkan peserta untuk memperbarui pemenang sementara. Akhirnya, program mencetak hasil pemenang dengan mencantumkan nama, jumlah soal yang selesai, dan total waktu yang digunakan, semua dalam alur yang terorganisir dengan baik.

### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n, jika bilangan n saat itu genap, maka suku berikutnya adalah n/2, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah:
### 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1
### Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.
### Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai suku awal.

```prosedure cetakDeret(in n : integer )```

### Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000.
### Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang sama dan dipisahkan oleh sebuah spasi.

| No | Masukan | Keluaran                                                              |
|---|----------|-------------------------------------------------------------------|
| 1 | 22       | 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1                         |


#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF1106

package main

import (
	"fmt"
)

// Prosedur untuk input bilangan
func inputBilangan() int {
	var n int
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)
	return n
}

// Prosedur untuk validasi input
func validasiInput(n int) bool {
	if n <= 0 || n >= 1000000 {
		fmt.Println("Nilai harus positif dan kurang dari 1000000.")
		return false
	}
	return true
}

// Prosedur untuk menghitung bilangan berikutnya dalam deret
func hitungBilanganBerikutnya(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

// Prosedur untuk menyimpan deret dalam slice
func buatDeret(n int) []int {
	deret := make([]int, 0)
	deret = append(deret, n)

	for n != 1 {
		n = hitungBilanganBerikutnya(n)
		deret = append(deret, n)
	}

	return deret
}

// Prosedur untuk mencetak deret
func cetakDeret(deret []int) {
	fmt.Print("Suku dan deret: ")
	for i, nilai := range deret {
		if i == len(deret)-1 {
			fmt.Printf("%d", nilai)
		} else {
			fmt.Printf("%d ", nilai)
		}
	}
	fmt.Println()
}

// Prosedur untuk analisis deret
func analisisDeret(deret []int) {
	panjangDeret := len(deret)
	nilaiTerbesar := deret[0]

	for _, nilai := range deret {
		if nilai > nilaiTerbesar {
			nilaiTerbesar = nilai
		}
	}

	fmt.Printf("\nAnalisis Deret:\n")
	fmt.Printf("Panjang deret: %d\n", panjangDeret)
	fmt.Printf("Nilai terbesar dalam deret: %d\n", nilaiTerbesar)
}

func main() {
	// Input bilangan
	n := inputBilangan()

	// Validasi input
	if validasiInput(n) {
		// Buat deret
		deret := buatDeret(n)

		// Cetak deret
		cetakDeret(deret)

		// Tampilkan analisis deret
		analisisDeret(deret)
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/8bcab542-3341-48ec-aa38-072dd939e399)

#### Alur Program

* **Input Bilangan:**
  + Program meminta pengguna untuk memasukkan bilangan awal.
    
* **Validasi Input:**
  + Memastikan bilangan yang dimasukkan positif dan kurang dari 1000000.

* **Buat Deret:**
  + Membuat slice kosong untuk menyimpan deret.
  + Menambahkan bilangan awal ke dalam slice.
  + Selama bilangan tidak sama dengan 1, hitung bilangan berikutnya berdasarkan aturan Collatz dan tambahkan ke dalam slice.

* **Cetak Deret:**
  + Mencetak semua elemen dalam slice deret.
    
* **Analisis Deret:**
  + Menghitung panjang deret.
  + Mencari nilai terbesar dalam deret.
  + Mencetak hasil analisis.

#### Deskripsi Program

Program ini menghitung dan menganalisis deret bilangan berdasarkan algoritma Collatz, di mana pengguna diminta memasukkan sebuah bilangan yang kemudian divalidasi agar positif dan kurang dari 1.000.000. Setelah validasi, program menghitung bilangan berikutnya dalam deret dengan aturan tertentu jika bilangan genap dibagi dua, dan jika ganjil, dikalikan tiga dan ditambah satu sambil menyimpan hasilnya dalam sebuah slice. Program juga mencetak deret tersebut dan melakukan analisis untuk menentukan panjang deret serta nilai terbesar di dalamnya. Dengan cara ini, program menerapkan konsep pemrograman terstruktur secara efektif untuk menghasilkan dan menganalisis deret bilangan.
