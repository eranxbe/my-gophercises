# Notes
html link parser, uses net/html, dfs and recursion to traverse the DOM.

## net/html package:
- parse html string
- using nodes and node types to determine the data and attributes
- use DFS to collect all a tags
- use DFS to aggregate the text of each a tag and its children