#include "exercise.h"

float get_average(int x, int y, int z) {
  return (float)(x + y + z)/3; 
}

void concat_strings(char *str1, const char *str2) {
  int i = 0;
  int j = 0;
  while (*(str1 + i) != '\0') {
    i++;
  };
  while (*(str2 + j) != '\0') {
    *(str1 + i) = *(str2 + j);
    j++;
    i++;
  };
  *(str1 + i) = '\0';
}
