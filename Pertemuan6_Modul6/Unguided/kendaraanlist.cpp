#include "kendaraanlist.h"
#include <iostream>

using namespace std;

void createList(list &L) {
    L.first = nullptr;
    L.last = nullptr;
}

address alokasi(infotype x) {
    address P = new elmlist;
    P->info = x;
    P->next = nullptr;
    P->prev = nullptr;
    return P;
}

void dealokasi(address &P) {
    delete P;
    P = nullptr;
}

void insertFirst(list &L, address P) {
    if (L.first == nullptr) {
        L.first = P;
        L.last = P;
    } else {
        P->next = L.first;
        L.first->prev = P;
        L.first = P;
    }
}

void insertLast(list &L, address P) {
    if (L.first == nullptr) {
        insertFirst(L, P);
    } else {
        L.last->next = P;
        P->prev = L.last;
        L.last = P;
    }
}

address findElm(list L, string nopol) {
    address P = L.first;
    while (P != nullptr) {
        if (P->info.nopol == nopol) {
            return P;
        }
        P = P->next;
    }
    return nullptr;
}

void deleteFirst(list &L, address &P) {
    if (L.first != nullptr) {
        P = L.first;
        if (L.first == L.last) {
            L.first = nullptr;
            L.last = nullptr;
        } else {
            L.first = L.first->next;
            L.first->prev = nullptr;
            P->next = nullptr;
        }
    }
}

void deleteLast(list &L, address &P) {
    if (L.last != nullptr) {
        P = L.last;
        if (L.first == L.last) {
            deleteFirst(L, P);
        } else {
            L.last = L.last->prev;
            L.last->next = nullptr;
            P->prev = nullptr;
        }
    }
}

void deleteAfter(list &L, address Prec, address &P) {
    if (Prec != nullptr && Prec->next != nullptr) {
        P = Prec->next;
        if (P == L.last) {
            deleteLast(L, P);
        } else {
            Prec->next = P->next;
            P->next->prev = Prec;
            P->next = nullptr;
            P->prev = nullptr;
        }
    }
}

void deleteSpecific(list &L, string nopol, address &P) {
    P = findElm(L, nopol);
    if (P != nullptr) {
        if (P == L.first) {
            deleteFirst(L, P);
        } else if (P == L.last) {
            deleteLast(L, P);
        } else {
            deleteAfter(L, P->prev, P);
        }
    }
}

void printInfo(list L) {
    address P = L.first;
    if (P == nullptr) {
        cout << "List kosong." << endl;
    } else {
        cout << "DATA LIST 1" << endl;
        while (P != nullptr) {
            cout << "nomor polisi : " << P->info.nopol << endl;
            cout << "warna        : " << P->info.warna << endl;
            cout << "tahun        : " << P->info.thnBuat << endl;
            cout << endl;
            P = P->next;
        }
    }
}