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
