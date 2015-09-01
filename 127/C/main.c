#include <stdio.h>
#include <time.h>

double term(int k) {
	if (k%2 == 0) {
		return 4.0 / (2.0*(double)(k) + 1.0);
	} else {
		return -4.0 / (2.0*(double)(k) + 1.0);
	}
}

// pi performs n iterations to compute an approximation of pi.
double pi(int n) {
	double f = 0.0;
	for (int k = 0; k <= n; k++) {
		f += term(k);
	}
	return f;
}

int main() {
	int n = 1000 * 1000 * 1000;
	printf("approximating pi with %d iterations.\n", n);
	printf("%.16f\n", pi(n));

	return 0;
}
