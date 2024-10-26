typedef enum SnekObjectKind {
  INTEGER
} snek_object_kind_t;

typedef union SnekObjectData {
  int v_int;
} snek_object_data_t;

typedef struct SnekObject {
  snek_object_kind_t kind;
  snek_object_data_t data;
} snek_object_t;

snek_object_t *new_snek_integer(int value);

#include "snekstack.h"
#include "stdlib.h"
#include <string.h>

void stack_push_multiple_types(stack_t *s) {
  float *pi = malloc(sizeof(float));
  if (pi != NULL) {
    *pi = 3.14;
  }
  char *snekchar = malloc(strlen("Sneklang is blazingly slow!") + 1);
  if (snekchar != NULL) {
    strcpy(snekchar, "Sneklang is blazingly slow!");
  }
  
  stack_push(s, pi);
  stack_push(s, snekchar);
  
}
