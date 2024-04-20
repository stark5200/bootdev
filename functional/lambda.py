type = {
    ".txt": "Text", 
    ".docx": "Document", 
    ".py": "Code"
}

def categorize_file(filename):
    get_category = lambda extension: type.get(extension, "Unknown")
    return get_category(filename[filename.rfind(".") :])

