#include <string.h>
#include <limits.h>
#include "../munit/munit.h"

typedef struct {
  char buffer[64];
  size_t length;
} TextBuffer;

int smart_append(TextBuffer* dest, const char* src);
