#ifndef SINGLYLIST_H
#ifndef NULL
#define NULL 0
#endif

typedef int infotype;
typedef struct ElmList *address;

struct ElmList {
    infotype info;
    address next;
};

struct List {
    address First;
};

#define Nil NULL
#define Info(P) (P)->info
#define Next(P) (P)->next
#define First(L) ((L).First)

void createList(List &L);
address alokasi(infotype x);
void dealokasi(address P);
void printInfo(List L);
void insertFirst(List &L, address P);

// Fungsi & Prosedur Tambahan untuk Soal 2
void deleteFirst(List &L, address &P);
void deleteLast(List &L, address &P);
void deleteAfter(address Prec, address &P);
int nbList(List L);
void deleteList(List &L);

#endif