package main

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	//"github.com/sirupsen/logrus"
	//"testing"
	//"time"
	_ "github.com/mattn/go-sqlite3" 
)

type MyFact struct {

	userId int
	gender int
	age int
	income int

	//IntAttribute     int64
	//StringAttribute  string
	//BooleanAttribute bool
	//FloatAttribute   float64
	//TimeAttribute    time.Time
	WhatToSay        string
}

func (mf *MyFact) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}

//func TestTutorial(t *testing.T) {
func main() {	
	//logrus.SetLevel(logrus.DebugLevel)
	fmt.Println("start rule check...")


	myFact := &MyFact{
		userId:     1,
		gender:  1,
		age: 35,
		income:   6000,

		//IntAttribute:     123,
		//StringAttribute:  "Some string value",
		//BooleanAttribute: true,
		//FloatAttribute:   1.234,
		//TimeAttribute:    time.Now(),
	}
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MF", myFact)
	if err != nil {
		panic(err)
	}

	workingMemory := ast.NewWorkingMemory()
	knowledgeBase := ast.NewKnowledgeBase("Tutorial", "0.0.1")
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase, workingMemory)

	drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.age == 35 && MF.income == 6000
    then
        MF.WhatToSay = MF.GetWhatToSay("pass");
		Retract("CheckValues");
}
`
	byteArr := pkg.NewBytesResource([]byte(drls))
	err = ruleBuilder.BuildRuleFromResource(byteArr)
	if err != nil {
		panic(err)
	}

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase, workingMemory)
	if err != nil {
		panic(err)
	}

	if myFact.WhatToSay != "Let say \"pass\"" {
		println("Not pass")
	} else {
		println(myFact.WhatToSay)
	}

}
