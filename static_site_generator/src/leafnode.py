from htmlnode import HtmlNode

class LeafNode(HtmlNode):
  def __init__(self, tag, value, props=None):
    if value != None:
      super().__init__(tag, value, None, props)
    else:
      raise ValueError("value is mandatory")
    
  def to_html(self):
    if self.value == None:
      raise ValueError("value is mandatory")
    else:
      if self.tag:
        if self.props==None:
          return f'<{self.tag}>{self.value}</{self.tag}>'
        else:
          return f'<{self.tag}{self.props_to_html()}>{self.value}</{self.tag}>'
      else:
        return f'{self.value}'
        
  
  def __eq__(self, leafNode):
    return (self.tag == leafNode.tag and self.value == leafNode.value and self.props == leafNode.props)
  
        
  def __repr__(self):
    print(f"Tag: {self.tag}\n Value: {self.value}\n Props:{self.props}")
    