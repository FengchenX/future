#ifndef MULTI_THREAD_H
#define MULTI_THREAD_H

#ifdef __golang__
extern "C" {
#endif

int multi_thread_setup(void);


void multi_thread_cleanup(void);

#ifdef __golang__
}
#endif

#endif
