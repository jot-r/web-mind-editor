-- name: CreateNodeContent :exec
INSERT INTO NODE_CONTENT(
	NODE_CONTENT_ID, 
	BACKGROUND_COLOR, 
	CREATED_TIMESTAMP, 
	MODIFIED_TIMESTAMP, 
	FREEMIND_ID, 
	TEXT, 
	NOTE, 
	URL, 
	STYLE, 
	ICONS,
	FONT_NAME, 
	FONT_BOLD, 
	FONT_SIZE) 
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateNode :exec
INSERT INTO NODE(
	NODE_ID,
	NODE_CONTENT_ID)
VALUES(?, ?);

-- name: CreateNodeRelation :exec
INSERT INTO NODE_RELATION(
	PARENT_NODE_ID,
	CHILD_NODE_ID,
	CHILD_NODE_KEY)
VALUES(?, ?, ?);