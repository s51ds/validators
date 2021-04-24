package test

import (
	"bufio"
	"fmt"
	"github.com/s51ds/validators/validate"
	"log"
	"os"
	"testing"
)

func TestCallSignsFromMasterSCP(t *testing.T) {
	f, err := os.Open("master.scp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		call := scanner.Text()
		if !validate.CallSign(call) {
			fmt.Println(call)
		}

	}

}
