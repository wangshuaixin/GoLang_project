package main

import(
	"flag"
	"fmt"
	"bufio" // use to readfile or readline from file
	"io"
	"os"
	"strconv" // use to convert type between str/int
	"time"
)

import "../algorithm/bubblesort"
import "../algorithm/qsort"

var(
	infile *string = flag.String("i","infile","File contains values for sorting")
	outfile *string = flag.String("o","outfile","File to receive sorted values")
	algorithm *string = flag.String("a","qsort","Sort algorithm")
)

func readValues(infile string)(values []int, err error){
	file, err := os.Open(infile)
	if err != nil{
		fmt.Println("Failed to open the input file.", infile)
		return
	}

	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil{
			if err1 != io.EOF{
				err = err1
			}
			break
		}

		if isPrefix{
			fmt.Println("A too long line, seems unexpexcted")
			return
		}

		str := string(line) // turn char to string

		value, err1 := strconv.Atoi(str)
		if err1 != nil{
			err = err1
			return
		}
		values = append(values,value)
	}
	return 
}

func writeValues(values []int, outfile string) error{
	file, err := os.Create(outfile)
	if err != nil{
		fmt.Println("File to create the output file: ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str+"\n")
	}
	return nil
}

func main(){
	flag.Parse()

	if infile != nil{
		fmt.Println("infile = ",*infile,"outfile =",*outfile,"algorithm =",*algorithm)
	}
	values, err := readValues(*infile)
	fmt.Println("Read values:", values)
	if err == nil{
		t1 := time.Now()
		switch *algorithm{
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.Bubblesort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknow or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs",t2.Sub(t1),"to complete.")
		
		writeValues(values, *outfile)
	}else{
		fmt.Println(err)
	}
}

