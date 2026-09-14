# Ethical AI Quest: Learning Amplifier vs. Shortcut

This project is about how i used AI during my programming work. The main goal was to use AI to help me learn and improve, not to just give me the answers or code.

---

## 1. Summary: Using AI to Learn, Not Skip Learning

During this quest,i  understood more the difference between using AI to **do the work** and using it to **help me understand**.

If i use AI just to generate code and copy it, i can miss the important part of learning. i need to struggle a bit with the problem, break it down and try to find the solution myself. Otherwise, i might think i understand something when actually i don't.

This can become a problem during a technical interview, an audit or when something breaks and i don't have AI available.

i think AI is more useful as a learning tool. First i try to solve the problem myself, make a hypothesis and write my own code. After that, i can ask AI to review it, find edge cases or point out things i didn't notice.

---

## 2. Most Important Insight (Exercise 1)

The most important thing i learned came from **Exercise 1**.

i learned that **AI can give an answer that sounds very correct, but it can still be wrong for the specific situation.**

For example, when i asked AI how to improve my `SplitWhiteSpaces` function, it suggested using:

```go
strings.Fields(s)
```

Normally this is a good and clean solution in Go. But for the Piscine, the `strings` package is not allowed for this exercise.

So if i had just copied the AI answer without checking the rules, my solution could fail the Piscine tests.

This showed me that i can't just trust AI because the answer sounds confident. i still need to understand the exercise and check if the solution follows the actual rules.

---

## 3. How I Will Use This for the Rest of the Curriculum

For the rest of the Piscine and my future programming studies, i want to keep using the **"Hypothesis Before Query"** method.

Before asking AI, i will first try to understand the problem, break it into smaller parts and test my own solution.

After i have a working solution, i can use AI to review it and ask about things like edge cases, performance or possible problems i didn't see.

Before submitting an exercise or doing an audit, I should also be able to explain my code without AI. I should know why I used a certain data structure or algorithm and be able to modify the code myself if someone asks me to.

The goal is not to stop using AI, but to make sure that **I am still the one doing the thinking.**

---

## Repository Structure

```text
ethical-ai/

├── README.md                  # Quest summary and main lessons

├── ex01_buggy.go              # First version with the whitespace bug

├── ex01_fixed.go              # My fixed version without AI

├── ex01_debug.md              # Bug explanation and AI review

├── ex02_fairness_contract.md  # Rules for using AI fairly

├── ex03_scenarios.md          # Problems caused by depending too much on AI

└── ex04_skills.md             # My skills evaluation and improvement plan
```
