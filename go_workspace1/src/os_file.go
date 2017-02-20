package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	//"compress/gzip"
	//"os"
	//"archive/zip"
	//"path/filepath"
	// "net/http"
)

func main() {

	//获取参数
	/*arg_num := len(os.Args)
	fmt.Println("args-number: %d", arg_num)

	for _, v := range os.Args {

		fmt.Println(v)
	}*/

	fmt.Println("----------------------------------")
	// 获取文件信息
	/*fileInfo, err := os.Stat("testfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.*/
	//Printf("System info: %+v\n\n", fileInfo.Sys())

	fmt.Println("---------------------------------")
	//rename and move a file 移动重命名

	/*file_path := "testfile.txt"
	new_path := "test2.txt"
	err1 := os.Rename(file_path, new_path)

	if err1 != nil {
		log.Fatal(err1)
	}*/

	// delete a file 删除文件
	// err2 := os.Remove("testfile.txt")
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	///打开文件

	/*file, err3 := os.Open("testfile.txt")
	if err3 != nil {
		fmt.Println(err)
	}
	file.Close()

	file, err3 = os.OpenFile("testfile.txt", os.O_APPEND, 0666)
	if err3 != nil {
		fmt.Println(err3)
	}
	file.Close()*/

	///文件是否存在

	/*	fileInfo, err := os.Stat("testfile.txt")
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("File does not exist.")
			}
		}
		fmt.Println("File does exist. File information:")
		fmt.Println(fileInfo)*/

	//复制文件
	// Open original file
	/*originalFile, err := os.Open("testfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	*/

	//搜索文件位置

	/*file, _ := os.Open("test.txt")
	  defer file.Close()

	  // Offset is how many bytes to move
	  // Offset can be positive or negative
	  var offset int64 = 5

	  // Whence is the point of reference for offset
	  // 0 = Beginning of file
	  // 1 = Current position
	  // 2 = End of file
	  var whence int = 0
	  newPosition, err := file.Seek(offset, whence)
	  if err != nil {
	      log.Fatal(err)
	  }
	  fmt.Println("Just moved to 5:", newPosition)

	  // Go back 2 bytes from current position
	  newPosition, err = file.Seek(-2, 1)
	  if err != nil {
	      log.Fatal(err)
	  }
	  fmt.Println("Just moved back two:", newPosition)

	  // Find the current position by getting the
	  // return value from Seek after moving 0 bytes
	  currentPosition, err := file.Seek(0, 1)
	  fmt.Println("Current position:", currentPosition)

	  // Go to beginning of file
	  newPosition, err = file.Seek(0, 0)
	  if err != nil {
	      log.Fatal(err)
	  }
	  fmt.Println("Position after seeking 0,0:", newPosition)
	*/

	////Write Bytes to a File(写入字节流到文件)：

	// Open a new file for writing only
	/*file, err := os.OpenFile(
		"testfile.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	byteSlice := []byte("Bytessdjflksj监控!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)*/

	//Quick Write to File(快速文件写入)：

	/*err := ioutil.WriteFile("test.txt", []byte("Hjksdfjjk健康的说法i\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}*/

	//Use Buffered Writer：

	// Open file for writing
	/* file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	   if err != nil {
	       log.Fatal(err)
	   }
	   defer file.Close()

	   // Create a buffered writer from the file
	   bufferedWriter := bufio.NewWriter(file)

	   // Write bytes to buffer
	   bytesWritten, err := bufferedWriter.Write(
	       []byte{65, 66, 67},
	   )
	   if err != nil {
	       log.Fatal(err)
	   }
	   log.Printf("Bytes written: %d\n", bytesWritten)

	   // Write string to buffer
	   // Also available are WriteRune() and WriteByte()
	   bytesWritten, err = bufferedWriter.WriteString(
	       "Buffered string\n",
	   )
	   if err != nil {
	       log.Fatal(err)
	   }
	   log.Printf("Bytes written: %d\n", bytesWritten)

	   // Check how much is stored in buffer waiting
	   unflushedBufferSize := bufferedWriter.Buffered()
	   log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	   // See how much buffer is available
	   bytesAvailable := bufferedWriter.Available()
	   if err != nil {
	       log.Fatal(err)
	   }
	   log.Printf("Available buffer: %d\n", bytesAvailable)

	   // Write memory buffer to disk
	   bufferedWriter.Flush()

	   // Revert any changes done to buffer that have
	   // not yet been written to file with Flush()
	   // We just flushed, so there are no changes to revert
	   // The writer that you pass as an argument
	   // is where the buffer will output to, if you want
	   // to change to a new writer
	   bufferedWriter.Reset(bufferedWriter)

	   // See how much buffer is available
	   bytesAvailable = bufferedWriter.Available()
	   if err != nil {
	       log.Fatal(err)
	   }
	   log.Printf("Available buffer: %d\n", bytesAvailable)

	   // Resize buffer. The first argument is a writer
	   // where the buffer should output to. In this case
	   // we are using the same buffer. If we chose a number
	   // that was smaller than the existing buffer, like 10
	   // we would not get back a buffer of size 10, we will
	   // get back a buffer the size of the original since
	   // it was already large enough (default 4096)
	   bufferedWriter = bufio.NewWriterSize(
	       bufferedWriter,
	       8000,
	   )

	   // Check available buffer size after resizing
	   bytesAvailable = bufferedWriter.Available()
	   if err != nil {
	       log.Fatal(err)
	   }
	   log.Printf("Available buffer: %d\n", bytesAvailable)

	*/

	//Read up to n Bytes from File(文件字节流读取)：

	// Open file for reading
	/* file, err := os.Open("testfile.txt")
	   if err != nil {
	       log.Fatal(err)
	   }
	   defer file.Close()

	   // Read up to len(b) bytes from the File
	   // Zero bytes written means end of file
	   // End of file returns error type io.EOF
	   byteSlice := make([]byte, 16)
	   bytesRead, err := file.Read(byteSlice)
	   if err != nil {
	       log.Fatal(err)
	   }
	   log.Printf("Number of bytes read: %d\n", bytesRead)
	   log.Printf("Data read: %s\n", byteSlice)*/

	//Read Exactly n Bytes(精确的读取字节流)：

	// Open file for reading
	/*file, err := os.Open("testfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// The file.Read() function will happily read a tiny file in to a large
	// byte slice, but io.ReadFull() will return an
	// error if the file is smaller than the byte slice.
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
	*/

	//Read All Bytes of File（读取文件所有字节）：
	// Open file for reading
	/* file, err := os.Open("test.txt")
	   if err != nil {
	       log.Fatal(err)
	   }

	   // os.File.Read(), io.ReadFull(), and
	   // io.ReadAtLeast() all work with a fixed
	   // byte slice that you make before you read

	   // ioutil.ReadAll() will read every byte
	   // from the reader (in this case a file),
	   // and return a slice of unknown slice
	   data, err := ioutil.ReadAll(file)
	   if err != nil {
	       log.Fatal(err)
	   }

	   fmt.Printf("Data as hex: %x\n", data)
	   fmt.Printf("Data as string: %s\n", data)
	   fmt.Println("Number of bytes read:", len(data))*/

	//Quick Read Whole File to Memory（快速读取文件到内存）：

	// Read file to byte slice
	/* data, err := ioutil.ReadFile("test.txt")
	   if err != nil {
	       log.Fatal(err)
	   }

	   log.Printf("Data read: %s\n", data)*/

	//Use Buffered Reader：

	// Open file and create a buffered reader on top
	/* file, err := os.Open("test.txt")
	   if err != nil {
	       log.Fatal(err)
	   }
	   bufferedReader := bufio.NewReader(file)

	   // Get bytes without advancing pointer
	   byteSlice := make([]byte, 5)
	   byteSlice, err = bufferedReader.Peek(5)
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	   // Read and advance pointer
	   numBytesRead, err := bufferedReader.Read(byteSlice)
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	   // Ready 1 byte. Error if no byte to read
	   myByte, err := bufferedReader.ReadByte()
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Read 1 byte: %c\n", myByte)

	   // Read up to and including delimiter
	   // Returns byte slice
	   dataBytes, err := bufferedReader.ReadBytes('\n')
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Read bytes: %s\n", dataBytes)

	   // Read up to and including delimiter
	   // Returns string
	   dataString, err := bufferedReader.ReadString('\n')
	   if err != nil {
	       log.Fatal(err)
	   }
	   fmt.Printf("Read string: %s\n", dataString)

	   // This example reads a few lines so test.txt
	   // should have a few lines of text to work correct
	*/

	//read with a scanner:
	// Open file and create scanner on top of it
	/*  file, err := os.Open("test.txt")
	    if err != nil {
	        log.Fatal(err)
	    }
	    scanner := bufio.NewScanner(file)

	    // Default scanner is bufio.ScanLines. Lets use ScanWords.
	    // Could also use a custom function of SplitFunc type
	    scanner.Split(bufio.ScanWords)

	    // Scan for next token.
	    success := scanner.Scan()
	    if success == false {
	        // False on error or EOF. Check error
	        err = scanner.Err()
	        if err == nil {
	            log.Println("Scan completed and reached EOF")
	        } else {
	            log.Fatal(err)
	        }
	    }

	    // Get data from scan with Bytes() or Text()
	    fmt.Println("First word found:", scanner.Text())

	    // Call scanner.Scan() again to find next token

	*/

	//Archive Files（文件归档）：
	/*   outFile, err := os.Create("test.zip")
	if err != nil {
	    log.Fatal(err)
	}
	defer outFile.Close()

	// Create a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outFile)


	// Add files to archive
	// We use some hard coded data to demonstrate,
	// but you could iterate through all the files
	// in a directory and pass the name and contents
	// of each file, or you can take data from your
	// program and write it write in to the archive
	// without
	var filesToArchive = []struct {
	    Name, Body string
	} {
	    {"test.txt", "String contents of file"},
	    {"test2.txt", "\x61\x62\x63\n"},
	}

	// Create and write files to the archive, which in turn
	// are getting written to the underlying writer to the
	// .zip file we created at the beginning
	for _, file := range filesToArchive {
	        fileWriter, err := zipWriter.Create(file.Name)
	        if err != nil {
	                log.Fatal(err)
	        }
	        _, err = fileWriter.Write([]byte(file.Body))
	        if err != nil {
	                log.Fatal(err)
	        }
	}

	// Clean up
	err = zipWriter.Close()
	if err != nil {
	        log.Fatal(err)
	}*/

	//Extract Archived Files（解压缩归档文件）：

	// Create a reader out of the zip archive
	/* zipReader, err := zip.OpenReader("test.zip")
	   if err != nil {
	       log.Fatal(err)
	   }
	   defer zipReader.Close()

	   // Iterate through each file/dir found in
	   for _, file := range zipReader.Reader.File {
	       // Open the file inside the zip archive
	       // like a normal file
	       zippedFile, err := file.Open()
	       if err != nil {
	           log.Fatal(err)
	       }
	       defer zippedFile.Close()

	       // Specify what the extracted file name should be.
	       // You can specify a full path or a prefix
	       // to move it to a different directory.
	       // In this case, we will extract the file from
	       // the zip to a file of the same name.
	       targetDir := "./"
	       extractedFilePath := filepath.Join(
	           targetDir,
	           file.Name,
	       )

	       // Extract the item (or create directory)
	       if file.FileInfo().IsDir() {
	           // Create directories to recreate directory
	           // structure inside the zip archive. Also
	           // preserves permissions
	           log.Println("Creating directory:", extractedFilePath)
	           os.MkdirAll(extractedFilePath, file.Mode())
	       } else {
	           // Extract regular file since not a directory
	           log.Println("Extracting file:", file.Name)

	           // Open an output file for writing
	           outputFile, err := os.OpenFile(
	               extractedFilePath,
	               os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
	               file.Mode(),
	           )
	           if err != nil {
	               log.Fatal(err)
	           }
	           defer outputFile.Close()

	           // "Extract" the file by copying zipped file
	           // contents to the output file
	           _, err = io.Copy(outputFile, zippedFile)
	           if err != nil {
	               log.Fatal(err)
	           }
	       }
	   }*/

	//Compress a File（压缩文件）：

	// Create .gz file to write to
	/* outputFile, err := os.Create("test.txt.gz")
	   if err != nil {
	       log.Fatal(err)
	   }

	   // Create a gzip writer on top of file writer
	   gzipWriter := gzip.NewWriter(outputFile)
	   defer gzipWriter.Close()

	   // When we write to the gzip writer
	   // it will in turn compress the contents
	   // and then write it to the underlying
	   // file writer as well
	   // We don't have to worry about how all
	   // the compression works since we just
	   // use it as a simple writer interface
	   // that we send bytes to
	   _, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	   if err != nil {
	       log.Fatal(err)
	   }

	   log.Println("Compressed data written to file.") */

	// Uncompress a File（文件解压）：
	// Open gzip file that we want to uncompress
	// The file is a reader, but we could use any
	// data source. It is common for web servers
	// to return gzipped contents to save bandwidth
	// and in that case the data is not in a file
	// on the file system but is in a memory buffer
	/* gzipFile, err := os.Open("test.txt.gz")
	   if err != nil {
	       log.Fatal(err)
	   }

	   // Create a gzip reader on top of the file reader
	   // Again, it could be any type reader though
	   gzipReader, err := gzip.NewReader(gzipFile)
	   if err != nil {
	       log.Fatal(err)
	   }
	   defer gzipReader.Close()

	   // Uncompress to a writer. We'll use a file writer
	   outfileWriter, err := os.Create("unzipped.txt")
	   if err != nil {
	       log.Fatal(err)
	   }
	   defer outfileWriter.Close()

	   // Copy contents of gzipped file to output file
	   _, err = io.Copy(outfileWriter, gzipReader)
	   if err != nil {
	       log.Fatal(err)
	   }*/

	//Downloading a File Over HTTP（Http文件下载）：

	// Create output file
	/* newFile, err := os.Create("devdungeon.html")
	   if err != nil {
	        log.Fatal(err)
	   }
	   defer newFile.Close()

	   // HTTP GET request devdungeon.com
	   url := "http://www.devdungeon.com/archive"
	   response, err := http.Get(url)
	   defer response.Body.Close()

	   // Write bytes from HTTP response to file.
	   // response.Body satisfies the reader interface.
	   // newFile satisfies the writer interface.
	   // That allows us to use io.Copy which accepts
	   // any type that implements reader and writer interface
	   numBytesWritten, err := io.Copy(newFile, response.Body)
	   if err != nil {
	        log.Fatal(err)
	   }
	   log.Printf("Downloaded %d byte file.\n", numBytesWritten)*/

}
