package gocockroachdb

import "database/sql"

type CockDbMock struct {
	Host        string
	User        string
	Password    string
	Database    string
	Port        string
	Sslmode     string
	Sslrootcert string
	db          *sql.DB
	err         error

	MockConnectSuccess  bool
	MockCloseSuccess    bool
	MockCommitSuccess   bool
	MockRollbackSuccess bool

	MockInsertSuccess1     bool
	MockInsertSuccess2     bool
	MockInsertSuccess3     bool
	MockInsertSuccess4     bool
	MockInsertSuccess5     bool
	MockInsertSuccess6     bool
	MockInsertSuccess7     bool
	MockInsertSuccess8     bool
	mockInsertSuccess1Used bool
	mockInsertSuccess2Used bool
	mockInsertSuccess3Used bool
	mockInsertSuccess4Used bool
	mockInsertSuccess5Used bool
	mockInsertSuccess6Used bool
	mockInsertSuccess7Used bool
	mockInsertSuccess8Used bool
	MockInsertID1          int64
	MockInsertID2          int64
	MockInsertID3          int64
	MockInsertID4          int64
	MockInsertID5          int64
	MockInsertID6          int64
	MockInsertID7          int64
	MockInsertID8          int64

	MockUpdateSuccess1     bool
	MockUpdateSuccess2     bool
	MockUpdateSuccess3     bool
	MockUpdateSuccess4     bool
	mockUpdateSuccess1Used bool
	mockUpdateSuccess2Used bool
	mockUpdateSuccess3Used bool
	mockUpdateSuccess4Used bool

	MockDeleteSuccess1     bool
	MockDeleteSuccess2     bool
	MockDeleteSuccess3     bool
	MockDeleteSuccess4     bool
	MockDeleteSuccess5     bool
	MockDeleteSuccess6     bool
	MockDeleteSuccess7     bool
	MockDeleteSuccess8     bool
	mockDeleteSuccess1Used bool
	mockDeleteSuccess2Used bool
	mockDeleteSuccess3Used bool
	mockDeleteSuccess4Used bool
	mockDeleteSuccess5Used bool
	mockDeleteSuccess6Used bool
	mockDeleteSuccess7Used bool
	mockDeleteSuccess8Used bool

	MockTestRow *DbRow

	mockRow1Used bool
	mockRow2Used bool
	mockRow3Used bool
	mockRow4Used bool
	mockRow5Used bool
	mockRow6Used bool
	mockRow7Used bool
	mockRow8Used bool
	MockRow1     *DbRow
	MockRow2     *DbRow
	MockRow3     *DbRow
	MockRow4     *DbRow
	MockRow5     *DbRow
	MockRow6     *DbRow
	MockRow7     *DbRow
	MockRow8     *DbRow

	mockRows1Used bool
	mockRows2Used bool
	mockRows3Used bool
	mockRows4Used bool
	mockRows5Used bool
	mockRows6Used bool
	mockRows7Used bool
	mockRows8Used bool
	MockRows1     *DbRows
	MockRows2     *DbRows
	MockRows3     *DbRows
	MockRows4     *DbRows
	MockRows5     *DbRows
	MockRows6     *DbRows
	MockRows7     *DbRows
	MockRows8     *DbRows

	MockExecSuccess bool
	MockExecError   error
}

// Connect Connect
func (c *CockDbMock) Connect() bool {
	return c.MockConnectSuccess
}

// New New
func (c *CockDbMock) New() Database {
	return c
}

// BeginTransaction BeginTransaction
func (c *CockDbMock) BeginTransaction() Transaction {
	var trans Transaction
	var ctx CockDbTxMock
	ctx.CockDbMock = c
	trans = &ctx
	return trans
}

// Test Test
func (c *CockDbMock) Test(query string, args ...any) *DbRow {
	return c.MockTestRow
}

// Insert Insert
func (c *CockDbMock) Insert(query string, args ...any) (bool, int64) {
	var rtn = false
	var id int64
	if !c.mockInsertSuccess1Used {
		c.mockInsertSuccess1Used = true
		rtn = c.MockInsertSuccess1
		id = c.MockInsertID1
	} else if !c.mockInsertSuccess2Used {
		c.mockInsertSuccess2Used = true
		rtn = c.MockInsertSuccess2
		id = c.MockInsertID2
	} else if !c.mockInsertSuccess3Used {
		c.mockInsertSuccess3Used = true
		rtn = c.MockInsertSuccess3
		id = c.MockInsertID3
	} else if !c.mockInsertSuccess4Used {
		c.mockInsertSuccess4Used = true
		rtn = c.MockInsertSuccess4
		id = c.MockInsertID4
	}
	return rtn, id
}

// Update Update
func (c *CockDbMock) Update(query string, args ...any) bool {
	var rtn = false
	if !c.mockUpdateSuccess1Used {
		c.mockUpdateSuccess1Used = true
		rtn = c.MockUpdateSuccess1
	} else if !c.mockUpdateSuccess2Used {
		c.mockUpdateSuccess2Used = true
		rtn = c.MockUpdateSuccess2
	} else if !c.mockUpdateSuccess3Used {
		c.mockUpdateSuccess3Used = true
		rtn = c.MockUpdateSuccess3
	} else if !c.mockUpdateSuccess4Used {
		c.mockUpdateSuccess4Used = true
		rtn = c.MockUpdateSuccess4
	}
	return rtn
}

// Get Get
func (c *CockDbMock) Get(query string, args ...any) *DbRow {
	//return m.MockRow
	var rtn *DbRow
	if !c.mockRow1Used {
		c.mockRow1Used = true
		rtn = c.MockRow1
	} else if !c.mockRow2Used {
		c.mockRow2Used = true
		rtn = c.MockRow2
	} else if !c.mockRow3Used {
		c.mockRow3Used = true
		rtn = c.MockRow3
	} else if !c.mockRow4Used {
		c.mockRow4Used = true
		rtn = c.MockRow4
	} else if !c.mockRow5Used {
		c.mockRow5Used = true
		rtn = c.MockRow5
	} else if !c.mockRow6Used {
		c.mockRow6Used = true
		rtn = c.MockRow6
	} else if !c.mockRow7Used {
		c.mockRow7Used = true
		rtn = c.MockRow7
	} else if !c.mockRow8Used {
		c.mockRow8Used = true
		rtn = c.MockRow8
	}
	return rtn
}

// GetList GetList
func (c *CockDbMock) GetList(query string, args ...any) *DbRows {
	var rtn *DbRows
	if !c.mockRows1Used {
		c.mockRows1Used = true
		rtn = c.MockRows1
	} else if !c.mockRows2Used {
		c.mockRows2Used = true
		rtn = c.MockRows2
	} else if !c.mockRows3Used {
		c.mockRows3Used = true
		rtn = c.MockRows3
	} else if !c.mockRows4Used {
		c.mockRows4Used = true
		rtn = c.MockRows4
	} else if !c.mockRows5Used {
		c.mockRows5Used = true
		rtn = c.MockRows5
	} else if !c.mockRows6Used {
		c.mockRows6Used = true
		rtn = c.MockRows6
	} else if !c.mockRows7Used {
		c.mockRows7Used = true
		rtn = c.MockRows7
	} else if !c.mockRows8Used {
		c.mockRows8Used = true
		rtn = c.MockRows8
	}
	return rtn
}

// Delete Delete
func (c *CockDbMock) Delete(query string, args ...any) bool {
	var rtn = false
	if !c.mockDeleteSuccess1Used {
		c.mockDeleteSuccess1Used = true
		rtn = c.MockDeleteSuccess1
	} else if !c.mockDeleteSuccess2Used {
		c.mockDeleteSuccess2Used = true
		rtn = c.MockDeleteSuccess2
	} else if !c.mockDeleteSuccess3Used {
		c.mockDeleteSuccess3Used = true
		rtn = c.MockDeleteSuccess3
	} else if !c.mockDeleteSuccess4Used {
		c.mockDeleteSuccess4Used = true
		rtn = c.MockDeleteSuccess4
	}
	return rtn
}

func (c *CockDbMock) Exec(query string) (bool, error) {
	return c.MockExecSuccess, c.MockExecError
}

// Close Close
func (c *CockDbMock) Close() bool {
	return c.MockCloseSuccess
}
