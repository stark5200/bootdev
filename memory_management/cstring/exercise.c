#include "exercise.h"

int smart_append(TextBuffer* dest, const char* src) {
  if (dest == NULL || src == NULL) {
    return 0;
  };
  const int max_buffer = 64;
  int src_len = strlen(src);
  int available_space = max_buffer - dest->length - 1;
  if (src_len > available_space) {
    strncat(dest->buffer, src, available_space);
    dest->length = 63;
    return 0;
  };
  strcat(dest->buffer, src);
  dest->length += src_len;
  return 1;
}


typedef union {
  int value;
  unsigned int err;
} val_or_err_t;

int main() {
  val_or_err_t lanes_score = {
    .value = -420
  };
  printf("value (set): %d\n", lanes_score.value);
  printf("err (unset): %u\n", lanes_score.err);

  
  val_or_err_t teejs_score = {
    .err = UINT_MAX
  };
  printf("value (unset): %d\n", teejs_score.value);
  printf("err (set): %u\n", teejs_score.err);
  
}
