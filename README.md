# go-learn
perbedaan antara kode yang saya buat dan kode nana:

1. **Tampilan Sambutan dengan ASCII Art**: Saya menambahkan fungsi displayWelcomeMessage dengan sambutan yang dilengkapi ASCII art.

2. **Looping Pemesanan:** Di kode saya, pemesanan berjalan dalam loop for, sehingga pengguna dapat terus memesan tiket hingga habis. Jika tiket sudah habis, program otomatis keluar dari loop dan memberi tahu pengguna bahwa konferensi sudah penuh.

3. **Penggunaan Goroutine untuk Pengiriman Tiket:** Di kode saya, saya menambahkan wg.Add(1) sebelum menjalankan goroutine sendTicket, yang membuat goroutine-goroutine tersebut dikelola oleh WaitGroup dan memastikan program menunggu semua goroutine selesai sebelum berakhir (melalui wg.Wait()). Waktu tunda di sendTicket juga saya atur menjadi 5 detik.