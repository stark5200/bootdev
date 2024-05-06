from textnode import TextNode
from htmlnode import HtmlNode
from leafnode import LeafNode

def main():
  tn1 = TextNode("This is a text node", "bold", "https://www.boot.dev")
  print(tn1.__repr__())
  
  hn1 = HtmlNode("div", "This is a div")
  print(hn1.__repr__())
  
  ln1 = HtmlNode("p", "This is a paragraph")
  ln2 = HtmlNode("a", "This is an anchor tag", {"src":"mypicture.com"})
  print(ln1.__repr__())
  print(ln2.__repr__())
  
main()