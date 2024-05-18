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
        node4 = TextNode("2nd test with default url", text_type_italic, "google.com")
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
          
    def test_regex(self):
        text1 = "This is text with an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and ![another](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png)"
        text1_matches = TextNode.extract_markdown_images(text1)
        expected1 = [("image", "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"), ("another", "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png")]
        
        self.assertEqual(text1_matches, expected1)
        
        text2 = "This is text with a [link](https://www.example.com) and [another](https://www.example.com/another)"
        text2_matches = TextNode.extract_markdown_links(text2)
        expected2 = [("link", "https://www.example.com"), ("another", "https://www.example.com/another")]
        self.assertEqual(text2_matches, expected2)

    
    def test_split_image_link(self):
      node1 = TextNode(
    "This is text with an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and another ![second image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png)",text_type_text,)
      new_nodes = TextNode.split_nodes_image([node1])
      expected = [
        TextNode("This is text with an ", text_type_text),
        TextNode("image", text_type_image, "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"),
        TextNode(" and another ", text_type_text),
        TextNode("second image", text_type_image, "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png"),
      ]
        



if __name__ == "__main__":
    unittest.main()
