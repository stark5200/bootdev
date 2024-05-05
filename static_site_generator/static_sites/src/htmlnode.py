class HTMLNode:
  def __init__(self, tag=None, value=None, children=None, props=None):
    self.tag = tag
    self.value = value
    self.children = children
    self.props = props
    
  def to_html(self):
    raise NotImplementedError
  
  def props_to_html(self):
    self_props = ''
    if self.props:
      for key, value in self.props.items():
        self_props += f' {key}="{value}"'
    return self_props
        
  def __eq__(self, htmlNode):
    return (self.tag == htmlNode.tag and self.value == htmlNode.value and self.children == htmlNode.children and self.props == htmlNode.props)
        
  def __repr__(self):
    print(f"Tag: {self.tag}\n Value: {self.value}\n Children: {self.children}\n Props:{self.props}")