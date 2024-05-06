from leafnode import LeafNode

class TextNode:
  text_type_text = "text"
  text_type_bold = "bold"
  text_type_italic = "italic"
  text_type_code = "code"
  text_type_link = "link"
  text_type_image = "image"
  
  text_types = [
    text_type_text,
    text_type_bold,
    text_type_italic,
    text_type_code,
    text_type_link,
    text_type_image
  ]
  
  def __init__(self, text, text_type, url="google.com"):
    self.text = text
    self.text_type = text_type
    self.url = url
    
  def text_node_to_html_node(self, text_node):
    if text_node.text_type not in self.text_types :
      raise Exception("text node type not valid")
    if text_node.text_type == self.text_type_text:
      return LeafNode(None, self.text)
    if text_node.text_type == self.text_type_bold:
      return LeafNode("b", self.text)
    if text_node.text_type == self.text_type_italic:
      return LeafNode("i", self.text)
    if text_node.text_type == self.text_type_code:
      return LeafNode("code", self.text)
    if text_node.text_type == self.text_type_link:
      return LeafNode("a", self.text, {"href": self.url})
    if text_node.text_type == self.text_type_image:
      return LeafNode(tag="img", value="", props={"src": self.url, "alt":self.text})
    
  def __eq__(self, textNode):
    return (self.text == textNode.text and self.text_type == textNode.text_type and self.url == textNode.url)
  
  def __repr__(self):
    return (f"TextNode({self.text},{self.text_type},{self.url})")