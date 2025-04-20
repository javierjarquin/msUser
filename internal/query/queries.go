package query

const (
	//	--
	//	usuarios
	//	--
	NEW_USER                            = "Insert into Usuario (name, lastName, userType, phone, address, email, pass) Values (?, ?, ?, ?, ?, ?, ?)"
	UPDATE_USER                         = "Update Usuario set name = ?, lastname = ?, userType = ?, phone = ?, address = ?, email = ?, pass = ? where id = ?"
	SEARCH_USER_BY_ID                   = "Select id, name, lastName, userType, phone, address, email, pass from Usuario where id = ?"
	SEARCH_USER_BY_EMAIL                = "Select id, pass from Usuario where email = ?"
	NEW_SESSION                         = "Insert into Session (userId, creationDate, ipAddres, comment) Values (?, ?, ?, ?)"
	SEARCH_USER_BY_SESSION              = "Select id, userId, creationDate, ipAddres, comment from Session where id = ?"
	NEW_TANDA                           = "Insert into Tanda (alias, poolAmount, period, members, starDate, endDate, totalEndPool, creationUserId) Values (?, ?, ?, ?, ?, ?, ?, ?)"
	UPDATE_TANDA                        = "Update Tanda set alias = ?, poolAmount = ?, period = ?, members = ?, starDate = ?, endDate = ?, totalEndPool = ? where id = ?"
	SEARCH_TANDA_BY_ID                  = "Select id, alias, poolAmount, period, members, starDate, endDate, totalEndPool, creationUserId from Tanda where id = ?"
	SEARCH_TANDA_BY_USER_ID             = "Select t.id, t.alias, t.poolAmount, CASE t.period WHEN 'SL' THEN 'Semanal' WHEN 'QL' THEN 'Quincenal' WHEN 'ML' THEN 'Mensual' WHEN 'BL' THEN 'Bimestral' WHEN 'TL' THEN 'Trimestral' ELSE 'Desconocido' END AS period, t.members, t.starDate, t.endDate, t.totalEndPool, concat(u.Name, ' ', u.LastName) usercreation from Tanda t, Usuario u where t.creationUserId = u.id and t.creationuserid = ?"
	NEW_TANDAUSUARIO                    = "Insert into TandaUsuario (tandaId, memberId, numberTicket, datePay, status) Values (?, ?, ?, ?, ?)"
	UPDATE_TANDAUSUARIO                 = "Update TandaUsuario set numberTicket = ?, datePay = ?, status = ? where id = ?"
	SEARCH_TANDAUSUARIO_BY_TANDA_ID     = "Select id, tandaId, memberId, numberTicket, datePay, status from TandaUsuario where tandaId = ?"
	NEW_TANDAPAGO                       = "Insert into TandaPago (tandaUsuarioId, periodNumber, paymentDate, amount, status) Values (?, ?, ?, ?, ?)"
	UPDATE_TANDAPAGO                    = "Update TandaPago set periodNumber = ?, paymentDate = ?, amount = ?, status = ? where tandaUsuarioId = ?"
	SEARCH_TANDAPAGO_BY_TANDAUSUARIO_ID = "Select tandaUsuarioId, periodNumber, paymentDate, amount, status from TandaPago where tandaUsuarioId = ?"
)
