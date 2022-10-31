package ports

type DBPort interface {
	CloseDBConnection()
	Save(result int32, operation string) error
}
