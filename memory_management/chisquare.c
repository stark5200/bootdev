#include<stdio.h>
#include<math.h>
#include <stdbool.h>

// at alpha = 0.05 and df = 5 --> cv = 11.07

bool chiGoodFit(float observed[], float expected[], float cv) {
    if (sizeof(observed) != sizeof(expected)) {
        printf("observed and expected arrays do not match in size\n");
        return false;
    }

    int size = sizeof(observed) / sizeof(observed[0]);
    float chi = 0.0;

    for (int i = 0; i < size; i++) {
        chi += ( pow( (observed[i]-expected[i]), 2) / expected[i] );
    }

    if (chi >= cv) {
        printf("observed and expected values fail the chi square goodness of fit test.\n");
        return false;
    }

    printf("observed and expected values pass the chi square goodness of fit test.\n");
    return true;
}


int main() {

  float RiggedResults[] = {3.0, 9.0, 11.0, 17.0, 24.0, 36.0};
  float NormalResults[] = {15.0, 12.0, 18.0, 25.0, 14.0, 16.0};
  float expectedResults[] = {16.66, 16.67, 16.67, 16.66, 16.67, 16.6};

  printf("testing Rigged Results:\n");
  bool result1 = chiGoodFit(RiggedResults, expectedResults, 11.07);
  printf("testing Normal Results:\n");
  bool result2 = chiGoodFit(NormalResults, expectedResults, 11.07);

  printf("does the first set follow a uniform distribution? %s\n", result1 ? "true" : "false");
  printf("does the second set follow a uniform distribution? %s\n", result2 ? "true" : "false");

  return 0;
}