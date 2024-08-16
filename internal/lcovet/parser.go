package lcovet

import (
	"bufio"
	"bytes"
	"io"
)

var eof = rune(0)

var Identifiers map[Token]string

type Parser struct {
	r *bufio.Reader
}

type Node struct {
	Token Token
	Id    string
	Value bytes.Buffer
}

type Record struct {
	File           bytes.Buffer   `json:"source_file"`
	Test           bytes.Buffer   `json:"test_name"`
	Lines          []string       `json:"line_data"`
	LinesFound     bytes.Buffer   `json:"lines_found"`
	LinesHit       bytes.Buffer   `json:"lines_hit"`
	Functions      []FunctionData `json:"functions_data"`
	FunctionsFound bytes.Buffer   `json:"functions_found"`
	FunctionsHit   bytes.Buffer   `json:"functions_hit"`
	Branches       []BranchData   `json:"branches_data"`
	BranchesFound  bytes.Buffer   `json:"branches_found"`
	BranchesHit    bytes.Buffer   `json:"branches_hit"`
}

func NewParser(r io.Reader) *Parser {
	return &Parser{r: bufio.NewReader(r)}
}

func (p *Parser) createNodes() *[]Node {
	var result []Node
	for {
		line, err := p.readLine()
		if err != nil {
			break
		}
		key, value := NewScanner(&line).Scan()
		node := Node{Token: key, Id: Identifiers[key], Value: value}
		result = append(result, node)
	}
	return &result
}

func (p *Parser) Parse() *[]Record {
	var result []Record
	nodes := p.createNodes()
	for record := range PartitionPerRecord(nodes) {
		rec := Record{}
		collectRecordData(&rec, record)
		result = append(result, rec)
	}
	return &result
}

type FunctionData struct {
	Line bytes.Buffer
	Name bytes.Buffer
}

type BranchData struct {
	Line   bytes.Buffer
	Block  bytes.Buffer
	Branch bytes.Buffer
	Hit    bytes.Buffer
}

func parseData(channel chan bytes.Buffer, data bytes.Buffer) {
	parts := bytes.FieldsFunc(data.Bytes(), func(ch rune) bool {
		return ch == ',' || ch == eof
	})
	go func() {
		for _, part := range parts {
			channel <- *bytes.NewBuffer(part)
		}
	}()
}

func collectRecordData(rec *Record, nodes []Node) {
	c := make(chan bytes.Buffer)
	for _, node := range nodes {
		switch node.Token {
		case SF:
			rec.File = node.Value
		case TN:
			rec.Test = node.Value
		case LF:
			rec.LinesFound = node.Value
		case LH:
			rec.LinesHit = node.Value
		case FN:
			parseData(c, node.Value)
			rec.Functions = append(rec.Functions, FunctionData{Line: <-c, Name: <-c})
		case FNF:
			rec.FunctionsFound = node.Value
		case FNH:
			rec.FunctionsHit = node.Value
		case BRDA:
			parseData(c, node.Value)
			rec.Branches = append(rec.Branches, BranchData{Line: <-c, Block: <-c, Branch: <-c, Hit: <-c})
		case BRF:
			rec.BranchesFound = node.Value
		case BRH:
			rec.BranchesHit = node.Value
		}
	}
	close(c)
}

func PartitionPerRecord(nodes *[]Node) chan []Node {
	channel := make(chan []Node)
	go func() {
		var result []Node
		for _, node := range *nodes {
			if node.Token == EOF {
				channel <- result
				result = nil
				continue
			}
			result = append(result, node)
		}
		close(channel)
	}()
	return channel
}

func (p *Parser) readLine() (buf bytes.Buffer, err error) {
	line, _, err := p.r.ReadLine()
	if err != nil {
		return
	}
	buf.Write(line)
	return
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func init() {
	Identifiers = make(map[Token]string)
	Identifiers[TN] = "TN"
	Identifiers[SF] = "SF"
	Identifiers[FN] = "FN"
	Identifiers[FNDA] = "FNDA"
	Identifiers[FNF] = "FNF"
	Identifiers[FNH] = "FNH"
	Identifiers[BRDA] = "BRDA"
	Identifiers[BRF] = "BRF"
	Identifiers[BRH] = "BRH"
	Identifiers[DA] = "DA"
	Identifiers[LF] = "LF"
	Identifiers[LH] = "LH"
	Identifiers[EOF] = "end_of_record"
}
