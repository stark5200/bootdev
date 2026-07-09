import os
from textnode import markdown_to_html_node
from parentnode import ParentNode

def extract_title(markdown: str) -> str:
    if markdown.startswith("# "):
        return markdown[2:].strip()
    raise Exception("invalid markdown, title not found.")
  
def generate_page(base_path: str, from_path: str, template_path: str, dest_path: str):
  
    print(f"Generating page from {from_path} to {dest_path} using {template_path}")
    
    with open(from_path, "r") as f:
        markdown = f.read()
        
    with open(template_path, "r") as f:
        template = f.read()

    content_node = markdown_to_html_node(markdown)
    content = content_node.to_html()
    title = extract_title(markdown)

    html = template.replace("{{ Title }}", title).replace("{{ Content }}", content)
    html = html.replace('href="/', f'href="{base_path}').replace('src="/', f'src="{base_path}')

    os.makedirs(os.path.dirname(dest_path), exist_ok=True)
    with open(dest_path, "w") as f:
        f.write(html)
        
def generate_pages_recursive(base_path: str, dir_path_content: str, template_path: str, dest_dir_path: str):
    for root, dirs, files in os.walk(dir_path_content):
        for file in files:
            if file.endswith(".md"):
                from_path = os.path.join(root, file)
                relative_path = os.path.relpath(from_path, dir_path_content)
                dest_path = os.path.join(dest_dir_path, relative_path[:-3] + ".html")
                generate_page(base_path, from_path, template_path, dest_path)