package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// defining structure for database
type Repository struct {
	DB *sql.DB
}

var repo Repository

// initialising database
func Init() {
	if db, err := sql.Open("sqlite3", "/tmp/olicooltown.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		log.Fatal("Database initialisation")
	}
}

func Create() int {
	const sql = "CREATE TABLE IF NOT EXISTS Tracks" +
		"(Id TEXT PRIMARY KEY, Audio TEXT)"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Clear() int {
	const sql = "DELETE FROM Tracks"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Update(track Track) int64 {
	const sql = "UPDATE Tracks SET Audio = ? WHERE id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(track.Audio, track.Id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Insert(track Track) int64 {
	const sql = "INSERT INTO Tracks(Id, Audio) VALUES (?, ?)"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(track.Id, track.Audio); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Read(id string) (Track, int64) {
	const sql = "SELECT * FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		var track Track
		row := stmt.QueryRow(id)
		if err := row.Scan(&track.Id, &track.Audio); err == nil {
			return track, 1
		} else {
			return Track{}, 0
		}
	}
	return Track{}, -1
}

func List() ([]string, bool) {
	const sql = "SELECT Id FROM Tracks Order By Id"
	var list []string
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if row, errQ := stmt.Query(); errQ == nil {
			defer row.Close()
			for row.Next() {
				var id string
				if errS := row.Scan(&id); errS == nil {
					list = append(list, id)
				} else {
					log.Println("List: Error in scanning.")
					return list, false
				}
			}
		} else {
			log.Println("List: Error in query.")	
			return list, false	
		}
	} else {
		log.Println("List: Error in prepare.")
		return list, false
	}
	return list, true
}

func Delete(id string) int64 {
	const sql = "DELETE FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if result, errE := stmt.Exec(id); errE == nil {
			if n, errR := result.RowsAffected(); errR == nil {
				return n
			} else {
				log.Println("Delete: Error in rows affected ", errR)
				return -1
			}
		} else {
			log.Println("Delete: Error in exec ", errE)
			return -1
		}
	} else {
		log.Println("Delete: Error in preparation ", err)
		return -1
	}
}
