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

/*
Low Level Stack
If you've taken our data structures course, you've already implemented a stack. We're going to implement a stack again, but this time we're going to do it while manually managing the memory of generic pointers!

We'll get to have our first deeper exploration of "generics" in C (remember, that just means void *) as well as creating a data structure we will later use in our mark-and-sweep garbage collector.

 Assignment
Take a look at snekstack.h, specifically the Stack struct.

count is the number of elements in the stack.
capacity is the number of elements the stack can hold before it needs to be resized in memory.
data is a pointer to all the generic data.
Implement the stack_new function:

Allocate memory for a new Stack struct on the heap.
If allocation fails, return NULL.
Initialize the count to 0.
Initialize the capacity to the given value.
Initialize the data by allocating enough memory for capacity number of void * pointers.
If alloction fails, free the Stack struct and return NULL.
Return the new Stack struct.
 */