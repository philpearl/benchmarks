
#include <mach/mach.h>
#include <stdlib.h>
#include "channel.h"

void channel_init(Channel *c, int capacity) {
    pthread_mutex_init(&c->mutex, NULL);
    pthread_cond_init(&c->cond, NULL);
    c->capacity = capacity;
    c->num_queued = 0;
    c->swaits = 0;
    c->rwaits = 0;
    c->closed = false;
    list_init(&c->l);
}

void channel_destroy(Channel *c) {
    list_destroy(&c->l);
    pthread_cond_destroy(&c->cond);
    pthread_mutex_destroy(&c->mutex);
}

void channel_close(Channel *c) {
    pthread_mutex_lock(&c->mutex);
    if (!c->closed) {
        c->closed = true;
        if (c->rwaits > 0) {
            pthread_cond_broadcast(&c->cond);
        }
    }
    pthread_mutex_unlock(&c->mutex);
}

void channel_send(Channel *c, void *value) {
    pthread_mutex_lock(&c->mutex);
    // Wait until there's space to send
    while (c->num_queued == c->capacity && !c->closed) {
        c->swaits++;
        pthread_cond_wait(&c->cond, &c->mutex);
        c->swaits--;
    }

    if (c->closed) {
        panic("send on closed channel");
    }
    
    // Push the value onto the list
    list_push(&c->l, value);
    c->num_queued++;

    // wake a waiting receiver as we've just added work
    if (c->rwaits > 0) {
        pthread_cond_signal(&c->cond);
    }
    
    pthread_mutex_unlock(&c->mutex);
}

bool channel_receive(Channel *c, void **value) {
    pthread_mutex_lock(&c->mutex);
    // Wait until there's something to read
    while (c->num_queued == 0 && !c->closed) {
        c->rwaits++;
        pthread_cond_wait(&c->cond, &c->mutex);
        c->rwaits--;
    }

    // Pull the next value off the list. Will be NULL if channel is closed
    // and empty
    bool have_result = false;
    if (c->num_queued > 0) {
        *value = list_pop(&c->l);
        c->num_queued--;
        have_result = true;    
    }

    // If senders are waiting, let them send
    if (c->swaits > 0) {
        pthread_cond_signal(&c->cond);
    }

    pthread_mutex_unlock(&c->mutex);

    return have_result;
}