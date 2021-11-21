package database

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
	//データ保存した最後のIDを取得
	LastInsertId() (int64, error)
	//更新したレコード数を取得
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
