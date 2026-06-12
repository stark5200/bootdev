import os

def get_files_info(working_directory: str, directory: str = ".") -> str:
    try:
        working_directory_abs_path = os.path.abspath(working_directory)
        target_directory_abs_path = os.path.normpath(os.path.join(working_directory_abs_path, directory))
        valid_target_directory = os.path.commonpath([working_directory_abs_path, target_directory_abs_path]) == working_directory_abs_path
        if not valid_target_directory:
            return f'Error: Cannot list "{directory}" as it is outside the permitted working directory'
        elif not os.path.isdir(target_directory_abs_path):
            return f'Error: "{directory}" is not a directory'
        else:
            return f'Success: "{directory}" is within the working directory'
    except Exception as e:
        return f"Error: unexpected error occurred while getting files info for {directory}: {str(e)}"