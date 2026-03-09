package gocockroachdb

import "database/sql"

// CockDbTxMock CockDbTxMock
type CockDbTxMock struct {
	Tx         *sql.Tx
	CockDbMock *CockDbMock
}

// Insert Insert
func (t *CockDbTxMock) Insert(query string, args ...any) (bool, int64) {
	var rtn = false
	var id int64
	if !t.CockDbMock.mockInsertSuccess1Used {
		t.CockDbMock.mockInsertSuccess1Used = true
		rtn = t.CockDbMock.MockInsertSuccess1
		id = t.CockDbMock.MockInsertID1
	} else if !t.CockDbMock.mockInsertSuccess2Used {
		t.CockDbMock.mockInsertSuccess2Used = true
		rtn = t.CockDbMock.MockInsertSuccess2
		id = t.CockDbMock.MockInsertID2
	} else if !t.CockDbMock.mockInsertSuccess3Used {
		t.CockDbMock.mockInsertSuccess3Used = true
		rtn = t.CockDbMock.MockInsertSuccess3
		id = t.CockDbMock.MockInsertID3
	} else if !t.CockDbMock.mockInsertSuccess4Used {
		t.CockDbMock.mockInsertSuccess4Used = true
		rtn = t.CockDbMock.MockInsertSuccess4
		id = t.CockDbMock.MockInsertID4
	} else if !t.CockDbMock.mockInsertSuccess5Used {
		t.CockDbMock.mockInsertSuccess5Used = true
		rtn = t.CockDbMock.MockInsertSuccess5
		id = t.CockDbMock.MockInsertID5
	} else if !t.CockDbMock.mockInsertSuccess6Used {
		t.CockDbMock.mockInsertSuccess6Used = true
		rtn = t.CockDbMock.MockInsertSuccess6
		id = t.CockDbMock.MockInsertID6
	} else if !t.CockDbMock.mockInsertSuccess7Used {
		t.CockDbMock.mockInsertSuccess7Used = true
		rtn = t.CockDbMock.MockInsertSuccess7
		id = t.CockDbMock.MockInsertID7
	} else if !t.CockDbMock.mockInsertSuccess8Used {
		t.CockDbMock.mockInsertSuccess8Used = true
		rtn = t.CockDbMock.MockInsertSuccess8
		id = t.CockDbMock.MockInsertID8
	}
	return rtn, id
}

// Update Update
func (t *CockDbTxMock) Update(query string, args ...any) bool {
	var rtn = false
	if !t.CockDbMock.mockUpdateSuccess1Used {
		t.CockDbMock.mockUpdateSuccess1Used = true
		rtn = t.CockDbMock.MockUpdateSuccess1
	} else if !t.CockDbMock.mockUpdateSuccess2Used {
		t.CockDbMock.mockUpdateSuccess2Used = true
		rtn = t.CockDbMock.MockUpdateSuccess2
	} else if !t.CockDbMock.mockUpdateSuccess3Used {
		t.CockDbMock.mockUpdateSuccess3Used = true
		rtn = t.CockDbMock.MockUpdateSuccess3
	} else if !t.CockDbMock.mockUpdateSuccess4Used {
		t.CockDbMock.mockUpdateSuccess4Used = true
		rtn = t.CockDbMock.MockUpdateSuccess4
	}
	return rtn
}

// Delete Delete
func (t *CockDbTxMock) Delete(query string, args ...any) bool {
	var rtn = false
	if !t.CockDbMock.mockDeleteSuccess1Used {
		t.CockDbMock.mockDeleteSuccess1Used = true
		rtn = t.CockDbMock.MockDeleteSuccess1
	} else if !t.CockDbMock.mockDeleteSuccess2Used {
		t.CockDbMock.mockDeleteSuccess2Used = true
		rtn = t.CockDbMock.MockDeleteSuccess2
	} else if !t.CockDbMock.mockDeleteSuccess3Used {
		t.CockDbMock.mockDeleteSuccess3Used = true
		rtn = t.CockDbMock.MockDeleteSuccess3
	} else if !t.CockDbMock.mockDeleteSuccess4Used {
		t.CockDbMock.mockDeleteSuccess4Used = true
		rtn = t.CockDbMock.MockDeleteSuccess4
	} else if !t.CockDbMock.mockDeleteSuccess5Used {
		t.CockDbMock.mockDeleteSuccess5Used = true
		rtn = t.CockDbMock.MockDeleteSuccess5
	} else if !t.CockDbMock.mockDeleteSuccess6Used {
		t.CockDbMock.mockDeleteSuccess6Used = true
		rtn = t.CockDbMock.MockDeleteSuccess6
	} else if !t.CockDbMock.mockDeleteSuccess7Used {
		t.CockDbMock.mockDeleteSuccess7Used = true
		rtn = t.CockDbMock.MockDeleteSuccess7
	} else if !t.CockDbMock.mockDeleteSuccess8Used {
		t.CockDbMock.mockDeleteSuccess8Used = true
		rtn = t.CockDbMock.MockDeleteSuccess8
	}
	return rtn
}

// Commit Commit
func (t *CockDbTxMock) Commit() bool {
	return t.CockDbMock.MockCommitSuccess
}

// Rollback Rollback
func (t *CockDbTxMock) Rollback() bool {
	return t.CockDbMock.MockRollbackSuccess
}
