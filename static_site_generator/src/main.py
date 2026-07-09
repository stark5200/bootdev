import os
import shutil
from textnode import TextNode
from htmlnode import HtmlNode
from leafnode import LeafNode
from pages import generate_page, generate_pages_recursive

from copystatic import copy_files_recursive

dir_path_static = "./static"
dir_path_public = "./public"

def main():
    print("hello world") 
    tn1 = TextNode("This is a text node", "bold", "https://www.boot.dev")
    print(tn1.__repr__())
    
    hn1 = HtmlNode("div", "This is a div")
    print(hn1.__repr__())
    
    ln1 = HtmlNode("p", "This is a paragraph")
    ln2 = HtmlNode("a", "This is an anchor tag", {"src":"mypicture.com"})
    print(ln1.__repr__())
    print(ln2.__repr__())
    
    print("Deleting public directory...")
    if os.path.exists(dir_path_public):
        shutil.rmtree(dir_path_public)

    print("Copying static files to public directory...")
    copy_files_recursive(dir_path_static, dir_path_public)
    
    
    generate_pages_recursive("./content", "./template.html", "./public")
  
main()