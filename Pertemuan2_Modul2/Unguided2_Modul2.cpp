#include <iostream>
using namespace std;

// Fungsi tukar menggunakan Pointer
void tukarPointer(int *x, int *y, int *z) {
    int temp = *x;
    *x = *y;
    *y = *z;
    *z = temp;
}

// Fungsi tukar menggunakan Reference
void tukarReference(int &x, int &y, int &z) {
    int temp = x;
    x = y;
    y = z;
    z = temp;
}

int main() {
    int a = 1, b = 2, c = 3;

    cout << "--- UJI COBA POINTER ---\n";
    cout << "Sebelum ditukar: a = " << a << ", b = " << b << ", c = " << c << endl;
    tukarPointer(&a, &b, &c);
    cout << "Setelah ditukar: a = " << a << ", b = " << b << ", c = " << c << endl;

   
    a = 1; b = 2; c = 3;

    cout << "\n--- UJI COBA REFERENCE ---\n";
    cout << "Sebelum ditukar: a = " << a << ", b = " << b << ", c = " << c << endl;
    tukarReference(a, b, c);
    cout << "Setelah ditukar: a = " << a << ", b = " << b << ", c = " << c << endl;

    return 0;
}