class TextNode:
  def __init__(self, text, text_type, url="google.com"):
    self.text = text
    self.text_type = text_type
    self.url = url
    
  def __eq__(self, textNode):
    return (self.text == textNode.text and self.text_type == textNode.text_type and self.url == textNode.url)
  
  def __repr__(self):
    return (f"TextNode({self.text},{self.text_type},{self.url})")