#include <iostream>
using namespace std;

int main() {
    float a, b;
    cout << "Masukkan dua bilangan: ";
    cin >> a >> b;
    
    cout << "Hasil + : " << a + b << endl;
    cout << "Hasil - : " << a - b << endl;
    cout << "Hasil * : " << a * b << endl;
    if (b != 0) cout << "Hasil / : " << a / b << endl;
    else cout << "Tidak bisa dibagi 0" << endl;
    
    return 0;
}
