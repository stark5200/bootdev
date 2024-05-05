import unittest

from textnode import TextNode


class TestTextNode(unittest.TestCase):
    def test_eq(self):
        node1 = TextNode("This is a text node", "bold")
        node2 = TextNode("This is a text node", "bold")
        self.assertEqual(node1, node2)
    
    def test_eq2(self):
        node3 = TextNode("2nd test with default url", "italic", "google.com")
        node4 = TextNode("2nd test with default url", "italic")
        self.assertEqual(node3, node4)


if __name__ == "__main__":
    unittest.main()
