import unittest

from textnode import (
  TextNode,
  text_type_text,  
)


class TestTextNode(unittest.TestCase):
    def test_eq(self):
        node1 = TextNode("This is a text node", "bold")
        node2 = TextNode("This is a text node", "bold")
        self.assertEqual(node1, node2)
    
    def test_eq2(self):
        node3 = TextNode("2nd test with default url", "italic", "google.com")
        node4 = TextNode("2nd test with default url", "italic")
        self.assertEqual(node3, node4)
        
    def test_split_nodes(self):
        node = TextNode("This `is text` with a `code block` word", TextNode.text_type_text)
        split_nodes_list1 = TextNode.split_nodes_delimiter([node], "`", TextNode.text_type_code)
        split_nodes_list2 = [
          TextNode("This is text with a ", TextNode.text_type_text),
          TextNode("code block", TextNode.text_type_code),
          TextNode(" word", TextNode.text_type_text),
        ]

        for i in range(3):
          self.assertEqual(split_nodes_list1[i], split_nodes_list2[i])


if __name__ == "__main__":
    unittest.main()
