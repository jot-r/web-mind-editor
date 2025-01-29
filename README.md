# web-mind-editor
web-mind-editor is designed with two objectives
- provide a full featured alternative to Freemind but with several improvements (see key features)
- learn/experiement/prove several technical concepts from computer science (see technical conecpts)

Since the Freemind file format *.mm is a quasi standard for mind mapping 
it heavily influences the data structures and features of web-mind-editor.

## Key features
- compatibilty: import and export of Freemind *.mm files

## Planned features
- editor: support for adding, deleting and modifying nodes and their relations
- full history: save every modification in persistent data structures, append only
- efficency: store thousands of nodes but load only the currently visible data
- versioning: direct access to any version in history
- collaboration: work with multiple users on the same document
- concurrency: resolve conflicts which occur by concurrent changes
- staging: changes need to be reviewed and accepted before a revision is created
- compare: show the differences between changes, commits, revisions or documents

## Technical concepts
- architecture: Go backend, Javascript frontend, SQLite data storage
- data management: git-like versioning, persistent data structures, path copying
- schema and query management: goose migrations + sqlc for queries
- performance for Freemind import: streaming parsers, journal_mode=WAL
- collaboration: path copying, three way merge
- self contained application: Go compiler links statically by default, Go embed for goose migrations and assets
- multi plattform: cross compiling backend Go code and deploy browser forntend
