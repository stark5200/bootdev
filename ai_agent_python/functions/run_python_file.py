import os
import subprocess


def run_python_file(working_directory: str, file_path: str, args: list[str] | None = None) -> str:
    try:
        working_directory_abs_path = os.path.abspath(working_directory)
        absolute_file_path = os.path.normpath(os.path.join(working_directory_abs_path, file_path))
        valid_target_directory = os.path.commonpath([working_directory_abs_path, absolute_file_path]) == working_directory_abs_path
        
        if not valid_target_directory:
            return f'Error: Cannot execute "{file_path}" as it is outside the permitted working directory'
        if not os.path.isfile(absolute_file_path):
            return f'Error: "{file_path}" does not exist or is not a regular file'
        if not file_path.endswith(".py"):
            return f'Error: "{file_path}" is not a Python file'
          
        command = ["python", absolute_file_path]
        command.extend(args or [])
        
        complete_process = subprocess.run(command, capture_output=True, text=True, cwd=working_directory_abs_path, timeout=30)
        
        output_string = ""
        
        if complete_process.returncode != 0:
            output_string += f"Process exited with code {complete_process.returncode}" + "\n"
            
        if len(complete_process.stdout + complete_process.stderr) == 0:
            output_string += "No output produced" + "\n"
        else:    
            output_string += "STDOUT: " + complete_process.stdout + "\n"
            output_string += "STDERR: " + complete_process.stderr + "\n"
            
        return output_string
    except Exception as e:
        return f"Error: executing Python file: {e}"
      
      
schema_run_python_file = {
    "type": "function",
    "function": {
        "name": "run_python_file",
        "description": "Executes a Python file in the specified working directory",
        "parameters": {
            "type": "object",
            "properties": {
                "file_path": {
                    "type": "string",
                    "description": "Path to the Python file to execute, relative to the working directory",
                },
                "args": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "description": "Arguments to pass to the Python file",
                },
            },
            "required": ["file_path"],
        },
    },
}