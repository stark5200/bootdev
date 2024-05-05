import unittest

from htmlnode import HTMLNode


class TestHTMLNode(unittest.TestCase):
    def test_eq(self):
        node1 = HTMLNode("a", "anchor tag")
        node2 = HTMLNode("a", "anchor tag", None, None)
        self.assertEqual(node1, node2)
        
    def test_props_to_html(self):
        node3 = HTMLNode("a", "anchor tag", None, {"href": "https://www.google.com", "target": "_blank"})
        self.assertEqual(' href="https://www.google.com" target="_blank"', node3.props_to_html())
        


if __name__ == "__main__":
    unittest.main()
