import os
from config import MAX_CHARS

def get_file_content(working_directory: str, file_path: str) -> str:      
    try:  
        file_content = ""
        working_directory_abs_path = os.path.abspath(working_directory)
        absolute_file_path = os.path.normpath(os.path.join(working_directory_abs_path, file_path))
        valid_target_directory = os.path.commonpath([working_directory_abs_path, absolute_file_path]) == working_directory_abs_path
        if not os.path.isfile(absolute_file_path):
            return f'Error: File not found or is not a regular file: "{file_path}"'
        if not valid_target_directory:
            return f'Error: Cannot read "{file_path}" as it is outside the permitted working directory'
        with open(absolute_file_path, 'r') as f:
            file_content += f.read(MAX_CHARS)
            if f.read(1):
                file_content += f'[...File "{file_path}" truncated at {MAX_CHARS} characters]'
        return file_content
    except Exception as e:
        return f"Error: unexpected error occurred while accessing {file_path}: {str(e)}"
      

schema_get_file_content = {
    "type": "function",
    "function": {
        "name": "get_file_content",
        "description": "Retrieves the content of a specified file relative to the working directory",
        "parameters": {
            "type": "object",
            "properties": {
                "file_path": {
                    "type": "string",
                    "description": "Path to the file to read, relative to the working directory",
                },
            },
            "required": ["file_path"],
        },
    },
}