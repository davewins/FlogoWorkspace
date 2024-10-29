package flogoXpath

import (
	"github.com/project-flogo/core/data/expression/function"
)

var in = &fnXPATH{}

func init() {
	function.ResolveAliases()
}

//func TestInt64Sample(t *testing.T) {
//	final, err := in.Eval(123, 456)
//	assert.Nil(t, err)
//	assert.Equal(t, int(579), final)
//}
