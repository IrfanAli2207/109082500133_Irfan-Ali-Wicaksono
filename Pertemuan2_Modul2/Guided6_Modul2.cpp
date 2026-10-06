#include <iostream>
using namespace std;

// Fungsi untuk mencari nilai maksimum dari tiga bilangan
int maks3(int a, int b, int c) {
    int temp_max = a;
    
    if (b > temp_max) {
        temp_max = b;
    }
    if (c > temp_max) {
        temp_max = c;
    }
    
    return temp_max;
}

int main() {
    int x, y, z;
    
    // Input dari pengguna
    cout << "Masukkan nilai 1: ";
    cin >> x;
    cout << "Masukkan nilai 2: ";
    cin >> y;
    cout << "Masukkan nilai 3: ";
    cin >> z;
    
    // Memanggil fungsi dan menampilkan hasil
    cout << "Nilai maksimum = " << maks3(x, y, z) << endl;
    
    return 0;
}