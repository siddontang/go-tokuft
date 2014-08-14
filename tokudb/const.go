package tokudb

// #include "tokudb.h"
import "C"

const (
	CREATE               = C.DB_CREATE
	PRIVATE              = C.DB_PRIVATE
	RDONLY               = C.DB_RDONLY
	RECOVER              = C.DB_RECOVER
	THREAD               = C.DB_THREAD
	TXN_NOSYNC           = C.DB_TXN_NOSYNC
	KEYFIRST             = C.DB_KEYFIRST
	KEYLAST              = C.DB_KEYLAST
	NOOVERWRITE          = C.DB_NOOVERWRITE
	NODUPDATA            = C.DB_NODUPDATA
	NOOVERWRITE_NO_ERROR = C.DB_NOOVERWRITE_NO_ERROR
	OPFLAGS_MASK         = C.DB_OPFLAGS_MASK
	AUTO_COMMIT          = C.DB_AUTO_COMMIT
	INIT_LOCK            = C.DB_INIT_LOCK
	INIT_LOG             = C.DB_INIT_LOG
	INIT_MPOOL           = C.DB_INIT_MPOOL
	INIT_TXN             = C.DB_INIT_TXN

	TXN_WRITE_NOSYNC  = C.DB_TXN_WRITE_NOSYNC
	TXN_NOWAIT        = C.DB_TXN_NOWAIT
	TXN_SYNC          = C.DB_TXN_SYNC
	TXN_SNAPSHOT      = C.DB_TXN_SNAPSHOT
	READ_UNCOMMITTED  = C.DB_READ_UNCOMMITTED
	READ_COMMITTED    = C.DB_READ_COMMITTED
	INHERIT_ISOLATION = C.DB_INHERIT_ISOLATION
	SERIALIZABLE      = C.DB_SERIALIZABLE
	TXN_READ_ONLY     = C.DB_TXN_READ_ONLY
)

//cursor
const (
	FIRST             = C.DB_FIRST
	LAST              = C.DB_LAST
	CURRENT           = C.DB_CURRENT
	NEXT              = C.DB_NEXT
	PREV              = C.DB_PREV
	SET               = C.DB_SET
	SET_RANGE         = C.DB_SET_RANGE
	SET_RANGE_REVERSE = C.DB_SET_RANGE_REVERSE
)
