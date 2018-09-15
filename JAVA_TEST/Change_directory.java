import java.lang.System;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.File;

///////////////////////////////////////////////////////////////////////////////////////////////////////////

class Change_directory {

  public static void change_into_directory(String dir, String filename) {
  
    StringBuffer output = new StringBuffer();
    StringBuffer error_output = new StringBuffer();  
    
		Process process_object;
		
		String current_dir = ".";
		
		try {			
      
			ProcessBuilder process_builder = new ProcessBuilder("cat", filename);
      process_builder.directory(new File(current_dir + "/" + dir));
      process_object = process_builder.start();
      process_object.waitFor();            
       
			BufferedReader reader = new BufferedReader(new InputStreamReader(process_object.getInputStream()));
      BufferedReader error_reader = new BufferedReader(new InputStreamReader(process_object.getErrorStream()));
      
      String line = "";
      String error_line = "";
      
			while ((line = reader.readLine()) != null) {
				output.append(line);
			}
			while ((error_line = error_reader.readLine()) != null) {
				error_output.append(error_line);
			}

		} catch (Exception e) {
			e.printStackTrace();
		}

		System.out.println(output.toString());
		System.out.println(error_output.toString());
  }
  
  ///////////////////////////////////////////////////////////////////////////////////////////////////////////
}

