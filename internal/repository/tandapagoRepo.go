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
type TandaPagoRepo interface {
	NewTandaPago(user domain.TandaPago) (domain.TandaPago, error)
	UpdateTandaPago(user domain.TandaPago) (domain.TandaPago, error)
	GetTandaPagoByTandaUsuarioId(Id int) ([]domain.TandaPago, error)
}

// --
// Definimos la estructura de conección, tenemos encuenta que el *sql.DB permite reutilizar el puntero de la conección a mysql
// --
func OpenTandaPagoDb(dns string) (TandaPagoRepo, error) {
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
func (r *userRepositoryDB) NewTandaPago(tandapago domain.TandaPago) (domain.TandaPago, error) {
	//	--
	//  Preparamos la consulta con la variable stmt y err
	//	--
	stmt, err := r.db.Prepare(query.NEW_TANDAPAGO)
	if err != nil {
		return domain.TandaPago{}, err
	}

	//	--
	//	 Retrasa la ejecución del codigo hasta que termine todo
	//	--
	defer stmt.Close()

	//	--
	//	Le pasamos los parametros que espera recibir la constante y validamos que no tenga error
	//	--
	res, err := stmt.Exec(tandapago.TandaUsuarioID, tandapago.PeriodNumber, tandapago.PaymentDate, tandapago.Amount, tandapago.Status)
	if err != nil {
		return domain.TandaPago{}, err
	}

	//	--
	//	Recuperamos el ultimo id insertado si ha ido bien todo y si ha ido mal devolvemos error.
	//	--
	RowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.TandaPago{}, err
	}

	if RowsAffected == 0 {
		return domain.TandaPago{}, err
	}

	return tandapago, nil

}

// --
// Creamos el repo para modificar los datos de una Tanda
// --
func (r *userRepositoryDB) UpdateTandaPago(tandapago domain.TandaPago) (domain.TandaPago, error) {
	//	--
	//	Importante la sintaxis _, ignora la respuesta de la consulta, con eso decimos que solo queremos el error en caso de que
	//	susceda
	//	--
	_, err := r.db.Exec(query.UPDATE_TANDAPAGO, tandapago.PeriodNumber, tandapago.PaymentDate, tandapago.Amount, tandapago.Status, tandapago.TandaUsuarioID)
	if err != nil {
		return domain.TandaPago{}, err
	}
	return tandapago, nil
}

// --
// Creamos el repo para recuperar todas las tandas de un usuario por Id
// --
func (r *userRepositoryDB) GetTandaPagoByTandaUsuarioId(Id int) ([]domain.TandaPago, error) {
	//	--
	//	En caso de consultas se usar un queryRow
	//	--
	var tandaspago []domain.TandaPago
	rows, err := r.db.Query(query.SEARCH_TANDAPAGO_BY_TANDAUSUARIO_ID, Id)
	if err != nil {
		return []domain.TandaPago{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var tandapago domain.TandaPago
		var DatePay []uint8
		err := rows.Scan(&tandapago.TandaUsuarioID, &tandapago.PeriodNumber, &DatePay, &tandapago.Amount, &tandapago.Status)
		if err != nil {
			return []domain.TandaPago{}, err
		}
		if len(DatePay) > 0 {
			tandapago.PaymentDate, err = time.Parse("2006-01-02 15:04:05", string(DatePay))
			if err != nil {
				return []domain.TandaPago{}, err
			}
		}

		if err != nil {
			return []domain.TandaPago{}, err
		}
		tandaspago = append(tandaspago, tandapago)
	}

	return tandaspago, nil
}
