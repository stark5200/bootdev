#include <stddef.h>

typedef struct Stack {
  size_t count;
  size_t capacity;
  void **data;
} stack_t;

stack_t *stack_new(size_t capacity);
void stack_push(stack_t *stack, void *obj);
void *stack_pop(stack_t *stack);
void stack_free(stack_t *stack);

/*
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
If allocation fails, free the Stack struct and return NULL.
Return the new Stack struct.
*/
