#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#include <iostream>

using namespace std;

#define Nil NULL
#define Info(P) (P)->info
#define Next(P) (P)->next
#define First(L) (L).first

typedef int infotype;
typedef struct elmlist *address;

struct elmlist {
    infotype info;
    address next;
};

struct List {
    address first;
};

// Prosedur dan Fungsi Dasar
void CreateList(List &L);
address alokasi(infotype x);
void dealokasi(address &P);
void printInfo(List L);
void insertFirst(List &L, address P);

// Fungsi Tambahan (Latihan No. 3 & No. 4)
address findElm(List L, infotype x);
int sumInfo(List L);

#endif