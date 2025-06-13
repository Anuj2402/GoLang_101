package main

import (
	// "fmt"
	// "bufio"
	"fmt"
	"os"
	// "strings"
)



func main() {
	/*

	f, err := os.Open("example.txt")
	if err != nil {
		// log the error 
		panic(err) 
	}
	fileInfo, err := f.Stat()
	if err != nil {
		// log the error 
		panic(err) 
	}

	fmt.Println("File name", fileInfo.Name())
	fmt.Println("Is Directory ?? -> ", fileInfo.IsDir())
	fmt.Println("File Size is", fileInfo.Size())
	fmt.Println("File Permissions", fileInfo.Mode())
	fmt.Println("File last modified", fileInfo.ModTime())
	*/
/*
	// Read the file 
	f2 , err :=os.Open("example.txt")
	
	if err!= nil {
		panic(err)
	}

	
	defer f2.Close()
	fileInfo, err := f2.Stat()
	if err != nil {
		panic(err)
	}
	filesize := fileInfo.Size()
	
	buf := make([]byte, filesize)

	d,err := f2.Read(buf)

	if err!=nil{
		panic(err)
	}


	for i:=0;i<len(buf); i++{
		println("data", d, string(buf[i]))
	}
	// println("data", d, string(buf)) // output will be like -> data 10 Hello Anuj-> here 10 is buffer size and Hello Anuj is only contain on that size \
	
*/

/*
// Read Folders i.e how many files are present in that aprticular folder or directory 

  dir, err := os.Open(".")

  if err!=nil{
	panic(err)

  }

  defer dir.Close()

     fileInfo,err:= dir.ReadDir(2) 

	 for _, fi := range fileInfo{
		fmt.Println(fi.Name(), fi.IsDir()) // outout -> files.go, example.txt
	 }

*/

/*

// creating a File 
    file,err:= os.Create("example2.txt")
	if err !=nil{
		panic(err)

	}

	defer file.Close() // Don't forget to close the file
	file.WriteString("Hi Golang\n")
	file.WriteString("nice language")


	// Step 2: Read the content back from the file
	data, err := os.ReadFile("example2.txt")
	if err != nil {
		panic(err)
	}

	// Step 3: Replace "language" with "Language"
	modified := strings.ReplaceAll(string(data), "language", "Language")

	// Step 4: Overwrite the file with the updated content
/*
		[]byte(modified) = converting string to bytes.

Required because os.WriteFile needs bytes.

It ensures the modified content (with replacements) is saved correctly to the file.

0644 -> is the file permssions

	err = os.WriteFile("example2.txt", []byte(modified), 0644)
	if err != nil {
		panic(err)
	}
	*/

/*
	//Read and write to another file (streaming fashion)

	sourceFile , err := os.Open("example.txt")

	if err!= nil{
		panic(err)
	}



	defer sourceFile.Close()


	destFile , err := os.Create("example3.txt")
	if err!=nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil{
			if err.Error() != "EOF"{
				panic(err)
			}
			break 
			
		}

		e:= writer.WriteByte(b)
		if e !=nil {
			panic(e)
		}



	}

	writer.Flush()

	fmt.Println("written to new file successfully")

	*/


	// delete a File 

	err:= os.Remove("example2.txt")

	if err!=nil {
		panic(err)
	}

	fmt.Println("File Deleted successfully ")

}

