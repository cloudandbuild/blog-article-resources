package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudandbuild/blog-article-resrouces/content/blog/article-3/ent"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("gremlin", "http://localhost:8182")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	bob, err := CreateUser(ctx, client, "bob")
	if err != nil {
		log.Fatalf("failed CreateUser err: %v", err)
	}

	_, err = CreateUser(ctx, client, "alice", bob)
	if err != nil {
		log.Fatalf("failed CreateUser err: %v", err)
	}

	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed User.Query().All err: %v", err)
	}
	log.Println(users)
}

func CreateUser(ctx context.Context, client *ent.Client, name string, friends ...*ent.User) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	var ids []string
	for _, friend := range friends {
		ids = append(ids, friend.ID)
	}
	if len(ids) > 0 {
		if _, err := u.Update().AddFriendIDs(ids...).Save(ctx); err != nil {
			return nil, fmt.Errorf("failed updating user: %w", err)
		}
	}
	return u, nil
}
