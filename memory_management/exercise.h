float get_average(int x, int y, int z);
void concat_strings(char *str1, const char *str2);
void printMessageOne();
void printMessageTwo();
void printMessageThree();
void printStackPointerDiff();


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

