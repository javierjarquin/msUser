package repository

import (
	"database/sql"
	"log"
	"time"

	//	--
	//	Agregamos las query's
	//
	"msUser/internal/query"
	//	--
	// 	Ruta de los archivos dominio, creo que podemos especificar los necesarios
	//	--
	"msUser/internal/domain"

	//	--
	//	los guiones bajo " _ " se usan para importar una librería al proyecto
	//	--
	_ "github.com/go-sql-driver/mysql"
)

// --
// Definimos el interfac de funciones del TandaRepo
// --
type TandaUsuarioRepo interface {
	NewTandaUsuario(user domain.TandaUsuario) (domain.TandaUsuario, error)
	UpdateTandaUsuario(user domain.TandaUsuario) (domain.TandaUsuario, error)
	GetTandaUsuarioByTandaId(Id int) ([]domain.TandaUsuario, error)
}

// --
// Definimos la estructura de conección, tenemos encuenta que el *sql.DB permite reutilizar el puntero de la conección a mysql
// --
func OpenTandaUsuarioDb(dns string) (TandaUsuarioRepo, error) {
	//	--
	//	Intenamos abrir la conección... pero si hay error la guardamos en err
	//	--
	db, err := sql.Open("mysql", dns)

	if err != nil {
		log.Fatal("Error al conectarse a la db:", err)
		return nil, err
	}

	//	--
	//	Retornamos el puntero del dominio userRepositoryDB e indicamos que en el campo erro no hubo y devolvemos nil
	//	--
	return &userRepositoryDB{db: db}, nil
}

// --
// Creamos el repo para crear una nueva tanda
// --
func (r *userRepositoryDB) NewTandaUsuario(tandausario domain.TandaUsuario) (domain.TandaUsuario, error) {
	//	--
	//  Preparamos la consulta con la variable stmt y err
	//	--
	stmt, err := r.db.Prepare(query.NEW_TANDAUSUARIO)
	if err != nil {
		return domain.TandaUsuario{}, err
	}

	//	--
	//	 Retrasa la ejecución del codigo hasta que termine todo
	//	--
	defer stmt.Close()

	//	--
	//	Le pasamos los parametros que espera recibir la constante y validamos que no tenga error
	//	--
	res, err := stmt.Exec(tandausario.TandaID, tandausario.MemberID, tandausario.NumberTicket, tandausario.DatePay, tandausario.Status)
	if err != nil {
		return domain.TandaUsuario{}, err
	}

	//	--
	//	Recuperamos el ultimo id insertado si ha ido bien todo y si ha ido mal devolvemos error.
	//	--
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return domain.TandaUsuario{}, err
	}

	tandausario.ID = int(lastInsertID)

	return tandausario, nil

}

// --
// Creamos el repo para modificar los datos de una Tanda
// --
func (r *userRepositoryDB) UpdateTandaUsuario(tandausuario domain.TandaUsuario) (domain.TandaUsuario, error) {
	//	--
	//	Importante la sintaxis _, ignora la respuesta de la consulta, con eso decimos que solo queremos el error en caso de que
	//	susceda
	//	--
	_, err := r.db.Exec(query.UPDATE_TANDAUSUARIO, tandausuario.NumberTicket, tandausuario.DatePay, tandausuario.Status, tandausuario.ID)
	if err != nil {
		return domain.TandaUsuario{}, err
	}
	return tandausuario, nil
}

// --
// Creamos el repo para recuperar todas las tandas de un usuario por Id
// --
func (r *userRepositoryDB) GetTandaUsuarioByTandaId(Id int) ([]domain.TandaUsuario, error) {
	//	--
	//	En caso de consultas se usar un queryRow
	//	--
	var tandasusuario []domain.TandaUsuario
	rows, err := r.db.Query(query.SEARCH_TANDAUSUARIO_BY_TANDA_ID, Id)
	if err != nil {
		return []domain.TandaUsuario{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var tandausuario domain.TandaUsuario
		var DatePay []uint8
		err := rows.Scan(&tandausuario.ID, &tandausuario.TandaID, &tandausuario.MemberID, &tandausuario.NumberTicket, &DatePay, &tandausuario.Status)
		if err != nil {
			return []domain.TandaUsuario{}, err
		}
		if len(DatePay) > 0 {
			tandausuario.DatePay, err = time.Parse("2006-01-02 15:04:05", string(DatePay))
			if err != nil {
				return []domain.TandaUsuario{}, err
			}
		}

		if err != nil {
			return []domain.TandaUsuario{}, err
		}
		tandasusuario = append(tandasusuario, tandausuario)
	}

	return tandasusuario, nil
}
