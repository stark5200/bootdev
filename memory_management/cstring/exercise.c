#include <string.h>
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
