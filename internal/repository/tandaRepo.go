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
type TandaRepo interface {
	NewTanda(user domain.Tanda) (domain.Tanda, error)
	UpdateTanda(user domain.Tanda) (domain.Tanda, error)
	GetTandaById(Id int) (domain.Tanda, error)
	GetTandaByUserId(Id int) ([]domain.Tanda, error)
}

// --
// Definimos la estructura de conección, tenemos encuenta que el *sql.DB permite reutilizar el puntero de la conección a mysql
// --
func OpenTandaDb(dns string) (TandaRepo, error) {
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
func (r *userRepositoryDB) NewTanda(tanda domain.Tanda) (domain.Tanda, error) {
	//	--
	//  Preparamos la consulta con la variable stmt y err
	//	--
	stmt, err := r.db.Prepare(query.NEW_TANDA)
	if err != nil {
		return domain.Tanda{}, err
	}

	//	--
	//	 Retrasa la ejecución del codigo hasta que termine todo
	//	--
	defer stmt.Close()

	//	--
	//	Le pasamos los parametros que espera recibir la constante y validamos que no tenga error
	//	--
	res, err := stmt.Exec(tanda.Alias, tanda.PoolAmount, tanda.Period, tanda.Members, tanda.StartDate, tanda.EndDate, tanda.TotalEndPool, tanda.CreationUserId)
	if err != nil {
		return domain.Tanda{}, err
	}

	//	--
	//	Recuperamos el ultimo id insertado si ha ido bien todo y si ha ido mal devolvemos error.
	//	--
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return domain.Tanda{}, err
	}

	tanda.ID = int(lastInsertID)

	return tanda, nil

}

// --
// Creamos el repo para modificar los datos de una Tanda
// --
func (r *userRepositoryDB) UpdateTanda(tanda domain.Tanda) (domain.Tanda, error) {
	//	--
	//	Importante la sintaxis _, ignora la respuesta de la consulta, con eso decimos que solo queremos el error en caso de que
	//	susceda
	//	--
	_, err := r.db.Exec(query.UPDATE_TANDA, tanda.Alias, tanda.PoolAmount, tanda.Period, tanda.Members, tanda.StartDate, tanda.EndDate, tanda.TotalEndPool, tanda.CreationUserId, tanda.ID)
	if err != nil {
		return domain.Tanda{}, err
	}
	return tanda, nil
}

//	--
//	Creamos el repo para recuperar los datos de una tanda por Id
//
// --
func (r *userRepositoryDB) GetTandaById(Id int) (domain.Tanda, error) {
	//	--
	//	En caso de consultas se usar un queryRow
	//	--
	var tanda domain.Tanda
	err := r.db.QueryRow(query.SEARCH_TANDA_BY_ID, Id).Scan(&tanda.ID, &tanda.Alias, &tanda.PoolAmount, &tanda.Period, &tanda.Members, &tanda.StartDate, &tanda.EndDate, &tanda.TotalEndPool, &tanda.CreationUserId)
	if err != nil {
		return domain.Tanda{}, err
	}
	return tanda, nil
}

// --
// Creamos el repo para recuperar todas las tandas de un usuario por Id
// --
func (r *userRepositoryDB) GetTandaByUserId(Id int) ([]domain.Tanda, error) {
	//	--
	//	En caso de consultas se usar un queryRow
	//	--
	var tandas []domain.Tanda
	rows, err := r.db.Query(query.SEARCH_TANDA_BY_USER_ID, Id)
	if err != nil {
		return []domain.Tanda{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var tanda domain.Tanda
		var startDate, endDate []uint8
		err := rows.Scan(&tanda.ID, &tanda.Alias, &tanda.PoolAmount, &tanda.Period, &tanda.Members, &startDate, &endDate, &tanda.TotalEndPool, &tanda.CreationUserId)
		if err != nil {
			return []domain.Tanda{}, err
		}
		if len(startDate) > 0 {
			tanda.StartDate, err = time.Parse("2006-01-02 15:04:05", string(startDate))
			if err != nil {
				return []domain.Tanda{}, err
			}
		}
		if len(endDate) > 0 {
			tanda.EndDate, err = time.Parse("2006-01-02 15:04:05", string(endDate))
			if err != nil {
				return []domain.Tanda{}, err
			}
		}
		if err != nil {
			return []domain.Tanda{}, err
		}
		tandas = append(tandas, tanda)
	}

	return tandas, nil
}
