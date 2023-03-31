package model

type MD5Sum struct {
	Name string `db:"name"`
	Sum  string `db:"sum"`
}
