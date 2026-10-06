#include <iostream>
using namespace std;

const int ukuran = 10;
int arrA[ukuran] = {11, 8, 5, 7, 12, 26, 3, 54, 33, 55};

// Function mencari nilai minimum
int cariMinimum() {
    int min = arrA[0];
    for(int i = 1; i < ukuran; i++) {
        if(arrA[i] < min) {
            min = arrA[i];
        }
    }
    return min;
}

// Function mencari nilai maksimum
int cariMaksimun() {
    int maks = arrA[0];
    for(int i = 1; i < ukuran; i++) {
        if(arrA[i] > maks) {
            maks = arrA[i];
        }
    }
    return maks;
}

// Prosedur menghitung nilai rata-rata (void)
void hitungRataRata() {
    float total = 0;
    for(int i = 0; i < ukuran; i++) {
        total += arrA[i];
    }
    float rata = total / ukuran;
    cout << "Nilai rata - rata = " << rata << endl;
}

int main() {
    int pilihan;

    do {
        cout << "\n--- Menu Program Array ---\n";
        cout << "1. Tampilkan isi array\n";
        cout << "2. Cari nilai maksimum\n";
        cout << "3. Cari nilai minimum\n";
        cout << "4. Hitung nilai rata - rata\n";
        cout << "5. Keluar\n";
        cout << "Pilihan Anda: ";
        cin >> pilihan;

        switch(pilihan) {
            case 1:
                cout << "Isi array: ";
                for(int i = 0; i < ukuran; i++) {
                    cout << arrA[i] << " ";
                }
                cout << endl;
                break;
            case 2:
                cout << "Nilai maksimum = " << cariMaksimun() << endl;
                break;
            case 3:
                cout << "Nilai minimum = " << cariMinimum() << endl;
                break;
            case 4:
                hitungRataRata();
                break;
            case 5:
                cout << "Keluar dari program.\n";
                break;
            default:
                cout << "Pilihan tidak valid!\n";
        }
    } while(pilihan != 5);

    return 0;
}