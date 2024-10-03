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


void format_object(snek_object_t obj, char *buffer) {
  switch (obj.kind) {
    case INTEGER:
      sprintf(buffer, "int:%d", obj.data.v_int);
      break;
    case STRING:
      sprintf(buffer, "string:%s", obj.data.v_int);
      break;
  }
}

// don't touch below this line'

snek_object_t new_integer(int i) {
  return (snek_object_t){
    .kind = INTEGER,
    .data = {.v_int = i}
  };
}

snek_object_t new_string(char *str) {
  // NOTE: We will learn how to copy this data later.
  return (snek_object_t){
    .kind = STRING,
    .data = {.v_string = str}
  };
}