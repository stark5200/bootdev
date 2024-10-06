#include "munit/munit.h"
#include <stdlib.h>
#include <string.h>

typedef struct Token {
    char* literal;
    int line;
    int column;
} token_t;

typedef short unsigned int size_t;

token_t** create_token_pointer_array(token_t* tokens, size_t count);
