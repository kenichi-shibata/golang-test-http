package utils

type SQLRecordNotFoundError struct {
	Record string
}

func (e *SQLRecordNotFoundError) Error() string { return "Record: " + e.Record + " not found" }
