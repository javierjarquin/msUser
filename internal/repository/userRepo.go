package repository

import (
	"database/sql"
	"log"

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
// Definimos la estructura de conección, tenemos encuenta que el *sql.DB permite reutilizar el puntero de la conección a mysql
// --
type userRepositoryDB struct {
	db *sql.DB
}

// Definimos el interfac de funciones del userRepo
type UserRepo interface {
	NewUser(user domain.User) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	GetUserById(Id int) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	NewSession(session domain.Session) (domain.Session, error)
}

// --
// Creamos la instacia mysql en donde recuperamos el dns y devolvermos un userRepo
// NOTA: al crear una función con este sintaxis func ejem(srt string)(ejerep, err) indicamos que esa funcion va a recibir un str
//
//	y que devolvera un objeto ejerep o error
//
// --
func OpenUserDb(dns string) (UserRepo, error) {
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
// Creamos la funcionalidad para crear un nuevo usuario
// iniciamos indicando que r es un receptor de la estrucura userRepositoryDB y le inidcamos que va a ser un puntero
// que se puede modificar * para evitar que se hagan copias innecesarias, definimos el nombre de la función y
// el parametro que espera recibir newUser(user domain.user) y retorna dos valores el objeto domain.User y el error
// (domain.user, error)
// --
func (r *userRepositoryDB) NewUser(user domain.User) (domain.User, error) {
	//	--
	//  Preparamos la consulta con la variable stmt y err
	//	--
	stmt, err := r.db.Prepare(query.NEW_USER)
	if err != nil {
		return domain.User{}, err
	}

	//	--
	//	 Retrasa la ejecución del codigo hasta que termine todo
	//	--
	defer stmt.Close()

	//	--
	//	Le pasamos los parametros que espera recibir la constante y validamos que no tenga error
	//	--
	res, err := stmt.Exec(user.Name, user.LastName, user.UserType, user.Phone, user.Address, user.Email, user.Pass)
	if err != nil {
		return domain.User{}, err
	}

	//	--
	//	Recuperamos el ultimo id insertado si ha ido bien todo y si ha ido mal devolvemos error.
	//	--
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	user.ID = int(lastInsertID)

	return user, nil

}

// --
// Creamos el repo para modificar los datos de un usuario
// --
func (r *userRepositoryDB) UpdateUser(user domain.User) (domain.User, error) {
	//	--
	//	Importante la sintaxis _, ignora la respuesta de la consulta, con eso decimos que solo queremos el error en caso de que
	//	susceda
	//	--
	_, err := r.db.Exec(query.UPDATE_USER, user.Name, user.LastName, user.UserType, user.Phone, user.Address, user.Email, user.Pass, user.ID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

//	--
//	Creamos el repo para recuperar los datos de un usuario por Id
//
// --
func (r *userRepositoryDB) GetUserById(Id int) (domain.User, error) {
	//	--
	//	En caso de consultas se usar un queryRow
	//	--
	var user domain.User
	err := r.db.QueryRow(query.SEARCH_USER_BY_ID, Id).Scan(&user.ID, &user.Name, &user.LastName, &user.UserType, &user.Phone, &user.Address, &user.Email, &user.Pass)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepositoryDB) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(query.SEARCH_USER_BY_EMAIL, email).Scan(&user.ID, &user.Pass)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepositoryDB) NewSession(session domain.Session) (domain.Session, error) {
	stmt, err := r.db.Prepare(query.NEW_SESSION)
	if err != nil {
		return domain.Session{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(session.UserID, session.CreationDate, session.IPAddres, session.Comments)
	if err != nil {
		return domain.Session{}, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return domain.Session{}, err
	}

	session.ID = int(lastInsertID)

	return session, nil
}

func (r *userRepositoryDB) GetUserBySessionId(Id int) (domain.Session, error) {
	var session domain.Session
	err := r.db.QueryRow(query.SEARCH_USER_BY_SESSION, Id).Scan(&session.ID, &session.UserID, &session.CreationDate, &session.IPAddres, &session.Comments)
	if err != nil {
		return domain.Session{}, err
	}
	return session, nil
}
