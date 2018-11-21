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
  
  current_directory, _ := os.Getwd()
  fmt.Println("Current directory from change_into_directory_1(): " + current_directory)
  
  cmd := exec.Command("cat", "file_1.txt")
  cmd.Dir = current_directory + "/DIR_1"

  output, error := cmd.Output()
  
  if error != nil {
    fmt.Println("<<< ERROR getting file 1 from directory 1 >>>")
  }
  
  fmt.Println(string(output))
}

func change_into_directory_2() {

  defer wait_group2.Done()
  
  current_directory, _ := os.Getwd()
  fmt.Println("Current directory from change_into_directory_2(): " + current_directory)
  
  cmd := exec.Command("cat", "file_2.txt")
  cmd.Dir = current_directory + "/DIR_2"

  output, error := cmd.Output()
  
  if error != nil {
    fmt.Println("<<< ERROR getting file 2 from directory 2 >>>")
  }
  
  fmt.Println(string(output))
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
