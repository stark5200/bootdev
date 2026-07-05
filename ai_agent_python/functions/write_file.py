import os

def write_file(working_directory: str, file_path: str, content: str) -> str:
    try:  
        working_directory_abs_path = os.path.abspath(working_directory)
        absolute_file_path = os.path.normpath(os.path.join(working_directory_abs_path, file_path))
        valid_target_directory = os.path.commonpath([working_directory_abs_path, absolute_file_path]) == working_directory_abs_path
        if not valid_target_directory:
            return f'Error: Cannot write to "{file_path}" as it is outside the permitted working directory'
        if os.path.isdir(absolute_file_path):
            return f'Error: Cannot write to "{file_path}" as it is a directory'
        os.makedirs(os.path.dirname(absolute_file_path), exist_ok=True)
        with open(absolute_file_path, 'w') as f:
            f.write(content)
        return f'Successfully wrote to "{file_path}" ({len(content)} characters written)'
    except Exception as e:
        return f"Error: unexpected error occurred while writing content {content} to {file_path}: {str(e)}"
      
      
schema_write_file = {
    "type": "function",
    "function": {
        "name": "write_file",
        "description": "Writes content to a specified file relative to the working directory",
        "parameters": {
            "type": "object",
            "properties": {
                "file_path": {
                    "type": "string",
                    "description": "Path to the file to write to, relative to the working directory",
                },
                "content": {
                    "type": "string",
                    "description": "Content to write to the file",
                },
            },
            "required": ["file_path", "content"],
        },
    },
}