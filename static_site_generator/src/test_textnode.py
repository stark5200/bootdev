import unittest
from textnode import *
from pages import *

class TestTextNode(unittest.TestCase):
    def test_eq(self):
        node = TextNode("This is a text node", TextType.BOLD)
        node2 = TextNode("This is a text node", TextType.BOLD)
        self.assertEqual(node, node2)
    
    def test_eq2(self):
        node3 = TextNode("2nd test with default url", TextType.ITALIC, "google.com")
        node4 = TextNode("2nd test with default url", TextType.ITALIC, "google.com")
        self.assertEqual(node3, node4)
        
    def test_eq3(self):
        node5 = TextNode("3rd test with None url", TextType.ITALIC, None)
        node6 = TextNode("3rd test with None url", TextType.ITALIC, None)
        self.assertEqual(node5, node6)
        
    def test_eq4(self):
        node7 = TextNode("4th should be different", TextType.TEXT, None)
        node8 = TextNode("4th will fail", TextType.ITALIC)
        self.assertNotEqual(node7, node8)
        
    def test_text_node_to_html_node(self):
        node = TextNode("This is a text node", TextType.TEXT)
        html_node = node.text_node_to_html_node(node)
        self.assertEqual(html_node.tag, None)
        self.assertEqual(html_node.value, "This is a text node")
        
    def test_code_delimiter(self):
        node = TextNode("This has `code` inside", TextType.TEXT)
        result = split_nodes_delimiter([node], "`", TextType.CODE)

        expected = [
            TextNode("This has ", TextType.TEXT),
            TextNode("code", TextType.CODE),
            TextNode(" inside", TextType.TEXT),
        ]

        self.assertEqual(result, expected)

    def test_bold_delimiter(self):
        node = TextNode("This has **bold** text", TextType.TEXT)
        result = split_nodes_delimiter([node], "**", TextType.BOLD)

        expected = [
            TextNode("This has ", TextType.TEXT),
            TextNode("bold", TextType.BOLD),
            TextNode(" text", TextType.TEXT),
        ]

        self.assertEqual(result, expected)

    def test_italic_delimiter(self):
        node = TextNode("This has _italic_ text", TextType.TEXT)
        result = split_nodes_delimiter([node], "_", TextType.ITALIC)

        expected = [
            TextNode("This has ", TextType.TEXT),
            TextNode("italic", TextType.ITALIC),
            TextNode(" text", TextType.TEXT),
        ]

        self.assertEqual(result, expected)

    def test_non_text_node_unchanged(self):
        node = TextNode("already bold", TextType.BOLD)
        result = split_nodes_delimiter([node], "**", TextType.BOLD)

        expected = [
            TextNode("already bold", TextType.BOLD),
        ]

        self.assertEqual(result, expected)

    def test_unmatched_delimiter_raises_exception(self):
        node = TextNode("This has `unclosed code", TextType.TEXT)

        with self.assertRaises(Exception):
            split_nodes_delimiter([node], "`", TextType.CODE)
            
    def test_split_nodes(self):
        node5 = TextNode("This is text with a `code block` word", TextType.TEXT)
        split_nodes_list1 = split_nodes_delimiter([node5], "`", TextType.CODE)
        split_nodes_list2 = [
          TextNode("This is text with a ", TextType.TEXT),
          TextNode("code block", TextType.CODE),
          TextNode(" word", TextType.TEXT),
        ]
        
        node6 = TextNode("Here is **Bold** part and here is **Thick** part", TextType.TEXT)
        split_nodes_list3 = split_nodes_delimiter([node6], "**", TextType.BOLD)
        split_nodes_list4 = [
          TextNode("Here is ", TextType.TEXT),
          TextNode("Bold", TextType.BOLD),
          TextNode(" part and here is ", TextType.TEXT),
          TextNode("Thick", TextType.BOLD),
          TextNode(" part", TextType.TEXT)
        ]
        
        for i in range(len(split_nodes_list2)):
          self.assertEqual(split_nodes_list1[i], split_nodes_list2[i])
          
        for i in range(len(split_nodes_list4)):
          self.assertEqual(split_nodes_list3[i], split_nodes_list4[i])
          
    """
    def test_regex(self):
        text1 = "This is text with an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and ![another](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png)"
        text1_matches = extract_markdown_images(text1)
        expected1 = [("image", "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"), ("another", "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png")]
        
        self.assertEqual(text1_matches, expected1)
        
        text2 = "This is text with a [link](https://www.example.com) and [another](https://www.example.com/another)"
        text2_matches = extract_markdown_links(text2)
        expected2 = [("link", "https://www.example.com"), ("another", "https://www.example.com/another")]
        self.assertEqual(text2_matches, expected2)

    
    def test_split_image_link(self):
        node1 = TextNode(
            "This is text with an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and another ![second image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png)",TextType.TEXT,)
        new_nodes = split_nodes_image([node1])
        expected = [
            TextNode("This is text with an ", TextType.TEXT),
            TextNode("image", TextType.IMAGE, "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"),
            TextNode(" and another ", TextType.TEXT),
            TextNode("second image", TextType.IMAGE, "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png"),
        ]
        self.assertEqual(new_nodes, expected)
"""

    def test_extract_markdown_images(self):
        matches = extract_markdown_images(
            "This is text with an ![image](https://i.imgur.com/zjjcJKZ.png)"
        )
        self.assertListEqual([("image", "https://i.imgur.com/zjjcJKZ.png")], matches)
        
    def test_split_images(self):
        node = TextNode(
            "This is text with an ![image](https://i.imgur.com/zjjcJKZ.png) and another ![second image](https://i.imgur.com/3elNhQu.png)",
            TextType.TEXT,
        )
        new_nodes = split_nodes_image([node])
        self.assertListEqual(
            [
                TextNode("This is text with an ", TextType.TEXT),
                TextNode("image", TextType.IMAGE, "https://i.imgur.com/zjjcJKZ.png"),
                TextNode(" and another ", TextType.TEXT),
                TextNode(
                    "second image", TextType.IMAGE, "https://i.imgur.com/3elNhQu.png"
                ),
            ],
            new_nodes,
        )
        
    def test_delim_bold(self):
        node = TextNode("This is text with a **bolded** word", TextType.TEXT)
        new_nodes = split_nodes_delimiter([node], "**", TextType.BOLD)
        self.assertListEqual(
            [
                TextNode("This is text with a ", TextType.TEXT),
                TextNode("bolded", TextType.BOLD),
                TextNode(" word", TextType.TEXT),
            ],
            new_nodes,
        )

    def test_delim_bold_double(self):
        node = TextNode(
            "This is text with a **bolded** word and **another**", TextType.TEXT
        )
        new_nodes = split_nodes_delimiter([node], "**", TextType.BOLD)
        self.assertListEqual(
            [
                TextNode("This is text with a ", TextType.TEXT),
                TextNode("bolded", TextType.BOLD),
                TextNode(" word and ", TextType.TEXT),
                TextNode("another", TextType.BOLD),
            ],
            new_nodes,
        )

    def test_delim_bold_multiword(self):
        node = TextNode(
            "This is text with a **bolded word** and **another**", TextType.TEXT
        )
        new_nodes = split_nodes_delimiter([node], "**", TextType.BOLD)
        self.assertListEqual(
            [
                TextNode("This is text with a ", TextType.TEXT),
                TextNode("bolded word", TextType.BOLD),
                TextNode(" and ", TextType.TEXT),
                TextNode("another", TextType.BOLD),
            ],
            new_nodes,
        )

    def test_delim_italic(self):
        node = TextNode("This is text with an _italic_ word", TextType.TEXT)
        new_nodes = split_nodes_delimiter([node], "_", TextType.ITALIC)
        self.assertListEqual(
            [
                TextNode("This is text with an ", TextType.TEXT),
                TextNode("italic", TextType.ITALIC),
                TextNode(" word", TextType.TEXT),
            ],
            new_nodes,
        )

    def test_delim_bold_and_italic(self):
        node = TextNode("**bold** and _italic_", TextType.TEXT)
        new_nodes = split_nodes_delimiter([node], "**", TextType.BOLD)
        new_nodes = split_nodes_delimiter(new_nodes, "_", TextType.ITALIC)
        self.assertEqual(
            [
                TextNode("bold", TextType.BOLD),
                TextNode(" and ", TextType.TEXT),
                TextNode("italic", TextType.ITALIC),
            ],
            new_nodes,
        )

    def test_delim_code(self):
        node = TextNode("This is text with a `code block` word", TextType.TEXT)
        new_nodes = split_nodes_delimiter([node], "`", TextType.CODE)
        self.assertListEqual(
            [
                TextNode("This is text with a ", TextType.TEXT),
                TextNode("code block", TextType.CODE),
                TextNode(" word", TextType.TEXT),
            ],
            new_nodes,
        )

    def test_extract_markdown_images(self):
        matches = extract_markdown_images(
            "This is text with an ![image](https://i.imgur.com/zjjcJKZ.png)"
        )
        self.assertListEqual([("image", "https://i.imgur.com/zjjcJKZ.png")], matches)

    def test_extract_markdown_links(self):
        matches = extract_markdown_links(
            "This is text with a [link](https://boot.dev) and [another link](https://wikipedia.org)"
        )
        self.assertListEqual(
            [
                ("link", "https://boot.dev"),
                ("another link", "https://wikipedia.org"),
            ],
            matches,
        )

    def test_split_image(self):
        node = TextNode(
            "This is text with an ![image](https://i.imgur.com/zjjcJKZ.png)",
            TextType.TEXT,
        )
        new_nodes = split_nodes_image([node])
        self.assertListEqual(
            [
                TextNode("This is text with an ", TextType.TEXT),
                TextNode("image", TextType.IMAGE, "https://i.imgur.com/zjjcJKZ.png"),
            ],
            new_nodes,
        )

    def test_split_image_single(self):
        node = TextNode(
            "![image](https://www.example.COM/IMAGE.PNG)",
            TextType.TEXT,
        )
        new_nodes = split_nodes_image([node])
        self.assertListEqual(
            [
                TextNode("image", TextType.IMAGE, "https://www.example.COM/IMAGE.PNG"),
            ],
            new_nodes,
        )

    def test_split_images(self):
        node = TextNode(
            "This is text with an ![image](https://i.imgur.com/zjjcJKZ.png) and another ![second image](https://i.imgur.com/3elNhQu.png)",
            TextType.TEXT,
        )
        new_nodes = split_nodes_image([node])
        self.assertListEqual(
            [
                TextNode("This is text with an ", TextType.TEXT),
                TextNode("image", TextType.IMAGE, "https://i.imgur.com/zjjcJKZ.png"),
                TextNode(" and another ", TextType.TEXT),
                TextNode(
                    "second image", TextType.IMAGE, "https://i.imgur.com/3elNhQu.png"
                ),
            ],
            new_nodes,
        )

    def test_split_links(self):
        node = TextNode(
            "This is text with a [link](https://boot.dev) and [another link](https://wikipedia.org) with text that follows",
            TextType.TEXT,
        )
        new_nodes = split_nodes_link([node])
        self.assertListEqual(
            [
                TextNode("This is text with a ", TextType.TEXT),
                TextNode("link", TextType.LINK, "https://boot.dev"),
                TextNode(" and ", TextType.TEXT),
                TextNode("another link", TextType.LINK, "https://wikipedia.org"),
                TextNode(" with text that follows", TextType.TEXT),
            ],
            new_nodes,
        )

    def test_text_to_textnodes(self):
        nodes = text_to_textnodes(
            "This is **text** with an _italic_ word and a `code block` and an ![image](https://i.imgur.com/zjjcJKZ.png) and a [link](https://boot.dev)"
        )
        self.assertListEqual(
            [
                TextNode("This is ", TextType.TEXT),
                TextNode("text", TextType.BOLD),
                TextNode(" with an ", TextType.TEXT),
                TextNode("italic", TextType.ITALIC),
                TextNode(" word and a ", TextType.TEXT),
                TextNode("code block", TextType.CODE),
                TextNode(" and an ", TextType.TEXT),
                TextNode("image", TextType.IMAGE, "https://i.imgur.com/zjjcJKZ.png"),
                TextNode(" and a ", TextType.TEXT),
                TextNode("link", TextType.LINK, "https://boot.dev"),
            ],
            nodes,
        )
        
    def test_markdown_to_blocks(self):
        md = """
This is **bolded** paragraph

This is another paragraph with _italic_ text and `code` here
This is the same paragraph on a new line

- This is a list
- with items
"""
        blocks = markdown_to_blocks(md)
        self.assertEqual(
            blocks,
            [
                "This is **bolded** paragraph",
                "This is another paragraph with _italic_ text and `code` here\nThis is the same paragraph on a new line",
                "- This is a list\n- with items",
            ],
        )
        
    def test_block_to_block_types(self):
        block = "# heading"
        self.assertEqual(block_to_block_type(block), BlockType.HEADING)
        block = "```\ncode\n```"
        self.assertEqual(block_to_block_type(block), BlockType.CODE)
        block = "> quote\n> more quote"
        self.assertEqual(block_to_block_type(block), BlockType.QUOTE)
        block = "- list\n- items"
        self.assertEqual(block_to_block_type(block), BlockType.ULIST)
        block = "1. list\n2. items"
        self.assertEqual(block_to_block_type(block), BlockType.OLIST)
        block = "paragraph"
        self.assertEqual(block_to_block_type(block), BlockType.PARAGRAPH)
        

    def test_headings(self):
        md = """
# this is an h1

this is paragraph text

## this is an h2
"""

        node = markdown_to_html_node(md)
        html = node.to_html()
        self.assertEqual(
            html,
            "<div><h1>this is an h1</h1><p>this is paragraph text</p><h2>this is an h2</h2></div>",
        )

    def test_code_block(self):
        md = """
```
This is text that _should_ remain
the **same** even with inline stuff
```
"""
        node = markdown_to_html_node(md)
        html = node.to_html()
        self.assertEqual(
            html,
            "<div><pre><code>This is text that _should_ remain\nthe **same** even with inline stuff\n</code></pre></div>",
        )
        
    def test_extract_title(self):
        md = "# This is a title"
        self.assertEqual(extract_title(md), "This is a title")
        md = "## This is a subtitle"
        self.assertEqual(extract_title(md), "This is a subtitle")
        md = "This is not a title"
        self.assertEqual(extract_title(md), "")


if __name__ == "__main__":
    unittest.main()