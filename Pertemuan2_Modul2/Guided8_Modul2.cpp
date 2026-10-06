#include <iostream>
using namespace std;

// 1. Call by Value (nilai asli TIDAK berubah)
void tukarByValue(int x, int y) {
    int temp = x;
    x = y;
    y = temp;
}

// 2. Call by Pointer (menggunakan pointer *, nilai asli BERUBAH)
void tukarByPointer(int *x, int *y) {
    int temp = *x;
    *x = *y;
    *y = temp;
}

// 3. Call by Reference (menggunakan alias &, nilai asli BERUBAH)
void tukarByReference(int &x, int &y) {
    int temp = x;
    x = y;
    y = temp;
}

int main() {
    // --- UJI COBA CALL BY VALUE ---
    int a = 4, b = 6;
    cout << "--- CALL BY VALUE ---" << endl;
    cout << "Sebelum: a = " << a << ", b = " << b << endl;
    tukarByValue(a, b);
    cout << "Setelah: a = " << a << ", b = " << b << " (TIDAK BERUBAH)\n\n";

    // --- UJI COBA CALL BY POINTER ---
    int c = 4, d = 6;
    cout << "--- CALL BY POINTER ---" << endl;
    cout << "Sebelum: c = " << c << ", d = " << d << endl;
    tukarByPointer(&c, &d); // Mengirim alamat memori pakai &
    cout << "Setelah: c = " << c << ", d = " << d << " (BERHASIL DITUKAR)\n\n";

    // --- UJI COBA CALL BY REFERENCE ---
    int e = 4, f = 6;
    cout << "--- CALL BY REFERENCE ---" << endl;
    cout << "Sebelum: e = " << e << ", f = " << f << endl;
    tukarByReference(e, f); // Mengirim variabel langsung tanpa &
    cout << "Setelah: e = " << e << ", f = " << f << " (BERHASIL DITUKAR)" << endl;

    return 0;
}