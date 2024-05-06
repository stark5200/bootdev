import unittest

from htmlnode import HtmlNode


class TestHTMLNode(unittest.TestCase):
    def test_eq(self):
        node1 = HtmlNode("a", "anchor tag")
        node2 = HtmlNode("a", "anchor tag", None, None)
        self.assertEqual(node1, node2)
        
    def test_props_to_html(self):
        node3 = HtmlNode("a", "anchor tag", None, {"href": "https://www.google.com", "target": "_blank"})
        self.assertEqual(' href="https://www.google.com" target="_blank"', node3.props_to_html())
        


if __name__ == "__main__":
    unittest.main()
