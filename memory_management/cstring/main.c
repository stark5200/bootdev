#include "exercise.h"

munit_case(RUN, test_return_0_for_null_value, {
  TextBuffer dest;
  const char* src; 
  int result = smart_append(&dest, src);
  munit_assert_int(result, ==, 0, "Should return 0 for null value");
});

munit_case(RUN, test_smart_append_empty_buffer, {
  TextBuffer dest;
  strcpy(dest.buffer, "");
  dest.length = 0;
  const char* src = "Hello";
  int result = smart_append(&dest, src);
  munit_assert_int(result, ==, 1, "Should return 1 for successful append");
  munit_assert_string_equal(dest.buffer, "Hello", "Buffer should contain 'Hello'");
  munit_assert_int(dest.length, ==, 5, "Length should be 5");
});

munit_case(SUBMIT, test_smart_append_full_buffer, {
  TextBuffer dest;
  strcpy(dest.buffer, "This is a very long string that will fill up the entire buffer.");
  dest.length = 63;
  const char* src = " Extra";
  int result = smart_append(&dest, src);
  munit_assert_int(result, ==, 0, "Should return 0 for unsuccessful append");
  munit_assert_string_equal(dest.buffer, "This is a very long string that will fill up the entire buffer.", "Buffer should remain unchanged");
  munit_assert_int(dest.length, ==, 63, "Length should remain 63");
});

munit_case(SUBMIT, test_smart_append_overflow, {
  TextBuffer dest;
  strcpy(dest.buffer, "This is a long string");
  dest.length = 21;
  const char* src = " that will fill the whole buffer and leave no space for some of the chars.";
  int result = smart_append(&dest, src);
  munit_assert_int(result, ==, 0, "Should return 0 for overflow append");
  munit_assert_string_equal(dest.buffer, "This is a long string that will fill the whole buffer and leave", "Buffer should be filled to capacity");
  munit_assert_int(dest.length, ==, 63, "Length should be 63 after overflow append");
});

int main() {
  MunitTest tests[] = {
    munit_test("/test_return_0_for_null_value", test_return_0_for_null_value),
    munit_test("/test_smart_append_empty_buffer", test_smart_append_empty_buffer),
    munit_test("/test_smart_append_full_buffer", test_smart_append_full_buffer),
    munit_test("/test_smart_append_overflow", test_smart_append_overflow),
    munit_null_test,
  };

  MunitSuite suite = munit_suite("smart_append", tests);

  return munit_suite_main(&suite, NULL, 0, NULL);
}

/*

///Union

Now that we understand structs and enums, we can learn aboutunions: a combination of the two concepts.

This is not the kind of union that $300k-earning Google employees fight for because they are "underpaid" and "don't have enough oat milk in the office kitchen". No, this feature is one that even Golang doesn't have (probably because they were worried about getting fired from Google for just mentioning the word!)
Unions in C can hold one of several types. They're like a less-strict sum type from the world of functional programming. Here's an example union:

typedef union AgeOrName {
  int age;
  char *name;
} age_or_name_t;
Copy icon
The age_or_name_t type can hold either an int or a char *, but not both at the same time (that would be a struct). We provide the list of possible types is so that the C compiler knows the maximum potential memory requirement, and can account for that. This is how the union is used:

age_or_name_t lane = { .age = 29 };
printf("age: %d\n", lane.age);
// age: 29
Copy icon
Here's where it gets interesting. What happens if we try to access the name field (even though we set the age field)?

printf("name: %s\n", lane.name);
// name:
Copy icon
We get... nothing? To be more specific, we get undefined behavior. A union only reserves enough space to hold the largest type in the union and then all of the fields use the same memory. So when we set .age to 29, we are writing the integer representation of 29 to the memory of the lane union:

0000 0000 0000 0000 0000 0000 0001 1101
Copy icon
Then if we try to access .name, we read from the same block of memory but try to interpret the bytes as a char *, which is why we get garbage (which is interpreted as nothing in this case). Put simply, setting the value of .age overwrites the value of .name and vice versa, and you should only access the field that you set.

 Assignment
Sneklang is going to need objects. We'll hand-code those objects, and Sneklang developers will use them to store dynamic variables, kinda like Python. Everything is an object, even simple integers and strings!

Take a look at the SnekObject struct in exercise.h. It has a kind field that stores the type of the object (like INTEGER or STRING) and a data field that stores the actual data.

Create a snek_object_kind_t enum type in exercise.h. It's the one used as the kind field of the provided SnekObject. It's an enum that can be an INTEGER (0) or a STRING (1).
Complete the format_object function in exercise.c that uses a switch on the .kind of a snek_object_t and writes a formatted string to the associated buffer.
For an integer, write int:n to the buffer, replacing n with the integer value
For a string, write string:str to the buffer, replacing str with the string value
You can use sprintf to write the formatted string to the buffer
 */