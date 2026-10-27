#include "Singlylist.h"

// Membuat list kosong
void CreateList(List &L) {
    First(L) = Nil;
}

// Alokasi memori untuk elemen baru
address alokasi(infotype x) {
    address P = new elmlist;
    if (P != Nil) {
        Info(P) = x;
        Next(P) = Nil;
    }
    return P;
}

// Dealokasi memori elemen
void dealokasi(address &P) {
    delete P;
    P = Nil;
}

// Menampilkan seluruh info dalam list (Gambar 5-2)
void printInfo(List L) {
    address P = First(L);
    while (P != Nil) {
        cout << Info(P) << " ";
        P = Next(P);
    }
    cout << endl;
}

// Menambahkan elemen di awal list
void insertFirst(List &L, address P) {
    Next(P) = First(L);
    First(L) = P;
}

// (Latihan No. 3) Mencari elemen dengan info tertentu
address findElm(List L, infotype x) {
    address P = First(L);
    while (P != Nil) {
        if (Info(P) == x) {
            return P;
        }
        P = Next(P);
    }
    return Nil;
}

// (Latihan No. 4) Menghitung jumlah total info seluruh elemen
int sumInfo(List L) {
    int total = 0;
    address P = First(L);
    while (P != Nil) {
        total += Info(P);
        P = Next(P);
    }
    return total;
}