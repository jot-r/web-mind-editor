package common

import (
	"bufio"
	"io"
)

type XmlStreamWriter struct {
	writer         bufio.Writer
	inStartElement bool
	firstElement   bool
}

func NewXmlStreamWriter(writer io.Writer) XmlStreamWriter {
	bufferedWriter := bufio.NewWriterSize(writer, 1024*20) // 20kByte
	return XmlStreamWriter{writer: *bufferedWriter, firstElement: true}
}

func (xmlStreamWriter *XmlStreamWriter) WriteStartElement(name string) {
	xmlStreamWriter.writeCloseStartElement()
	xmlStreamWriter.writeNewLine()
	xmlStreamWriter.writer.WriteString("<" + name)
	xmlStreamWriter.inStartElement = true
}

func (xmlStreamWriter *XmlStreamWriter) writeCloseStartElement() {
	if xmlStreamWriter.inStartElement {
		xmlStreamWriter.writer.WriteString(">")
		xmlStreamWriter.inStartElement = false
	}
}

func (xmlStreamWriter *XmlStreamWriter) writeNewLine() {
	if xmlStreamWriter.firstElement {
		xmlStreamWriter.firstElement = false
	} else {
		xmlStreamWriter.writer.WriteString("\n")
	}
}

func (xmlStreamWriter *XmlStreamWriter) WriteEndElement(name string) {
	if xmlStreamWriter.inStartElement {
		xmlStreamWriter.writer.WriteString("/>")
		xmlStreamWriter.inStartElement = false
	} else {
		xmlStreamWriter.writeNewLine()
		xmlStreamWriter.writer.WriteString("</" + name + ">")
	}
	xmlStreamWriter.writer.Flush()
}

func (xmlStreamWriter *XmlStreamWriter) WriteAttribute(name string, value string) {
	xmlStreamWriter.writer.WriteString(" " + name + `="` + value + `"`)
}

func (xmlStreamWriter *XmlStreamWriter) WriteComment(comment string) {
	xmlStreamWriter.writeCloseStartElement()
	xmlStreamWriter.writeNewLine()
	xmlStreamWriter.writer.WriteString("<!-- " + comment + " -->")
}

func (xmlStreamWriter *XmlStreamWriter) WriteInnerXml(innerXml string) {
	xmlStreamWriter.writeCloseStartElement()
	xmlStreamWriter.writer.WriteString(innerXml)
}
