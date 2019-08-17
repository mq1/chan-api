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

type Thread struct {
	UUID     uuid.UUID
	Board    Board
	Name     string
	Subject  string
	Comment  string
	FileHash string
}

type NewThread struct {
	BoardUUID string
	Name      string
	Subject   string
	Comment   string
	FileHash  string
}

func CreateThreadTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS thread (
			uuid uuid NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
			board_uuid uuid NOT NULL references board(uuid),
			name text,
			subject text,
			comment text,
			file_hash text
		)
	`)

	return err
}

func Threads() ([]*Thread, error) {
	var threads []*Thread
	rows, err := db.Query(`SELECT uuid, board_uuid, name, subject, comment, file_hash FROM thread`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.UUID, &thread.Board.UUID, &thread.Name, &thread.Subject, &thread.Comment, &thread.FileHash); err != nil {
			return nil, err
		}
		threads = append(threads, &thread)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return threads, nil
}

func CreateThread(newThread NewThread) (*Thread, error) {
	uuid, err := uuid.FromString(newThread.BoardUUID)
	if err != nil {
		return nil, err
	}

	thread := Thread{
		Board:    Board{UUID: uuid},
		Name:     newThread.Name,
		Subject:  newThread.Subject,
		Comment:  newThread.Comment,
		FileHash: newThread.FileHash,
	}
	err = db.QueryRow(`INSERT INTO thread (board_uuid, name, subject, comment, file_hash) VALUES ($1, $2, $3, $4, $5) RETURNING uuid`, thread.Board.UUID, thread.Name, thread.Subject, thread.Comment, thread.FileHash).Scan(&thread.UUID)
	if err != nil {
		return new(Thread), err
	}

	return &thread, nil
}
