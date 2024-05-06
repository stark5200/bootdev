import unittest

from leafnode import LeafNode


class TestLeafNode(unittest.TestCase):
    def test_eq(self):
        node1 = LeafNode("p", "paragraph")
        node2 = LeafNode("p", "paragraph", None)
        self.assertEqual(node1, node2)
        
    def test_props_to_html(self):
        node3 = LeafNode("a", "anchor tag", {"src": "https://www.google.com"})
        self.assertEqual(' src="https://www.google.com"', node3.props_to_html())
        
    def test_to_html(self):
        node4 = LeafNode("a", "anchor tag", {"src": "https://www.google.com"})
        self.assertEqual('<a src="https://www.google.com">anchor tag</a>', node4.to_html())
        
        node5 = LeafNode("p", "Lets go for round two.")
        self.assertEqual('<p>Lets go for round two.</p>', node5.to_html())
        
        node6 = LeafNode(None, "No fluff.")
        self.assertEqual('No fluff.', node6.to_html())
        


if __name__ == "__main__":
    unittest.main()
