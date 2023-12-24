# Note

## go generate

Menurut Google: go:generate is a directive. It is used primarily to generate Go
source code. But note that you can use it for any command. To use it, you have
to put a line in any Go file : //go:generate command argument. command argument
will tell Go what you want to run.

simpel nya go:generate sebuah baris comment khusus yang digunakan untuk
menjalankan perintah command line yang dapat menghasilkan file atau kode
tambahan yang diperlukan oleh kode program utama aplikasi atau hanya sekedar
menjalankan sebuah perintah tanpa menghasilkan apa-apa.

untuk menggunakan fitur ini dapat menggunakan perintah `go generate` diikuti
dengan nama file yang akan dijalankan dengan perintah `go generate` yaitu file
yang terdapat baris comment `//go:generate command argument` di dalamnya.
