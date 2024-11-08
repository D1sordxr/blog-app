package sqlite

//
//func NewStorage() *gorm.DB {
//	var cfg config.Config
//	gormConfig := gorm.Config{}
//
//	db, err := gorm.Open(sqlite.Open(cfg.StoragePath), &gormConfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	migrate(db)
//
//	return db
//}
//
//func NewSQLDB(storagePath string) (*Storage, error) {
//	const op = "storage.sqlite.New"
//
//	db, err := sql.Open("sqlite3", "./blog_db")
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	stmt, err := db.Prepare(`
//	CREATE TABLE IF NOT EXIST posts(
//		id INTEGER PRIMARY KEY,
//		title TEXT NOT NULL,
//		content TEXT NOT NULL);
//	CREATE INDEX IF	NOT EXISTS idx_title ON url(title);
//	)`)
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	_, err = stmt.Exec()
//	if err != nil {
//		return nil, fmt.Errorf("%s: %w", op, err)
//	}
//
//	//return &Storage{DB: db}, nil
//	return nil, nil
//}
