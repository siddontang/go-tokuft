#ifndef TOKUDB_EXT
#define TOKUDB_EXT

#include "tokudb.h"

#ifdef __cplusplus
extern "C" {
#endif

int env_open(DB_ENV *, const char *, uint32_t, int);
int env_close(DB_ENV *, uint32_t);
int env_txn_begin(DB_ENV *, DB_TXN *, DB_TXN **, uint32_t);

int txn_commit(DB_TXN*, uint32_t);
int txn_abort(DB_TXN*);

int db_open(DB*, DB_TXN*, const char*, uint32_t, int);
int db_close(DB*, uint32_t);

int db_get(DB *, DB_TXN *, DBT *, DBT *, uint32_t);
int db_put(DB *, DB_TXN *, DBT *, DBT *, uint32_t);
int db_del(DB *, DB_TXN *, DBT *, uint32_t);

int db_cursor(DB *, DB_TXN *, DBC **, uint32_t);

int cursor_close(DBC*);
int cursor_get(DBC *, DBT *, DBT *, uint32_t);

#ifdef __cplusplus
}
#endif

#endif