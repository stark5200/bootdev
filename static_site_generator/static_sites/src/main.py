from textnode import TextNode
from htmlnode import HTMLNode

def main():
  TN1 = TextNode("This is a text node", "bold", "https://www.boot.dev")
  print(TN1.__repr__())
  
  HN1 = HTMLNode("div", "This is a div")
  print(HN1.__repr__())
  
main()