#include <iostream>
#include "Singlylist.h"
using namespace std;

void createList(List &L) {
    First(L) = Nil;
}

address alokasi(infotype x) {
    address P = new ElmList;
    if (P != Nil) {
        Info(P) = x;
        Next(P) = Nil;
    }
    return P;
}

void dealokasi(address P) {
    delete P;
}

void printInfo(List L) {
    address P = First(L);
    while (P != Nil) {
        cout << Info(P) << " ";
        P = Next(P);
    }
    cout << endl;
}

void insertFirst(List &L, address P) {
    if (P != Nil) {
        Next(P) = First(L);
        First(L) = P;
    }
}

void deleteFirst(List &L, address &P) {
    if (First(L) != Nil) {
        P = First(L);
        First(L) = Next(First(L));
        Next(P) = Nil;
    } else {
        P = Nil;
    }
}

void deleteLast(List &L, address &P) {
    if (First(L) != Nil) {
        address Last = First(L);
        address Prec = Nil;
        while (Next(Last) != Nil) {
            Prec = Last;
            Last = Next(Last);
        }
        P = Last;
        if (Prec == Nil) {
            First(L) = Nil;
        } else {
            Next(Prec) = Nil;
        }
    } else {
        P = Nil;
    }
}

void deleteAfter(address Prec, address &P) {
    if (Prec != Nil && Next(Prec) != Nil) {
        P = Next(Prec);
        Next(Prec) = Next(P);
        Next(P) = Nil;
    } else {
        P = Nil;
    }
}

int nbList(List L) {
    int count = 0;
    address P = First(L);
    while (P != Nil) {
        count++;
        P = Next(P);
    }
    return count;
}

void deleteList(List &L) {
    address P;
    while (First(L) != Nil) {
        deleteFirst(L, P);
        dealokasi(P);
    }
    cout << "- List Berhasil Terhapus -" << endl;
}