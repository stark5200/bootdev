#include "bootlib.h"
#include "munit.h"
#include "snekstack.h"

munit_case(RUN, create_stack_small, {
  stack_t *s = stack_new(3);
  assert_int(s->capacity, ==, 3, "Sets capacity to 3");
  assert_int(s->count, ==, 0, "No elements in the stack yet");
  assert_ptr_not_null(s->data, "Allocates the stack data");

  free(s->data);
  free(s);

  assert(boot_all_freed());
});

munit_case(SUBMIT, create_stack_large, {
  stack_t *s = stack_new(100);
  assert_int(s->capacity, ==, 100, "Sets capacity to 100");
  assert_int(s->count, ==, 0, "No elements in the stack yet");
  assert_ptr_not_null(s->data, "Allocates the stack data");

  free(s->data);
  free(s);

  assert(boot_all_freed());
});

int main() {
  MunitTest tests[] = {
    munit_test("/create_stack_small", create_stack_small),
    munit_test("/create_stack_large", create_stack_large),
    munit_null_test,
  };

  MunitSuite suite = munit_suite("snekstack", tests);

  return munit_suite_main(&suite, NULL, 0, NULL);
}


#include <stdlib.h>

#include "munit.h"
#include "snekobject.h"

munit_case(RUN, test_integer_constant, {
  assert_int(INTEGER, ==, 0, "INTEGER is defined as 0");
});

munit_case(RUN, test_integer_obj, {
  snek_object_t *obj = malloc(sizeof(snek_object_t));
  obj->kind = INTEGER;
  obj->data.v_int = 0;
  assert_int(obj->kind, ==, INTEGER, "must be INTEGER type");
  assert_int(obj->data.v_int, ==, 0, "must equal zero");

  free(obj);
});

int main() {
  MunitTest tests[] = {
    munit_test("/integer_constant", test_integer_constant),
    munit_test("/integer_obj", test_integer_obj),
    munit_null_test,
  };

  MunitSuite suite = munit_suite("object-integer-def", tests);

  return munit_suite_main(&suite, NULL, 0, NULL);
}

#include <stdlib.h>

#include "exercise.h"

token_t** create_token_pointer_array(token_t* tokens, size_t count) {
  token_t** token_pointers = malloc(count * sizeof(token_t*));
  if (token_pointers == NULL) {
    exit(1);
  }

  for (size_t i = 0; i < count; ++i) {
    token_pointers[i] = (token_t*)malloc(sizeof(token_t));
    if (token_pointers[i] == NULL) {
      exit(1);
    }
    token_pointers[i]->literal = tokens[i].literal;
    token_pointers[i]->line = tokens[i].line;
    token_pointers[i]->column = tokens[i].column;
  }

  return token_pointers;
}
