package main

import "fmt"

/*
	임베딩당한 Struct의 메서드를 임베딩한 Struct에서 호출한다면 Receiver는 누가 될까?
 */

type Child struct {
	child string
	*Parent
}

func (c Child) loadFile() {
	// c에 파일 읽어오는거, 이거는 함수 하나로 빼서 하면 되니깐
}

type Parent struct {
	parent string
}

func (c *Child) Print() {
	fmt.Printf("%+v\n", c)
}


type People interface {
	loadFile()
}

func GetAgentConfig() People {
	var agentType = 1
	switch agentType {
	case 1 :
		c := Child{}
		c.loadFile()
		return c
	}

	return nil
}



func main() {
	p := GetAgentConfig()
	fmt.Printf("%+v\n", p)
	//p.(Child)

}