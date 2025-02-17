package query

const (
	//	--
	//	usuarios
	//	--
	NEW_USER             = "Insert into Usuario (name, lastName, userType, phone, address, email, pass) Values (?, ?, ?, ?, ?, ?, ?)"
	UPDATE_USER          = "Update Usuario set name = ?, lastname = ?, userType = ?, phone = ?, address = ?, email = ?, pass = ? where id = ?"
	SEARCH_USER_BY_ID    = "Select id, name, lastName, userType, phone, address, email, pass from Usuario where id = ?"
	SEARCH_USER_BY_EMAIL = "Select id, pass from Usuario where email = ?"
	NEW_SESSION          = "Insert into Session (userId, creationDate, ipAddres, comment) Values (?, ?, ?, ?)"
)
