package db

type (

	// Connection representa os dados de uma conexao com banco de dados
	Connection struct{}

	// Database representa os metodos basicos para um banco de dados
	Database interface {

		// ValidateDsn realiza a validacao da string de conexao para o banco em questao
		ValidateDsn(dsn string) error

		// Connect realiza uma conexao com o banco de dados
		Connect(dsn string) (*Connection, error)

		// Select busca registros no banco de dados
		Select(table string, where string) ([]string, error)

		// Insert insere um novo registro no banco de dados
		Insert(table string, values []string) (int, error)

		// Update atualiza um registro existente no banco de dados
		Update(table string, values []string, where string) (int, error)

		// Delete exclui um registro existente no banco de dados
		Delete(table string, where string) (int, error)

		// Disconnect finaliza uma conexao com o banco de dados
		Disconnect() error
	}
)
