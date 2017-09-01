
#include <stdio.h>
#include "channel.h"
#include "channel_test.h"

void test_channel() {
    Channel c;

    channel_init(&c, 10);

    for (int i = 0; i < 10; i++) {
        channel_send(&c, (void *)(i+1));
    }

    channel_close(&c);
    
    void *v;
    while (channel_receive(&c, &v)) {
        printf("received %d\n", (int)(v));
    }

    channel_destroy(&c);
}