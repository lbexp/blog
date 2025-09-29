package core_db

import "os"

type database struct {
	file *os.File
}

func NewConnection(path string) (*database, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return &database{file}, nil
}

func (db *database) Close() error {
	err := db.file.Close()

	return err
}

func (db *database) Read() {
	// TODO: Create sort of query language for read
	// Utilize db.file.Read and only read specific line
	// Utilize json.Unmarshal
}

func (db *database) Write() {
	// TODO: Create sort of query language for write
	// Utilize db.file.Read and only read specific line
	// Utilize json.Unmarshal
}
