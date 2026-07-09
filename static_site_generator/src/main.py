import os
import sys
import shutil
from textnode import TextNode
from htmlnode import HtmlNode
from leafnode import LeafNode
from pages import generate_page, generate_pages_recursive
from copystatic import copy_files_recursive

dir_path_static = "./static"
dir_path_docs = "./docs"

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
    
    if os.path.exists(dir_path_docs):
        print("Deleting docs directory...")
        shutil.rmtree(dir_path_docs)

    print("Copying static files to docs directory...")
    copy_files_recursive(dir_path_static, dir_path_docs)
    
    basepath = sys.argv[1] if len(sys.argv) > 1 else "/"
    
    generate_pages_recursive(basepath, "./content", "./template.html", "./docs")
  
main()