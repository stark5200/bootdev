#include <stdio.h>
#include "exercise.h"

float get_average(int x, int y, int z) {
  return (float)(x + y + z)/3; 
}

void concat_strings(char *str1, const char *str2) {
  int e;
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

/// stack and heap

int main() {
  printMessageOne();
  printMessageTwo();
  printMessageThree();
  return 0;
}

void printMessageOne() {
  const char *message = "Dark mode?\n";
  printStackPointerDiff();
  printf("%s\n", message);
  //printMessageTwo();
}

void printMessageTwo() {
  const char *message = "More like...\n";
  printStackPointerDiff();
  printf("%s\n", message);
  //printMessageThree();
}

void printMessageThree() {
  const char *message = "dark roast.\n";
  printStackPointerDiff();
  printf("%s\n", message);
}

// don't touch below this line

void printStackPointerDiff() {
  static void *last_sp = NULL;
  void *current_sp;
  current_sp = __builtin_frame_address(0);
  long diff = (char*)last_sp - (char*)current_sp;
  if (last_sp == NULL){
    last_sp = current_sp;
    diff = 0;
  }
  printf("---------------------------------\n");
  printf("Stack pointer offset: %ld bytes\n", diff);
  printf("---------------------------------\n");
}

#include "snekstack.h"
#include <assert.h>
#include <stddef.h>
#include <stdlib.h>

void stack_push(stack_t *stack, void *obj) {
  // ?
}

// don't touch below this line

stack_t *stack_new(size_t capacity) {
  stack_t *stack = malloc(sizeof(stack_t));
  if (stack == NULL) {
    return NULL;
  }

  stack->count = 0;
  stack->capacity = capacity;
  stack->data = malloc(stack->capacity * sizeof(void *));
  if (stack->data == NULL) {
    free(stack);
    return NULL;
  }

  return stack;
}