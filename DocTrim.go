// 将docx文件进行精简
// 可以接受xml文件，也可以接受docx文件
// 返回精简后的主文档xml

package DocTrim

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/nbio/xml"
)

const (
	refTag  = "_r"
	hashTag = "_h"
)

// Slimmer 精简器
type DocTrim struct {
	dict     map[uint64]*Node
	seq      uint64
	hashDict map[uint64]uint64
}

func (slim *DocTrim) Reset() {
	slim.dict = make(map[uint64]*Node)
	slim.seq = 1
	slim.hashDict = make(map[uint64]uint64)
}

func (slim *DocTrim) RegHash(hash uint64, node *Node) (uint64, bool) {
	if seq, ok := slim.hashDict[hash]; ok {
		node.isCompat = true
		node.hash = seq

		exists := slim.dict[seq]
		exists.refCount++
		return seq, false
	}

	slim.hashDict[hash] = slim.seq
	slim.dict[slim.seq] = node
	node.hash = slim.seq
	slim.seq++
	return slim.hashDict[hash], true
}

// MakeReader 根据URL创建zip.ReadCloser对象
// 如果URL以http://或https://开头，则下载文件并返回ReadCloser对象
// 否则，打开本地文件并返回ReadCloser对象
func (s *DocTrim) MakeReader(url string) (*zip.ReadCloser, error) {
	// download file
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("error downloading file: %v", err)
			return nil, err
		}
		defer resp.Body.Close()

		reader, err := zip.NewReader(resp.Body.(io.ReaderAt), resp.ContentLength)
		if err != nil {
			log.Fatalf("error reading zip file: %v", err)
			return nil, err
		}

		return &zip.ReadCloser{Reader: *reader}, nil
	} else {
		// open file
		return zip.OpenReader(url)
	}
}

// Process 处理文档
// 根据URL打开或下载文件，并遍历zip文件中的文件
// 如果找到"word/document.xml"文件，则读取其内容并返回
// 否则，返回空字符串
func (s DocTrim) Process(url string) ([]byte, error) {
	// open or download file
	if url == "" {
		return nil, errors.New("url is empty")
	}

	r, err := s.MakeReader(url)
	if err != nil {
		log.Fatalf("error opening zip file: %v", err)
		return nil, err
	}

	for _, f := range r.File {
		log.Printf("File: %s\n", f.Name)
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				log.Fatalf("error opening file: %v", err)
				return nil, err
			}
			defer rc.Close()
			return s.Pack(rc)
		}
	}

	return nil, errors.New("document.xml not found")
}

// xmlNodeToJson 将XML节点转换为JSON对象
// 将XML节点的属性转换为JSON对象的属性
// 将XML节点的文本内容转换为JSON对象的"_byT"属性
// 将XML节点的子节点转换为JSON对象的"_byC"属性（如果有多个子节点）或直接作为属性添加到JSON对象中
func (s DocTrim) xmlNodeToJson(node *Node) (map[string]interface{}, error) {
	jsonObj := make(map[string]interface{})
	for _, attr := range node.Attrs {
		jsonObj[attr.Name.Local] = attr.Value
	}

	textContent := strings.TrimSpace(string(node.Content))
	if textContent != "" {
		jsonObj["_byT"] = textContent
	}

	if len(node.Children) > 0 {
		childrenList := []interface{}{}
		names := []string{}
		for _, child := range node.Children {
			childJSON, _ := s.xmlNodeToJson(child)

			if strings.HasSuffix(child.XMLName.Local, "Pr") {
				delete(childJSON, "_byN")
				jsonObj[child.XMLName.Local] = childJSON
			} else {
				childrenList = append(childrenList, childJSON)
				names = append(names, child.XMLName.Local)
			}
		}

		// 将子节点添加到json对象中
		switch len(childrenList) {
		case 0:
			break
		case 1:
			delete(childrenList[0].(map[string]interface{}), "_byN")
			jsonObj[names[0]] = childrenList[0]
		default:
			if strings.HasSuffix(node.XMLName.Local, "Pr") {
				for i, child := range childrenList {
					delete(child.(map[string]interface{}), "_byN")
					jsonObj[names[i]] = child
				}
			} else {
				jsonObj["_byC"] = childrenList
			}
		}
	}
	jsonObj["_byN"] = node.XMLName.Local

	return jsonObj, nil
}

// Node 表示XML节点
type Node struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Content  []byte     `xml:",chardata"`
	Children []*Node    `xml:",any"`

	hash     uint64
	refCount int
	isCompat bool
}

func EqualXml(l, r []byte) bool {
	var lNode, rNode Node
	xml.Unmarshal(l, &lNode)
	xml.Unmarshal(r, &rNode)
	return NodeEquals(&lNode, &rNode)
}

func NodeEquals(l, r *Node) bool {
	if l.XMLName.Local == "sectPr" &&
		r.XMLName.Local == "sectPr" {
		return true
	}
	if l.XMLName.Local != r.XMLName.Local ||
		len(l.Content) != len(r.Content) ||
		!bytes.Equal(l.Content, r.Content) ||
		l.XMLName.Space != r.XMLName.Space {
		log.Printf("l: %s, r: %s", l.XMLName.Local, r.XMLName.Local)
		return false
	}

	if len(l.Attrs) != len(r.Attrs) {
		log.Printf("l: %s, r: %s", l.XMLName.Local, r.XMLName.Local)
		return false
	}

	for i, attr := range l.Attrs {
		if attr.Name.Local != r.Attrs[i].Name.Local || attr.Value != r.Attrs[i].Value || attr.Name.Space != r.Attrs[i].Name.Space {
			log.Printf("l: %s, r: %s", l.XMLName.Local, r.XMLName.Local)
			return false
		}
	}

	if len(l.Children) != len(r.Children) {
		log.Printf("l: %s, r: %s", l.XMLName.Local, r.XMLName.Local)
		return false
	}

	for i, child := range l.Children {
		if !NodeEquals(child, r.Children[i]) {
			log.Printf("l: %s, r: %s", l.XMLName.Local, r.XMLName.Local)
			return false
		}
	}

	return true
}

// ComputeHash 计算节点的哈希值
// 使用fnv算法计算节点的哈希值，并将结果存储在hash字段中
// 如果哈希值已存在于字典中，则将节点的isCompat字段设置为true
func (node *Node) ComputeHash(slim *DocTrim) uint64 {

	hash := fnv.New64a()
	hash.Write([]byte(strconv.Itoa(len(node.Children))))
	for _, child := range node.Children {

		child.ComputeHash(slim)
		childHash := []byte(strconv.FormatUint(child.hash, 16)) // Convert child.hash to []byte
		hash.Write(childHash)
	}

	if len(node.Content) > 0 {
		hash.Write([]byte(strconv.Itoa(len(node.Content))))
		hash.Write(node.Content)
	}

	hash.Write([]byte(strconv.Itoa(len(node.Attrs))))
	for _, attr := range node.Attrs {

		hash.Write([]byte(strconv.Itoa(len(attr.Name.Local))))
		hash.Write([]byte(attr.Name.Local))
		hash.Write([]byte(strconv.Itoa(len(attr.Value))))
		hash.Write([]byte(attr.Value))
	}

	hash.Write([]byte(strconv.Itoa(len(node.XMLName.Local))))
	hash.Write([]byte(node.XMLName.Local))

	seq, _ := slim.RegHash(hash.Sum64(), node)
	return seq
}

// Compact 压缩节点
// 如果节点的isCompat字段为true，则将节点的属性、内容和子节点清空
// 否则，如果节点的refCount大于0，则将节点的哈希值作为属性添加到节点中
// 最后，递归压缩子节点
func (node *Node) Compact() error {
	if node.isCompat {
		refAttr := xml.Attr{
			Name:  xml.Name{Local: refTag},
			Value: strconv.FormatUint(node.hash, 16),
		}
		node.Attrs = []xml.Attr{}
		node.Attrs = append(node.Attrs, refAttr)
		node.Content = []byte{}
		node.Children = []*Node{}

		fmt.Sprintf("COMPACTED: %s", node.XMLName.Local)

		return nil
	} else if node.refCount > 0 {
		refAttr := xml.Attr{
			Name:  xml.Name{Local: hashTag},
			Value: strconv.FormatUint(node.hash, 16),
		}
		node.Attrs = append(node.Attrs, refAttr)
	}

	for i := range node.Children {
		node.Children[i].Compact()
	}

	return nil
}

func (node *Node) UndoCompact(dict map[uint64]*Node) error {
	for i := range node.Children {
		node.Children[i].UndoCompact(dict)
	}

	for i, attr := range node.Attrs {
		if attr.Name.Local == refTag {
			hash, _ := strconv.ParseUint(attr.Value, 16, 64)
			if exist, ok := dict[hash]; ok {
				node.Attrs = []xml.Attr{}
				node.Attrs = append(node.Attrs, exist.Attrs...)
				node.Content = exist.Content
				node.Children = exist.Children
				break
			}
		} else if attr.Name.Local == hashTag {
			hash, _ := strconv.ParseUint(attr.Value, 16, 64)
			if _, ok := dict[hash]; ok {
				return errors.New("hash attribute found")
			}
			// remove hash attribute
			node.Attrs = append(node.Attrs[:i], node.Attrs[i+1:]...)

			dict[hash] = node
			break
		}
	}

	return nil
}

// Marshal 将节点转换为XML字节数组
func (node *Node) Marshal() ([]byte, error) {
	return xml.Marshal(node)
}

func EmptyToSelfClosing(from []byte) []byte {
	re := regexp.MustCompile(`<([\w:]+)([^<>]*)></([\w:]+)>`)
	to := re.ReplaceAllFunc(from, func(text []byte) []byte {
		allSubStrings := re.FindAllSubmatch(text, -1)
		if len(allSubStrings) > 0 && len(allSubStrings[0]) > 3 && bytes.Equal(allSubStrings[0][1], allSubStrings[0][3]) {
			return append(append([]byte("<"), allSubStrings[0][1]...), append(allSubStrings[0][2], []byte(" />")...)...)
		}
		return text
	})
	return to
}

const defaultHeader = `<w:document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" xmlns:w16cex="http://schemas.microsoft.com/office/word/2018/wordml/cex" xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid" xmlns:w16="http://schemas.microsoft.com/office/word/2018/wordml" xmlns:w16sdtdh="http://schemas.microsoft.com/office/word/2020/wordml/sdtdatahash" xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" mc:Ignorable="w14 w15 w16se w16cid w16 w16cex w16sdtdh wp14">`

func (node *Node) OmitNode() {
	if node.XMLName.Local == "sectPr" {
		node.Attrs = []xml.Attr{}
		node.Children = []*Node{}
	}

	for _, child := range node.Children {
		child.OmitNode()
	}

}

// Pack 压缩XML
// 使用xml.Decoder将XML解码为Node对象
// 然后计算节点的哈希值，并压缩节点及其子节点
// 最后，使用xml.Marshal将Node对象转换为XML字节数组并返回
func (slim *DocTrim) Pack(xmlData io.Reader) ([]byte, error) {
	decoder := xml.NewDecoder(xmlData)
	var root Node
	if err := decoder.Decode(&root); err != nil {
		log.Fatalf("error decoding xml: %v", err)
		return nil, err
	}

	// 将内容重复的节点使用引用标注
	slim.Reset()
	root.ComputeHash(slim)
	root.Compact()

	// 名字空间优化
	root.OmitNode()

	xml, _ := root.Marshal()

	xml = bytes.Replace(xml, []byte(defaultHeader), []byte("<w:document>"), 1)

	xml = EmptyToSelfClosing(xml)
	//fmt.Println(string(xml))

	return xml, nil
}

// Unpack 解压缩XML
// 使用xml.Decoder将XML解码为Node对象
// 然后计算节点的哈希值，并压缩节点及其子节点
// 最后，使用xml.Marshal将Node对象转换为XML字节数组并返回
func (s DocTrim) Unpack(reader io.Reader) ([]byte, error) {
	var dict = make(map[uint64]*Node)

	// replace <w:document> with defaultHeader
	xmldata, _ := ioutil.ReadAll(reader)
	xmldata = bytes.ReplaceAll(xmldata, []byte("<w:document>"), []byte(defaultHeader))

	decoder := xml.NewDecoder(bytes.NewReader(xmldata))
	var root Node
	if err := decoder.Decode(&root); err != nil {
		log.Fatalf("error decoding xml: %v", err)
		return nil, err
	}

	root.UndoCompact(dict)
	xml, _ := root.Marshal()
	//fmt.Println(string(xml))

	return xml, nil
}

func main() {
	s := DocTrim{}
	f := "test.docx"
	if len(os.Args) > 1 {
		f = os.Args[1]
	}

	if strings.HasSuffix(f, ".docx") {
		xml, err := s.Process(f)
		if err != nil {
			log.Fatalf("error processing file: %v", err)
		}
		fmt.Println(xml)
	} else if strings.HasSuffix(f, ".xml") {
		xmlFile, err := os.Open(f)
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}
		// json := s.XmlToJson(xmlFile)
		// fmt.Println(json)
		data, err := s.Pack(xmlFile)
		if err != nil {
			log.Fatalf("error packing xml: %v", err)
			return
		}
		fmt.Println(string(data))
	} else {
		log.Fatalf("unsupported file type: %s", f)
	}

}
