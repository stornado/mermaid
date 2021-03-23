package main

import (
	"io/ioutil"

	"github.com/stornado/mermaid/pkg/mermaid"
)

func main() {
	sequenceDiagram := `
    sequenceDiagram
    participant Alice
    participant Bob
    participant John as John<br />Second Line
    rect rgb(200, 220, 100)
    rect rgb(200, 255, 200)
    Alice ->> Bob: Hello Bob, how are you?
    Bob-->>John: How about you John?
    end
    Bob--x Alice: I am good thanks!
    Bob-x John: I am good thanks!
    Note right of John: John thinks a long<br />long time, so long<br />that the text does<br />not fit on a row.
    Bob-->Alice: Checking with John...
    Note over John:wrap: John looks like he's still thinking, so Bob prods him a bit.
    Bob-x John: Hey John - we're still waiting to know<br />how you're doing
    Note over John:nowrap: John's trying hard not to break his train of thought.
    Bob-x John:wrap: John! Are you still debating about how you're doing? How long does it take??
    Note over John: After a few more moments, John<br />finally snaps out of it.
    end
    alt either this
    Alice->>John: Yes
    else or this
    Alice->>John: No
    else or this will happen
    Alice->John: Maybe
    end
    par this happens in parallel
    Alice -->> Bob: Parallel message 1
    and
    Alice -->> John: Parallel message 2
    end	
`
	svg, png, err := mermaid.Render(sequenceDiagram)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("mermaid.svg", []byte(svg), 0644)
	ioutil.WriteFile("mermaid.png", png, 0644)
}
