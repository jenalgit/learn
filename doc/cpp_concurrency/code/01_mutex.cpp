#include <iostream>
#include <vector>
#include <thread>
#include <mutex>
using namespace std;

int accum = 0;
mutex accum_mutex;

void square(int x) {
	int temp = x * x;
	accum_mutex.lock();
	accum += temp;
	accum_mutex.unlock();
}

int main() {
	vector<thread> ths;
	for (int i = 1; i <= 20; i++) {
		ths.push_back(thread(&square, i));
	}

	// & to retrieve a reference and not a copy of the object
	// , since join changes the nature of the object.
	// for (auto& th : ths) {
	for (vector<thread>::iterator th = ths.begin(); th != ths.end(); th++)
		(*th).join(); // join() on each thread
		// blocks until the thread finishes
	cout << "accum = " << accum << endl;
}

/*
$ g++ -std=c++11 01_mutex.cpp -pthread
accum = 2870


2870 is the correct answer
but:

$ for i in {1..1000}; do ./a.out; done | sort | uniq -c
   1000 accum = 2870

no race conditions!
*/
