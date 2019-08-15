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
	uuid "github.com/satori/go.uuid"
)

type Board struct {
	UUID uuid.UUID
	Name string
}

type NewBoard struct {
	Name string
}

func CreateBoardTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS board (
			uuid uuid NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
			name text NOT NULL UNIQUE
		)
	`)

	return err
}

func Boards() ([]*Board, error) {
	var boards []*Board
	rows, err := db.Query(`SELECT uuid, name FROM board`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var board Board
		if err := rows.Scan(&board.UUID, &board.Name); err != nil {
			return nil, err
		}
		boards = append(boards, &board)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return boards, nil
}

func CreateBoard(newBoard NewBoard) (*Board, error) {
	board := Board{Name: newBoard.Name}
	err := db.QueryRow(`INSERT INTO board (name) VALUES ($1) RETURNING uuid`, newBoard.Name).Scan(&board.UUID)
	if err != nil {
		return new(Board), err
	}

	return &board, nil
}
