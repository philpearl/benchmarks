#ifndef LIST_H_INCLUDED
#define LIST_H_INCLUDED

typedef struct elt {
    struct elt *next;
    void *value;
} Elt;

typedef struct list {
    Elt *head;
    Elt *tail;

    Elt *free;
} List;

void list_init(List *l);
void list_destroy(List *l);
void list_push(List *l, void *v);
void *list_pop(List *l);

#endif