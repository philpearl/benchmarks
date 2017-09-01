#include <stdio.h>
#include <time.h>
#include <mach/mach_time.h>
#include <pthread.h>

#include "channel_test.h"
#include "channel.h"

void channel_benchmark(char *name, int N);
void basic_work(char *name, int N, int wwork);
void benchmark_done(char *name, int N, uint64_t start);
void run_benchmark(char * name, int N, void (*f)(int N));
void* work_thread(void *data);
void work_on_thread(char *name, int N, int wwork);
void work_on_thread_pool(char *name, int N, int work, int threads);
void *work_pool_thread(void *data);

int main() {
    test_channel();

    basic_work("basic_work_1000", 10000, 1000);
    basic_work("basic_work_10000", 1000, 10000);
    basic_work("basic_work_100000", 1000, 100000);

    work_on_thread("work_on_thread_0", 1000, 0);
    work_on_thread("work_on_thread_1000", 1000, 1000);
    work_on_thread("work_on_thread_10000", 100, 10000);
    work_on_thread("work_on_thread_100000", 100, 100000);

    channel_benchmark("channel", 1000);

    work_on_thread_pool("work_on_1_thread_pool_0", 100, 0, 1);
    work_on_thread_pool("work_on_8_thread_pool_0", 100, 0, 8);
    work_on_thread_pool("work_on_1_thread_pool_1000", 100, 1000, 1);
    work_on_thread_pool("work_on_8_thread_pool_1000", 100, 1000, 8);
    work_on_thread_pool("work_on_1_thread_pool_100000", 10, 100000, 1);        
    work_on_thread_pool("work_on_8_thread_pool_100000", 100, 100000, 8);
}

/*
"basic work", 100000 iterations in 233258021ns, 2332ns per iteration
"work on thread", 100000 iterations in 1983044612ns, 19830ns per iteration
"channel", 100000 iterations in 5165249052ns, 51652ns per iteration
"work on thread pool", 100 iterations in 1139876322ns, 11398763ns per iteration

Starting & joining a thread to do work takes ~17 µs. The work takes 2 µs, so the starting & joining and context switch takes ~15 µs.
Writing and reading from a channel 1000 times takes 51 µs, so one send & receive takes 51 ns
1000 requests to do work on a pool of threads takes 11 ms, so 11 µs per request. The work takes 2 µs, so 9 µs must be for the context switches. 9 µs per pair of context switches

With work 100,000 adds
"basic work", 100 iterations in 29342364ns, 293423ns per iteration
"work on thread", 100 iterations in 31753460ns, 317534ns per iteration
"channel", 100 iterations in 5132274ns, 51322ns per iteration
"work on thread pool", 100 iterations in 11515749329ns, 115157493ns per iteration

With 4 threads
"work on thread pool", 100 iterations in 5761974186ns, 57619741ns per iteration

basic work scales as you expect, taking ~290µs.
Starting & joining a thread to do work takes ~317µs. The work takes ~290µs, so the overhead is ~27µs
1000 requests to do work on a pool of 2 threads takes 115 ms, so 115 µs per request. The work takes 290 µs, so we've magicked it away.
With 4 threads things look better again.
Better with 8
No worse with 100
*/

void channel_benchmark(char *name, int N) {
    Channel c;
    channel_init(&c, 1000);

    uint64_t start = mach_absolute_time();
    int n;
    for (n = 0; n < N; n++) {
        for (int i = 0; i < 1000; i ++) {
            channel_send(&c, (void *)1);
        }
        for (int i = 0; i < 1000; i ++) {
            void *v;
            channel_receive(&c, &v);
        }
    }
    benchmark_done(name, N, start);
    channel_destroy(&c);
}

// basic_work is a benchmark for some baseline work. The work is just adding to a 
// variable 1000 times
void basic_work(char *name, int N, int wwork) {
    uint64_t start = mach_absolute_time();
    int n;
    for (n = 0; n < N; n++) {
        int total = 0;
        for (int i = 0; i < wwork; i++) {
            total++;
        }
        if (total != wwork) {
            printf("total is %d\n", total);
        }        
    }
    benchmark_done(name, N, start);
}

// work_on_thread is a benchmark that starts a thread, does the baseline work,
// then joins the thread and gets the result
void work_on_thread(char *name, int N, int wwork) {
    uint64_t start = mach_absolute_time();
    int n;
    long total = 0;
    for (n = 0; n < N; n++) {
        pthread_t thread;

        int err = pthread_create(&thread, NULL, &work_thread, (void *)wwork);
        if (err != 0) {
            printf("Error creating thread. %d\n", err);
            return;
        }

        void * return_val;
        err = pthread_join(thread, &return_val);
        if (err != 0) {
            printf("Error joining thread. %d\n", err);
            return;      
        }
        total += (long)(return_val);
    }
    benchmark_done(name, N, start);
}

void* work_thread(void *data) {
    int wwork = (int)(data);
    long total = 0;
    int i;
    for (i = 0; i < wwork; i++) {
        total++;
    }
    if (total != wwork) {
        printf("total is %ld\n", total);
    }        
    return (void *)(total);
}


typedef struct poolData {
    Channel in;
    Channel out;
    int work;
} PoolData;

void work_on_thread_pool(char *name, int N, int work, int num_threads) {
    PoolData pd;

    // Start the thread pool
    channel_init(&pd.in, 1000);
    channel_init(&pd.out, 1000);
    pd.work = work;

    pthread_t threads[100];
    for (int i = 0; i < num_threads; i++) {
        pthread_create(&threads[i], NULL, &work_pool_thread, &pd);
    }

    uint64_t start = mach_absolute_time();
    for (int k = 0; k < N; k++) {
        // Send a 1000 requests for work
        for (int i = 0; i < 1000; i++) {
            channel_send(&pd.in, (void *)1);
        }
        int total = 0;
        for (int i = 0; i < 1000; i++) {
            void *v;
            channel_receive(&pd.out, &v);
            total += (int)(v);
        }
        if (total != 1000*work) {
            printf("total not as expected. %d\n", total);
        }
    }
    benchmark_done(name, N, start);

    // Terminate the thread pool
    channel_close(&pd.in);
    for (int i = 0; i < num_threads; i++) {
        pthread_join(threads[i], NULL);
    }
    channel_close(&pd.out);
    channel_destroy(&pd.in);
    channel_destroy(&pd.out);
}

void *work_pool_thread(void *data) {
    PoolData *pd = (PoolData *)(data);
    void *v;
    while (channel_receive(&pd->in, &v)) {
        int total = 0;
        for (int l = 0; l < pd->work; l++) {
            total++;
        }

        channel_send(&pd->out, (void *)total);
    }
    return NULL;
}

char* units[5] = {
    "ns",
    "µs",
    "ms",
    "s",
    "ks"
};

void benchmark_done(char *name, int N, uint64_t start) {
    uint64_t duration = mach_absolute_time() - start;
    double perIter = (double)(duration)/N;
    int u;
    for (u = 0; perIter > 1000; u++) {
        perIter = perIter/1000;
    }

    printf("\"%s\", %d iterations in %lluns, %.3f %s per iteration\n", name, N, duration, perIter, units[u]);
}