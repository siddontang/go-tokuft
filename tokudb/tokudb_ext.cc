#include "tokudb_ext.h"
#include <sys/stat.h>

extern "C" {

int env_open(DB_ENV *env, const char *dir, uint32_t flags, int mode) {
    return env->open(env, dir, flags, mode);
}

int env_close(DB_ENV *env, uint32_t flags) {
    return env->close(env, flags);
}

int env_txn_begin(DB_ENV * env, DB_TXN * parentTx, DB_TXN **tx, uint32_t flags){
    return env->txn_begin(env, parentTx, tx, flags);
}

int txn_commit(DB_TXN *tx, uint32_t flags) {
    return tx->commit(tx, flags);
}

int txn_abort(DB_TXN *tx) {
    return tx->abort(tx);
}

int db_open(DB* db, DB_TXN *tx, const char* name, uint32_t flags, int mode) {
    return db->open(db, tx, name, NULL, DB_BTREE, flags, mode);
}

int db_close(DB *db, uint32_t flags) {
    return db->close(db, flags);
}

int db_get(DB *db, DB_TXN *tx, DBT *key, DBT *value, uint32_t flags) {
    return db->get(db, tx, key, value, flags);
}

int db_put(DB *db, DB_TXN *tx, DBT *key, DBT *value, uint32_t flags) {
    return db->put(db, tx, key, value, flags);
}

int db_del(DB *db, DB_TXN *tx, DBT *key, uint32_t flags) {
    return db->del(db, tx, key, flags);
}


int db_cursor(DB *db, DB_TXN *tx, DBC** c, uint32_t flags) {
    return db->cursor(db, tx, c, flags);
}

int cursor_close(DBC *c) {
    return c->c_close(c);
}

int cursor_get(DBC *c, DBT *key, DBT *value, uint32_t flags) {
    return c->c_get(c, key, value, flags);
}

}