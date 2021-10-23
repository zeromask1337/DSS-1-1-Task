package cos418_hw1_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func sumWorker(nums chan int, out chan int) {
	sum := 0
	for number := range nums {
		sum += number
	}
	out <- sum
}

func sum(num int, fileName string) int {

	fileReader, err := os.Open(fileName)
	checkError(err)
	ints, err := readInts(fileReader)
	checkError(err)

	sum := 0
	chIn := make(chan int, len(ints))
	chOut := make(chan int)

	for _, i := range ints {
		chIn <- i
	}

	close(chIn)

	for i := 1; i <= num; i++ {
		go sumWorker(chIn, chOut)
	}

	for i := 1; i <= num; i++ {
		sum += <-chOut
	}
	return sum
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
