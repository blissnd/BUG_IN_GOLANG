import java.lang.System;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.File;

///////////////////////////////////////////////////////////////////////////////////////////////////////////

class Concurrency_Test implements Runnable {   
  
  public void run() {
  
    Change_directory.change_into_directory("DIR_1", "file_1.txt");
  }
  
  ////////////////////////////////////////////////  
  
  public static void main(String[] args) {
            
    Thread thread_object1 = new Thread(new Concurrency_Test());
    thread_object1.start();
    
    Thread thread_object2 = new Thread(new Concurrency_Test2());
    thread_object2.start();
    
    Thread thread_object3 = new Thread(new Concurrency_Test3());
    thread_object3.start();
    
    Thread thread_object4 = new Thread(new Concurrency_Test());
    thread_object4.start();
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

class Concurrency_Test2 implements Runnable {   
  
  public void run() {
      
    Change_directory.change_into_directory("DIR_2", "file_2.txt");  
  }
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

class Concurrency_Test3 implements Runnable {   
  
  public void run() {
  
    Change_directory.change_into_directory("DIR_3", "file_3.txt");
  }
}
            
///////////////////////////////////////////////////////////////////////////////////////////////////////////

