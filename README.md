# Link Collector using Go
> An easy but not practical way to collect links from a string.

## This file in other languages
[简体中文](zh-simplified.md)  
[繁體中文](zh-traditional.md)  
[한국어](kr.md)  
[日本語](jp.md)  

## Brief Introduction
This program can collect links from sources that can be transformed into a string. It's not a huge program and the logic inside isn't hard to comprehend. However, it doesn't promise to collect all links correctly and the reasons are as follow.
- When writing this little program, I was supposing that all links are surrounded by single/double quotes. So there is a chance that it won't cover all the links.
- I also decided to collect specific types of links including "http://..." and "https://". All others are uninteresting for this program.
