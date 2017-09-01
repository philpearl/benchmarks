
#include <stdlib.h>

#include "list.h"

void list_init(List *l) {
    l->free = NULL;
    l->head = NULL;
    l->tail = NULL;
}

void list_destroy(List *l) {
    while (list_pop(l) != NULL);

    while (l->free != NULL) {
        Elt *elt  = l->free;
        l->free = elt->next;

        free(elt);
    }
}

void list_push(List *l, void *v) {
    Elt *elt = l->free;
    if (elt == NULL) {
        // Allocate a new element
        elt = (Elt *)malloc(sizeof(Elt));        
    } else {
        l->free = elt->next;
    }

    elt->value = v;
    elt->next = NULL;

    if (l->tail == NULL) {
        l->head = elt;
        l->tail = elt;
    } else {
        l->tail->next = elt;
        l->tail = elt;
    }
}

void *list_pop(List *l) {
    Elt *elt = l->head;
    if (elt == NULL) {
        return NULL;
    }

    l->head = elt->next;
    elt->next = NULL;
    if (l->head == NULL) {
        l->tail = NULL;
    }

    void *v = elt->value;

    // Add elt to free list
    elt->next = l->free;
    l->free = elt;

    return v;
}