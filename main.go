package main

import (
    "fmt"
    "strings"
)

type paketInternet struct{
    nama string
    biayaDasar int
    batasPerangkat int
    denda int
}

func (pi paketInternet) hitungDenda(perangakatTerhubung int) int{
    if perangakatTerhubung > pi.batasPerangkat{
        return pi.denda
    }
    return 0
}

func main() {

    daftarPaket := map[string]paketInternet{
        "STARTER":{batasPerangkat: 100, biayaDasar: 150000, denda: 75000},
        "BUSINESS":{batasPerangkat: 500, biayaDasar: 500000, denda: 200000},
        "ENTERPRISE":{batasPerangkat: 2000, biayaDasar: 1500000, denda: 0},
    }


    var paket string
    var pemakaian float64
    var perangkatTerhubung int
    var pembayaranLancar string
    var biayaProgresif float64 
    var diskon int
    var gantiPaket string
    var paketBaru string
    
    fmt.Println("SELAMAT DATANG DI PROGRAM TAGIHAN INTERNET DWI LAKSMANA")

    fmt.Print("Masukkan nama paket internet (STARTER/BUSINESS/ENTERPRISE): ")
    fmt.Scanln(&paket)
    fmt.Print("Masukkan jumlah pemakaian data (dalam GB): ")
    fmt.Scanln(&pemakaian)
    fmt.Print("Masukkan jumlah perangkat terhubung: ")
    fmt.Scanln(&perangkatTerhubung)
    fmt.Print("Apakah pembayaran lancar? (Y/N): ")
    fmt.Scanln(&pembayaranLancar)

    paket = strings.ToUpper(strings.TrimSpace(paket))
    pembayaranLancar = strings.ToUpper(strings.TrimSpace(pembayaranLancar))
    
    dataLama, paketValid := daftarPaket[paket]
    if !paketValid {
        fmt.Println("Paket Tidak Valid")
        return
    }

    dendaLama := dataLama.hitungDenda(perangkatTerhubung)

    if (pemakaian >=0 && pemakaian <=100){
        biayaProgresif = 1000
    } else if (pemakaian >100 && pemakaian <=500){
        biayaProgresif = 1500
    } else if (pemakaian >500){
        biayaProgresif = 2500
    }
    var kuotaProgresif = biayaProgresif * pemakaian
    var tagihanKotor = float64(dataLama.biayaDasar) + kuotaProgresif + float64(dendaLama)
    if (tagihanKotor >= 2000000 && pembayaranLancar == "Y") {
        diskon = int(tagihanKotor * 10 / 100)
    } else if (tagihanKotor < 2000000 && pembayaranLancar == "N") {
        diskon = 0
    }
    var tagihanSetelahDiskon = tagihanKotor - float64(diskon)
    var tagihanBersih = tagihanSetelahDiskon + (tagihanSetelahDiskon * 11 / 100)

    //struk
    fmt.Println("================STRUK TAGIHAN==============")
    fmt.Println("Abodemen                               : Rp", dataLama.biayaDasar)
    fmt.Printf("Kuota Progresif                        : Rp %.2f\n", kuotaProgresif)
    fmt.Println("Denda                                  : Rp", dendaLama)
    fmt.Println("-------------------------------------------------------")
    fmt.Printf ("Tagihan Kotor                          : Rp %.2f\n", tagihanKotor)
    fmt.Println("Diskon                                 : Rp", diskon)
    fmt.Printf("Total Sesudah PPN                      : Rp %.2f\n", tagihanBersih)
    if (pemakaian>50 && pembayaranLancar == "N"){
        fmt.Println("AKSES ANDA DIBLOKIR. SEGERA LAKUKAN PEMBAYARAN UNTUK MENGHINDARI PEMBLOKIRAN AKSES INTERNET ANDA.")
    } 

    var opsi []string
    semuaPaket := []string{"STARTER", "BUSINESS", "ENTERPRISE"}
    for _, p := range semuaPaket {
        if p != paket{
            opsi = append(opsi, p)
        }
    }

    fmt.Println("Apakah anda ingin memperbarui langganan internet anda? (Y/N): ")
    fmt.Scanln(&gantiPaket)
    gantiPaket = strings.ToUpper(strings.TrimSpace(gantiPaket))

    var dataBaru paketInternet

    if gantiPaket == "Y" {
        for{
            fmt.Println("Pilih paket untuk memperbarui langganan intenet anda(" + strings.Join(opsi, "/") + "): ")
            fmt.Scanln(&paketBaru)
            paketBaru = strings.ToUpper(paketBaru)

            if paketBaru == paket{
                fmt.Println("Anda tidak dapat memilih paket yang sama dengan langganan anda sekarang\n")
                continue
            }

            var paketBaruValid bool
            dataBaru, paketBaruValid = daftarPaket[paketBaru]
            if !paketBaruValid {
                fmt.Println("Paket yang anda pilih tidak valid\n")
                continue
            }
            break
        }
        dendaBaru := dataBaru.hitungDenda(perangkatTerhubung)

        fmt.Println("===== PERUBAHAN DATA LANGGANAN INTERNET =====")
        fmt.Println("PAKET LAMA         : ",paket)
        fmt.Println("Biaya dasar        : ", dataLama.biayaDasar)
        fmt.Println("Batas Perangkat    : ", dataLama.batasPerangkat)
        fmt.Println("Denda Lebih        : ", dendaLama)

        fmt.Println("---------------------------------------------")

        fmt.Println("===== PERUBAHAN DATA LANGGANAN INTERNET =====")
        fmt.Println("PAKET Baru         : ",paketBaru)
        fmt.Println("Biaya dasar        : ", dataBaru.biayaDasar)
        fmt.Println("Batas Perangkat    : ", dataBaru.batasPerangkat)
        fmt.Println("Denda Lebih        : ", dendaBaru)
    
    }else{
        fmt.Println("Terimakasih telah menggunakan layanan kami")
    }


}