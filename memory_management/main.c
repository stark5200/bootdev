#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
/*
#include "munit.h"
#include "exercise.h"

munit_case(RUN, test_get_average, {
  float result = get_average(3, 4, 5);
  munit_assert_double_equal(result, 4.0, "Average of 3, 4, 5 is 4");
});

munit_case(RUN, test_non_integer, {
  float result = get_average(3, 3, 5);
  munit_assert_double_equal(result, 11.0 / 3.0, "Average of 3, 3, 5 is 3.66667");
});

munit_case(SUBMIT, test_average_of_same, {
  float result2 = get_average(10, 10, 10);
  munit_assert_double_equal(result2, 10.0, "Average of 10s... is 10");
});

munit_case(SUBMIT, test_average_of_big_numbers, {
  float result3 = get_average(1050, 2050, 2075);
  munit_assert_double_equal(
      result3, 1725.0, "Bigger numbers can still get averaged, duh!"
  );
});

*/
int main() {
  /*
  MunitTest tests[] = {
      munit_test("/get_average", test_get_average),
      munit_test("/get_average_float", test_non_integer),
      munit_test("/get_average_same", test_average_of_same),
      munit_test("/get_average_big", test_average_of_big_numbers),
      munit_null_test,
  };

  MunitSuite suite = munit_suite("get_average", tests);
  */

  // Use %zu is for printing `sizeof` result
  printf("sizeof(char)   = %zu\n", sizeof(char));
  printf("sizeof(bool)   = %zu\n", sizeof(bool));
  printf("sizeof(int)   = %zu\n", sizeof(int));
  printf("sizeof(float)   = %zu\n", sizeof(float));
  printf("sizeof(double)   = %zu\n", sizeof(double));
  printf("sizeof(size_t)   = %zu\n", sizeof(size_t));
  return 0;

  //return munit_suite_main(&suite, NULL, 0, NULL);


}

/*
Pragma Once and Header Guards
We saw how .h header files are used in a previous lesson, but before we go further let's talk about a potential issue you might run into: multiple inclusions. If the same header file gets included more than once, you can end up with some nasty errors caused by redefining things like functions or structs.

 Pragma once
One simple solution (and the one we'll use for the rest of this course) is #pragma once. Adding this line to the top of a header file tells the compiler to include the file only once, even if it's referenced multiple times across your program.

// my_header.h

#pragma once

struct Point {
    int x;
    int y;
};
Copy icon
 Header Guards
Another common way to avoid multiple inclusions is with include guards, which use preprocessor directives like this:

#ifndef MY_HEADER_H
#define MY_HEADER_H

// some cool code

#endif
Copy icon
This method works by defining a unique macro for the header file. If it’s already been included, the guard prevents it from being processed again.

Throughout this course, you’ll see #pragma once in our header files. It's quicker and less error-prone than traditional include guards, and it works well with most modern compilers.
 */

struct Coordinate {
  int x;
  int y;
  int z;
};

struct Coordinate new_coord(int x, int y, int z) {
  struct Coordinate c = {
  .x = x, 
  .y = y,
  .z = z,
  };

  // or type def

  /*
    #pragma once

    typedef struct Coordinate {
    int x;
    int y;
    int z;
    } coordinate_t;

    coordinate_t new_coord(int x, int y, int z);
    coordinate_t scale_coordinate(coordinate_t coord, int factor);
   */

  return c;
}

