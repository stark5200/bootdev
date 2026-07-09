from htmlnode import HtmlNode

class ParentNode(HtmlNode):
  def __init__(self, tag, children, props=None):
    if children != None and tag != None:
      super().__init__(tag, None, children, props)
    else:
      raise ValueError("children/tag are mandatory")
    
  def to_html(self):
    if self.tag == None:
      raise ValueError("tag is mandatory")
    if self.children == None:
      raise ValueError("children are mandatory")
    to_html_message = ''
    for c in self.children:
      to_html_message += c.to_html()  
    return f'<{self.tag}{self.props_to_html()}>{to_html_message}</{self.tag}>'  
  
  def __eq__(self, parentNode):
    return (self.tag == parentNode.tag and self.children == parentNode.children and self.props == parentNode.props)
  
  
  def __repr__(self):
    print(f"Tag: {self.tag}\n Children: {self.children}\n Props:{self.props}")
    