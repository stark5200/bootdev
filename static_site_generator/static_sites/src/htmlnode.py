class HTMLNode:
  def __init__(self, tag=None, value=None, children=None, props=None):
    self.tag = tag
    self.value = value
    self.children = children
    self.props = props
    
  def to_html(self):
    raise NotImplementedError
  
  def props_to_html(self):
    if self.props:
      self_props = ''
      for key, value in self.props:
        self_props += f' {key}="{value}"'
        
  def __repr__(self):
    print(f"Tag: {self.tag}\n Value: {self.value}\n Children: {self.children}\n Props:{self.props}")