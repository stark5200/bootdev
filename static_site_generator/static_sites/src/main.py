from textnode import TextNode

def main():
  TN1 = TextNode("This is a text node", "bold", "https://www.boot.dev")
  print(TN1.__repr__())
  
main()