package main

import (
	"fmt"
	"os"
  "os/exec"
  "sync"
  "io/ioutil"
  "strings"
  "strconv"
)

var wait_group sync.WaitGroup

///////////////////////////////////

func generate_text_file(target_dir string, filename string) {
    
  binary_output := []byte("This is " + filename + "\n")  
  
  path_name := target_dir + "/" + filename

  ioutil.WriteFile(path_name, binary_output, 0755)
}

///////////////////////////////////

func generate_bash_script(target_dir string, command_string string) string {
  
  exec.Command("mkdir", "-p", target_dir).Run()
  
  uuid, _ := exec.Command("uuidgen").Output()
  uuid_string := strings.TrimSpace(string(uuid))
    
  binary_output := []byte("#!/bin/bash\ncd " + target_dir + "\n" + command_string + "\n")  
  
  script_name := uuid_string + ".sh" 
  path_name := target_dir + "/" + script_name

  ioutil.WriteFile(path_name, binary_output, 0755)
  
  return script_name
}

///////////////////////////////////

func change_into_directory_and_run(target_dir string, target_file string) {
  fmt.Println(target_dir)  
  defer wait_group.Done()
  
  command_string := "cat " + target_file
  
  script_name := generate_bash_script(target_dir, command_string)
  generate_text_file(target_dir, target_file)
    
  cmd := exec.Command(target_dir + "/" + script_name)
  output, error := cmd.Output()
  os.Remove(target_dir + "/" + script_name)
  
  if error != nil {
    fmt.Println("<<< ERROR getting file from directory >>>")
    //panic(error)
  }
  
  fmt.Println(string(output))
}

///////////////////////////////////
///////////////////////////////////

func main() {
  
  fmt.Println("\n--------------- Running the functions in parallel ---------------\n")
  
  for loop := 0; loop < 10; loop++ {
    wait_group.Add(1)
    loop_as_string := strconv.Itoa(loop)
    go change_into_directory_and_run("./DIR_" + loop_as_string, "file_" + loop_as_string + ".txt")
  }
  
  wait_group.Wait()
}
