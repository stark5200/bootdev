#include "munit/munit.h"

float get_average(int x, int y, int z);
void concat_strings(char *str1, const char *str2);
void printMessageOne();
void printMessageTwo();
void printMessageThree();
void printStackPointerDiff();

typedef struct Token {
    char* literal;
    int line;
    int column;
} token_t;

token_t** create_token_pointer_array(token_t* tokens, size_t count);



typedef enum SnekObjectKind {
  INTEGER, 
  STRING, 
} snek_object_kind_t;

// don't touch below this line'

typedef union SnekObjectData {
  int v_int;
  char *v_string;
} snek_object_data_t;

typedef struct SnekObject {
  snek_object_kind_t kind;
  snek_object_data_t data;
} snek_object_t;

snek_object_t new_integer(int);
snek_object_t new_string(char *str);
void format_object(snek_object_t obj, char *buffer);

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
