package freemind

import (
	"io"
	"iter"

	"strconv"

	"web-mind-editor/common"

	"github.com/google/uuid"
	"github.com/orisano/gosax"
)

type FreemindParser struct {
	GosaxReader *gosax.Reader
}

type ParsedFreemindNode struct {
	NodeId            common.NodeId
	ParentNodeId      *common.NodeId
	NodeText          *string
	NodeNote          *string
	BackgroundColor   *string
	CreatedTimestamp  *int64
	ModifiedTimestamp *int64
	FreemindId        *string
	Link              *string
	Icons             []string
	FontName          *string
	FontSize          *int64
	FontBold          *int64
	//Style
	//HGAP
	//VSHIFT
	//POSITION=left/right
}

type AttributeInRegistry struct {
	Name *string
}

func ParseFreemindFile(reader io.Reader) FreemindParser {
	gosaxReader := gosax.NewReader(reader)
	gosaxReader.EmitSelfClosingTag = true
	return FreemindParser{GosaxReader: gosaxReader}
}

func (freemindParser *FreemindParser) AtrributesInRegistry() iter.Seq[AttributeInRegistry] {
	return func(yield func(AttributeInRegistry) bool) {

	}
}

func (p *FreemindParser) Nodes() iter.Seq[ParsedFreemindNode] {
	return func(yield func(ParsedFreemindNode) bool) {
		nodeIsWritten := true

		var newNodeId = common.NodeId(uuid.New())
		parsedNode := ParsedFreemindNode{NodeId: newNodeId}

		writePreviousNode := func() {
			if !nodeIsWritten {
				yield(parsedNode)
				var newNodeId = common.NodeId(uuid.New())
				parsedNode = ParsedFreemindNode{NodeId: newNodeId}
				nodeIsWritten = true
			}
		}

		hierarchyTracker := common.NewStack[common.NodeId]()

		for {
			event, err := p.GosaxReader.Event()
			common.CheckErr(err)
			if event.Type() == gosax.EventEOF {
				break
			}

			switch event.Type() {
			case gosax.EventStart:
				elementName, nameRemainingBytes := gosax.Name(event.Bytes)
				elementNameString := string(elementName)

				switch elementNameString {
				case "node":
					writePreviousNode()

					if !hierarchyTracker.IsEmpty() {
						var parentNode = hierarchyTracker.Peek()
						parsedNode.ParentNodeId = &parentNode
					}

					hierarchyTracker.Push(parsedNode.NodeId)

					nodeIsWritten = false
					attrBytes := nameRemainingBytes
					for len(attrBytes) > 0 {
						attr, attrRemainingBytes, err := gosax.NextAttribute(attrBytes)
						common.CheckErr(err)

						rawValue := string(attr.Value)
						attrValue := rawValue[1 : len(rawValue)-1]

						switch string(attr.Key) {
						case "BACKGROUND_COLOR":
							parsedNode.BackgroundColor = &attrValue
						case "ID":
							parsedNode.FreemindId = &attrValue
						case "TEXT":
							parsedNode.NodeText = &attrValue
						case "CREATED":
							temp, err := strconv.ParseInt(attrValue, 10, 64)
							common.CheckErr(err)
							if err == nil {
								parsedNode.CreatedTimestamp = &temp
							}
						case "LINK":
							parsedNode.Link = &attrValue
						case "MODIFIED":
							temp, err := strconv.ParseInt(attrValue, 10, 64)
							common.CheckErr(err)
							if err == nil {
								parsedNode.ModifiedTimestamp = &temp
							}
						}

						attrBytes = attrRemainingBytes
					}
				case "richcontent":
					attrBytes := nameRemainingBytes
					for len(attrBytes) > 0 {
						attr, attrRemainingBytes, err := gosax.NextAttribute(attrBytes)
						common.CheckErr(err)

						if string(attr.Key) == "TYPE" && string(attr.Value) == "\"NODE\"" {
							parsedNode.NodeText = getInnerXml(p.GosaxReader, "richcontent")
						} else if string(attr.Key) == "TYPE" && string(attr.Value) == "\"NOTE\"" {
							parsedNode.NodeNote = getInnerXml(p.GosaxReader, "richcontent")
						}

						attrBytes = attrRemainingBytes
					}
				case "font":
					attrBytes := nameRemainingBytes
					for len(attrBytes) > 0 {
						attr, attrRemainingBytes, err := gosax.NextAttribute(attrBytes)
						common.CheckErr(err)

						rawValue := string(attr.Value)
						attrValue := rawValue[1 : len(rawValue)-1]

						if string(attr.Key) == "BOLD" && string(attr.Value) == "\"true\"" {
							var bold int64 = 1
							parsedNode.FontBold = &bold
						} else if string(attr.Key) == "NAME" {
							if len(attrValue) > 0 {
								parsedNode.FontName = &attrValue
							}
						} else if string(attr.Key) == "SIZE" {
							temp, err := strconv.ParseInt(attrValue, 10, 64)
							common.CheckErr(err)
							if err == nil {
								parsedNode.FontSize = &temp
							}
						}

						attrBytes = attrRemainingBytes
					}
				case "icon":
					attrBytes := nameRemainingBytes
					for len(attrBytes) > 0 {
						attr, attrRemainingBytes, err := gosax.NextAttribute(attrBytes)
						common.CheckErr(err)

						rawValue := string(attr.Value)
						attrValue := rawValue[1 : len(rawValue)-1]

						if string(attr.Key) == "BUILTIN" {
							parsedNode.Icons = append(parsedNode.Icons, attrValue)
						}

						attrBytes = attrRemainingBytes
					}
				}
			case gosax.EventEnd:
				elementName, _ := gosax.Name(event.Bytes)
				if string(elementName) == "node" {
					hierarchyTracker.Pop()
				}
			}
		}
		writePreviousNode()
	}
}

func getInnerXml(reader *gosax.Reader, endTag string) *string {
	var innerXml = ""
	inBody := false
	for {
		event, err := reader.Event()
		if err != nil {
			break
		}

		switch event.Type() {
		case gosax.EventEOF:
			break
		case gosax.EventStart:
			elementName, _ := gosax.Name(event.Bytes)
			elementNameString := string(elementName)
			if elementNameString == "body" {
				inBody = true
			} else {
				if inBody {
					innerXml += string(event.Bytes)
				}
			}
		case gosax.EventEnd:
			elementName, _ := gosax.Name(event.Bytes)
			elementNameString := string(elementName)
			if elementNameString == endTag {
				return &innerXml
			} else if elementNameString == "body" {
				inBody = false
			}

			if inBody {
				innerXml += string(event.Bytes)
			}
		default:
			if inBody {
				innerXml += string(event.Bytes)
			}
		}
	}
	return nil
}
