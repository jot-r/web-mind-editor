-- +goose Up
CREATE TABLE NODE(
	NODE_ID TEXT PRIMARY KEY,
	NODE_CONTENT_ID TEXT NOT NULL REFERENCES NODE_CONTENT);
				
-- +goose Down
DROP TABLE NODE;