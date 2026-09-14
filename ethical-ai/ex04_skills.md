# Exercise 4: Building Irreplaceable Skills

## Self-Evaluation (1-5 Scale)

### 1. Problem decomposition — 4/5

I think I'm pretty good at breaking a big problem into smaller parts. For example, when I worked on binary tree level-order traversal, I separated the problem into using a queue, visiting the nodes and adding the children. It makes the problem easier to understand and solve step by step.

### 2. Systems thinking — 3/5

This is one of the skills I need to improve. I can understand individual data structures, but sometimes it's harder for me to understand what happens when many components or files work together. Pointer changes can also be confusing when the code becomes more complex.

### 3. Critical evaluation — 3/5

I can usually check if a solution makes sense, but I still have some things to improve. For example, I noticed that the AI suggested `strings.Fields`, but this is not allowed in the Piscine. However, I didn't notice at first that `word += string(c)` can also be inefficient with long strings.

### 4. Debugging mindset — 4/5

I think I'm good at finding bugs when I take the time to test things properly. For the whitespace bug, I checked the characters one by one and tested spaces, tabs and new lines to understand what was happening.

### 5. Conceptual understanding — 4/5

I understand the main idea behind the algorithms I work with. For example, I understand why BST deletion can need the in-order successor instead of just replacing the node randomly, because we still need to keep the BST rules.

---

## Lowest-Rated Skill: Systems Thinking (3/5)

I chose **Systems Thinking** as the main skill I want to work on this week.

### Action 1: Draw the structure before coding

Before coding operations with multiple nodes or pointers, like tree deletion or linked list operations, I will draw the structure on paper first.

I will try to show what the structure looks like before, during and after the operation. I will also check things like parent/child links and if the head and tail are still correct.

This should help me understand better what is actually changing in memory instead of just writing the code and hoping it works.

### Action 2: Test different components together

For every new function I work on, I will make a small `main.go` to test it together with other functions.

For example, instead of only testing one function, I can do:

`insert -> search -> delete -> traverse`

I will also try some difficult cases, like an empty list, a tree with only one node, or unexpected inputs.

The goal is to understand how all the parts work together and not only check if one function works by itself.
