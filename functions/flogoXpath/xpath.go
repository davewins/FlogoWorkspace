package flogoXpath

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnXPATH struct {
}

func init() {
	function.Register(&fnXPATH{})
}

func (s *fnXPATH) Name() string {
	return "xpath"
}

func (s *fnXPATH) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *fnXPATH) Eval(in ...interface{}) (interface{}, error) {

	xpath, err := coerce.ToString(in[0])
	if err != nil {
		return nil, fmt.Errorf("xpath function first parameter [%+v] must be string", in[0])
	}
	xml, err := coerce.ToString(in[1])
	if err != nil {
		return nil, fmt.Errorf("xpath function second parameter [%+v] must be string", in[1])
	}
	fmt.Println("XPATH:", xpath)
	fmt.Println("XML:", xml)
	doc, err := xmlquery.Parse(strings.NewReader(xml))
	fmt.Println("Completed the parse step")
	if err != nil {
		fmt.Println("The parse step failed")
		panic(err)
	}
	fmt.Println("About to execute the queryAll function")
	result := xmlquery.FindOne(doc, xpath)
	//fmt.Printf("result: %s\n", result.InnerText())

	return result.InnerText(), nil

}
