#include "kendaraanlist.h"
#include <iostream>

using namespace std;

int main() {
    list L;
    createList(L);
    infotype x;
    address P;
    string nopolCari, nopolHapus;

    // 1. Input data kendaraan & validasi duplikat (Sesuai Gambar 6-23)
    x = {"D001", "hitam", 90};
    cout << "masukkan nomor polisi: " << x.nopol << endl;
    cout << "masukkan warna kendaraan: " << x.warna << endl;
    cout << "masukkan tahun kendaraan: " << x.thnBuat << endl;
    if (findElm(L, x.nopol) == nullptr) {
        P = alokasi(x);
        insertFirst(L, P);
    } else {
        cout << "nomor polisi sudah terdaftar" << endl;
    }
    cout << endl;

    x = {"D003", "putih", 70};
    cout << "masukkan nomor polisi: " << x.nopol << endl;
    cout << "masukkan warna kendaraan: " << x.warna << endl;
    cout << "masukkan tahun kendaraan: " << x.thnBuat << endl;
    if (findElm(L, x.nopol) == nullptr) {
        P = alokasi(x);
        insertFirst(L, P);
    } else {
        cout << "nomor polisi sudah terdaftar" << endl;
    }
    cout << endl;

    x = {"D001", "merah", 80};
    cout << "masukkan nomor polisi: " << x.nopol << endl;
    cout << "masukkan warna kendaraan: " << x.warna << endl;
    cout << "masukkan tahun kendaraan: " << x.thnBuat << endl;
    if (findElm(L, x.nopol) == nullptr) {
        P = alokasi(x);
        insertFirst(L, P);
    } else {
        cout << "nomor polisi sudah terdaftar" << endl;
    }
    cout << endl;

    x = {"D004", "kuning", 90};
    cout << "masukkan nomor polisi: " << x.nopol << endl;
    cout << "masukkan warna kendaraan: " << x.warna << endl;
    cout << "masukkan tahun kendaraan: " << x.thnBuat << endl;
    if (findElm(L, x.nopol) == nullptr) {
        P = alokasi(x);
        insertFirst(L, P);
    } else {
        cout << "nomor polisi sudah terdaftar" << endl;
    }
    cout << endl;

    printInfo(L);

    // 2. Pencarian Elemen (Sesuai Gambar 6-24)
    nopolCari = "D001";
    cout << "Masukkan Nomor Polisi yang dicari : " << nopolCari << endl;
    address found = findElm(L, nopolCari);
    if (found != nullptr) {
        cout << "Nomor Polisi : " << found->info.nopol << endl;
        cout << "Warna        : " << found->info.warna << endl;
        cout << "Tahun        : " << found->info.thnBuat << endl;
    } else {
        cout << "Data tidak ditemukan." << endl;
    }
    cout << endl;

    // 3. Penghapusan Elemen (Sesuai Gambar 6-25)
    nopolHapus = "D003";
    cout << "Masukkan Nomor Polisi yang akan dihapus : " << nopolHapus << endl;
    deleteSpecific(L, nopolHapus, P);
    if (P != nullptr) {
        cout << "Data dengan nomor polisi " << nopolHapus << " berhasil dihapus." << endl;
        dealokasi(P);
    } else {
        cout << "Data tidak ditemukan." << endl;
    }
    cout << endl;

    printInfo(L);

    return 0;
}