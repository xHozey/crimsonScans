package database

import "database/sql"

func InitDB(db *sql.DB) error {
	user := `CREATE TABLE IF NOT EXISTS user(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT,
		profilePic TEXT
	)`

	manhua := `CREATE TABLE IF NOT EXISTS manhua(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		description TEXT,
		status TEXT CHECK(status IN ('ongoing', 'completed', 'hiatus')),
		coverImage TEXT,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	chapter := `CREATE TABLE IF NOT EXISTS chapter(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		manhuaId INTEGER,
		chapterNb INTEGER,
		title TEXT,
		chapterFile TEXT,
		FOREIGN KEY (manhuaId) REFERENCES manhua(id) ON DELETE CASCADE,
		UNIQUE(manhuaId, chapterNb)
	)`

	chapterPages := `CREATE TABLE IF NOT EXISTS page(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chapterId INTEGER,
		pageNb INTEGER DEFAULT 1,
		imageURL TEXT,
		FOREIGN KEY (chapterId) REFERENCES chapter(id) ON DELETE CASCADE
	)`

	geners := `CREATE TABLE IF NOT EXISTS genres(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		gener TEXT UNIQUE
	)`

	manhuaGener := `CREATE TABLE IF NOT EXISTS manhuaGener(
		manhuaId INTEGER,
		generId INTEGER,
		UNIQUE (manhuaId, generId),
		FOREIGN KEY (manhuaId) REFERENCES manhua(id) ON DELETE CASCADE,
		FOREIGN KEY (generId) REFERENCES geners(id) ON DELETE CASCADE
	)`

	favorites := `CREATE TABLE IF NOT EXISTS favorite(
		userId INTEGER,
		manhuaId INTEGER,
		UNIQUE(userId, manhuaId),
		FOREIGN KEY (manhuaId) REFERENCES manhua(id) ON DELETE CASCADE,
		FOREIGN KEY (userId) REFERENCES user(id) ON DELETE CASCADE
	)`

	comments := `CREATE TABLE IF NOT EXISTS comment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER,
		manhuaId INTEGER,
		chapterId INTEGER,
		content TEXT,
		FOREIGN KEY (manhuaId) REFERENCES manhua(id) ON DELETE CASCADE,
		FOREIGN KEY (userId) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (chapterId) REFERENCES chapter(id) ON DELETE CASCADE
	)`

	roles := `CREATE TABLE IF NOT EXISTS roles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		role TEXT UNIQUE
	)`

	userRole := `CREATE TABLE IF NOT EXISTS userRole(
		userId INTEGER,
		roleId INTEGER,
		PRIMARY KEY(userId, roleId),
		FOREIGN KEY (userId) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (roleId) REFERENCES roles(id) ON DELETE CASCADE
	)`

	tables := []string{user, manhua, chapter, chapterPages, geners, manhuaGener, favorites, comments, roles, userRole}
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return err
		}
	}

	return nil
}
