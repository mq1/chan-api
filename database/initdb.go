/* Copyright (C) 2019  Manuel Quarneti
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		log.Fatal(err)
	}

	if err := CreateBoardTable(); err != nil {
		log.Fatal(err)
	}

	if err := CreateThreadTable(); err != nil {
		log.Fatal(err)
	}

	if err := CreateReplyTable(); err != nil {
		log.Fatal(err)
	}

	return err
}
