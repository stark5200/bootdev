import unittest

from textnode import (
  TextNode,
  text_type_text,  
  text_type_bold,
  text_type_italic,
  text_type_code,
  text_type_link,
  text_type_image,
)


class TestTextNode(unittest.TestCase):
    def test_eq(self):
        node1 = TextNode("This is a text node", "bold")
        node2 = TextNode("This is a text node", "bold")
        self.assertEqual(node1, node2)
    
    def test_eq2(self):
        node3 = TextNode("2nd test with default url", text_type_italic, "google.com")
        node4 = TextNode("2nd test with default url", text_type_italic)
        self.assertEqual(node3, node4)
        
    def test_split_nodes(self):
        node5 = TextNode("This is text with a `code block` word", text_type_text)
        split_nodes_list1 = TextNode.split_nodes_delimiter([node5], "`", text_type_code)
        split_nodes_list2 = [
          TextNode("This is text with a ", text_type_text),
          TextNode("code block", text_type_code),
          TextNode(" word", text_type_text),
        ]
        
        node6 = TextNode("Here is **Bold** part and here is **Thick** part", text_type_text)
        split_nodes_list3 = TextNode.split_nodes_delimiter([node6], "**", text_type_bold)
        split_nodes_list4 = [
          TextNode("Here is ", text_type_text),
          TextNode("Bold", text_type_bold),
          TextNode(" part and here is ", text_type_text),
          TextNode("Thick", text_type_bold),
          TextNode(" part", text_type_text)
        ]
        
        for i in range(len(split_nodes_list2)):
          self.assertEqual(split_nodes_list1[i], split_nodes_list2[i])
          
        for i in range(len(split_nodes_list4)):
          self.assertEqual(split_nodes_list3[i], split_nodes_list4[i])


if __name__ == "__main__":
    unittest.main()
