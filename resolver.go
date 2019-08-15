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

package chan_api

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/mquarneti/chan-api/database"
)

type Resolver struct{}

func (r *Resolver) Board() BoardResolver {
	return &boardResolver{r}
}
func (r *Resolver) Thread() ThreadResolver {
	return &threadResolver{r}
}
func (r *Resolver) Reply() ReplyResolver {
	return &replyResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type boardResolver struct{ *Resolver }

func (r *boardResolver) UUID(ctx context.Context, obj *database.Board) (string, error) {
	return obj.UUID.String(), nil
}

type threadResolver struct{ *Resolver }

func (r *threadResolver) UUID(ctx context.Context, obj *database.Thread) (string, error) {
	return obj.UUID.String(), nil
}

type replyResolver struct{ *Resolver }

func (r *replyResolver) UUID(ctx context.Context, obj *database.Reply) (string, error) {
	return obj.UUID.String(), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Boards(ctx context.Context) ([]*database.Board, error) {
	return database.Boards()
}

func (r *queryResolver) Threads(ctx context.Context, input string) ([]*database.Thread, error) {
	return database.Threads()
}

func (r *queryResolver) Replies(ctx context.Context, input string) ([]*database.Reply, error) {
	return database.Replies()
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateBoard(ctx context.Context, input database.NewBoard) (*database.Board, error) {
	return database.CreateBoard(input)
}

func (r *mutationResolver) CreateThread(ctx context.Context, input database.NewThread) (*database.Thread, error) {
	return database.CreateThread(input)
}

func (r *mutationResolver) CreateReply(ctx context.Context, input database.NewReply) (*database.Reply, error) {
	return database.CreateReply(input)
}
