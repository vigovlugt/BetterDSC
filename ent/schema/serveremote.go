package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ServerEmote holds the schema definition for the ServerEmote entity.
type ServerEmote struct {
	ent.Schema
}

// Fields of the ChannelEmote.
func (ServerEmote) Fields() []ent.Field {
	return []ent.Field{
		field.String("server_id"),
		field.String("emote_id"),
		field.String("code"),
		field.String("image_type"),
	}
}

// Edges of the ChannelEmote.
func (ServerEmote) Edges() []ent.Edge {
	return nil
}
