package main

import (
	"fmt"
	"os"
  "os/exec"
  "sync"
)

var wait_group1 sync.WaitGroup
var wait_group2 sync.WaitGroup

func change_into_directory_1() {

  defer wait_group1.Done()
  
  os.Chdir("./DIR_1")
  current_directory, _ := os.Getwd()
  fmt.Println("Current directory from change_into_directory_1(): " + current_directory)
  
  cmd := exec.Command("cat", "file_1.txt")
  output, error := cmd.Output()
  
  if error != nil {
    fmt.Println("<<< ERROR >>>")
  }
  
  fmt.Println(string(output))
  os.Chdir("..") 
}

func change_into_directory_2() {

  defer wait_group2.Done()
  
  os.Chdir("./DIR_2")
  current_directory, _ := os.Getwd()
  fmt.Println("Current directory from change_into_directory_2(): " + current_directory)
  
  cmd := exec.Command("cat", "file_2.txt")
  output, error := cmd.Output()
  
  if error != nil {
    fmt.Println("<<< ERROR >>>")
  }
  
  fmt.Println(string(output))
  os.Chdir("..") 
}

///////////////////////////////////

func main() {
  
  fmt.Println("Running the functions sequentially:\n")
  
  wait_group1.Add(1)
  wait_group2.Add(1)
  
  change_into_directory_1()
  change_into_directory_2()
  
  fmt.Println("\n--------------- Running the functions in parallel ---------------\n")
  
  wait_group1.Add(1)
  wait_group2.Add(1)
  
  go change_into_directory_1()
  go change_into_directory_2()
  
  wait_group1.Wait()
  wait_group2.Wait()
}
