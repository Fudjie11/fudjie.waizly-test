package enum

type CRUDAction int

const (
	CRUDActionInsert CRUDAction = iota
	CRUDActionUpdate
	CRUDActionDelete
)
