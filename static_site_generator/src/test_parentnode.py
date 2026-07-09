import unittest

from parentnode import ParentNode
from leafnode import LeafNode


class TestLeafNode(unittest.TestCase):
    def test_eq(self):
        ln1 = LeafNode("p", "paragraph")
        ln2 = LeafNode("a", "anchor tag", {"src": "https://www.google.com"})
        node1 = ParentNode("div", [ln1, ln2])
        node2 = ParentNode("div", [ln1, ln2], None)
        self.assertEqual(node1, node2)
        
    def test_to_html(self):
        ln1 = LeafNode("p", "paragraph")
        ln2 = LeafNode("a", "anchor tag", {"src": "https://www.google.com"})
        
        node3 = ParentNode("div", [ln1, ln2])
        self.assertEqual('<div><p>paragraph</p><a src="https://www.google.com">anchor tag</a></div>', node3.to_html())
        
    def test_to_html_with_children(self):
        child_node = LeafNode("span", "child")
        parent_node = ParentNode("div", [child_node])
        self.assertEqual(parent_node.to_html(), "<div><span>child</span></div>")

    def test_to_html_with_grandchildren(self):
        grandchild_node = LeafNode("b", "grandchild")
        child_node = ParentNode("span", [grandchild_node])
        parent_node = ParentNode("div", [child_node])
        self.assertEqual(
            parent_node.to_html(),
            "<div><span><b>grandchild</b></span></div>",
        )

if __name__ == "__main__":
    unittest.main()