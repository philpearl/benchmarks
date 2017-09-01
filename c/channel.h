#ifndef CHANNEL_H_INCLUDED
#define CHANNEL_H_INCLUDED

#include <pthread.h>
#include <stdbool.h>
#include "list.h"

typedef struct channel {
    int capacity;
    int num_queued;
    List l;

    pthread_cond_t cond;
    pthread_mutex_t mutex;

    int rwaits;
    int swaits;
    bool closed;

} Channel;

void channel_init(Channel *c, int capacity);
void channel_destroy(Channel *c);
void channel_send(Channel *c, void *value);
void channel_close(Channel *c);
bool channel_receive(Channel *c, void **value);

#endif