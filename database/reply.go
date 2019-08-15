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

import uuid "github.com/satori/go.uuid"

type Reply struct {
	UUID    uuid.UUID
	Thread  Thread
	Name    string
	Comment string
}

type NewReply struct {
	ThreadUUID string
	Name       string
	Comment    string
}

func CreateReplyTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS reply (
			uuid uuid NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
			name text,
			comment text
		)
	`)

	return err
}

func Replies() ([]*Reply, error) {
	var replies []*Reply
	rows, err := db.Query(`SELECT uuid, thread_uuid, name, comment FROM reply`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reply Reply
		if err := rows.Scan(&reply.UUID, &reply.Thread.UUID, &reply.Name, &reply.Comment); err != nil {
			return nil, err
		}
		replies = append(replies, &reply)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return replies, nil
}

func CreateReply(newReply NewReply) (*Reply, error) {
	uuid, err := uuid.FromString(newReply.ThreadUUID)
	if err != nil {
		return nil, err
	}
	reply := Reply{
		Thread:  Thread{UUID: uuid},
		Name:    newReply.Name,
		Comment: newReply.Comment,
	}
	err = db.QueryRow(`INSERT INTO reply (thread_uuid, name, comment) VALUES ($1, $2, $3, $4) RETURNING uuid`, reply.Thread.UUID, reply.Name, reply.Comment).Scan(&reply.UUID)
	if err != nil {
		return new(Reply), err
	}

	return &reply, nil
}
