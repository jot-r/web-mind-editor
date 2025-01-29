# web-mind-editor
web-mind-editor is designed with two objectives
- provide a full featured alternative to Freemind but with several improvements (see Key features)
- learn/experiement/prove several academic concepts from computer science e.g. persistent data structures, parsers, concurrent changes, three way merge, ...

Since the Freemind file format *.mm is a quasi standard for mind mapping 
it heavily influences the data structures and features of web-mind-editor.

## Key features
- compatibilty: import and export of Freemind *.mm files

## planned features
- editor: support for adding, deleting and modifying nodes and their relations
- full history: save every modification in persistent data structures, append only
- efficency: store thousands of nodes but only the currently visible data
- versioning: direct access to any version in history
- concurrency: work with multiple users on the same document
- staging: changes need to be reviewed and accepted before a revision is created
- compare: show the differences between changes, commits, revisions or documents
