package main

import (
	"fmt"
	"log"
)

func main() {
	master := []byte("sample master passwd")
	info, err := NewPasswordInfo("gmail")
	if err != nil {
		log.Fatal(err)
	}

	info.AddRequiredChars(&RequiredChars{
		Chars: "!&@#",
		Num:   2,
	})

	err = GenerateSalt(master, info)
	if err != nil {
		log.Fatal(err)
	}

	pw, err := Generate(master, info)
	fmt.Println("pw", pw, "err", err)
}
