package freemind

import (
	_ "database/sql"
	"io"
	_ "strconv"
	_ "strings"

	"github.com/google/uuid"

	"web-mind-editor/common"
)

var closeNodeUUID common.NodeId = common.NodeId(uuid.New())

func ExportFileFromDatabase(writer io.Writer, filename string) {
	/*datalayer.Connect()
	defer datalayer.Disconnect()

	xmlWriter := NewXmlStreamWriter(writer)

	var formatVersion sql.NullString
	var rootNodeId sql.NullString
	result, err := db.Query(`SELECT FORMAT_VERSION,
							 ROOT_NODE_ID
					  FROM FILE JOIN REVISION ON REVISION.REVISION_ID = FILE.LATEST_REVISION_ID
					  WHERE FILE.FILE_NAME = ?`, filename)
	checkErr(err)
	result.Next()
	result.Scan(&formatVersion, &rootNodeId)

	if !formatVersion.Valid || !rootNodeId.Valid {
		return
	}

	parsedRootNodeId := NodeId(uuid.MustParse(rootNodeId.String))

	xmlWriter.WriteStartElement("map")
	xmlWriter.WriteAttribute("version", formatVersion.String)
	xmlWriter.WriteComment("To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net")
	xmlWriter.WriteStartElement("attribute_registry")
	xmlWriter.WriteEndElement("attribute_registry")

	hierarchyTracker := NewStack[NodeId]()
	hierarchyTracker.Push(closeNodeUUID) // for closing element
	hierarchyTracker.Push(parsedRootNodeId)

	for !hierarchyTracker.IsEmpty() {
		currentId := hierarchyTracker.Pop()

		if currentId == closeNodeUUID {
			xmlWriter.WriteEndElement("node")
			continue
		}

		nodeRows, err := db.Query(`SELECT BACKGROUND_COLOR,
								CREATED_TIMESTAMP,
								MODIFIED_TIMESTAMP,
								TEXT,
								NOTE,
								FREEMIND_ID,
								URL
						FROM NODE
						JOIN NODE_CONTENT ON NODE.NODE_CONTENT_ID = NODE_CONTENT.NODE_CONTENT_ID
						WHERE NODE.NODE_ID = ?`, currentId)
		checkErr(err)

		childRows, err := db.Query(`SELECT CHILD_NODE_ID
						FROM NODE_RELATION
						WHERE PARENT_NODE_ID = ?
						ORDER BY rowid DESC`, currentId)
		checkErr(err)

		defer nodeRows.Close()
		defer childRows.Close()

		for nodeRows.Next() {
			var backgroundColor sql.NullString
			var created sql.NullInt64
			var modified sql.NullInt64
			var text sql.NullString
			var note sql.NullString
			var freemindId sql.NullString
			var url sql.NullString
			err := nodeRows.Scan(&backgroundColor, &created, &modified, &text, &note, &freemindId, &url)
			checkErr(err)

			xmlWriter.WriteStartElement("node")
			if backgroundColor.Valid {
				xmlWriter.WriteAttribute("BACKGROUND_COLOR", backgroundColor.String)
			}
			if created.Valid {
				xmlWriter.WriteAttribute("CREATED", strconv.FormatInt(created.Int64, 10))
			}
			xmlWriter.WriteAttribute("FOLDED", "true")

			// HGAP="number"

			if freemindId.Valid {
				xmlWriter.WriteAttribute("ID", freemindId.String)
			}

			if url.Valid {
				xmlWriter.WriteAttribute("LINK", url.String)
			}

			if modified.Valid {
				xmlWriter.WriteAttribute("MODIFIED", strconv.FormatInt(modified.Int64, 10))
			}

			xmlWriter.WriteAttribute("POSITION", "right")

			// STYLE = bubble (abgerundet), fork(nur unterschtrichen)
			// VSHIFT = "number"

			if text.Valid {
				if strings.Contains(text.String, "<") {
					xmlWriter.WriteStartElement("richcontent")
					xmlWriter.WriteAttribute("TYPE", "NODE")
					xmlWriter.WriteInnerXml(text.String)
					xmlWriter.WriteEndElement("richcontent")
				} else {
					xmlWriter.WriteAttribute("TEXT", text.String)
				}
			}

			if note.Valid {
				xmlWriter.WriteStartElement("richcontent")
				xmlWriter.WriteAttribute("TYPE", "NOTE")
				xmlWriter.WriteInnerXml(note.String)
				xmlWriter.WriteEndElement("richcontent")
			}

			for childRows.Next() {
				var childNodeId sql.NullString
				childRows.Scan(&childNodeId)
				if childNodeId.Valid {
					parsedChildNodeId := NodeId(uuid.MustParse(childNodeId.String))
					checkErr(err)
					hierarchyTracker.Push(closeNodeUUID)
					hierarchyTracker.Push(parsedChildNodeId)
				}
			}
		}
	}
	xmlWriter.WriteEndElement("map")*/
}
