#include "exercise.h"
#include <string.h>


token_t** create_token_pointer_array(token_t* tokens, size_t count) {
  token_t** token_pointers = (token_t**)malloc(count * sizeof(token_t*));
  if (token_pointers == NULL) {
    exit(1);
  }
  for (size_t i = 0; i < count; ++i) {
    token_t* token_pointer = (token_t*)malloc(sizeof(token_t));
    token_pointer->literal = tokens[i].literal;
    token_pointer->line = tokens[i].line;
    token_pointer->column = tokens[i].column;
    token_pointers[i] = token_pointer;
  }
  
  return token_pointers;
}

void swap_strings(char **a, char **b) {
  char *temp = *a;
  *a = *b;
  *b = temp;
}

/// Generic swap

void swap(void *vp1, void *vp2, size_t size) {
  void *temp_buffer = (void*)malloc(size);
  memcpy(temp_buffer, vp1, size);
  memcpy(vp1, vp2, size);
  memcpy(vp2, temp_buffer, size);
  free(temp_buffer);
}



/*
token_t** create_token_pointer_array(token_t* tokens, size_t count) {
  token_t** token_pointers = (token_t**)malloc(count * sizeof(token_t*));
  if (token_pointers == NULL) {
    exit(1);
  }
  for (size_t i = 0; i < count; ++i) {
    token_pointers[i] = &tokens[i];
  }
  return token_pointers;
}
*/