#include "snekstack.h"
#include <assert.h>
#include <stddef.h>
#include <stdlib.h>

void stack_push(stack_t *stack, void *obj) {
  if (stack->count == stack->capacity) {
    stack->data = realloc(stack->data, stack->capacity * 2 * sizeof(void *));
    if (stack->data == NULL) {
      return;
    }
    stack->capacity = stack->capacity * 2;
  }
  stack->data[stack->count] = obj;
  stack->count++;
}

void *stack_pop(stack_t *stack) {
  if (stack->count == 0) {
    return NULL;
  }
  stack->count--;
  return stack->data[stack->count];
}

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

void stack_free(stack_t *stack) {
  if (stack == NULL) {
    return;
  }
  if (stack->data == NULL) {
    return;
  }
  free(stack->data);
  free(stack);
}
