from leafnode import LeafNode

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
class TextNode:
  
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
  
  
  
  def split_nodes_delimiter(old_nodes: list, delimiter: str, text_type: str) -> list:
    new_text_nodes = []
    for node in old_nodes:
      split_node = node.text.split(delimiter)
      if len(split_node) % 2 != 1:
        continue
      
      for i in len(split_node):
        type = text_type_text
        if i % 2 == 1:
          type = text_type
        new_text_nodes.append(TextNode(split_node[i], type))
    return new_text_nodes