import os

def get_files_info(working_directory: str, directory: str = ".") -> str:
    if directory == ".":
        return_statement = f"Result for current directory: \n"
    else:
        return_statement = f"Result for '{directory}' directory: \n"
    try:
        working_directory_abs_path = os.path.abspath(working_directory)
        target_directory_abs_path = os.path.normpath(os.path.join(working_directory_abs_path, directory))
        valid_target_directory = os.path.commonpath([working_directory_abs_path, target_directory_abs_path]) == working_directory_abs_path
        if not valid_target_directory:
            return return_statement + f'Error: Cannot list "{directory}" as it is outside the permitted working directory'
        elif not os.path.isdir(target_directory_abs_path):
            return return_statement + f'Error: "{directory}" is not a directory'
        else:
            try:
                files = os.listdir(target_directory_abs_path)
                for file in files:
                    full_path = os.path.join(target_directory_abs_path, file)
                    size = os.path.getsize(full_path)
                    is_dir = os.path.isdir(full_path)
                    return_statement += f"- {file}: file_size={size} bytes, is_dir={is_dir}\n"
            except Exception as e:
                return return_statement + f"Error: unexpected error occurred while listing files in {directory}: {str(e)}"
            return return_statement
    except Exception as e:
        return return_statement + f"Error: unexpected error occurred while getting files info for {directory}: {str(e)}"