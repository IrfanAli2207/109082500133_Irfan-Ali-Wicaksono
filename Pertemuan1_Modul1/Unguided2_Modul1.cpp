#include <iostream>
using namespace std;

int main() {
    int n;
    string s[] = {"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"};
    
    cout << "Masukkan angka: ";
    cin >> n;
    cout << n << " : ";

    if (n == 0) {
        cout << "nol";
    } else if (n == 100) {
        cout << "seratus";
    } else if (n < 10) {
        cout << s[n];
    } else if (n == 10) {
        cout << "sepuluh";
    } else if (n == 11) {
        cout << "sebelas";
    } else if (n < 20) {
        cout << s[n % 10] << " belas";
    } else {
        int p = n / 10;
        int b = n % 10;
        cout << s[p] << " puluh";
        if (b != 0) cout << " " << s[b];
    }
    
    cout << endl;
    return 0;
}
